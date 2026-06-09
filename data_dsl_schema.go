package axhub

// Schema definitions and DefineSchema (mirrors node dsl/schema.ts).
//
// Column defs are a primitive type string ("uuid" | "string" | "number" |
// "integer" | "boolean" | "timestamp" | "json") or an enum descriptor
// (EnumColumn). DefineSchema builds the Cols accessor map used by the typed
// Where(schema.Cols["x"]) path.

// EnumColumn is a column whose def is one of a fixed set of string values.
type EnumColumn struct {
	Values []string
}

// ColumnDef is either a primitive type string or an EnumColumn.
type ColumnDef any

// SchemaShape maps column name -> ColumnDef.
type SchemaShape = map[string]ColumnDef

// DataColumn is a single column handle (used by the typed Where path).
type DataColumn struct {
	Table string
	Name  string
	Def   ColumnDef
}

// Validator is the duck-typed validation hook (see data_dsl_validation.go).
type Validator interface {
	SafeParse(data any) ValidationResult
}

// DataTableSchema is a defined table schema with a column accessor map.
type DataTableSchema struct {
	Table    string
	Columns  SchemaShape
	Cols     map[string]DataColumn
	Validate Validator
}

// DefineSchema builds a DataTableSchema from a table name + column shape.
// Mirrors node defineSchema; the optional validator is attached for the
// insert/update pre-flight validation hook.
func DefineSchema(table string, columns SchemaShape, validate ...Validator) *DataTableSchema {
	shape := make(SchemaShape, len(columns))
	cols := make(map[string]DataColumn, len(columns))
	for name, def := range columns {
		shape[name] = def
		cols[name] = DataColumn{Table: table, Name: name, Def: def}
	}
	s := &DataTableSchema{Table: table, Columns: shape, Cols: cols}
	if len(validate) > 0 {
		s.Validate = validate[0]
	}
	return s
}
