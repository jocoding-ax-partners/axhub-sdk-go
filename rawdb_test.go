package axhub

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRawDbTables_ParsesTypedColumns(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/apps/app_x/db/tables" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		// snake_case wire shape; the transport camelCases keys before parsing.
		_ = json.NewEncoder(w).Encode(map[string]any{
			"tables": []map[string]any{{
				"name":    "posts",
				"managed": false,
				"columns": []map[string]any{
					{"name": "id", "data_type": "uuid", "nullable": false},
					{"name": "title", "data_type": "text", "nullable": true},
				},
			}},
		})
	}))
	defer srv.Close()
	client := NewClient(Config{BaseURL: srv.URL, Token: "t", TokenType: TokenTypePAT})

	tables, err := client.Apps.RawDb().Tables(context.Background(), "app_x")
	if err != nil {
		t.Fatalf("Tables: %v", err)
	}
	if len(tables) != 1 || tables[0].Name != "posts" || len(tables[0].Columns) != 2 {
		t.Fatalf("unexpected tables: %+v", tables)
	}
	if tables[0].Columns[0].DataType != "uuid" || !tables[0].Columns[1].Nullable {
		t.Fatalf("columns not parsed typed: %+v", tables[0].Columns)
	}
}

// F-3: a successful empty read must be distinguishable from an auth failure.
// Empty slice + nil error = genuinely empty (raw DB not enabled or 0 tables).
func TestRawDbTables_EmptyMeansGenuinelyEmpty(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{"tables": []any{}})
	}))
	defer srv.Close()
	client := NewClient(Config{BaseURL: srv.URL, Token: "t", TokenType: TokenTypePAT})

	tables, err := client.Apps.RawDb().Tables(context.Background(), "app_x")
	if err != nil {
		t.Fatalf("empty tables must not error (empty != auth failure): %v", err)
	}
	if len(tables) != 0 {
		t.Fatalf("expected empty, got %+v", tables)
	}
}

func TestRawDbTableRows_ParsesPageAndForwardsPerPage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("per_page"); got != "100" {
			t.Errorf("per_page not forwarded: %q", got)
		}
		_ = json.NewEncoder(w).Encode(map[string]any{
			"rows":     []map[string]any{{"id": "1"}, {"id": "2"}},
			"page":     1,
			"per_page": 100,
			"has_more": true,
		})
	}))
	defer srv.Close()
	client := NewClient(Config{BaseURL: srv.URL, Token: "t", TokenType: TokenTypePAT})

	page, err := client.Apps.RawDb().TableRows(context.Background(), "app_x", "posts", &RawDbTableRowsOptions{PerPage: 100})
	if err != nil {
		t.Fatalf("TableRows: %v", err)
	}
	if len(page.Rows) != 2 || page.PerPage != 100 || !page.HasMore {
		t.Fatalf("page not parsed: %+v", page)
	}
}

func TestRawDbTableRows_RequiresAppAndTable(t *testing.T) {
	client := NewClient(Config{BaseURL: "https://example.invalid", Token: "t", TokenType: TokenTypePAT})
	if _, err := client.Apps.RawDb().TableRows(context.Background(), "app_x", "", nil); err == nil {
		t.Fatal("empty table should error before any request")
	}
	if _, err := client.Apps.RawDb().Tables(context.Background(), ""); err == nil {
		t.Fatal("empty appID should error before any request")
	}
}
