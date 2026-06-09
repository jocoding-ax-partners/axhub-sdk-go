package axhub

import "strings"

// Column projection: select serialization, validation, and client-side row
// narrowing (mirrors node projection.ts).

// serializeSelect joins columns with commas into the _select query param. A nil
// select yields ("", false) so callers can omit the param.
func serializeSelect(selectCols []string) (string, bool) {
	if selectCols == nil {
		return "", false
	}
	return strings.Join(selectCols, ","), true
}

// validateSelectColumns rejects an empty select and, when a schema is known,
// unknown columns.
func validateSelectColumns(schema *DataTableSchema, selectCols []string) *AxHubError {
	if selectCols == nil {
		return nil
	}
	if len(selectCols) == 0 {
		return newValidationError("select must include at least one column; omit select to fetch full rows", "select_empty")
	}
	if schema == nil {
		return nil
	}
	var invalid []string
	for _, c := range selectCols {
		if _, ok := schema.Columns[c]; !ok {
			invalid = append(invalid, c)
		}
	}
	if len(invalid) == 0 {
		return nil
	}
	plural := ""
	if len(invalid) != 1 {
		plural = "s"
	}
	return newValidationError("select contains unknown column"+plural+": "+strings.Join(invalid, ", "), "select_unknown_column")
}

// projectRow narrows a returned row to the selected keys client-side.
func projectRow(row map[string]any, selectCols []string) map[string]any {
	if selectCols == nil {
		return row
	}
	out := make(map[string]any, len(selectCols))
	for _, k := range selectCols {
		if v, ok := row[k]; ok {
			out[k] = v
		}
	}
	return out
}

// projectRows narrows each returned row to the selected keys client-side.
func projectRows(rows []map[string]any, selectCols []string) []map[string]any {
	if selectCols == nil {
		return rows
	}
	out := make([]map[string]any, len(rows))
	for i, r := range rows {
		out[i] = projectRow(r, selectCols)
	}
	return out
}
