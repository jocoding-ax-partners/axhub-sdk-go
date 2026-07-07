package axhub

import (
	"context"
	"encoding/json"
	"strconv"
)

// RawDbColumn describes one column of a raw (physical Postgres) DB table.
type RawDbColumn struct {
	Name     string `json:"name"`
	DataType string `json:"dataType"`
	Nullable bool   `json:"nullable"`
}

// RawDbTable is a physical table exposed by an app's raw DB (opt-in Postgres).
type RawDbTable struct {
	Name    string        `json:"name"`
	Managed bool          `json:"managed"`
	Columns []RawDbColumn `json:"columns"`
}

// RawDbTableRows is one page of rows read from a raw DB table. HasMore is a
// per_page+1 probe — there is no exact total.
type RawDbTableRows struct {
	Rows    []map[string]any `json:"rows"`
	Page    int              `json:"page"`
	PerPage int              `json:"perPage"`
	HasMore bool             `json:"hasMore"`
}

// RawDbTableRowsOptions are optional pagination controls for TableRows. A nil
// *RawDbTableRowsOptions (or zero values) uses the backend default page size.
type RawDbTableRowsOptions struct {
	PerPage int
	Page    int
}

// RawDbClient is the typed, read-only raw-DB facade, mirroring the Node SDK's
// sdk.apps.rawDb.*. Reach it via client.Apps.RawDb().
//
// Raw DB is an opt-in physical Postgres surface. This facade is read-only
// (introspection + row reads); enabling raw DB and writing rows are separate
// concerns (writes go through the deployed app's DATABASE_URL, not this SDK).
type RawDbClient struct{ apps *AppsClient }

// RawDb returns the typed raw-DB read facade.
func (a *AppsClient) RawDb() *RawDbClient { return &RawDbClient{apps: a} }

// Tables lists the raw DB tables for an app, with typed column metadata.
//
// A successful 2xx call that returns an empty slice means the app genuinely has
// no raw DB tables — either raw DB is not enabled for the app, or it has zero
// tables. A 4xx authentication or permission failure returns a non-nil error,
// so across the expected 2xx/4xx responses an empty slice with a nil error
// means "empty", not "auth failed" (resolving the ambiguity the raw map
// response leaves).
func (r *RawDbClient) Tables(ctx context.Context, appID string) ([]RawDbTable, error) {
	if appID == "" {
		return nil, &AxHubError{Category: "validation", Code: "required", Message: "appID is required"}
	}
	resp, err := r.apps.client.Request(ctx, "schemaGetApiV1AppsByAppIDDbTables", map[string]string{"appID": appID}, nil, nil)
	if err != nil {
		return nil, err
	}
	var wrap struct {
		Tables []RawDbTable `json:"tables"`
	}
	if err := decodeRawDb(resp, &wrap); err != nil {
		return nil, err
	}
	return wrap.Tables, nil
}

// TableRows reads one page of rows from a raw DB table. Pass nil opts for the
// backend default page size.
func (r *RawDbClient) TableRows(ctx context.Context, appID, table string, opts *RawDbTableRowsOptions) (*RawDbTableRows, error) {
	if appID == "" || table == "" {
		return nil, &AxHubError{Category: "validation", Code: "required", Message: "appID and table are required"}
	}
	var query map[string]string
	if opts != nil {
		query = map[string]string{}
		if opts.PerPage > 0 {
			query["per_page"] = strconv.Itoa(opts.PerPage)
		}
		if opts.Page > 0 {
			query["page"] = strconv.Itoa(opts.Page)
		}
	}
	resp, err := r.apps.client.Request(ctx, "schemaGetApiV1AppsByAppIDDbTablesByTableRows", map[string]string{"appID": appID, "table": table}, query, nil)
	if err != nil {
		return nil, err
	}
	var page RawDbTableRows
	if err := decodeRawDb(resp, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

// decodeRawDb re-marshals a decoded response map into a typed struct via JSON.
// The operation-id transport already camelCases response keys, so the DTO json
// tags use camelCase.
func decodeRawDb(src map[string]any, dst any) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}
