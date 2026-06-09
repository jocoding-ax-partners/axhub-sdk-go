package axhub

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"time"
)

// Runtime schema introspection via the table /inspect endpoint, with an
// appId-resolution fallback and error normalization (mirrors node discover.ts).
//
// Primary:  GET /api/v1/tenants/{t}/apps/{a}/tables/{table}/inspect
// Fallback: on 404, resolve appId by scanning GET /api/v1/apps?tenant_slug=...,
//           then GET /api/v1/apps/{appId}/tables/{table}.
// Neither endpoint has a generated operation-id, so discover goes through the
// raw-path transport. camelizeResponse=true so table_name/tableName both
// resolve (inspect payload is metadata, not user row data).

const (
	appLookupPageSize  = 100
	appLookupMaxPages  = 10
	appLookupBudgetMS  = 5_000
)

var (
	forbiddenColumnNames = map[string]bool{"__proto__": true, "constructor": true, "prototype": true}
	columnNameRe         = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)
)

func fetchDiscoveredSchema(ctx context.Context, c *Client, tenantSlug, appSlug, table string) (*DataTableSchema, error) {
	// The appId path is the route the `axhub` CLI uses and is verified to work with
	// a data-ring PAT (2026-06). The slug /inspect route rejects a slug in the
	// {tenant} path segment on the live backend ("tenant_id 형식이 잘못됐어요", HTTP
	// 400) — a 400 not a 404, so the old slug-first order never reached this working
	// path. appId is primary; slug inspect is a best-effort fallback. The appId error
	// is the meaningful one, so it is what surfaces.
	schema, err := fetchAppIDInspect(ctx, c, tenantSlug, appSlug, table)
	if err == nil {
		return schema, nil
	}
	if fallback, fbErr := fetchSlugInspect(ctx, c, tenantSlug, appSlug, table); fbErr == nil {
		return fallback, nil
	}
	return nil, normalizeDiscoverError(err, tenantSlug, appSlug, table)
}

func fetchSlugInspect(ctx context.Context, c *Client, tenantSlug, appSlug, table string) (*DataTableSchema, error) {
	path := fmt.Sprintf("/api/v1/tenants/%s/apps/%s/tables/%s/inspect",
		url.PathEscape(tenantSlug), url.PathEscape(appSlug), url.PathEscape(table))
	raw, err := c.requestRaw(ctx, "GET", path, nil, nil, true)
	if err != nil {
		return nil, err
	}
	return schemaFromInspectResult(table, raw), nil
}

func fetchAppIDInspect(ctx context.Context, c *Client, tenantSlug, appSlug, table string) (*DataTableSchema, error) {
	appID, err := resolveAppID(ctx, c, tenantSlug, appSlug)
	if err != nil {
		return nil, err
	}
	if appID == "" {
		return nil, newTableNotFoundError(fmt.Sprintf("Dynamic data table '%s' was not found", table), "")
	}
	path := fmt.Sprintf("/api/v1/apps/%s/tables/%s", url.PathEscape(appID), url.PathEscape(table))
	raw, err := c.requestRaw(ctx, "GET", path, nil, nil, true)
	if err != nil {
		return nil, err
	}
	return schemaFromInspectResult(table, raw), nil
}

func resolveAppID(ctx context.Context, c *Client, tenantSlug, appSlug string) (string, error) {
	startedAt := time.Now()
	cursor := ""
	for page := 0; page < appLookupMaxPages; page++ {
		if time.Since(startedAt) > appLookupBudgetMS*time.Millisecond {
			return "", newIntrospectFailedError(
				fmt.Sprintf("app lookup budget exceeded (%dms) while searching for slug '%s' in tenant '%s'", appLookupBudgetMS, appSlug, tenantSlug),
				0, false, "")
		}
		query := url.Values{}
		query.Set("tenant_slug", tenantSlug)
		query.Set("limit", fmt.Sprintf("%d", appLookupPageSize))
		if cursor != "" {
			query.Set("cursor", cursor)
		}
		raw, err := c.requestRaw(ctx, "GET", "/api/v1/apps", query, nil, true)
		if err != nil {
			return "", err
		}
		items, _ := raw["items"].([]any)
		for _, it := range items {
			app, ok := it.(map[string]any)
			if !ok {
				continue
			}
			id, idOK := app["id"].(string)
			if app["slug"] == appSlug && idOK && id != "" {
				return id, nil
			}
		}
		// Empty page on the first request means the tenant truly has no apps.
		if page == 0 && len(items) == 0 {
			return "", nil
		}
		next := stringFromAny(raw["next_cursor"])
		if next == "" {
			next = stringFromAny(raw["nextCursor"])
		}
		if next == "" {
			return "", nil
		}
		cursor = next
	}
	return "", newScanLimitExceededError(
		fmt.Sprintf("App lookup exceeded %d pages x %d apps without finding slug '%s'", appLookupMaxPages, appLookupPageSize, appSlug))
}

func normalizeDiscoverError(err error, tenantSlug, appSlug, table string) error {
	if ax, ok := err.(*AxHubError); ok {
		switch ax.Code {
		case "table_not_found", "introspect_failed", "scan_limit_exceeded":
			return ax
		}
		if isNotFound(err) {
			return newTableNotFoundError(fmt.Sprintf("Dynamic data table '%s' was not found", table), ax.RequestID)
		}
		if ax.Status >= 500 {
			return newIntrospectFailedError(fmt.Sprintf("Failed to introspect dynamic data table '%s'", table), ax.Status, ax.Retryable, ax.RequestID)
		}
		return ax
	}
	return err
}

func schemaFromInspectResult(table string, raw map[string]any) *DataTableSchema {
	shape := SchemaShape{}
	if columns, ok := raw["columns"].([]any); ok {
		for _, col := range columns {
			column, ok := col.(map[string]any)
			if !ok {
				continue
			}
			name, _ := column["name"].(string)
			if name == "" || forbiddenColumnNames[name] || !columnNameRe.MatchString(name) {
				continue
			}
			colType, _ := column["type"].(string)
			shape[name] = columnTypeToDef(colType)
		}
	}
	tableName := table
	for _, key := range []string{"tableName", "table_name", "name"} {
		if v, ok := raw[key].(string); ok && v != "" {
			tableName = v
			break
		}
	}
	return DefineSchema(tableName, shape)
}

func columnTypeToDef(colType string) string {
	switch colType {
	case "uuid":
		return "uuid"
	case "int", "integer", "bigint":
		return "integer"
	case "float", "numeric", "double precision", "real":
		return "number"
	case "bool", "boolean":
		return "boolean"
	case "timestamp", "timestamptz", "timestamp with time zone":
		return "timestamp"
	case "json", "jsonb":
		return "json"
	default: // text / varchar / character varying / unknown
		return "string"
	}
}

func isNotFound(err error) bool {
	ax, ok := err.(*AxHubError)
	if !ok {
		return false
	}
	return ax.Category == "not_found" || ax.Status == 404
}

func stringFromAny(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
