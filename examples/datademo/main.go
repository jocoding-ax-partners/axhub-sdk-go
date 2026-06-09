// Data CRUD demo mirroring node examples/data-crud-demo.ts.
//
// Shows the ergonomic data layer: DefineSchema + the fluent
// Tenant().App().Data().Table() chain + a count() call over the offset-only
// dynamic-table API. Runs against AXHUB_BASE_URL when set, otherwise spins up a
// local mock server so the example is self-contained for CI smoke.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	axhub "github.com/jocoding-ax-partners/axhub-sdk-go"
)

func main() {
	baseURL := os.Getenv("AXHUB_BASE_URL")
	var srv *httptest.Server
	if baseURL == "" {
		srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// /data/acme/crm/orders/_count
			_ = json.NewEncoder(w).Encode(map[string]any{"count": 3})
		}))
		defer srv.Close()
		baseURL = srv.URL
	}

	token := os.Getenv("AXHUB_TOKEN")
	if token == "" {
		token = "demo"
	}

	sdk := axhub.NewClient(axhub.Config{BaseURL: baseURL, Token: token, TokenType: axhub.TokenTypePAT})

	// Define a schema (enables select validation + typed Where via schema.Cols).
	orders := axhub.DefineSchema("orders", axhub.SchemaShape{
		"id":    "uuid",
		"total": "number",
	})

	table := sdk.Tenant("acme").App("crm").Data().TableSchema(orders)

	count, err := table.Count(context.Background(), &axhub.CountOptions{
		Where: ptr(axhub.Where("status").Eq("paid")),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "go data demo failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("go data demo ok: orders(status=paid) count=%d base=%s\n", count, sdk.BaseURL())
}

func ptr(e axhub.QueryExpr) *axhub.QueryExpr { return &e }
