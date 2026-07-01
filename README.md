# AX Hub Go SDK

AX Hub Go SDK for `https://api.axhub.ai`. It gives agents one small client, generated backend route metadata, bounded-context route facades, typed error metadata, conformance tests, and a live-testable app/data workflow.

## Install

```bash
go get github.com/jocoding-ax-partners/axhub-sdk-go@v0.5.0
```

## Required environment for agent work

```bash
export AXHUB_TOKEN="<short-lived PAT>"
export AXHUB_TENANT_ID="cc1e58f1-8e46-4ac7-96c1-190c4cdd5b70"   # test tenant
export AXHUB_TENANT_SLUG="test"
```

PAT mode is explicit: `TokenTypePAT` sends `X-Api-Key`. JWT mode is `TokenTypeJWT` and sends `Authorization: Bearer`.

## Agent quickstart: create a disposable app and enable its database

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	axhub "github.com/jocoding-ax-partners/axhub-sdk-go"
)

func main() {
	ctx := context.Background()
	tenantID := os.Getenv("AXHUB_TENANT_ID")
	tenantSlug := os.Getenv("AXHUB_TENANT_SLUG")
	if tenantID == "" || tenantSlug == "" {
		log.Fatal("AXHUB_TENANT_ID and AXHUB_TENANT_SLUG are required")
	}

	client := axhub.NewClient(axhub.Config{
		BaseURL:           "https://api.axhub.ai",
		Token:             os.Getenv("AXHUB_TOKEN"),
		TokenType:         axhub.TokenTypePAT,
		DefaultTenantID:   tenantID,
		DefaultTenantSlug: tenantSlug,
	})

	me, err := client.Request(ctx, "authGetApiV1Me", nil, nil, nil)
	if err != nil { log.Fatal(err) }
	userID, _ := me["userId"].(string)
	if userID == "" {
		if user, ok := me["user"].(map[string]any); ok {
			userID, _ = user["id"].(string)
		}
	}
	if userID == "" { log.Fatal("authGetApiV1Me did not return a user id") }

	suffix := fmt.Sprintf("%d", time.Now().UnixMilli())
	slug := "agent-go-" + suffix[len(suffix)-8:]

	app, err := client.Apps.Create(ctx, map[string]any{
		"slug": slug, "name": "Agent Go README QA", "visibility": "private",
		"auth_mode": "anonymous", "resource_preset": "S", "deploy_method": "docker", "subdomain": slug,
	})
	if err != nil { log.Fatal(err) }
	appID := app["id"].(string)

	// Enable raw DB mode: a dedicated Postgres role is issued and DATABASE_URL is
	// injected into the app on its next deploy. The app then does row CRUD over
	// direct SQL with its own pg driver.
	_, err = client.Request(ctx, "appsPostApiV1AppsByAppIDRawDb",
		map[string]string{"appID": appID}, nil, map[string]any{})
	if err != nil { log.Fatal(err) }

	// Admin introspection/browse of the physical DB.
	tables, err := client.Request(ctx, "schemaGetApiV1AppsByAppIDDbTables",
		map[string]string{"appID": appID}, nil, nil)
	if err != nil { log.Fatal(err) }

	fmt.Println("created", appID, tables)
}
```

## How to call the full API surface

- High-level app create: `client.Apps.Create(ctx, body)` uses `defaultTenantID`.
- Any route by operation id: `client.Request(ctx, operationID, pathParams, query, body)`.
- Generated facade: `client.Data().SchemaGetApiV1AppsByAppIDDbTables(ctx, axhub.OperationParams{...})`.
- Route inventory: `axhub.Routes`, `axhub.ContextRoutes`, and `axhub.ErrorCodes`.
- Errors: use `var axErr *axhub.AxHubError; errors.As(err, &axErr)` and branch on `Code`, `Category`, `Status`, and `Retryable`.

## Dynamic app, schema, and data operations

Use the high-level `apps.create` helper for the first app, then use generated operation IDs for every backend route. Request bodies use backend wire keys, usually `snake_case`. Responses are normalized to camelCase in this SDK family, so read `tableName`, `requestId`, `revokedAt`, and similar keys from responses.

| Task | Operation ID | Required path params | Success assertion |
|------|--------------|----------------------|-------------------|
| Create env var | `appsPostApiV1AppsByAppIDEnvVars` | `appID` | `env.list` includes `key` |
| Delete env var | `appsDeleteApiV1AppsByAppIDEnvVarsByKey` | `appID`, `key` | `env.list` no longer includes `key` |
| Enable raw DB | `appsPostApiV1AppsByAppIDRawDb` | `appID` | dedicated Postgres role issued; `DATABASE_URL` injected on next deploy |
| Disable raw DB | `appsDeleteApiV1AppsByAppIDRawDb` | `appID` | raw DB mode turned off (data preserved) |
| List DB tables | `schemaGetApiV1AppsByAppIDDbTables` | `appID` | response lists physical `information_schema` tables |
| Browse DB rows | `schemaGetApiV1AppsByAppIDDbTablesByTableRows` | `appID`, `table` | response has `rows` array |
| Delete app | `appsDeleteApiV1AppsByAppID`, then `appsDeleteApiV1AppsByAppIDPermanent` | `appID` | app is soft-deleted, then permanently deleted |

Important semantics from live QA:

- App database access is raw: `appsPostApiV1AppsByAppIDRawDb` issues a dedicated Postgres role and injects `DATABASE_URL` on the next deploy (the connection string is never returned). The app does its own row CRUD via direct SQL; the SDK exposes admin introspection/browse only (`schemaGetApiV1AppsByAppIDDbTables` / `...DbTablesByTableRows`).
- Deployment creation without a connected git/bootstrap source can return a precondition-style 4xx. That verifies SDK error handling, not a deploy bug.


## Live QA evidence agents can trust

The SDK behavior documented here reflects live production QA against the AX Hub `test` tenant on 2026-06-08.

- Tenant used for destructive QA: slug `test`, id `cc1e58f1-8e46-4ac7-96c1-190c4cdd5b70`.
- Go, Java, Kotlin, Python, and Ruby each ran the generated all-operation sweep against 189 backend routes: SDK exceptions `0`, backend 5xx `0`.
- Go, Java, Kotlin, Python, and Ruby each passed strict destructive DB QA: the flow exercised the raw-DB enable/reset lifecycle (instead of the removed dynamic-table flow), then deleted the app and re-read to prove deletion semantics.
- Node ran the full production mutation suite and a real app bootstrap/deploy wait. Deployment id `d3a48ce3-0f9c-4bab-aa07-863c31c44460` finished `succeeded`, then the app was deleted permanently.

Do not print tokens. Use short-lived PATs for agent QA and revoke them after the run.


## Verification commands

Use local tests for every docs/code change. Run live tests only when you intentionally want destructive QA against `test`.


```bash
go test ./... -count=1

# Destructive live all-operation sweep, only with a disposable PAT.
AXHUB_LIVE_ALL_METHODS=1 \
AXHUB_TOKEN="$AXHUB_TOKEN" \
AXHUB_LIVE_TENANT_ID="$AXHUB_TENANT_ID" \
AXHUB_LIVE_TENANT_SLUG="$AXHUB_TENANT_SLUG" \
go test ./... -run TestLiveAllGeneratedOperationFacadesHitProd -count=1 -v
```

## Troubleshooting for agents

- `tenant_id_required`: pass `defaultTenantId` / `AXHUB_TENANT_ID` before calling `apps.create`.
- `tokenType must be explicit`: set PAT mode when using a PAT. PATs are sent as `X-Api-Key`; JWTs are sent as `Authorization: Bearer`.
- `slug_taken` or `schema_name_taken`: append a timestamp suffix and retry. Never reuse fixture names in live destructive QA.
- `permission_denied` / `not_admin`: the SDK is working. The token lacks the role for that route.
- `precondition_failed` on deploy: connect git or use the app bootstrap flow first.
- 4xx responses are expected for negative assertions. SDK bugs are unexpected exceptions, response decode failures, or backend 5xx during a valid call.


## Release

See `RELEASE.md` for tag order, environment approvals, registry prerequisites, and smoke-test handling.

## License

Apache-2.0.
