package axhub

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestRegressionAppsCreateConformance(t *testing.T) {
	var seenMethod, seenPath, seenAPIKey, seenRequestID string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seenMethod, seenPath = r.Method, r.URL.Path
		seenAPIKey = r.Header.Get("X-Api-Key")
		seenRequestID = r.Header.Get("X-Request-ID")
		_ = json.NewEncoder(w).Encode(map[string]any{"id": "app_1", "tenant_id": "tnt_1", "slug": "my-app", "schema_name": "app_my-app"})
	}))
	defer srv.Close()

	client := NewClient(Config{BaseURL: srv.URL, Token: "pat_x", TokenType: TokenTypePAT, DefaultTenantID: "tnt_1"})
	got, err := client.Apps.Create(context.Background(), map[string]any{"slug": "my-app", "name": "My App"})
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}
	if got["id"] != "app_1" || got["schemaName"] != "app_my-app" {
		t.Fatalf("unexpected response: %#v", got)
	}
	if seenMethod != "POST" || seenPath != "/api/v1/tenants/tnt_1/apps" {
		t.Fatalf("wrong request %s %s", seenMethod, seenPath)
	}
	if seenAPIKey != "pat_x" {
		t.Fatalf("missing PAT auth header: %q", seenAPIKey)
	}
	if seenRequestID == "" {
		t.Fatalf("missing X-Request-ID")
	}
}

func TestRegressionTenantRequiredBeforeRequest(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { called = true }))
	defer srv.Close()
	client := NewClient(Config{BaseURL: srv.URL, Token: "pat_x", TokenType: TokenTypePAT})
	_, err := client.Apps.Create(context.Background(), map[string]any{"slug": "my-app"})
	if err == nil {
		t.Fatal("expected tenant required error")
	}
	axErr, ok := err.(*AxHubError)
	if !ok || axErr.Category != "tenant_id_required" || axErr.Code != "tenant_id_required" {
		t.Fatalf("wrong error: %#v", err)
	}
	if called {
		t.Fatalf("request should not be sent when tenant context is missing")
	}
}

func TestRegressionErrorMappingAndRouteCoverage(t *testing.T) {
	if len(Routes) != 222 {
		t.Fatalf("route coverage drift: got %d", len(Routes))
	}
	if len(ErrorCodes) != 86 {
		t.Fatalf("error code drift: got %d", len(ErrorCodes))
	}
	info, ok := ErrorCodes["slug_taken"]
	if !ok || info.Category != "conflict" || info.Status != 409 {
		t.Fatalf("slug_taken mapping drift: %#v", info)
	}
}

func TestRegressionErrorMetadataAndRedaction(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(401)
		_ = json.NewEncoder(w).Encode(map[string]any{"error": map[string]any{"category": "unauthenticated", "code": "token_expired", "message": "expired", "request_id": "req_go"}})
	}))
	defer srv.Close()
	client := NewClient(Config{BaseURL: srv.URL, Token: "pat_secret", TokenType: TokenTypePAT, DefaultTenantID: "tnt_1"})
	if client.RedactedToken() != "***REDACTED***" {
		t.Fatalf("token prefix leaked: %s", client.RedactedToken())
	}
	_, err := client.Request(context.Background(), "appsGetApiV1AppsByAppID", map[string]string{"appID": "app_1"}, nil, nil)
	if err == nil {
		t.Fatal("expected error")
	}
	axErr := err.(*AxHubError)
	if axErr.RequestID != "req_go" || !axErr.Retryable {
		t.Fatalf("error metadata drift: %#v", axErr)
	}
}

func TestRegressionEightContextCoverage(t *testing.T) {
	want := []string{"apps", "identity", "tenants", "authz", "audit", "gateway", "cost", "data", "deployments"}
	for _, name := range want {
		if len(ContextRoutes[name]) == 0 {
			t.Fatalf("missing context routes for %s", name)
		}
	}
}

func TestRegressionNonJSONSuccessAndScalarErrorBodies(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/auth/google_oauth2/start":
			w.Header().Set("Content-Type", "text/html")
			_, _ = w.Write([]byte("<html>oauth redirect target</html>"))
		case "/oauth/token":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			_, _ = w.Write([]byte(`"invalid_request"`))
		default:
			w.WriteHeader(404)
		}
	}))
	defer srv.Close()
	client := NewClient(Config{BaseURL: srv.URL})
	got, err := client.Request(context.Background(), "authGetAuthGoogleOauth2Start", nil, nil, nil)
	if err != nil {
		t.Fatalf("non-json success should not fail: %v", err)
	}
	if got["raw"] != "<html>oauth redirect target</html>" {
		t.Fatalf("unexpected raw response: %#v", got)
	}
	_, err = client.Request(context.Background(), "authPostOauthToken", nil, nil, map[string]any{"noop": true})
	axErr, ok := err.(*AxHubError)
	if !ok || axErr.Status != 400 || axErr.Code != "http_400" {
		t.Fatalf("scalar JSON error should become typed HTTP error: %#v", err)
	}
}

func TestRegressionOAuthFormEncodingAndRedirectPolicy(t *testing.T) {
	var tokenContentType, tokenBody string
	redirectTargetHit := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/oauth/token":
			tokenContentType = r.Header.Get("Content-Type")
			raw, _ := io.ReadAll(r.Body)
			tokenBody = string(raw)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "tok_go", "token_type": "Bearer", "expires_in": 3600})
		case "/auth/google_oauth2/start":
			http.Redirect(w, r, "/redirect-target", http.StatusFound)
		case "/redirect-target":
			redirectTargetHit = true
			w.WriteHeader(500)
		default:
			w.WriteHeader(404)
		}
	}))
	defer srv.Close()

	client := NewClient(Config{BaseURL: srv.URL, Token: "pat_secret", TokenType: TokenTypePAT})
	got, err := client.Request(context.Background(), "authPostOauthToken", nil, nil, map[string]any{"grant_type": "client_credentials", "client_id": "cid"})
	if err != nil {
		t.Fatalf("oauth token form request failed: %v", err)
	}
	if got["accessToken"] != "tok_go" {
		t.Fatalf("oauth token response drift: %#v", got)
	}
	if !strings.HasPrefix(tokenContentType, "application/x-www-form-urlencoded") || !strings.Contains(tokenBody, "grant_type=client_credentials") || strings.Contains(tokenBody, "{") {
		t.Fatalf("oauth token was not form-encoded content-type=%q body=%q", tokenContentType, tokenBody)
	}
	redirect, err := client.Request(context.Background(), "authGetAuthGoogleOauth2Start", nil, nil, nil)
	if err != nil {
		t.Fatalf("redirect response should be returned: %v", err)
	}
	if redirect["status"] != 302 || redirect["location"] != "/redirect-target" {
		t.Fatalf("redirect response drift: %#v", redirect)
	}
	if redirectTargetHit {
		t.Fatalf("redirect was followed; auth headers could leak to redirect target")
	}
}

func TestRegressionErrorFullFieldDecode(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		_, _ = w.Write([]byte(`{"error":{"category":"validation","code":"invalid_value","message":"잘못된 값","request_id":"req_full","resource":"app","retryable":false,"doc_url":"https://docs.axhub.ai/errors/invalid_value","fields":[{"name":"slug","code":"invalid_format","message":"형식 오류"},{"name":"name","code":"required"}],"retry":{"after_ms":1500}}}`))
	}))
	defer srv.Close()
	client := NewClient(Config{BaseURL: srv.URL})
	_, err := client.doHTTP(context.Background(), "GET", "/api/v1/apps/app_1", nil, nil, false)
	axErr, ok := err.(*AxHubError)
	if !ok {
		t.Fatalf("expected *AxHubError, got %T: %v", err, err)
	}
	if axErr.Category != "validation" || axErr.Code != "invalid_value" || axErr.RequestID != "req_full" || axErr.Status != 422 {
		t.Fatalf("base field drift: %#v", axErr)
	}
	if axErr.Resource != "app" {
		t.Fatalf("resource drift: %q", axErr.Resource)
	}
	if axErr.DocURL != "https://docs.axhub.ai/errors/invalid_value" {
		t.Fatalf("doc_url drift: %q", axErr.DocURL)
	}
	if axErr.Retry == nil || axErr.Retry.AfterMs != 1500 {
		t.Fatalf("retry drift: %#v", axErr.Retry)
	}
	if len(axErr.Fields) != 2 {
		t.Fatalf("fields length drift: %#v", axErr.Fields)
	}
	if axErr.Fields[0].Name != "slug" || axErr.Fields[0].Code != "invalid_format" || axErr.Fields[0].Message != "형식 오류" {
		t.Fatalf("fields[0] drift: %#v", axErr.Fields[0])
	}
	if axErr.Fields[1].Name != "name" || axErr.Fields[1].Code != "required" || axErr.Fields[1].Message != "" {
		t.Fatalf("fields[1] drift: %#v", axErr.Fields[1])
	}
}

func TestRegression428NonEnvelopeErrorPreservesCode(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(428)
		// Non-envelope: bare {code,message} at the root, no "error" wrapper (W2).
		_, _ = w.Write([]byte(`{"code":"confirm_required","message":"confirmation required"}`))
	}))
	defer srv.Close()
	client := NewClient(Config{BaseURL: srv.URL})
	_, err := client.doHTTP(context.Background(), "GET", "/api/v1/apps/app_1", nil, nil, false)
	axErr, ok := err.(*AxHubError)
	if !ok {
		t.Fatalf("expected *AxHubError, got %T: %v", err, err)
	}
	if axErr.Code != "confirm_required" {
		t.Fatalf("W2 non-envelope code not preserved: got %q, want confirm_required (masked as http_428?)", axErr.Code)
	}
	if axErr.Status != 428 {
		t.Fatalf("status drift: got %d, want 428", axErr.Status)
	}
	if axErr.Message != "confirmation required" {
		t.Fatalf("message drift: %q", axErr.Message)
	}
}

func TestRegression429RetryAfterBackoffSucceeds(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&calls, 1) <= 2 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(429)
			_, _ = w.Write([]byte(`{"error":{"category":"unavailable","code":"temporarily_unavailable","message":"too many","retryable":true,"retry":{"after_ms":0}}}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()
	client := NewClient(Config{BaseURL: srv.URL})
	got, err := client.doHTTP(context.Background(), "GET", "/api/v1/ping", nil, nil, false)
	if err != nil {
		t.Fatalf("expected retry to succeed, got error: %v", err)
	}
	m, ok := got.(map[string]any)
	if !ok || m["ok"] != true {
		t.Fatalf("unexpected body after retry: %#v", got)
	}
	if n := atomic.LoadInt32(&calls); n != 3 {
		t.Fatalf("expected 3 attempts (1 initial + 2 retries), got %d", n)
	}
}

func TestRegression429RetriesExhaustReturnError(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.Header().Set("Retry-After", "0")
		w.WriteHeader(429)
		_, _ = w.Write([]byte(`{"error":{"category":"unavailable","code":"temporarily_unavailable","message":"too many","retryable":true}}`))
	}))
	defer srv.Close()
	client := NewClient(Config{BaseURL: srv.URL})
	_, err := client.doHTTP(context.Background(), "GET", "/api/v1/ping", nil, nil, false)
	axErr, ok := err.(*AxHubError)
	if !ok || axErr.Status != 429 {
		t.Fatalf("expected 429 AxHubError after exhausting retries, got %T: %v", err, err)
	}
	if n := atomic.LoadInt32(&calls); n != maxRateLimitRetries+1 {
		t.Fatalf("expected %d attempts before giving up, got %d", maxRateLimitRetries+1, n)
	}
}

func TestRetryAfterDurationParsing(t *testing.T) {
	cases := []struct {
		header string
		want   time.Duration
	}{
		{"", defaultRetryAfter},
		{"0", 0},
		{"2", 2 * time.Second},
		{"  5 ", 5 * time.Second},
		{"-1", defaultRetryAfter},
		{"garbage", defaultRetryAfter},
	}
	for _, tc := range cases {
		if got := retryAfterDuration(tc.header); got != tc.want {
			t.Fatalf("retryAfterDuration(%q) = %v, want %v", tc.header, got, tc.want)
		}
	}
}
