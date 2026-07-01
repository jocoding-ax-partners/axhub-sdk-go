package axhub

type ErrorInfo struct {
	Category  string
	Status    int
	Retryable bool
}

type FieldError struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
}

type RetryInfo struct {
	AfterMs int `json:"after_ms"`
}

type AxHubError struct {
	Category, Code, Message, RequestID string
	Status                             int
	Retryable                          bool
	Resource                           string
	Fields                             []FieldError
	Retry                              *RetryInfo
	DocURL                             string
}

func (e *AxHubError) Error() string {
	if e.Message != "" {
		return e.Code + ": " + e.Message
	}
	return e.Code
}

var ErrorCodes = map[string]ErrorInfo{
	"action_denied": {Category: "permission_denied", Status: 403, Retryable: false},
	"action_invalid": {Category: "validation", Status: 400, Retryable: false},
	"already_accessed": {Category: "conflict", Status: 409, Retryable: false},
	"already_active": {Category: "conflict", Status: 409, Retryable: false},
	"already_deleted": {Category: "conflict", Status: 409, Retryable: false},
	"already_exists": {Category: "conflict", Status: 409, Retryable: false},
	"already_inactive": {Category: "conflict", Status: 409, Retryable: false},
	"already_member": {Category: "conflict", Status: 409, Retryable: false},
	"already_promoted": {Category: "conflict", Status: 409, Retryable: false},
	"already_revoked": {Category: "conflict", Status: 409, Retryable: false},
	"already_settled": {Category: "conflict", Status: 409, Retryable: false},
	"already_suspended": {Category: "conflict", Status: 409, Retryable: false},
	"already_terminal": {Category: "conflict", Status: 409, Retryable: false},
	"app_unavailable": {Category: "conflict", Status: 409, Retryable: false},
	"auth_expired": {Category: "unavailable", Status: 503, Retryable: false},
	"bad_request": {Category: "validation", Status: 400, Retryable: false},
	"build_env_no_override": {Category: "validation", Status: 400, Retryable: false},
	"cannot_reactivate": {Category: "conflict", Status: 409, Retryable: false},
	"charge_failed": {Category: "CategoryPaymentRequired", Status: 402, Retryable: false},
	"confirm_required": {Category: "precondition_failed", Status: 412, Retryable: false},
	"conflict": {Category: "conflict", Status: 409, Retryable: false},
	"connector_inactive": {Category: "permission_denied", Status: 403, Retryable: false},
	"cross_tenant": {Category: "validation", Status: 400, Retryable: false},
	"domain_blocked": {Category: "precondition_failed", Status: 422, Retryable: false},
	"domain_taken": {Category: "conflict", Status: 409, Retryable: false},
	"duplicate": {Category: "validation", Status: 400, Retryable: false},
	"empty": {Category: "validation", Status: 400, Retryable: false},
	"exceeds_max": {Category: "conflict", Status: 409, Retryable: false},
	"expiry_in_past": {Category: "validation", Status: 400, Retryable: false},
	"feature_not_in_plan": {Category: "permission_denied", Status: 403, Retryable: false},
	"final_visibility_too_wide": {Category: "validation", Status: 400, Retryable: false},
	"forbidden": {Category: "permission_denied", Status: 403, Retryable: false},
	"github_device_flow_disabled": {Category: "unavailable", Status: 503, Retryable: false},
	"grant_already_terminal": {Category: "conflict", Status: 409, Retryable: false},
	"grant_conflict": {Category: "conflict", Status: 409, Retryable: false},
	"grant_expired": {Category: "permission_denied", Status: 403, Retryable: false},
	"grant_revoked": {Category: "permission_denied", Status: 403, Retryable: false},
	"internal_error": {Category: "internal", Status: 500, Retryable: false},
	"invalid_drive": {Category: "validation", Status: 400, Retryable: false},
	"invalid_entitlement": {Category: "validation", Status: 400, Retryable: false},
	"invalid_expiry": {Category: "validation", Status: 400, Retryable: false},
	"invalid_format": {Category: "validation", Status: 400, Retryable: false},
	"invalid_oauth_state": {Category: "validation", Status: 400, Retryable: false},
	"invalid_state_transition": {Category: "conflict", Status: 409, Retryable: false},
	"invalid_target": {Category: "conflict", Status: 409, Retryable: false},
	"invalid_value": {Category: "validation", Status: 400, Retryable: false},
	"invitation_expired": {Category: "not_found", Status: 410, Retryable: false},
	"kind_engine_mismatch": {Category: "validation", Status: 400, Retryable: false},
	"last_admin": {Category: "conflict", Status: 409, Retryable: false},
	"link_invalid": {Category: "not_found", Status: 404, Retryable: false},
	"no_active_grant": {Category: "not_found", Status: 404, Retryable: false},
	"no_available_seat": {Category: "conflict", Status: 409, Retryable: false},
	"no_billing_key": {Category: "CategoryPaymentRequired", Status: 402, Retryable: false},
	"no_payment_method": {Category: "CategoryPaymentRequired", Status: 402, Retryable: false},
	"not_admin": {Category: "permission_denied", Status: 403, Retryable: false},
	"not_allowed": {Category: "validation", Status: 400, Retryable: false},
	"not_deleted": {Category: "conflict", Status: 409, Retryable: false},
	"not_deployed": {Category: "conflict", Status: 409, Retryable: false},
	"not_found": {Category: "not_found", Status: 404, Retryable: false},
	"not_member": {Category: "permission_denied", Status: 403, Retryable: false},
	"not_promotable": {Category: "precondition_failed", Status: 412, Retryable: false},
	"not_suspended": {Category: "conflict", Status: 409, Retryable: false},
	"oauth_denied": {Category: "validation", Status: 400, Retryable: false},
	"payment_failed": {Category: "CategoryPaymentRequired", Status: 402, Retryable: false},
	"payment_required": {Category: "CategoryPaymentRequired", Status: 402, Retryable: false},
	"pending_exists": {Category: "conflict", Status: 409, Retryable: false},
	"pending_review_exists": {Category: "precondition_failed", Status: 412, Retryable: false},
	"permanently_deleted": {Category: "not_found", Status: 410, Retryable: false},
	"plan_version_exists": {Category: "conflict", Status: 409, Retryable: false},
	"precondition_failed": {Category: "precondition_failed", Status: 412, Retryable: false},
	"preset_in_use": {Category: "conflict", Status: 409, Retryable: false},
	"preset_mismatch": {Category: "validation", Status: 400, Retryable: false},
	"preset_not_in_plan": {Category: "CategoryPaymentRequired", Status: 402, Retryable: false},
	"prod_deploy_required": {Category: "precondition_failed", Status: 412, Retryable: false},
	"promote_in_progress": {Category: "precondition_failed", Status: 412, Retryable: true},
	"quota_exceeded": {Category: "CategoryPaymentRequired", Status: 402, Retryable: false},
	"required": {Category: "validation", Status: 400, Retryable: false},
	"resource_quota_exceeded": {Category: "CategoryPaymentRequired", Status: 402, Retryable: false},
	"schema_name_taken": {Category: "conflict", Status: 409, Retryable: false},
	"seat_in_use": {Category: "conflict", Status: 409, Retryable: false},
	"seat_unassigned": {Category: "CategoryPaymentRequired", Status: 402, Retryable: false},
	"seats_not_supported": {Category: "conflict", Status: 409, Retryable: false},
	"session_ended": {Category: "unauthenticated", Status: 401, Retryable: true},
	"session_expired": {Category: "unauthenticated", Status: 401, Retryable: true},
	"slug_taken": {Category: "conflict", Status: 409, Retryable: false},
	"staging_already_enabled": {Category: "conflict", Status: 409, Retryable: false},
	"staging_mismatch": {Category: "precondition_failed", Status: 412, Retryable: false},
	"staging_not_enabled": {Category: "precondition_failed", Status: 412, Retryable: false},
	"staging_required": {Category: "precondition_failed", Status: 412, Retryable: false},
	"static_release_in_use": {Category: "conflict", Status: 409, Retryable: false},
	"static_release_not_ready": {Category: "precondition_failed", Status: 412, Retryable: false},
	"temporarily_unavailable": {Category: "unavailable", Status: 429, Retryable: true},
	"token_expired": {Category: "unauthenticated", Status: 401, Retryable: true},
	"token_invalid": {Category: "unauthenticated", Status: 401, Retryable: true},
	"token_missing": {Category: "unauthenticated", Status: 401, Retryable: true},
	"too_long": {Category: "validation", Status: 400, Retryable: false},
	"unknown_plan": {Category: "not_found", Status: 404, Retryable: false},
	"unpaid_balance": {Category: "CategoryPaymentRequired", Status: 402, Retryable: false},
	"unsupported_for_static_app": {Category: "conflict", Status: 409, Retryable: false},
	"version_not_approved": {Category: "permission_denied", Status: 403, Retryable: false},
	"visibility_widen_not_allowed": {Category: "conflict", Status: 409, Retryable: false},
}
