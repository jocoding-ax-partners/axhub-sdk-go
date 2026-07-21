# Changelog

## v0.13.0 — 2026-07-21

### Added
- 라우트 87→97 — notifications(사용자 알림함 3·앱 발송 2)·access-requests 4·배포별 진단 1 op facade 추가 (backend 8714f5cd re-pin, allowlist 87→97, 신규 notifications context).
- 에러 카탈로그 117→133 — `feature_disabled`·`promote_snapshot_missing`·`rate_limited`·`staging_namespace_too_long` 등 (backend spec 128/133/134 + PR #618).

### Fixed
- 에러 카탈로그 category 오염 교정 — payment 계열 10종의 category 가 Go 식별자 `CategoryPaymentRequired` 문자열로 잘못 들어가던 것을 wire 값 `payment_required` 로 복구.

### Changed
- conformance corpus 45→49 — typed-error 3종(feature_disabled·promote_snapshot_missing·402 payment) + check-availability `reason:"invalid"` 회귀 벡터.

## v0.12.0 — 2026-07-08

### Added
- 에러 카탈로그 106→117 — Google Workspace 조직도 연동(spec 109/113)·디렉터리 단일소스(spec 114) 에러 코드 11종 추가 (`google_domain_taken`, `directory_source_conflict`, `group_scim_managed` 등; backend 73e89024 re-pin).

## v0.11.0 — 2026-07-08

### Added
- `apps.RawDb.Tables` / `apps.RawDb.TableRows` — raw DB 조회용 typed helper (read-only introspection). node SDK `sdk.apps.rawDb.*` 와 표면 정렬.

### Changed
- live QA harness 를 유지 표면에 정렬 (미노출 op 참조 제거).
## v0.10.1 — 2026-07-08

### Changed
- 테스트: live QA harness 를 현행 route surface 에 정렬 (기능 변화 없음).

## v0.10.0 — 2026-07-03

### Added
- `TenantsGetApiV1TenantsByTenantIDMembersDirectory` — `GET /api/v1/tenants/{tenantID}/members/directory`. 활성 멤버 디렉토리(가입 멤버만, email 제외 PII 축소 뷰). 각 행의 `groupId` 로 그룹별 필터가 가능하다. 호출 권한 tenant_member. node 에만 있던 표면을 6개 SDK 로 정렬.

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