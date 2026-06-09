package axhub

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Serialize the predicate DSL into backend filter query params (mirrors node
// where-serializer.ts). Each pushable atom becomes column=<op>.<value>
// (PostgREST-style). Repeated columns collapse into multiple values (q.Add), so
// the transport emits repeated query params. Only top-level And(...) of pushable
// atoms and bare atoms are accepted; or/not/raw and nested-and return a
// ValidationError — matching the live backend's filter grammar (gap-matrix S7-S9).

type pushableFilter struct {
	column string
	value  string
}

// serializeWhere returns the query params for a where expr, or an error. A nil
// expr yields empty params. Any deferred BuildErr (e.g. ReDoS guard) surfaces
// here. url.Values carries repeated columns as multiple values.
func serializeWhere(expr *QueryExpr) (url.Values, *AxHubError) {
	out := url.Values{}
	if expr == nil {
		return out, nil
	}
	filters, err := collectPushableFilters(expr, true)
	if err != nil {
		return nil, err
	}
	for _, f := range filters {
		out.Add(f.column, f.value)
	}
	return out, nil
}

func collectPushableFilters(expr *QueryExpr, allowAnd bool) ([]pushableFilter, *AxHubError) {
	if expr.BuildErr != nil {
		return nil, expr.BuildErr
	}
	switch expr.Op {
	case "eq", "ne", "gt", "gte", "lt", "lte", "like":
		return []pushableFilter{{column: expr.Column, value: expr.Op + "." + stringifyFilterValue(expr.Value)}}, nil
	case "in":
		values := make([]string, 0, len(expr.Values))
		for _, v := range expr.Values {
			s := stringifyFilterValue(v)
			if strings.Contains(s, ",") {
				return nil, newValidationError(
					fmt.Sprintf("IN filter values cannot contain commas because the live backend uses comma-separated IN lists (bad value: %s)", s),
					"filter_in_comma",
				)
			}
			values = append(values, s)
		}
		return []pushableFilter{{column: expr.Column, value: "in." + strings.Join(values, ",")}}, nil
	case "and":
		if allowAnd {
			out := make([]pushableFilter, 0, len(expr.Clauses))
			for i := range expr.Clauses {
				sub, err := collectPushableFilters(&expr.Clauses[i], false)
				if err != nil {
					return nil, err
				}
				out = append(out, sub...)
			}
			return out, nil
		}
	}
	return nil, newValidationError(
		fmt.Sprintf("Data where clause '%s' cannot be pushed to the live backend; use top-level and(eq/ne/gt/gte/lt/lte/in/like) only", expr.Op),
		"unsupported_filter",
	)
}

func stringifyFilterValue(value any) string {
	switch v := value.(type) {
	case nil:
		return "null"
	case time.Time:
		return v.UTC().Format(time.RFC3339Nano)
	case bool:
		if v {
			return "true"
		}
		return "false"
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'g', -1, 64)
	case fmt.Stringer:
		return v.String()
	default:
		raw, err := json.Marshal(v)
		if err != nil {
			return fmt.Sprint(v)
		}
		return string(raw)
	}
}
