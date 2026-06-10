package axhub

import "errors"

// Typed data-layer error constructors. Go has no error subclassing, so these
// return *AxHubError with the exact (category, code) pairs the node/python
// ports use, letting callers match on Code/Category. Mirrors node's
// ValidationError / LegacyCursorError / InvalidCursorError / TableNotFoundError
// / IntrospectFailedError / ScanLimitExceededError hierarchy.

func newValidationError(message, code string) *AxHubError {
	return &AxHubError{Category: "validation", Code: code, Message: message}
}

func newLegacyCursorError(message string) *AxHubError {
	return &AxHubError{Category: "validation", Code: "legacy_cursor", Message: message}
}

func newInvalidCursorError(message string) *AxHubError {
	return &AxHubError{Category: "validation", Code: "invalid_cursor", Message: message}
}

func newTableNotFoundError(message, requestID string) *AxHubError {
	return &AxHubError{Category: "not_found", Code: "table_not_found", Status: 404, Message: message, RequestID: requestID}
}

func newIntrospectFailedError(message string, status int, retryable bool, requestID string) *AxHubError {
	return &AxHubError{Category: "internal", Code: "introspect_failed", Status: status, Message: message, Retryable: retryable, RequestID: requestID}
}

func newScanLimitExceededError(message string) *AxHubError {
	return &AxHubError{Category: "internal", Code: "scan_limit_exceeded", Message: message}
}

func newConfigurationError(message, code string) *AxHubError {
	return &AxHubError{Category: "configuration", Code: code, Message: message}
}

// mapWhereRequired normalizes the backend's where-required 400 into the same
// actionable ValidationError the old client-side pre-check threw. The backend
// 400s an unfiltered list/count on NON-owner-scoped tables ("최소 1개의 WHERE
// 필터가 필요해요") but ACCEPTS it on owner-scoped tables (rows auto-scope to
// the caller) — both confirmed live 2026-06. A pre-check cannot tell the two
// apart (0.3.0 regression), so the request goes through and only the 400 is
// rewritten.
func mapWhereRequired(op string, err error) error {
	var ae *AxHubError
	if errors.As(err, &ae) && ae.Code == "required" && ae.Status == 400 {
		return newValidationError("AxHub data "+op+" requires at least one WHERE filter on this table (the backend rejects unfiltered scans on non-owner-scoped tables). Pass a Where(...).", "where_required")
	}
	return err
}
