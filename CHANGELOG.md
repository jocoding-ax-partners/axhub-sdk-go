# Changelog

## v0.9.0 — 2026-07-03

### Added
- `TenantsGetApiV1TenantsByTenantIDOrgDirectory` — `GET /api/v1/tenants/{tenantID}/org-directory`. SCIM 사내 조직도 조회: 부서(그룹)→인원 구조, 아직 가입하지 않은 인원 포함(`joined` 플래그로 구분), 표시 이름 없으면 이메일 폴백. 호출 권한 tenant_member.

## v0.8.0 — 2026-07-02

### Removed (developer-surface reduction)
The SDK now exposes only the developer surface (85 operations). 132 operations were removed:
platform admin (templates admin, revoke-all, internal, SCIM, webhooks), tenant/org management
(members, seats, groups, invitations, invite-links, icons, categories, email-domains, identity
providers, tenant create/update/delete), authorization management (grant issuance, presets,
subjects), audit, cost (context removed), review-requests, static-site releases/staging, and all
browser auth / OAuth flows (`/auth/*`, `/oauth/*`) including `authPostOauthToken`.
Auth model: inject an issued PAT or JWT; token exchange happens outside the SDK.
New backend routes are excluded by default until added to `conformance/sdk-allowlist.json`.

### Changed
- Bounded contexts: 9 -> 8 (`cost` removed; `audit` remains registered with no operations). Requests are always JSON (form-encoding path removed).

## v0.7.0 — 2026-07-02

### Changed
- OAuth token-style responses (`authPostOauthToken`, `authPostOauthDeviceAuthorization`) now preserve RFC 6749/8628 standard keys in snake_case (`access_token`, `token_type`, `expires_in`, `refresh_token`, `id_token`, `scope`, `resource`, `tenant`) instead of camelizing them. Code reading `accessToken` from these responses must switch to `access_token`. Non-standard keys in the same responses are still camelized. Enforced by conformance vector expanded-004.
- No route surface changes (no operationId added, removed, or renamed).
