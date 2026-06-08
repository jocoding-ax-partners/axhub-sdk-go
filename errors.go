package axhub

type ErrorInfo struct {
	Category  string
	Status    int
	Retryable bool
}

type AxHubError struct {
	Category, Code, Message, RequestID string
	Status                             int
	Retryable                          bool
}

func (e *AxHubError) Error() string {
	if e.Message != "" {
		return e.Code + ": " + e.Message
	}
	return e.Code
}

var ErrorCodes = map[string]ErrorInfo{
	"already_accessed":         {Category: "conflict", Status: 409, Retryable: false},
	"already_active":           {Category: "conflict", Status: 409, Retryable: false},
	"already_deleted":          {Category: "conflict", Status: 409, Retryable: false},
	"already_exists":           {Category: "conflict", Status: 409, Retryable: false},
	"already_inactive":         {Category: "conflict", Status: 409, Retryable: false},
	"already_member":           {Category: "conflict", Status: 409, Retryable: false},
	"already_revoked":          {Category: "conflict", Status: 409, Retryable: false},
	"already_settled":          {Category: "conflict", Status: 409, Retryable: false},
	"already_suspended":        {Category: "conflict", Status: 409, Retryable: false},
	"app_unavailable":          {Category: "conflict", Status: 409, Retryable: false},
	"bad_request":              {Category: "validation", Status: 400, Retryable: false},
	"cannot_reactivate":        {Category: "conflict", Status: 409, Retryable: false},
	"conflict":                 {Category: "conflict", Status: 409, Retryable: false},
	"cross_tenant":             {Category: "validation", Status: 400, Retryable: false},
	"domain_blocked":           {Category: "precondition_failed", Status: 422, Retryable: false},
	"domain_taken":             {Category: "conflict", Status: 409, Retryable: false},
	"duplicate":                {Category: "validation", Status: 400, Retryable: false},
	"empty":                    {Category: "validation", Status: 400, Retryable: false},
	"forbidden":                {Category: "permission_denied", Status: 403, Retryable: false},
	"internal_error":           {Category: "internal", Status: 500, Retryable: false},
	"invalid_expiry":           {Category: "validation", Status: 400, Retryable: false},
	"invalid_format":           {Category: "validation", Status: 400, Retryable: false},
	"invalid_state_transition": {Category: "conflict", Status: 409, Retryable: false},
	"invalid_value":            {Category: "validation", Status: 400, Retryable: false},
	"invitation_expired":       {Category: "not_found", Status: 410, Retryable: false},
	"last_admin":               {Category: "conflict", Status: 409, Retryable: false},
	"not_admin":                {Category: "permission_denied", Status: 403, Retryable: false},
	"not_allowed":              {Category: "validation", Status: 400, Retryable: false},
	"not_deleted":              {Category: "conflict", Status: 409, Retryable: false},
	"not_found":                {Category: "not_found", Status: 404, Retryable: false},
	"not_member":               {Category: "permission_denied", Status: 403, Retryable: false},
	"not_suspended":            {Category: "conflict", Status: 409, Retryable: false},
	"pending_exists":           {Category: "conflict", Status: 409, Retryable: false},
	"permanently_deleted":      {Category: "not_found", Status: 410, Retryable: false},
	"precondition_failed":      {Category: "precondition_failed", Status: 412, Retryable: false},
	"required":                 {Category: "validation", Status: 400, Retryable: false},
	"schema_name_taken":        {Category: "conflict", Status: 409, Retryable: false},
	"slug_taken":               {Category: "conflict", Status: 409, Retryable: false},
	"temporarily_unavailable":  {Category: "unavailable", Status: 429, Retryable: true},
	"token_expired":            {Category: "unauthenticated", Status: 401, Retryable: true},
	"token_invalid":            {Category: "unauthenticated", Status: 401, Retryable: true},
	"token_missing":            {Category: "unauthenticated", Status: 401, Retryable: true},
	"too_long":                 {Category: "validation", Status: 400, Retryable: false},
}
