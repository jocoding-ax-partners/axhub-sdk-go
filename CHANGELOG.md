# Changelog

## v0.7.0 — 2026-07-02

### Changed
- OAuth token-style responses (`authPostOauthToken`, `authPostOauthDeviceAuthorization`) now preserve RFC 6749/8628 standard keys in snake_case (`access_token`, `token_type`, `expires_in`, `refresh_token`, `id_token`, `scope`, `resource`, `tenant`) instead of camelizing them. Code reading `accessToken` from these responses must switch to `access_token`. Non-standard keys in the same responses are still camelized. Enforced by conformance vector expanded-004.
- No route surface changes (no operationId added, removed, or renamed).
