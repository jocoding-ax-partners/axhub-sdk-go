package axhub

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"strconv"
)

// Ergonomic data layer: fluent builder + dynamic-table CRUD + offset pagination
// (mirrors node index.ts DataClient / TenantDataFactory / AppDataFactory /
// DataTableClient).
//
// Public chain: client.Tenant(slug).App(slug).Data().Table(name) and
// .Data().Discover(ctx, table). Named distinctly from the route-table
// Client.Data()/DataClient facade (which is the generated Schema operation
// surface) to avoid collision.
//
// Wire paths (EXACTLY as node, via requestRaw so row bodies and the list
// envelope return verbatim, no snake->camel rewriting):
//
//	list / insert          GET|POST          /data/{tenant}/{app}/{table}
//	get / update / delete   GET|PATCH|DELETE  /data/{tenant}/{app}/{table}/{id}
//	count                   GET               /data/{tenant}/{app}/{table}/_count

// TenantContext is the tenant-scoped entry to the ergonomic data layer.
type TenantContext struct {
	client     *Client
	tenantSlug string
}

// AppContext is the app-scoped entry to the ergonomic data layer.
type AppContext struct {
	client     *Client
	tenantSlug string
	appSlug    string
}

// DataFactory mints table clients and runs discover() under one tenant/app.
type DataFactory struct {
	client     *Client
	tenantSlug string
	appSlug    string
}

// Tenant scopes the ergonomic data layer to a tenant slug.
func (c *Client) Tenant(tenantSlug string) *TenantContext {
	return &TenantContext{client: c, tenantSlug: tenantSlug}
}

// App scopes to an app slug under the tenant.
func (t *TenantContext) App(appSlug string) *AppContext {
	return &AppContext{client: t.client, tenantSlug: t.tenantSlug, appSlug: appSlug}
}

// Data returns the data factory for this tenant/app.
func (a *AppContext) Data() *DataFactory {
	return &DataFactory{client: a.client, tenantSlug: a.tenantSlug, appSlug: a.appSlug}
}

// Table binds a DataResource to a table name.
func (d *DataFactory) Table(table string) *DataResource {
	return &DataResource{client: d.client, tenantSlug: d.tenantSlug, appSlug: d.appSlug, tableName: table}
}

// TableSchema binds a DataResource to a defined schema (enables select +
// insert/update validation).
func (d *DataFactory) TableSchema(schema *DataTableSchema) *DataResource {
	return &DataResource{client: d.client, tenantSlug: d.tenantSlug, appSlug: d.appSlug, tableName: schema.Table, schema: schema}
}

// Discover introspects the table schema (cached on the root client) and returns
// a schema-bound DataResource.
func (d *DataFactory) Discover(ctx context.Context, table string, opts ...DiscoverOptions) (*DataResource, error) {
	var opt DiscoverOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	key := schemaCacheKey(d.tenantSlug, d.appSlug, table)
	schema, err := d.client.dataSchemaCache.getOrSet(
		key,
		func() (*DataTableSchema, error) {
			return fetchDiscoveredSchema(ctx, d.client, d.tenantSlug, d.appSlug, table)
		},
		opt.Fresh,
		opt.TTLMS,
	)
	if err != nil {
		return nil, err
	}
	return &DataResource{client: d.client, tenantSlug: d.tenantSlug, appSlug: d.appSlug, tableName: schema.Table, schema: schema}, nil
}

// InvalidateSchema drops a discovered schema from the cache (one table, or all
// when table == "").
func (d *DataFactory) InvalidateSchema(table string) {
	if table == "" {
		d.client.dataSchemaCache.Invalidate("")
		return
	}
	d.client.dataSchemaCache.Invalidate(schemaCacheKey(d.tenantSlug, d.appSlug, table))
}

// DiscoverOptions controls schema cache behavior on Discover.
type DiscoverOptions struct {
	Fresh bool
	TTLMS int
}

// DataResource is bound to one {tenant}/{app}/{table} with CRUD + pagination.
type DataResource struct {
	client     *Client
	tenantSlug string
	appSlug    string
	tableName  string
	schema     *DataTableSchema
}

// Schema returns the schema this resource was bound to (nil for a bare Table()).
func (r *DataResource) Schema() *DataTableSchema { return r.schema }

func (r *DataResource) path(id string) string {
	base := fmt.Sprintf("/data/%s/%s/%s", url.PathEscape(r.tenantSlug), url.PathEscape(r.appSlug), url.PathEscape(r.tableName))
	if id == "" {
		return base
	}
	return base + "/" + url.PathEscape(id)
}

// ListOptions configures List/ListAll. WhereExpr/OrderBy/Select are optional;
// pagination uses offset-only cursor/page (after/before/direction are rejected).
type ListOptions struct {
	Where     *QueryExpr
	OrderBy   any // string or []OrderField
	Select    []string
	Page      int
	PageSize  int
	Limit     int
	Cursor    string
	// NOTE: mirrors node behavior — offset-only API. Setting any of these is a
	// LegacyCursorError (see gap-matrix S7-S9).
	After     string
	Before    string
	Direction string
}

// List fetches one offset page (mirrors node DataTableClient.list).
func (r *DataResource) List(ctx context.Context, opts *ListOptions) (*PaginatedList, error) {
	if opts == nil {
		opts = &ListOptions{}
	}
	if err := validateSelectColumns(r.schema, opts.Select); err != nil {
		return nil, err
	}
	if err := rejectLegacyPageOptions(opts, r.tableName); err != nil {
		return nil, err
	}
	page, perr := resolveOffsetPage(opts.Cursor, opts.Page, r.tableName)
	if perr != nil {
		return nil, perr
	}
	var perPageSrc int
	if opts.PageSize != 0 {
		perPageSrc = opts.PageSize
	} else {
		perPageSrc = opts.Limit
	}
	perPage, hasPerPage := clampPerPage(perPageSrc, opts.PageSize != 0 || opts.Limit != 0)

	query, werr := serializeWhere(opts.Where)
	if werr != nil {
		return nil, werr
	}
	if hasPerPage {
		query.Set("per_page", strconv.Itoa(perPage))
	}
	if page != 1 {
		query.Set("page", strconv.Itoa(page))
	}
	if sort, ok := serializeOrderBy(opts.OrderBy); ok {
		query.Set("sort", sort)
	}
	if sel, ok := serializeSelect(opts.Select); ok {
		query.Set("_select", sel)
	}

	raw, err := r.client.requestRaw(ctx, "GET", r.path(""), query, nil, false)
	if err != nil {
		return nil, mapWhereRequired("list", err)
	}
	items := projectRows(rowsFromAny(raw["items"]), opts.Select)
	// NOTE: mirrors node behavior — currentPage falls back to the requested
	// page, hasNext reads the backend `has_more` flag verbatim, hasPrev is
	// derived client-side from the page number (see gap-matrix S7-S9).
	currentPage := page
	if p, ok := intFromAny(raw["page"]); ok {
		currentPage = p
	}
	hasNext := boolFromAny(raw["has_more"])
	hasPrev := currentPage > 1
	result := &PaginatedList{
		Items:        items,
		HasNext:      hasNext,
		HasPrev:      hasPrev,
		TotalIsExact: false,
	}
	if hasNext {
		result.NextCursor = intToCursorPtr(currentPage + 1)
	}
	if hasPrev {
		result.FirstCursor = intToCursorPtr(currentPage - 1)
	}
	return result, nil
}

// ListAll drives List to exhaustion, returning every item plus drift markers
// when the backend total grows mid-scan (mirrors node listAll).
func (r *DataResource) ListAll(ctx context.Context, opts *ListOptions) ([]ListAllItem, error) {
	if opts == nil {
		opts = &ListOptions{}
	}
	base := *opts
	fetcher := func(cursor *string) (*PaginatedList, error) {
		page := base
		page.Page = 0
		page.Cursor = ""
		if cursor != nil {
			page.Cursor = *cursor
		}
		return r.List(ctx, &page)
	}
	var start *string
	if opts.Cursor != "" {
		start = &opts.Cursor
	}
	return listAll(fetcher, start)
}

// CountOptions configures Count.
type CountOptions struct {
	Where *QueryExpr
}

// Count returns the row count for an optional where filter.
func (r *DataResource) Count(ctx context.Context, opts *CountOptions) (int, error) {
	var where *QueryExpr
	if opts != nil {
		where = opts.Where
	}
	query, werr := serializeWhere(where)
	if werr != nil {
		return 0, werr
	}
	raw, err := r.client.requestRaw(ctx, "GET", r.path("")+"/_count", query, nil, false)
	if err != nil {
		return 0, mapWhereRequired("count", err)
	}
	count, _ := intFromAny(raw["count"])
	return count, nil
}

// GetOptions configures Get.
type GetOptions struct {
	Select []string
}

// Get fetches a single row by id (mirrors node DataTableClient.get).
func (r *DataResource) Get(ctx context.Context, id string, opts *GetOptions) (map[string]any, error) {
	var selectCols []string
	if opts != nil {
		selectCols = opts.Select
	}
	if err := validateSelectColumns(r.schema, selectCols); err != nil {
		return nil, err
	}
	var query url.Values
	if sel, ok := serializeSelect(selectCols); ok {
		query = url.Values{}
		query.Set("_select", sel)
	}
	raw, err := r.client.requestRaw(ctx, "GET", r.path(id), query, nil, false)
	if err != nil {
		return nil, err
	}
	return projectRow(raw, selectCols), nil
}

// Insert creates one row (mirrors node DataTableClient.insert).
func (r *DataResource) Insert(ctx context.Context, row map[string]any) (map[string]any, error) {
	if err := runSchemaValidation(r.schema, row, "insert"); err != nil {
		return nil, err
	}
	return r.client.requestRaw(ctx, "POST", r.path(""), nil, row, false)
}

// BulkResult is the {items, count} shape returned by InsertMany.
type BulkResult struct {
	Items []map[string]any
	Count int
}

// InsertMany loops single inserts (no bulk endpoint exists) and returns
// {items, count} (mirrors node DataTableClient.insertMany / gap-matrix).
func (r *DataResource) InsertMany(ctx context.Context, rows []map[string]any) (*BulkResult, error) {
	for _, row := range rows {
		if err := runSchemaValidation(r.schema, row, "insert"); err != nil {
			return nil, err
		}
	}
	items := make([]map[string]any, 0, len(rows))
	for _, row := range rows {
		// NOTE: mirrors node behavior — context cancellation aborts mid-loop in
		// place of node's AbortError, surfacing how many rows committed.
		if err := ctx.Err(); err != nil {
			return nil, &AxHubError{Category: "abort", Code: "aborted", Message: fmt.Sprintf("insertMany aborted after %d of %d rows", len(items), len(rows))}
		}
		out, err := r.Insert(ctx, row)
		if err != nil {
			return nil, err
		}
		items = append(items, out)
	}
	return &BulkResult{Items: items, Count: len(items)}, nil
}

// Update patches a row by id (mirrors node DataTableClient.update).
func (r *DataResource) Update(ctx context.Context, id string, patch map[string]any) (map[string]any, error) {
	if err := runSchemaValidation(r.schema, patch, "update"); err != nil {
		return nil, err
	}
	return r.client.requestRaw(ctx, "PATCH", r.path(id), nil, patch, false)
}

// Delete removes a row by id (mirrors node DataTableClient.delete).
func (r *DataResource) Delete(ctx context.Context, id string) error {
	_, err := r.client.requestRaw(ctx, "DELETE", r.path(id), nil, nil, false)
	return err
}

// ----------------------------- pagination guards -----------------------------

func clampPerPage(value int, present bool) (int, bool) {
	if !present {
		return 0, false
	}
	v := value
	if v > 100 {
		v = 100
	}
	if v < 1 {
		v = 1
	}
	return v, true
}

func rejectLegacyPageOptions(opts *ListOptions, tableName string) *AxHubError {
	if opts.After != "" || opts.Before != "" || opts.Direction != "" {
		return newLegacyCursorError("after/before keyset cursors are not supported by the live AX Hub data API; use cursor/page numeric offset pagination")
	}
	return nil
}

func resolveOffsetPage(cursor string, page int, tableName string) (int, *AxHubError) {
	if cursor != "" {
		if err := validatePlainCursor(cursor, tableName); err != nil {
			return 0, err
		}
		n, _ := strconv.Atoi(cursor)
		return n, nil
	}
	if page == 0 {
		return 1, nil
	}
	if page < 1 {
		return 0, newInvalidCursorError("page must be a positive integer")
	}
	return page, nil
}

func validatePlainCursor(cursor, tableName string) *AxHubError {
	if len(cursor) > MaxCursorTokenLength {
		return newInvalidCursorError(fmt.Sprintf("Cursor token exceeds maximum size (%d chars)", MaxCursorTokenLength))
	}
	if len(cursor) >= 3 && cursor[:3] == "v1:" {
		return newLegacyCursorError("Legacy v1: cursor token is not compatible with AX Hub offset-only pagination; restart pagination without cursor")
	}
	if isV2Cursor(cursor) {
		return newLegacyCursorError("v2 keyset cursors are not supported by the live AX Hub data API; restart pagination and use the numeric cursor returned by list()")
	}
	n, err := strconv.Atoi(cursor)
	if err != nil || n < 1 {
		return newInvalidCursorError("Plain cursor must be a positive integer page or a v2: keyset token")
	}
	return nil
}

// ------------------------------- any coercion --------------------------------

func rowsFromAny(v any) []map[string]any {
	arr, ok := v.([]any)
	if !ok {
		return nil
	}
	out := make([]map[string]any, 0, len(arr))
	for _, item := range arr {
		if m, ok := item.(map[string]any); ok {
			out = append(out, m)
		}
	}
	return out
}

func intFromAny(v any) (int, bool) {
	switch n := v.(type) {
	case float64:
		if math.Trunc(n) == n {
			return int(n), true
		}
		return int(n), true
	case int:
		return n, true
	case int64:
		return int(n), true
	}
	return 0, false
}

func boolFromAny(v any) bool {
	b, _ := v.(bool)
	return b
}
