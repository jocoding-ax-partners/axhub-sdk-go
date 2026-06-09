package axhub

// Unit + wire tests for the ergonomic data layer. These mirror the node/python
// data-layer behavior. The conformance runner only exercises the operation-id
// route-table surface (sdk.operation), NOT this fluent layer — so these tests
// are the real proof of the port. Covers: where serialization incl. IN-comma
// guard + pushable rejection, order-by id tiebreaker, per_page clamp, select
// validate/project, cursor rejection (legacy/v1/v2/invalid/oversized/bad-page),
// CRUD wire paths + envelope, verbatim (no-camelize) rows, list_all drift,
// discover slug+appId-fallback+TableNotFound, discover cache across chains,
// schema-cache LRU/TTL, and LIKE escaping + ReDoS guard.

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"
)

// ----------------------------- mock data server ------------------------------

type capturedRequest struct {
	method  string
	path    string
	query   url.Values
	body    map[string]any
	headers http.Header
}

type dataServer struct {
	srv     *httptest.Server
	mu      sync.Mutex
	last    *capturedRequest
	status  int
	body    any
}

func newDataServer() *dataServer {
	ds := &dataServer{status: 200, body: map[string]any{}}
	ds.srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var parsed map[string]any
		if r.Body != nil {
			dec := json.NewDecoder(r.Body)
			_ = dec.Decode(&parsed)
		}
		ds.mu.Lock()
		ds.last = &capturedRequest{
			method:  r.Method,
			path:    r.URL.Path,
			query:   r.URL.Query(),
			body:    parsed,
			headers: r.Header.Clone(),
		}
		status, body := ds.status, ds.body
		ds.mu.Unlock()
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(body)
	}))
	return ds
}

func (ds *dataServer) setResponse(body any, status int) {
	ds.mu.Lock()
	ds.body, ds.status = body, status
	ds.mu.Unlock()
}

func (ds *dataServer) reset() {
	ds.mu.Lock()
	ds.last = nil
	ds.mu.Unlock()
}

func (ds *dataServer) lastReq() *capturedRequest {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	return ds.last
}

func (ds *dataServer) close() { ds.srv.Close() }

func (ds *dataServer) client() *Client {
	return NewClient(Config{BaseURL: ds.srv.URL, Token: "pat_x", TokenType: TokenTypePAT})
}

func (ds *dataServer) table() *DataResource {
	return ds.client().Tenant("acme").App("crm").Data().Table("orders")
}

// ------------------------------- where serializer ----------------------------

func mustSerialize(t *testing.T, expr QueryExpr) url.Values {
	t.Helper()
	v, err := serializeWhere(&expr)
	if err != nil {
		t.Fatalf("serializeWhere err: %v", err)
	}
	return v
}

func TestWhereAtomAndAndIn(t *testing.T) {
	if got := mustSerialize(t, Where("status").Eq("paid")).Get("status"); got != "eq.paid" {
		t.Fatalf("eq got %q", got)
	}
	out := mustSerialize(t, And(Where("total").Gte(10), Where("status").Ne("void")))
	if out.Get("total") != "gte.10" || out.Get("status") != "ne.void" {
		t.Fatalf("and got %v", out)
	}
	if got := mustSerialize(t, Where("id").In("a", "b")).Get("id"); got != "in.a,b" {
		t.Fatalf("in got %q", got)
	}
}

func TestWhereBoolAndNilStringifyLikeJS(t *testing.T) {
	if got := mustSerialize(t, Where("active").Eq(true)).Get("active"); got != "eq.true" {
		t.Fatalf("bool got %q", got)
	}
	if got := mustSerialize(t, Where("deleted").Eq(nil)).Get("deleted"); got != "eq.null" {
		t.Fatalf("nil got %q", got)
	}
}

func TestWhereRepeatedColumnCollapsesToList(t *testing.T) {
	out := mustSerialize(t, And(Where("tag").Eq("a"), Where("tag").Eq("b")))
	vals := out["tag"]
	if len(vals) != 2 || vals[0] != "eq.a" || vals[1] != "eq.b" {
		t.Fatalf("repeated column got %v", vals)
	}
}

func TestWhereInCommaGuard(t *testing.T) {
	_, err := serializeWhere(&QueryExpr{Op: "in", Column: "name", Values: []any{"a,b"}})
	assertCode(t, err, "filter_in_comma")
}

func TestWhereUnsupportedFiltersRejected(t *testing.T) {
	exprs := []QueryExpr{
		Or(Where("a").Eq(1)),
		Not(Where("a").Eq(1)),
		Raw("1=1"),
	}
	for _, e := range exprs {
		_, err := serializeWhere(&e)
		assertCode(t, err, "unsupported_filter")
	}
}

func TestWhereNestedAndIsNotPushable(t *testing.T) {
	_, err := serializeWhere(ptr(And(And(Where("a").Eq(1)))))
	assertCode(t, err, "unsupported_filter")
}

func TestWhereDeferredBuildErrorSurfaces(t *testing.T) {
	// like.Raw ReDoS guard defers an error onto the expr; it must surface from
	// serializeWhere (and therefore List/Count), not the builder.
	_, err := serializeWhere(ptr(Where("name").Like.Raw("%%%%x")))
	assertCode(t, err, "like_pattern_redos")
}

// --------------------------------- order by ----------------------------------

func TestOrderByStringFormAppendsIDTiebreaker(t *testing.T) {
	if got, _ := serializeOrderBy("-total"); got != "-total,id" {
		t.Fatalf("got %q", got)
	}
	if got, _ := serializeOrderBy("name"); got != "name,id" {
		t.Fatalf("got %q", got)
	}
}

func TestOrderByFieldListForm(t *testing.T) {
	if got, _ := serializeOrderBy([]OrderField{{Field: "total", Dir: "desc"}}); got != "-total,id" {
		t.Fatalf("got %q", got)
	}
}

func TestOrderByEmptyIsNone(t *testing.T) {
	if _, ok := serializeOrderBy(nil); ok {
		t.Fatalf("nil order-by should yield no sort")
	}
}

// -------------------------------- clamp per page -----------------------------

func TestClampPerPage(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{{0, 1}, {50, 50}, {1000, 100}, {-5, 1}}
	for _, c := range cases {
		if got, ok := clampPerPage(c.in, true); !ok || got != c.want {
			t.Fatalf("clamp(%d) got %d", c.in, got)
		}
	}
	if _, ok := clampPerPage(0, false); ok {
		t.Fatalf("absent per-page should be omitted")
	}
}

// ---------------------------------- select -----------------------------------

func TestSelectSerialize(t *testing.T) {
	if got, ok := serializeSelect([]string{"id", "total"}); !ok || got != "id,total" {
		t.Fatalf("got %q", got)
	}
	if _, ok := serializeSelect(nil); ok {
		t.Fatalf("nil select should be omitted")
	}
}

func TestSelectEmptyRejected(t *testing.T) {
	assertCode(t, validateSelectColumns(nil, []string{}), "select_empty")
}

func TestSelectUnknownColumnRejectedWithSchema(t *testing.T) {
	schema := DefineSchema("orders", SchemaShape{"id": "uuid", "total": "number"})
	assertCode(t, validateSelectColumns(schema, []string{"id", "nope"}), "select_unknown_column")
}

func TestProjectRowNarrows(t *testing.T) {
	out := projectRow(map[string]any{"id": "x", "total": 5, "extra": 1}, []string{"id"})
	if len(out) != 1 || out["id"] != "x" {
		t.Fatalf("project got %v", out)
	}
}

// ----------------------------- cursor rejection ------------------------------

func TestCursorAfterBeforeDirectionRejected(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	ctx := context.Background()
	for _, opt := range []*ListOptions{{After: "x"}, {Before: "x"}, {Direction: "forward"}} {
		_, err := ds.table().List(ctx, opt)
		assertCode(t, err, "legacy_cursor")
	}
}

func TestCursorV1AndV2Rejected(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	ctx := context.Background()
	for _, c := range []string{"v1:abc", "v2:abc"} {
		_, err := ds.table().List(ctx, &ListOptions{Cursor: c})
		assertCode(t, err, "legacy_cursor")
	}
}

func TestCursorNonIntegerRejected(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	ctx := context.Background()
	for _, c := range []string{"abc", "0"} {
		_, err := ds.table().List(ctx, &ListOptions{Cursor: c})
		assertCode(t, err, "invalid_cursor")
	}
}

func TestCursorOversizedRejected(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	_, err := ds.table().List(context.Background(), &ListOptions{Cursor: strings.Repeat("1", 5000)})
	assertCode(t, err, "invalid_cursor")
}

func TestBadPageRejected(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	_, err := ds.table().List(context.Background(), &ListOptions{Page: -1})
	assertCode(t, err, "invalid_cursor")
}

func TestIsV2CursorHelper(t *testing.T) {
	if !isV2Cursor("v2:x") || isV2Cursor("3") {
		t.Fatalf("isV2Cursor wrong")
	}
}

// ----------------------------------- list wire -------------------------------

func TestListQueryAndEnvelope(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	ds.setResponse(map[string]any{
		"items":    []any{map[string]any{"id": "1", "created_at": "t"}},
		"page":     2,
		"per_page": 10,
		"has_more": true,
	}, 200)
	result, err := ds.table().List(context.Background(), &ListOptions{
		Where:    ptr(Where("status").Eq("paid")),
		OrderBy:  "-total",
		Select:   []string{"id", "created_at"},
		Page:     2,
		PageSize: 10,
	})
	if err != nil {
		t.Fatalf("list err: %v", err)
	}
	last := ds.lastReq()
	if last.method != "GET" || last.path != "/data/acme/crm/orders" {
		t.Fatalf("wrong wire %s %s", last.method, last.path)
	}
	if last.query.Get("status") != "eq.paid" || last.query.Get("per_page") != "10" ||
		last.query.Get("page") != "2" || last.query.Get("sort") != "-total,id" ||
		last.query.Get("_select") != "id,created_at" {
		t.Fatalf("wrong query %v", last.query)
	}
	if last.headers.Get("X-Api-Key") != "pat_x" {
		t.Fatalf("missing PAT header")
	}
	if result.NextCursor == nil || *result.NextCursor != "3" {
		t.Fatalf("next cursor %v", result.NextCursor)
	}
	if result.FirstCursor == nil || *result.FirstCursor != "1" {
		t.Fatalf("first cursor %v", result.FirstCursor)
	}
	if !result.HasNext || !result.HasPrev || result.TotalIsExact {
		t.Fatalf("envelope flags wrong %+v", result)
	}
}

func TestListRowDataReturnedVerbatimNoCamelize(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	ds.setResponse(map[string]any{
		"items":    []any{map[string]any{"id": "1", "created_at": "2020", "is_active": true}},
		"has_more": false,
	}, 200)
	result, err := ds.table().List(context.Background(), &ListOptions{Where: wptr(Where("id").Eq("1"))})
	if err != nil {
		t.Fatalf("list err: %v", err)
	}
	row := result.Items[0]
	if row["created_at"] != "2020" || row["is_active"] != true {
		t.Fatalf("snake_case keys must NOT be rewritten, got %v", row)
	}
	if _, camelized := row["createdAt"]; camelized {
		t.Fatalf("row was camelized: %v", row)
	}
	if result.NextCursor != nil {
		t.Fatalf("expected nil next cursor")
	}
}

func TestListPage1OmitsPageQuery(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	ds.setResponse(map[string]any{"items": []any{}, "has_more": false}, 200)
	_, err := ds.table().List(context.Background(), &ListOptions{Page: 1, Where: wptr(Where("id").Eq("1"))})
	if err != nil {
		t.Fatalf("list err: %v", err)
	}
	if _, ok := ds.lastReq().query["page"]; ok {
		t.Fatalf("page=1 should omit page query")
	}
}

func TestListSelectProjectsClientSide(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	ds.setResponse(map[string]any{
		"items":    []any{map[string]any{"id": "1", "total": 9, "secret": "x"}},
		"has_more": false,
	}, 200)
	result, err := ds.table().List(context.Background(), &ListOptions{Select: []string{"id", "total"}, Where: wptr(Where("id").Eq("1"))})
	if err != nil {
		t.Fatalf("list err: %v", err)
	}
	row := result.Items[0]
	if _, ok := row["secret"]; ok || len(row) != 2 {
		t.Fatalf("client-side projection failed: %v", row)
	}
}

// ----------------------------------- crud wire -------------------------------

func TestCountWire(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	ds.setResponse(map[string]any{"count": 42}, 200)
	n, err := ds.table().Count(context.Background(), &CountOptions{Where: ptr(Where("status").Eq("paid"))})
	if err != nil || n != 42 {
		t.Fatalf("count got %d err %v", n, err)
	}
	last := ds.lastReq()
	if last.path != "/data/acme/crm/orders/_count" || last.query.Get("status") != "eq.paid" {
		t.Fatalf("count wire wrong %s %v", last.path, last.query)
	}
}

func TestGetWire(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	ds.setResponse(map[string]any{"id": "abc", "total": 5}, 200)
	row, err := ds.table().Get(context.Background(), "abc", &GetOptions{Select: []string{"id", "total"}})
	if err != nil {
		t.Fatalf("get err %v", err)
	}
	last := ds.lastReq()
	if last.method != "GET" || last.path != "/data/acme/crm/orders/abc" || last.query.Get("_select") != "id,total" {
		t.Fatalf("get wire wrong %s %s %v", last.method, last.path, last.query)
	}
	if row["id"] != "abc" {
		t.Fatalf("get row wrong %v", row)
	}
}

func TestInsertWire(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	ds.setResponse(map[string]any{"id": "new", "total": 7}, 200)
	out, err := ds.table().Insert(context.Background(), map[string]any{"total": 7})
	if err != nil {
		t.Fatalf("insert err %v", err)
	}
	last := ds.lastReq()
	if last.method != "POST" || last.path != "/data/acme/crm/orders" {
		t.Fatalf("insert wire wrong %s %s", last.method, last.path)
	}
	if last.body["total"] != float64(7) {
		t.Fatalf("insert body wrong %v", last.body)
	}
	if out["id"] != "new" {
		t.Fatalf("insert resp wrong %v", out)
	}
}

func TestInsertManyLoops(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	ds.setResponse(map[string]any{"id": "x"}, 200)
	out, err := ds.table().InsertMany(context.Background(), []map[string]any{{"a": 1}, {"a": 2}})
	if err != nil {
		t.Fatalf("insertMany err %v", err)
	}
	if out.Count != 2 || len(out.Items) != 2 {
		t.Fatalf("insertMany got %+v", out)
	}
}

func TestUpdateWire(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	ds.setResponse(map[string]any{"id": "abc", "total": 9}, 200)
	out, err := ds.table().Update(context.Background(), "abc", map[string]any{"total": 9})
	if err != nil {
		t.Fatalf("update err %v", err)
	}
	last := ds.lastReq()
	if last.method != "PATCH" || last.path != "/data/acme/crm/orders/abc" {
		t.Fatalf("update wire wrong %s %s", last.method, last.path)
	}
	if out["id"] != "abc" {
		t.Fatalf("update resp wrong %v", out)
	}
}

func TestDeleteWire(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	ds.setResponse(map[string]any{}, 204)
	if err := ds.table().Delete(context.Background(), "abc"); err != nil {
		t.Fatalf("delete err %v", err)
	}
	last := ds.lastReq()
	if last.method != "DELETE" || last.path != "/data/acme/crm/orders/abc" {
		t.Fatalf("delete wire wrong %s %s", last.method, last.path)
	}
}

// ---------------------------------- list_all ---------------------------------

func TestListAllDrivesPagesAndEmitsDrift(t *testing.T) {
	total2, total3 := 2, 3
	pages := []*PaginatedList{
		{Items: []map[string]any{{"id": float64(1)}}, NextCursor: strptr("2"), Total: &total2},
		{Items: []map[string]any{{"id": float64(2)}}, NextCursor: nil, Total: &total3},
	}
	i := 0
	fetcher := func(_ *string) (*PaginatedList, error) {
		p := pages[i]
		i++
		return p, nil
	}
	out, err := listAll(fetcher, nil)
	if err != nil {
		t.Fatalf("listAll err %v", err)
	}
	if len(out) != 3 {
		t.Fatalf("want 3 emissions got %d: %+v", len(out), out)
	}
	if out[0].Type != "item" || out[1].Type != "drift" || out[1].AddedSince != 1 || out[2].Type != "item" {
		t.Fatalf("drift sequence wrong %+v", out)
	}
}

// ---------------------------------- discover ---------------------------------

func TestDiscoverSlugInspect(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	ds.setResponse(map[string]any{
		"tableName": "orders",
		"columns": []any{
			map[string]any{"name": "id", "type": "uuid"},
			map[string]any{"name": "total", "type": "numeric"},
			map[string]any{"name": "__proto__", "type": "text"}, // skipped
			map[string]any{"name": "ok name", "type": "text"},   // invalid ident, skipped
		},
	}, 200)
	tc, err := ds.client().Tenant("acme").App("crm").Data().Discover(context.Background(), "orders")
	if err != nil {
		t.Fatalf("discover err %v", err)
	}
	if ds.lastReq().path != "/api/v1/tenants/acme/apps/crm/tables/orders/inspect" {
		t.Fatalf("discover path %s", ds.lastReq().path)
	}
	cols := tc.Schema().Columns
	if len(cols) != 2 || cols["id"] != "uuid" || cols["total"] != "number" {
		t.Fatalf("schema columns wrong %v", cols)
	}
}

func TestDiscoverCachesAcrossChains(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	ds.setResponse(map[string]any{"tableName": "orders", "columns": []any{map[string]any{"name": "id", "type": "uuid"}}}, 200)
	// One client; the schema cache lives on it, so two discover() calls from
	// separate tenant().app() chains hit /inspect exactly once.
	client := ds.client()
	ds.reset()
	if _, err := client.Tenant("acme").App("crm").Data().Discover(context.Background(), "orders"); err != nil {
		t.Fatalf("first discover err %v", err)
	}
	first := ds.lastReq()
	if first == nil || first.path != "/api/v1/tenants/acme/apps/crm/tables/orders/inspect" {
		t.Fatalf("first discover path %v", first)
	}
	ds.reset()
	if _, err := client.Tenant("acme").App("crm").Data().Discover(context.Background(), "orders"); err != nil {
		t.Fatalf("second discover err %v", err)
	}
	if ds.lastReq() != nil {
		t.Fatalf("second discover should be served from cache, server saw %v", ds.lastReq().path)
	}
}

func TestDiscoverAppIDFallback(t *testing.T) {
	// slug inspect 404 -> appId fallback resolves via /api/v1/apps, then hits
	// /api/v1/apps/{appId}/tables/{table}.
	var paths []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		switch {
		case strings.HasSuffix(r.URL.Path, "/inspect"):
			w.WriteHeader(404)
			_ = json.NewEncoder(w).Encode(map[string]any{"error": map[string]any{"code": "not_found", "category": "not_found"}})
		case r.URL.Path == "/api/v1/apps":
			_ = json.NewEncoder(w).Encode(map[string]any{"items": []any{map[string]any{"id": "app_123", "slug": "crm"}}})
		default: // /api/v1/apps/app_123/tables/orders
			_ = json.NewEncoder(w).Encode(map[string]any{"tableName": "orders", "columns": []any{map[string]any{"name": "id", "type": "uuid"}}})
		}
	}))
	defer srv.Close()
	c := NewClient(Config{BaseURL: srv.URL, Token: "pat_x", TokenType: TokenTypePAT})
	tc, err := c.Tenant("acme").App("crm").Data().Discover(context.Background(), "orders")
	if err != nil {
		t.Fatalf("discover fallback err %v", err)
	}
	if tc.Schema().Columns["id"] != "uuid" {
		t.Fatalf("fallback schema wrong %v", tc.Schema().Columns)
	}
	joined := strings.Join(paths, " ")
	if !strings.Contains(joined, "/api/v1/apps") || !strings.Contains(joined, "/api/v1/apps/app_123/tables/orders") {
		t.Fatalf("fallback did not resolve appId: %v", paths)
	}
}

func TestDiscover404BecomesTableNotFound(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	ds.setResponse(map[string]any{"error": map[string]any{"code": "not_found", "category": "not_found"}}, 404)
	_, err := ds.client().Tenant("acme").App("crm").Data().Discover(context.Background(), "ghosts")
	assertCode(t, err, "table_not_found")
}

// ------------------------------- schema cache --------------------------------

func TestSchemaCacheGetOrSetCaches(t *testing.T) {
	cache := NewSchemaCache(SchemaCacheOptions{})
	calls := 0
	loader := func() (*DataTableSchema, error) {
		calls++
		return DefineSchema("orders", SchemaShape{"id": "uuid"}), nil
	}
	_, _ = cache.getOrSet("k", loader, false, 0)
	_, _ = cache.getOrSet("k", loader, false, 0)
	if calls != 1 {
		t.Fatalf("expected 1 load, got %d", calls)
	}
	cache.Invalidate("k")
	_, _ = cache.getOrSet("k", loader, false, 0)
	if calls != 2 {
		t.Fatalf("expected 2 loads after invalidate, got %d", calls)
	}
}

func TestSchemaCacheLRUEviction(t *testing.T) {
	cache := NewSchemaCache(SchemaCacheOptions{MaxEntries: 2})
	for _, k := range []string{"a", "b", "c"} {
		cache.Set(k, DefineSchema(k, SchemaShape{"id": "uuid"}), 0)
	}
	if cache.Get("a") != nil {
		t.Fatalf("a should be evicted")
	}
	if cache.Get("c") == nil {
		t.Fatalf("c should be present")
	}
}

func TestSchemaCacheTTLExpiry(t *testing.T) {
	cache := NewSchemaCache(SchemaCacheOptions{TTLMS: 1})
	cache.Set("k", DefineSchema("k", SchemaShape{"id": "uuid"}), 0)
	time.Sleep(5 * time.Millisecond)
	if cache.Get("k") != nil {
		t.Fatalf("entry should have expired")
	}
}

// ------------------------------- like guards ---------------------------------

func TestLikeContainsEscapesWildcards(t *testing.T) {
	expr := Where("name").Like.Contains("50%_off")
	if expr.Value != "%50\\%\\_off%" {
		t.Fatalf("escape wrong: %q", expr.Value)
	}
}

func TestLikeRawRedosGuardDeferred(t *testing.T) {
	expr := Where("name").Like.Raw("%%%%x")
	if expr.BuildErr == nil || expr.BuildErr.Code != "like_pattern_redos" {
		t.Fatalf("expected deferred redos build error, got %v", expr.BuildErr)
	}
}

// --------------------------------- helpers -----------------------------------

func ptr(e QueryExpr) *QueryExpr { return &e }
func strptr(s string) *string    { return &s }

func assertCode(t *testing.T, err error, wantCode string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error with code %q, got nil", wantCode)
	}
	ax, ok := err.(*AxHubError)
	if !ok {
		t.Fatalf("expected *AxHubError, got %T: %v", err, err)
	}
	if ax.Code != wantCode {
		t.Fatalf("error code got %q want %q (%v)", ax.Code, wantCode, ax)
	}
}

func wptr(e QueryExpr) *QueryExpr { return &e }

func TestListCountRequireWhereFilter(t *testing.T) {
	ds := newDataServer()
	defer ds.close()
	// The live data ring 400s an unfiltered list/count (mass-scan guard); the SDK
	// fails fast with where_required before any HTTP request.
	_, listErr := ds.table().List(context.Background(), nil)
	if ax, ok := listErr.(*AxHubError); !ok || ax.Code != "where_required" {
		t.Fatalf("filterless List: expected where_required, got %v", listErr)
	}
	_, countErr := ds.table().Count(context.Background(), nil)
	if ax, ok := countErr.(*AxHubError); !ok || ax.Code != "where_required" {
		t.Fatalf("filterless Count: expected where_required, got %v", countErr)
	}
}
