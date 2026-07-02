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
	"strconv"
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
	// conformance vectors can assert camelCase fields. OAuth token-style
	// responses preserve RFC 6749 standard keys in snake_case (contract shared
	// with the python SDK's _camelize_oauth_response).
	oauthSnake := isOauthSnakeCaseResponseOperation(operationID)
	if m, ok := decoded.(map[string]any); ok {
		if oauthSnake {
			return camelizeOauthMap(m), nil
		}
		return camelizeMap(m), nil
	}
	if decoded == nil {
		return map[string]any{}, nil
	}
	if oauthSnake {
		return map[string]any{"value": camelizeOauth(decoded)}, nil
	}
	return map[string]any{"value": camelize(decoded)}, nil
}

// doHTTP performs auth + request-id + body encoding + status handling for the
// operation-id transport (Request). It returns the decoded JSON body (any),
// leaving camelize policy to the caller.
func (c *Client) doHTTP(ctx context.Context, method, path string, query url.Values, body any, formEncoded bool) (any, error) {
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, err
	}
	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}
	// Encode the body once so it can be replayed across 429 retries (an
	// io.Reader is single-use). contentType empty means no body.
	var bodyBytes []byte
	var contentType string
	if body != nil {
		if formEncoded {
			bodyBytes = []byte(formValues(body).Encode())
			contentType = "application/x-www-form-urlencoded"
		} else {
			raw, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}
			bodyBytes = raw
			contentType = "application/json"
		}
	}

	// 429 backoff: honor the Retry-After header (seconds) and retry up to
	// maxRateLimitRetries times, falling back to defaultRetryAfter when the
	// header is absent. Mirrors the node core http.ts/ratelimit.ts sleep path.
	for attempt := 0; ; attempt++ {
		var reader io.Reader
		if bodyBytes != nil {
			reader = bytes.NewReader(bodyBytes)
		}
		req, err := http.NewRequestWithContext(ctx, method, u.String(), reader)
		if err != nil {
			return nil, err
		}
		if contentType != "" {
			req.Header.Set("Content-Type", contentType)
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
		raw, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode == http.StatusTooManyRequests && attempt < maxRateLimitRetries {
			if err := sleepWithContext(ctx, retryAfterDuration(resp.Header.Get("Retry-After"))); err != nil {
				return nil, err
			}
			continue
		}
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
}

const (
	// maxRateLimitRetries bounds the 429 retry loop (initial attempt + this
	// many retries). Matches the node core's bounded retry policy.
	maxRateLimitRetries = 2
	// defaultRetryAfter is the backoff used when a 429 omits Retry-After;
	// the backend's ratelimit middleware hardcodes 60s (ratelimit.go:34).
	defaultRetryAfter = 60 * time.Second
)

// retryAfterDuration parses an HTTP Retry-After header (RFC 7231 §7.1.3):
// either numeric seconds or an HTTP-date. Falls back to defaultRetryAfter when
// the header is missing or unparseable. Mirrors node core parseRetryAfter.
func retryAfterDuration(header string) time.Duration {
	header = strings.TrimSpace(header)
	if header == "" {
		return defaultRetryAfter
	}
	if secs, err := strconv.ParseFloat(header, 64); err == nil {
		if secs < 0 {
			return defaultRetryAfter
		}
		return time.Duration(secs * float64(time.Second))
	}
	if t, err := http.ParseTime(header); err == nil {
		if d := time.Until(t); d > 0 {
			return d
		}
		return 0
	}
	return defaultRetryAfter
}

// sleepWithContext waits for d or until ctx is cancelled, whichever comes
// first. A non-positive d returns immediately.
func sleepWithContext(ctx context.Context, d time.Duration) error {
	if d <= 0 {
		return nil
	}
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

var pathParamRe = regexp.MustCompile(`\{[^}]+\}`)

func unresolvedPathParam(path string) bool { return pathParamRe.MatchString(path) }

var formEncodedOperations = map[string]bool{
	"authPostOauthDeviceAuthorization": true,
	"authPostOauthRevoke":              true,
	"authPostOauthToken":               true,
}

func isFormEncodedOperation(operationID string) bool { return formEncodedOperations[operationID] }

// oauthSnakeCaseResponseOperations lists token-style OAuth operations whose
// responses keep RFC 6749/8628 standard keys in snake_case. Must match the
// python SDK's _OAUTH_RESPONSE_SNAKE_CASE_OPERATIONS.
var oauthSnakeCaseResponseOperations = map[string]bool{
	"authPostOauthDeviceAuthorization": true,
	"authPostOauthToken":               true,
}

func isOauthSnakeCaseResponseOperation(operationID string) bool {
	return oauthSnakeCaseResponseOperations[operationID]
}

// oauthResponseSnakeKeys are the RFC 6749 standard response keys preserved
// verbatim. Must match the python SDK's _OAUTH_RESPONSE_SNAKE_KEYS.
var oauthResponseSnakeKeys = map[string]bool{
	"access_token": true, "token_type": true, "expires_in": true, "refresh_token": true,
	"id_token": true, "scope": true, "resource": true, "tenant": true,
}

func camelizeOauthMap(in map[string]any) map[string]any {
	out := map[string]any{}
	for k, v := range in {
		key := k
		if !oauthResponseSnakeKeys[k] {
			key = snakeToCamel(k)
		}
		out[key] = camelizeOauth(v)
	}
	return out
}

func camelizeOauth(v any) any {
	switch t := v.(type) {
	case map[string]any:
		return camelizeOauthMap(t)
	case []any:
		for i, x := range t {
			t[i] = camelizeOauth(x)
		}
		return t
	default:
		return v
	}
}

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

type errorPayload struct {
	Category       string       `json:"category"`
	Code           string       `json:"code"`
	Message        string       `json:"message"`
	RequestID      string       `json:"request_id"`
	RequestIDCamel string       `json:"requestId"`
	Retryable      *bool        `json:"retryable"`
	Resource       string       `json:"resource"`
	Fields         []FieldError `json:"fields"`
	Retry          *RetryInfo   `json:"retry"`
	DocURL         string       `json:"doc_url"`
}

func parseError(status int, raw []byte) error {
	var env struct {
		Error errorPayload `json:"error"`
	}
	_ = json.Unmarshal(raw, &env)
	// W2: non-envelope errors arrive as a bare {code,message,...} at the root
	// (no "error" wrapper, e.g. some 428 preconditions). Fall back to a root-level
	// parse so the real code/category survive instead of masking as http_<status>.
	if env.Error.Code == "" {
		var root errorPayload
		if json.Unmarshal(raw, &root) == nil && root.Code != "" {
			env.Error = root
		}
	}
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
	return &AxHubError{
		Category:  env.Error.Category,
		Code:      env.Error.Code,
		Status:    status,
		Message:   env.Error.Message,
		RequestID: requestID,
		Retryable: retryable,
		Resource:  env.Error.Resource,
		Fields:    env.Error.Fields,
		Retry:     env.Error.Retry,
		DocURL:    env.Error.DocURL,
	}
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
