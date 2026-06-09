package axhub

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
