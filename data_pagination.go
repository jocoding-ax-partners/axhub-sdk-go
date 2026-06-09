package axhub

import (
	"strconv"
	"strings"
)

// Offset pagination helpers used by the data layer (subset of node core
// pagination.ts). Keyset encode/decode is intentionally NOT ported: the live AX
// Hub data API is offset-only, so the data layer only needs the order-by
// normalizer and the cursor-shape guards that reject legacy keyset tokens (see
// node, gap-matrix S7-S9).

const MaxCursorTokenLength = 4096

// OrderField is one sort field with a direction ("asc" | "desc").
type OrderField struct {
	Field string
	Dir   string
}

// PaginatedList is the offset-pagination envelope returned by List.
type PaginatedList struct {
	Items       []map[string]any
	NextCursor  *string
	FirstCursor *string
	HasNext     bool
	HasPrev     bool
	Total       *int
	TotalIsExact bool
}

// ListAllItem is either an item (Type=="item") or a drift marker
// (Type=="drift") emitted when the backend total grows mid-scan.
type ListAllItem struct {
	Type       string // "item" | "drift"
	Value      map[string]any
	AddedSince int
}

// normalizeOrderBy parses a string or []OrderField into normalized fields,
// appending an "id" asc tiebreaker when fields are present and none is "id"
// (mirrors node/python normalize_order_by).
func normalizeOrderBy(orderBy any) []OrderField {
	var fields []OrderField
	switch ob := orderBy.(type) {
	case string:
		for _, part := range strings.Split(ob, ",") {
			trimmed := strings.TrimSpace(part)
			var f OrderField
			if strings.HasPrefix(trimmed, "-") {
				f = OrderField{Field: trimmed[1:], Dir: "desc"}
			} else if strings.HasPrefix(trimmed, "+") {
				f = OrderField{Field: trimmed[1:], Dir: "asc"}
			} else {
				f = OrderField{Field: trimmed, Dir: "asc"}
			}
			if f.Field != "" {
				fields = append(fields, f)
			}
		}
	case []OrderField:
		for _, p := range ob {
			dir := p.Dir
			if dir == "" {
				dir = "asc"
			}
			fields = append(fields, OrderField{Field: p.Field, Dir: dir})
		}
	}
	if len(fields) > 0 {
		hasID := false
		for _, f := range fields {
			if f.Field == "id" {
				hasID = true
				break
			}
		}
		if !hasID {
			fields = append(fields, OrderField{Field: "id", Dir: "asc"})
		}
	}
	return fields
}

// serializeOrderBy serializes order-by into the `sort` query param value, or
// ("", false) when empty. A bare string with no parseable fields round-trips.
func serializeOrderBy(orderBy any) (string, bool) {
	normalized := normalizeOrderBy(orderBy)
	if len(normalized) == 0 {
		if s, ok := orderBy.(string); ok && s != "" {
			return s, true
		}
		return "", false
	}
	parts := make([]string, len(normalized))
	for i, f := range normalized {
		if f.Dir == "desc" {
			parts[i] = "-" + f.Field
		} else {
			parts[i] = f.Field
		}
	}
	return strings.Join(parts, ","), true
}

// isV2Cursor reports whether a token is a v2 keyset cursor (rejected by the
// offset-only data API).
func isV2Cursor(token string) bool {
	return strings.HasPrefix(token, "v2:")
}

// listAll drives a paginated fetcher to exhaustion, yielding each item and a
// drift marker when the backend total grows mid-iteration (mirrors node listAll).
func listAll(fetcher func(cursor *string) (*PaginatedList, error), startCursor *string) ([]ListAllItem, error) {
	var out []ListAllItem
	cursor := startCursor
	var initialTotal *int
	var lastTotal *int
	for {
		page, err := fetcher(cursor)
		if err != nil {
			return out, err
		}
		if page.Total != nil {
			if initialTotal == nil {
				t := *page.Total
				initialTotal = &t
				lt := *page.Total
				lastTotal = &lt
			} else {
				base := *initialTotal
				if lastTotal != nil {
					base = *lastTotal
				}
				if *page.Total > base {
					out = append(out, ListAllItem{Type: "drift", AddedSince: *page.Total - base})
					lt := *page.Total
					lastTotal = &lt
				}
			}
		}
		for _, item := range page.Items {
			out = append(out, ListAllItem{Type: "item", Value: item})
		}
		if page.NextCursor == nil {
			return out, nil
		}
		cursor = page.NextCursor
	}
}

func intToCursorPtr(n int) *string {
	s := strconv.Itoa(n)
	return &s
}
