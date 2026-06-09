package axhub

import "strings"

// Predicate DSL: Where(col).Eq(v), And(...), Or/Not/Raw, plus LIKE escaping and
// ReDoS guards (mirrors node dsl/ops.ts).
//
// IDIOMATIC-GO DEVIATION: node/python builders throw synchronously (e.g.
// like.Raw() ReDoS guard). Go cannot return an error from a chained
// Where(col).Raw(v) call cleanly, so a builder that would have thrown instead
// stamps a deferred BuildErr onto the QueryExpr. serializeWhere checks BuildErr
// first and returns it, so the error surfaces from List/Count — not the builder.

const (
	MaxLikePatternLength      = 1024
	MaxConsecutiveWildcards   = 4
	MaxLikeAlternationSegment = 6
)

// QueryExpr is a node in the predicate tree. Plain struct (no interface) so the
// where-serializer can walk it without type switches on dynamic types.
type QueryExpr struct {
	Op string // eq | ne | gt | gte | lt | lte | like | in | and | or | not | raw

	Column string
	Value  any   // for binary ops
	Values []any // for "in"

	Clauses []QueryExpr // for and/or
	Clause  *QueryExpr  // for not

	SQL    string // for raw
	Params []any  // for raw

	// BuildErr carries an error stamped at build time (e.g. ReDoS guard) so it
	// surfaces from serializeWhere -> List/Count rather than the builder.
	BuildErr *AxHubError
}

// EscapeLike escapes the LIKE metacharacters \\, %, and _.
func EscapeLike(value string) string {
	if value == "" {
		return value
	}
	var b strings.Builder
	for _, r := range value {
		if r == '\\' || r == '%' || r == '_' {
			b.WriteByte('\\')
		}
		b.WriteRune(r)
	}
	return b.String()
}

// AssertSafeLikePattern rejects LIKE patterns that translate to
// catastrophic-backtracking regex shapes (mirrors node assertSafeLikePattern).
func AssertSafeLikePattern(pattern string) *AxHubError {
	if len(pattern) > MaxLikePatternLength {
		return newValidationError("LIKE pattern exceeds 1024 chars; refuse to compile", "like_pattern_too_long")
	}
	runOfWildcards := 0
	segments := 0
	runes := []rune(pattern)
	for i := 0; i < len(runes); i++ {
		ch := runes[i]
		if ch == '\\' {
			i++ // skip the escaped char
			runOfWildcards = 0
			continue
		}
		if ch == '%' {
			runOfWildcards++
			if runOfWildcards >= MaxConsecutiveWildcards {
				return newValidationError("LIKE pattern has consecutive '%'; refuse to compile (ReDoS guard)", "like_pattern_redos")
			}
		} else {
			if runOfWildcards == 1 {
				segments++
			}
			runOfWildcards = 0
		}
	}
	if segments > MaxLikeAlternationSegment {
		return newValidationError("LIKE pattern has too many '%X%' alternation segments; refuse to compile (ReDoS guard)", "like_pattern_redos")
	}
	return nil
}

// Raw builds a non-pushable raw SQL clause (rejected by the where-serializer).
func Raw(sql string, params ...any) QueryExpr {
	expr := QueryExpr{Op: "raw", SQL: sql}
	if len(params) > 0 {
		expr.Params = params
	}
	return expr
}

// And combines clauses (only top-level And of pushable atoms is serializable).
func And(clauses ...QueryExpr) QueryExpr {
	return QueryExpr{Op: "and", Clauses: clauses}
}

// Or combines clauses (non-pushable; rejected by the where-serializer).
func Or(clauses ...QueryExpr) QueryExpr {
	return QueryExpr{Op: "or", Clauses: clauses}
}

// Not negates a clause (non-pushable; rejected by the where-serializer).
func Not(clause QueryExpr) QueryExpr {
	return QueryExpr{Op: "not", Clause: &clause}
}

// WhereBuilder builds a single-column predicate. Like is the LIKE sub-builder.
type WhereBuilder struct {
	name string
	Like LikeBuilder
}

// LikeBuilder builds LIKE predicates with escaping + ReDoS guard.
type LikeBuilder struct{ name string }

func (b WhereBuilder) binary(op string, value any) QueryExpr {
	return QueryExpr{Op: op, Column: b.name, Value: value}
}

func (b WhereBuilder) Eq(value any) QueryExpr  { return b.binary("eq", value) }
func (b WhereBuilder) Ne(value any) QueryExpr  { return b.binary("ne", value) }
func (b WhereBuilder) Gt(value any) QueryExpr  { return b.binary("gt", value) }
func (b WhereBuilder) Gte(value any) QueryExpr { return b.binary("gte", value) }
func (b WhereBuilder) Lt(value any) QueryExpr  { return b.binary("lt", value) }
func (b WhereBuilder) Lte(value any) QueryExpr { return b.binary("lte", value) }

// In builds an IN predicate. The comma guard fires in the where-serializer
// (mirrors node, where stringified IN values containing commas are rejected).
func (b WhereBuilder) In(values ...any) QueryExpr {
	return QueryExpr{Op: "in", Column: b.name, Values: values}
}

func (l LikeBuilder) Contains(value string) QueryExpr {
	return QueryExpr{Op: "like", Column: l.name, Value: "%" + EscapeLike(value) + "%"}
}

func (l LikeBuilder) StartsWith(value string) QueryExpr {
	return QueryExpr{Op: "like", Column: l.name, Value: EscapeLike(value) + "%"}
}

func (l LikeBuilder) EndsWith(value string) QueryExpr {
	return QueryExpr{Op: "like", Column: l.name, Value: "%" + EscapeLike(value)}
}

// Raw is the trusted-SQL LIKE escape hatch. The ReDoS guard runs here; on
// failure the error is deferred onto the expr and surfaces from List/Count.
func (l LikeBuilder) Raw(value string) QueryExpr {
	if err := AssertSafeLikePattern(value); err != nil {
		return QueryExpr{Op: "like", Column: l.name, Value: value, BuildErr: err}
	}
	return QueryExpr{Op: "like", Column: l.name, Value: value}
}

// Where starts a predicate for a column. Accepts a column name (string) or a
// DataColumn from a defined schema (its Name is used).
func Where(column any) WhereBuilder {
	var name string
	switch c := column.(type) {
	case string:
		name = c
	case DataColumn:
		name = c.Name
	case *DataColumn:
		name = c.Name
	default:
		name = ""
	}
	return WhereBuilder{name: name, Like: LikeBuilder{name: name}}
}
