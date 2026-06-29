package axhub

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestAllGeneratedOperationFacadesMakeHTTPRequests(t *testing.T) {
	if len(Routes) != 228 {
		t.Fatalf("route coverage drift: %d", len(Routes))
	}

	expectedIndex := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if expectedIndex >= len(Routes) {
			t.Fatalf("unexpected extra request %s %s", r.Method, r.URL.String())
		}
		route := Routes[expectedIndex]
		params := testPathParamsFor(route.Path)
		wantPath := testRenderPath(route.Path, params)
		if r.Method != route.Method {
			t.Errorf("%s used method %s, want %s", route.OperationID, r.Method, route.Method)
		}
		if r.URL.Path != wantPath {
			t.Errorf("%s used path %s, want %s", route.OperationID, r.URL.Path, wantPath)
		}
		if got := r.URL.Query().Get("e2e"); got != "ok" {
			t.Errorf("%s missing e2e query, got %q", route.OperationID, got)
		}
		if got := r.Header.Get("X-Api-Key"); got != "pat_e2e" {
			t.Errorf("%s missing PAT header, got %q", route.OperationID, got)
		}
		if got := r.Header.Get("X-Request-ID"); got == "" {
			t.Errorf("%s missing request id", route.OperationID)
		}
		expectedIndex++
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]any{"operation_id": route.OperationID, "ok": true}); err != nil {
			t.Errorf("encode response for %s: %v", route.OperationID, err)
		}
	}))
	defer server.Close()

	c := NewClient(Config{BaseURL: server.URL, Token: "pat_e2e", TokenType: TokenTypePAT})
	contexts := map[string]any{
		"apps": c.Apps, "identity": c.Identity(), "tenants": c.Tenants(), "authz": c.Authz(),
		"audit": c.Audit(), "gateway": c.Gateway(), "cost": c.Cost(), "data": c.Data(), "deployments": c.Deployments(),
	}
	for _, route := range Routes {
		target := contexts[contextName(route)]
		if target == nil {
			t.Fatalf("missing context for %s", route.OperationID)
		}
		method := reflect.ValueOf(target).MethodByName(testGoMethodName(route.OperationID))
		if !method.IsValid() {
			t.Fatalf("missing generated method %s on %s", testGoMethodName(route.OperationID), contextName(route))
		}
		out := method.Call([]reflect.Value{
			reflect.ValueOf(context.Background()),
			reflect.ValueOf(OperationParams{PathParams: testPathParamsFor(route.Path), Query: map[string]string{"e2e": "ok"}, Body: testBodyFor(route)}),
		})
		if !out[1].IsNil() {
			t.Fatalf("%s returned error: %v", route.OperationID, out[1].Interface())
		}
		got := out[0].Interface().(map[string]any)
		if got["operationId"] != route.OperationID {
			t.Fatalf("%s response was not parsed/camelized: %v", route.OperationID, got)
		}
	}
	if expectedIndex != len(Routes) {
		t.Fatalf("expected %d HTTP requests, saw %d", len(Routes), expectedIndex)
	}
}

var testPathParamPattern = regexp.MustCompile(`\{([^}]+)\}`)

func testPathParamsFor(path string) map[string]string {
	params := map[string]string{}
	for _, match := range testPathParamPattern.FindAllStringSubmatch(path, -1) {
		params[match[1]] = testPathParamValue(match[1])
	}
	return params
}

func testPathParamValue(name string) string {
	switch name {
	case "tenantID":
		return "tnt_1"
	case "tenantSlug":
		return "test-tenant"
	case "appID":
		return "app_1"
	case "appSlug":
		return "app-slug"
	case "table", "tableName":
		return "table_1"
	case "path":
		return "resource-path"
	case "domain":
		return "example.com"
	default:
		return strings.ToLower(name) + "_1"
	}
}

func testRenderPath(path string, params map[string]string) string {
	return testPathParamPattern.ReplaceAllStringFunc(path, func(token string) string {
		name := token[1 : len(token)-1]
		return url.PathEscape(params[name])
	})
}

func testBodyFor(route Route) any {
	if route.Method == http.MethodGet || route.Method == http.MethodDelete {
		return nil
	}
	return map[string]any{"operationId": route.OperationID, "ok": true}
}

func testGoMethodName(operationID string) string {
	if operationID == "" {
		return operationID
	}
	return strings.ToUpper(operationID[:1]) + operationID[1:]
}
