package axhub

// Optional schema validation hook (mirrors node dsl/zod.ts / python validation.py).
//
// The SDK duck-types a validator via the Validator interface so the validation
// library stays optional. On "update", a PartialValidator (if implemented) is
// used so partial patches are accepted.

// ValidationIssue is one field-level validation failure.
type ValidationIssue struct {
	Path    []any
	Code    string
	Message string
}

// ValidationResult is the duck-typed SafeParse result.
type ValidationResult struct {
	Success bool
	Issues  []ValidationIssue
}

// PartialValidator is optionally implemented by validators to relax required
// fields on update (mirrors zod .partial()).
type PartialValidator interface {
	Partial() Validator
}

// runSchemaValidation validates data against schema.Validate before any network
// request. mode is "insert" or "update" (update uses Partial() when available).
func runSchemaValidation(schema *DataTableSchema, data any, mode string) *AxHubError {
	if schema == nil || schema.Validate == nil {
		return nil
	}
	validator := schema.Validate
	if mode == "update" {
		if pv, ok := validator.(PartialValidator); ok {
			validator = pv.Partial()
		}
	}
	result := validator.SafeParse(data)
	if result.Success {
		return nil
	}
	return newValidationError("validation failure before network request", "validation_failed")
}
