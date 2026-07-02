package axhub

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"testing"
	"time"
)

const liveDeadUUID = "00000000-0000-4000-8000-00000000dead"

var liveHighRiskTenantOps = map[string]bool{
	"tenantsDeleteApiV1TenantsByTenantID":                             true,
	"tenantsPatchApiV1TenantsByTenantID":                              true,
	"tenantsDeleteApiV1TenantsByTenantIDIcon":                         true,
	"gatewayGetApiV1TenantsByTenantIDConnectorsByConnectorIDDiscover": true,
	"gatewayPostApiV1TenantsByTenantIDConnectors":                     true,
}

var liveHighRiskAppOps = map[string]bool{
	"appsDeleteApiV1AppsByAppID":                         true,
	"appsDeleteApiV1AppsByAppIDPermanent":                true,
	"deployPostApiV1AppsByAppIDDeploymentsByDidCancel":   true,
	"deployPostApiV1AppsByAppIDDeploymentsByDidRollback": true,
}

type liveOperationResult struct {
	OperationID string `json:"operationId"`
	Method      string `json:"method"`
	Kind        string `json:"kind"`
	Status      int    `json:"status,omitempty"`
	Code        string `json:"code,omitempty"`
	Category    string `json:"category,omitempty"`
	ServerError bool   `json:"server_error,omitempty"`
	Error       string `json:"error,omitempty"`
}

func TestLiveAllGeneratedOperationFacadesHitProd(t *testing.T) {
	if os.Getenv("AXHUB_LIVE_ALL_METHODS") != "1" {
		t.Skip("live prod all-method sweep is opt-in")
	}
	token := os.Getenv("AXHUB_TOKEN")
	if token == "" {
		t.Fatal("AXHUB_TOKEN is required")
	}
	tenantID := getenv("AXHUB_LIVE_TENANT_ID", "cc1e58f1-8e46-4ac7-96c1-190c4cdd5b70")
	tenantSlug := getenv("AXHUB_LIVE_TENANT_SLUG", "test")
	baseURL := getenv("AXHUB_LIVE_BASE_URL", "https://api.axhub.ai")
	if len(Routes) != 85 {
		t.Fatalf("route coverage drift: %d", len(Routes))
	}
	ctx := context.Background()
	c := NewClient(Config{BaseURL: baseURL, Token: token, TokenType: TokenTypePAT, DefaultTenantID: tenantID, DefaultTenantSlug: tenantSlug})
	fixture := map[string]string{}
	createdFixture := false
	if created, err := c.Apps.Create(ctx, map[string]any{"slug": "sdk-e2e-destructive-go-" + time.Now().Format("20060102150405"), "name": "SDK destructive E2E disposable"}); err == nil {
		if id, ok := created["id"].(string); ok && id != "" {
			fixture["appID"] = id
			createdFixture = true
		}
		if slug, ok := created["slug"].(string); ok {
			fixture["appSlug"] = slug
		}
	}
	defer func() {
		if createdFixture {
			_, _ = c.Request(ctx, "appsDeleteApiV1AppsByAppID", map[string]string{"appID": fixture["appID"]}, nil, nil)
			_, _ = c.Request(ctx, "appsDeleteApiV1AppsByAppIDPermanent", map[string]string{"appID": fixture["appID"]}, nil, nil)
		}
	}()

	contexts := map[string]any{
		"apps": c.Apps, "identity": c.Identity(), "tenants": c.Tenants(), "authz": c.Authz(),
		"audit": c.Audit(), "gateway": c.Gateway(), "data": c.Data(), "deployments": c.Deployments(),
	}
	results := []liveOperationResult{}
	for _, route := range Routes {
		target := contexts[contextName(route)]
		method := reflect.ValueOf(target).MethodByName(testGoMethodName(route.OperationID))
		if !method.IsValid() {
			t.Fatalf("missing generated method %s", route.OperationID)
		}
		result := liveOperationResult{OperationID: route.OperationID, Method: route.Method, Kind: "unknown"}
		out := method.Call([]reflect.Value{
			reflect.ValueOf(ctx),
			reflect.ValueOf(OperationParams{
				PathParams: livePathParamsFor(route, fixture, tenantID, tenantSlug),
				Query:      map[string]string{"sdk_e2e": "live_all_methods"},
				Body:       liveBodyFor(route),
			}),
		})
		if !out[1].IsNil() {
			err := out[1].Interface().(error)
			var ax *AxHubError
			if errors.As(err, &ax) {
				result.Kind = "axhub_error"
				result.Status = ax.Status
				result.Code = ax.Code
				result.Category = ax.Category
				result.ServerError = ax.Status >= 500
			} else {
				result.Kind = "exception"
				result.Error = err.Error()
			}
		} else {
			result.Kind = "success"
		}
		results = append(results, result)
	}
	summary := map[string]any{
		"sdk":          "go",
		"baseUrl":      baseURL,
		"tenantId":     tenantID,
		"fixture":      map[string]any{"created": createdFixture, "appID": fixture["appID"], "appSlug": fixture["appSlug"]},
		"total":        len(results),
		"destructive":  countResults(results, func(r liveOperationResult) bool { return r.Method != "GET" }),
		"success":      countResults(results, func(r liveOperationResult) bool { return r.Kind == "success" }),
		"axhub_error":  countResults(results, func(r liveOperationResult) bool { return r.Kind == "axhub_error" }),
		"exception":    countResults(results, func(r liveOperationResult) bool { return r.Kind == "exception" }),
		"serverErrors": filterResults(results, func(r liveOperationResult) bool { return r.ServerError }),
		"exceptions":   filterResults(results, func(r liveOperationResult) bool { return r.Kind == "exception" }),
		"results":      results,
	}
	if path := os.Getenv("AXHUB_LIVE_RESULT_PATH"); path != "" {
		raw, _ := json.MarshalIndent(summary, "", "  ")
		if err := os.WriteFile(path, raw, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	if len(results) != 85 {
		t.Fatalf("total drift %d", len(results))
	}
	expectedDestructive := 0
	for _, route := range Routes {
		if route.Method != "GET" {
			expectedDestructive++
		}
	}
	if got := countResults(results, func(r liveOperationResult) bool { return r.Method != "GET" }); got != expectedDestructive {
		t.Fatalf("destructive method count drift %d != %d", got, expectedDestructive)
	}
	if bad := filterResults(results, func(r liveOperationResult) bool { return r.Kind == "exception" || r.ServerError }); len(bad) > 0 {
		t.Fatalf("live failures: %+v", bad)
	}
}

func livePathParamsFor(route Route, fixture map[string]string, tenantID, tenantSlug string) map[string]string {
	params := map[string]string{}
	for _, match := range testPathParamPattern.FindAllStringSubmatch(route.Path, -1) {
		params[match[1]] = livePathParamValue(match[1], route.OperationID, fixture, tenantID, tenantSlug)
	}
	return params
}

func livePathParamValue(name, operationID string, fixture map[string]string, tenantID, tenantSlug string) string {
	switch name {
	case "tenantID":
		if liveHighRiskTenantOps[operationID] {
			return liveDeadUUID
		}
		return tenantID
	case "tenantSlug":
		return tenantSlug
	case "appID":
		if liveHighRiskAppOps[operationID] {
			return liveDeadUUID
		}
		if fixture["appID"] != "" {
			return fixture["appID"]
		}
		return liveDeadUUID
	case "appSlug":
		if fixture["appSlug"] != "" {
			return fixture["appSlug"]
		}
		return "sdk-e2e-missing-app"
	case "table", "tableName":
		return "sdk_e2e_missing_table"
	case "path":
		return "sdk/e2e/missing"
	case "domain":
		return "sdk-e2e.invalid"
	case "providerID":
		if operationID == "authGetAuthByProviderIDStart" {
			return "github"
		}
		return "sdk-e2e-provider"
	case "patID":
		return liveDeadUUID
	case "key":
		return "SDK_E2E_NOOP"
	case "connector":
		return "sdk-e2e-connector"
	default:
		return liveDeadUUID
	}
}

func liveBodyFor(route Route) any {
	if route.Method == "GET" || route.Method == "DELETE" {
		return nil
	}
	return map[string]any{"sdk_e2e": true, "operation_id": route.OperationID}
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func countResults(results []liveOperationResult, pred func(liveOperationResult) bool) int {
	count := 0
	for _, result := range results {
		if pred(result) {
			count++
		}
	}
	return count
}

func filterResults(results []liveOperationResult, pred func(liveOperationResult) bool) []liveOperationResult {
	out := []liveOperationResult{}
	for _, result := range results {
		if pred(result) {
			out = append(out, result)
		}
	}
	return out
}
