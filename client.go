package axhub

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"
)

type TokenType string

const (
	TokenTypePAT TokenType = "pat"
	TokenTypeJWT TokenType = "jwt"
)

type Config struct {
	BaseURL, Token, DefaultTenantID, DefaultTenantSlug string
	TokenType                                          TokenType
	HTTPClient                                         *http.Client
}

type Client struct {
	baseURL, token, defaultTenantID, defaultTenantSlug string
	tokenType                                          TokenType
	http                                               *http.Client
	Apps                                               *AppsClient
}
type AppsClient struct{ client *Client }

func NewClient(cfg Config) *Client {
	base := strings.TrimRight(cfg.BaseURL, "/")
	if base == "" {
		base = "https://api.axhub.ai"
	}
	hc := cfg.HTTPClient
	if hc == nil {
		hc = http.DefaultClient
	}
	hc = noRedirectClient(hc)
	c := &Client{baseURL: base, token: cfg.Token, tokenType: cfg.TokenType, defaultTenantID: cfg.DefaultTenantID, defaultTenantSlug: cfg.DefaultTenantSlug, http: hc}
	c.Apps = &AppsClient{client: c}
	return c
}
func (c *Client) BaseURL() string { return c.baseURL }
func (c *Client) RedactedToken() string {
	if c.token == "" {
		return ""
	}
	return "***REDACTED***"
}

func (a *AppsClient) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	tenant := a.client.defaultTenantID
	if tenant == "" {
		return nil, &AxHubError{Category: "tenant_id_required", Code: "tenant_id_required", Message: "default tenant id is required"}
	}
	return a.client.Request(ctx, "appsPostApiV1TenantsByTenantIDApps", map[string]string{"tenantID": tenant}, nil, body)
}

func (c *Client) Request(ctx context.Context, operationID string, pathParams map[string]string, query map[string]string, body any) (map[string]any, error) {
	route, ok := routeByOperation[operationID]
	if !ok {
		return nil, &AxHubError{Category: "validation", Code: "unknown_operation", Message: operationID}
	}
	path := route.Path
	for k, v := range pathParams {
		path = strings.ReplaceAll(path, "{"+k+"}", url.PathEscape(v))
	}
	if unresolvedPathParam(path) {
		return nil, &AxHubError{Category: "validation", Code: "required", Message: "missing path parameter for " + path}
	}
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, err
	}
	if len(query) > 0 {
		q := u.Query()
		for k, v := range query {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}
	var reader io.Reader
	if body != nil {
		if isFormEncodedOperation(operationID) {
			reader = strings.NewReader(formValues(body).Encode())
		} else {
			raw, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}
			reader = bytes.NewReader(raw)
		}
	}
	req, err := http.NewRequestWithContext(ctx, route.Method, u.String(), reader)
	if err != nil {
		return nil, err
	}
	if body != nil {
		if isFormEncodedOperation(operationID) {
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		} else {
			req.Header.Set("Content-Type", "application/json")
		}
	}
	req.Header.Set("X-Request-ID", newRequestID())
	if c.token != "" {
		if c.tokenType == TokenTypePAT {
			req.Header.Set("X-Api-Key", c.token)
		} else if c.tokenType == TokenTypeJWT {
			req.Header.Set("Authorization", "Bearer "+c.token)
		} else {
			return nil, &AxHubError{Category: "validation", Code: "required", Message: "tokenType must be pat or jwt"}
		}
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		return map[string]any{"status": resp.StatusCode, "location": resp.Header.Get("Location")}, nil
	}
	if resp.StatusCode >= 400 {
		return nil, parseError(resp.StatusCode, raw)
	}
	if len(strings.TrimSpace(string(raw))) == 0 {
		return map[string]any{}, nil
	}
	var decoded any
	if err := json.Unmarshal(raw, &decoded); err != nil {
		return map[string]any{"raw": string(raw)}, nil
	}
	if m, ok := decoded.(map[string]any); ok {
		return camelizeMap(m), nil
	}
	return map[string]any{"value": camelize(decoded)}, nil
}

var pathParamRe = regexp.MustCompile(`\{[^}]+\}`)

func unresolvedPathParam(path string) bool { return pathParamRe.MatchString(path) }

var formEncodedOperations = map[string]bool{
	"authPostOauthDeviceAuthorization": true,
	"authPostOauthRevoke":              true,
	"authPostOauthToken":               true,
}

func isFormEncodedOperation(operationID string) bool { return formEncodedOperations[operationID] }

func formValues(body any) url.Values {
	values := url.Values{}
	rv := reflect.ValueOf(body)
	if rv.Kind() == reflect.Map {
		for _, key := range rv.MapKeys() {
			values.Set(fmt.Sprint(key.Interface()), fmt.Sprint(rv.MapIndex(key).Interface()))
		}
	}
	return values
}

func noRedirectClient(hc *http.Client) *http.Client {
	clone := *hc
	clone.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return &clone
}

func parseError(status int, raw []byte) error {
	var env struct {
		Error struct {
			Category       string `json:"category"`
			Code           string `json:"code"`
			Message        string `json:"message"`
			RequestID      string `json:"request_id"`
			RequestIDCamel string `json:"requestId"`
			Retryable      *bool  `json:"retryable"`
		} `json:"error"`
	}
	_ = json.Unmarshal(raw, &env)
	if env.Error.Code == "" {
		env.Error.Code = fmt.Sprintf("http_%d", status)
	}
	info, hasInfo := ErrorCodes[env.Error.Code]
	if env.Error.Category == "" {
		if hasInfo {
			env.Error.Category = info.Category
		} else {
			env.Error.Category = "unknown"
		}
	}
	requestID := env.Error.RequestID
	if requestID == "" {
		requestID = env.Error.RequestIDCamel
	}
	retryable := hasInfo && info.Retryable
	if env.Error.Retryable != nil {
		retryable = *env.Error.Retryable
	}
	return &AxHubError{Category: env.Error.Category, Code: env.Error.Code, Status: status, Message: env.Error.Message, RequestID: requestID, Retryable: retryable}
}

func camelizeMap(in map[string]any) map[string]any {
	out := map[string]any{}
	for k, v := range in {
		out[snakeToCamel(k)] = camelize(v)
	}
	return out
}
func camelize(v any) any {
	switch t := v.(type) {
	case map[string]any:
		return camelizeMap(t)
	case []any:
		for i, x := range t {
			t[i] = camelize(x)
		}
		return t
	default:
		return v
	}
}
func snakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := 1; i < len(parts); i++ {
		if parts[i] != "" {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

func newRequestID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	enc := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(b)
	return fmt.Sprintf("%010d%s", time.Now().UnixMilli()%10000000000, enc)[:26]
}
