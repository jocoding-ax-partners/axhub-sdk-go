package axhub

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
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
	if len(Routes) != 177 {
		t.Fatalf("route coverage drift: got %d", len(Routes))
	}
	if len(ErrorCodes) != 42 {
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
	want := []string{"apps", "identity", "tenants", "authz", "audit", "gateway", "data", "deployments"}
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
