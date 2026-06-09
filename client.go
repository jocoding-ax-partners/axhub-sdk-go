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
	// schemaCache backs the ergonomic data layer's Data().Discover(); it lives
	// on the root client so discover() across separate Tenant().App() chains
	// shares one cache (mirrors the node single per-SDK DataClient cache).
	dataSchemaCache *SchemaCache
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
	c.dataSchemaCache = NewSchemaCache(SchemaCacheOptions{})
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
	var values url.Values
	if len(query) > 0 {
		values = url.Values{}
		for k, v := range query {
			values.Set(k, v)
		}
	}
	decoded, err := c.doHTTP(ctx, route.Method, path, values, body, isFormEncodedOperation(operationID))
	if err != nil {
		return nil, err
	}
	// The operation-id route table camelizes snake_case response keys so the
	// conformance vectors can assert camelCase fields.
	if m, ok := decoded.(map[string]any); ok {
		return camelizeMap(m), nil
	}
	if decoded == nil {
		return map[string]any{}, nil
	}
	return map[string]any{"value": camelize(decoded)}, nil
}

// requestRaw is the data-layer transport: a path-based request that returns the
// decoded JSON WITHOUT snake->camel rewriting, so dynamic-table row data and the
// list envelope come back verbatim (mirrors python request_raw). When
// camelizeResponse is true the response IS camelized — discover() uses that so
// `tableName`/`table_name` both resolve on the inspect metadata payload.
func (c *Client) requestRaw(ctx context.Context, method, path string, query url.Values, body any, camelizeResponse bool) (map[string]any, error) {
	decoded, err := c.doHTTP(ctx, method, path, query, body, false)
	if err != nil {
		return nil, err
	}
	m, ok := decoded.(map[string]any)
	if !ok {
		if decoded == nil {
			return map[string]any{}, nil
		}
		return map[string]any{"value": decoded}, nil
	}
	if camelizeResponse {
		return camelizeMap(m), nil
	}
	return m, nil
}

// doHTTP performs auth + request-id + body encoding + status handling shared by
// the operation-id transport (Request) and the data-layer transport
// (requestRaw). It returns the decoded JSON body (any), leaving camelize policy
// to the caller.
func (c *Client) doHTTP(ctx context.Context, method, path string, query url.Values, body any, formEncoded bool) (any, error) {
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, err
	}
	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}
	var reader io.Reader
	if body != nil {
		if formEncoded {
			reader = strings.NewReader(formValues(body).Encode())
		} else {
			raw, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}
			reader = bytes.NewReader(raw)
		}
	}
	req, err := http.NewRequestWithContext(ctx, method, u.String(), reader)
	if err != nil {
		return nil, err
	}
	if body != nil {
		if formEncoded {
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
	return decoded, nil
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
