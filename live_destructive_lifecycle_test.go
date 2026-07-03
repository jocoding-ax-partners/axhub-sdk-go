package axhub

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
)

// TestLiveDestructiveLifecycleHitProd drives a full member-surface DESTRUCTIVE
// lifecycle against live prod (create/update/delete real resources), going
// beyond the read-oriented facade sweep. Opt-in + disposable fixtures + LIFO
// cleanup + prefix "sdke2e-go-" so the orchestrator can orphan-sweep.
//
// Canonical sequence — later translated to java/kotlin/python/ruby. Mirrors
// node prod-sdk-full-mutation.e2e.test.ts and ADDS rows + raw-db (node gaps).
// SUCCESS on member-mutable ops; TYPED-FAILURE where a precondition is absent.
// Admin-sdk-scoped ops (tenant/members/categories/authz/audit/groups/scim/
// connectors/static) are intentionally out of scope (ADR-0043; covered as 403
// by the read sweep).
func TestLiveDestructiveLifecycleHitProd(t *testing.T) {
	if os.Getenv("AXHUB_LIVE_DESTRUCTIVE") != "1" {
		t.Skip("destructive live prod lifecycle is opt-in (set AXHUB_LIVE_DESTRUCTIVE=1)")
	}
	token := os.Getenv("AXHUB_TOKEN")
	if token == "" {
		t.Skip("AXHUB_TOKEN required")
	}
	base := dlEnv("AXHUB_LIVE_BASE_URL", "https://api.axhub.ai")
	tenantID := dlEnv("AXHUB_LIVE_TENANT_ID", "cc1e58f1-8e46-4ac7-96c1-190c4cdd5b70")

	c := NewClient(Config{BaseURL: base, Token: token, TokenType: TokenTypePAT, DefaultTenantID: tenantID})
	ctx, cancel := context.WithTimeout(context.Background(), 240*time.Second)
	defer cancel()

	suffix := fmt.Sprintf("%d", time.Now().UnixNano())
	appSlug := "sdke2e-go-" + suffix

	var cleanups []func()
	defer func() {
		for i := len(cleanups) - 1; i >= 0; i-- {
			cleanups[i]()
		}
	}()
	addCleanup := func(f func()) { cleanups = append(cleanups, f) }

	// must: member-mutable op that MUST succeed.
	must := func(label, opID string, pp map[string]string, body any) map[string]any {
		res, err := c.Request(ctx, opID, pp, nil, body)
		if err != nil {
			t.Fatalf("MUST %s (%s): %v", label, opID, err)
		}
		return res
	}
	// tolerate: op whose precondition may be unavailable — accept success OR a
	// typed AxHubError with an allowed 4xx status; fail on 5xx/transport/other.
	tolerate := func(label, opID string, pp map[string]string, body any, allowed ...int) {
		_, err := c.Request(ctx, opID, pp, nil, body)
		if err == nil {
			return
		}
		var ae *AxHubError
		if !errors.As(err, &ae) {
			t.Errorf("TOLERATE %s (%s): expected *AxHubError, got %T: %v", label, opID, err, err)
			return
		}
		if !dlContains(allowed, ae.Status) {
			t.Errorf("TOLERATE %s (%s): status %d not in %v (%s)", label, opID, ae.Status, allowed, ae.Code)
		}
	}
	// expectFail: precondition genuinely unavailable — MUST return a typed 4xx.
	expectFail := func(label, opID string, pp map[string]string, body any, allowed ...int) {
		_, err := c.Request(ctx, opID, pp, nil, body)
		if err == nil {
			t.Errorf("EXPECTFAIL %s (%s): expected typed failure, got success", label, opID)
			return
		}
		var ae *AxHubError
		if !errors.As(err, &ae) {
			t.Errorf("EXPECTFAIL %s (%s): expected *AxHubError, got %T", label, opID, err)
			return
		}
		if !dlContains(allowed, ae.Status) {
			t.Errorf("EXPECTFAIL %s (%s): status %d not in %v (%s)", label, opID, ae.Status, allowed, ae.Code)
		}
	}

	// --- app create (+ cleanup registered immediately) ---
	appRes := must("create app", "appsPostApiV1TenantsByTenantIDApps",
		map[string]string{"tenantID": tenantID},
		map[string]any{"slug": appSlug, "name": "SDK destructive E2E " + suffix, "description": "sdke2e disposable"})
	appID := dlStr(appRes, "id", "appId", "appID")
	if appID == "" {
		t.Fatalf("create app: no id in response %v", dlKeys(appRes))
	}
	addCleanup(func() {
		_, _ = c.Request(ctx, "appsDeleteApiV1AppsByAppID", map[string]string{"appID": appID}, nil, nil)
	})

	// --- app update ---
	must("update app", "appsPatchApiV1AppsByAppID", map[string]string{"appID": appID},
		map[string]any{"name": "SDK destructive E2E " + suffix + " renamed"})

	// --- env vars ---
	must("set env var", "appsPostApiV1AppsByAppIDEnvVars", map[string]string{"appID": appID},
		map[string]any{"key": "SDK_E2E_SECRET", "value": "sekret-" + suffix})

	// --- comments ---
	cRes := must("add comment", "appsPostApiV1AppsByAppIDComments", map[string]string{"appID": appID},
		map[string]any{"body": "sdke2e comment " + suffix})
	commentID := dlStr(cRes, "id", "commentId", "commentID")
	if commentID != "" {
		must("delete comment", "appsDeleteApiV1CommentsByCommentID", map[string]string{"commentID": commentID}, nil)
	}

	// --- likes (idempotent) ---
	must("like", "appsPostApiV1AppsByAppIDLikes", map[string]string{"appID": appID}, map[string]any{})
	must("unlike", "appsDeleteApiV1AppsByAppIDLikes", map[string]string{"appID": appID}, nil)

	// --- raw-db (node MISSES; body contract uncertain → tolerate both POST + DELETE; app is disposable) ---
	tolerate("raw-db exec", "appsPostApiV1AppsByAppIDRawDb", map[string]string{"appID": appID},
		map[string]any{"sql": "select 1"}, 400, 403, 404, 409, 422, 501)
	tolerate("raw-db reset", "appsDeleteApiV1AppsByAppIDRawDb", map[string]string{"appID": appID}, nil,
		400, 403, 404, 409, 422, 501)

	// --- oauth client (clientSecret surfaced once) ---
	ocRes := must("create oauth client", "authPostApiV1AppsByAppIDOauthClients", map[string]string{"appID": appID},
		map[string]any{
			"name":                       "SDK E2E OAuth " + suffix,
			"type":                       "confidential",
			"token_endpoint_auth_method": "client_secret_post",
			"redirect_uris":              []any{"https://example.com/callback"},
			"allowed_scopes":             []any{"read"},
			"allowed_grant_types":        []any{"authorization_code", "refresh_token"},
		})
	if dlStr(ocRes, "clientId", "client_id", "id") == "" || dlStr(ocRes, "clientSecret", "client_secret") == "" {
		t.Errorf("oauth client: missing clientId/clientSecret in %v", dlKeys(ocRes))
	}

	// --- PAT lifecycle (account-scoped: explicit revoke, survives app deletion) ---
	patRes := must("issue PAT", "schemaPostApiV1MePersonalAccessTokens", nil,
		map[string]any{"name": "SDK E2E " + suffix, "expires_in_days": 1})
	patID := dlStr(patRes, "id", "patId", "patID")
	if patID != "" {
		addCleanup(func() {
			_, _ = c.Request(ctx, "schemaDeleteApiV1MePersonalAccessTokensByPatID", map[string]string{"patID": patID}, nil, nil)
		})
		if dlStr(patRes, "rawToken", "raw_token") == "" {
			t.Errorf("issue PAT: missing rawToken in %v", dlKeys(patRes))
		}
		must("revoke PAT", "schemaDeleteApiV1MePersonalAccessTokensByPatID", map[string]string{"patID": patID}, nil)
	}

	// --- TYPED-FAILURE: preconditions genuinely unavailable ---
	expectFail("deployment create (no commit)", "deployPostApiV1AppsByAppIDDeployments", map[string]string{"appID": appID},
		map[string]any{"commit_sha": "a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0"}, 400, 404, 409, 412)
	expectFail("git connect (no install)", "deployPostApiV1AppsByAppIDGitGithubConnect", map[string]string{"appID": appID},
		map[string]any{"repo_full_name": "jocoding/sdke2e-nonexistent", "branch": "main", "installation_id": 0}, 400, 403, 404, 409)
	tolerate("access grant (self)", "appsPostApiV1AppsByAppIDAccess", map[string]string{"appID": appID}, map[string]any{}, 400, 403, 409)

	// --- explicit teardown (cleanup stack also covers on failure) ---
	must("delete app", "appsDeleteApiV1AppsByAppID", map[string]string{"appID": appID}, nil)
}

func dlEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func dlContains(xs []int, x int) bool {
	for _, v := range xs {
		if v == x {
			return true
		}
	}
	return false
}

func dlStr(m map[string]any, keys ...string) string {
	for _, k := range keys {
		if v, ok := m[k]; ok {
			if s, ok := v.(string); ok && s != "" {
				return s
			}
		}
	}
	return ""
}

func dlKeys(m map[string]any) []string {
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
