package axhub

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type vector struct {
	Name   string `json:"name"`
	Client struct {
		TokenType         string `json:"tokenType"`
		Token             string `json:"token"`
		DefaultTenantID   string `json:"defaultTenantId"`
		DefaultTenantSlug string `json:"defaultTenantSlug"`
	} `json:"client"`
	Call struct {
		Symbol     string            `json:"symbol"`
		OperationID string            `json:"operationId"`
		Args       map[string]any     `json:"args"`
		PathParams map[string]string  `json:"pathParams"`
		Query      map[string]string  `json:"query"`
		Body       any               `json:"body"`
	} `json:"call"`
	HTTPExpect *struct {
		Method, Path   string
		HeadersInclude []string          `json:"headersInclude"`
		HeadersExact   map[string]string `json:"headersExact"`
	} `json:"httpExpect"`
	MockResponse *struct {
		Status int
		Body   any
	} `json:"mockResponse"`
	Expect map[string]map[string]any `json:"expect"`
}

func TestConformanceVectors(t *testing.T) {
	files := conformanceFiles(t)
	for _, file := range files {
		t.Run(filepath.Base(file), func(t *testing.T) {
			raw, _ := os.ReadFile(file)
			var v vector
			if err := json.Unmarshal(raw, &v); err != nil {
				t.Fatal(err)
			}
			var capturedMethod, capturedPath string
			capturedHeaders := map[string]string{}
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				capturedMethod, capturedPath = r.Method, r.URL.Path
				for k, vals := range r.Header {
					if len(vals) > 0 {
						capturedHeaders[strings.ToLower(k)] = vals[0]
					}
				}
				status := 200
				body := any(map[string]any{})
				if v.MockResponse != nil {
					status = v.MockResponse.Status
					body = v.MockResponse.Body
				}
				w.WriteHeader(status)
				_ = json.NewEncoder(w).Encode(body)
			}))
			defer srv.Close()
			c := NewClient(Config{BaseURL: srv.URL, Token: v.Client.Token, TokenType: TokenType(v.Client.TokenType), DefaultTenantID: v.Client.DefaultTenantID, DefaultTenantSlug: v.Client.DefaultTenantSlug})
			got, err := dispatchVector(context.Background(), c, v)
			if want, ok := v.Expect["error"]; ok {
				if err == nil {
					t.Fatalf("expected error")
				}
				ax, ok := err.(*AxHubError)
				if !ok {
					t.Fatalf("expected AxHubError got %T", err)
				}
				if ax.Category != wantString(want, "category") || ax.Code != wantString(want, "code") {
					t.Fatalf("wrong error %#v want %#v", ax, want)
				}
				if requestID := wantString(want, "requestId"); requestID != "" && ax.RequestID != requestID {
					t.Fatalf("request id got %s want %s", ax.RequestID, requestID)
				}
				if retryable, ok := want["retryable"].(bool); ok && ax.Retryable != retryable {
					t.Fatalf("retryable got %v want %v", ax.Retryable, retryable)
				}
				if v.HTTPExpect == nil && capturedMethod != "" {
					t.Fatalf("expected no request")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error %v", err)
				}
				for k, want := range v.Expect["ok"] {
					if got[k] != want {
						t.Fatalf("%s got %#v want %#v in %#v", k, got[k], want, got)
					}
				}
			}
			if v.HTTPExpect != nil {
				if capturedMethod != v.HTTPExpect.Method || capturedPath != v.HTTPExpect.Path {
					t.Fatalf("request got %s %s want %s %s", capturedMethod, capturedPath, v.HTTPExpect.Method, v.HTTPExpect.Path)
				}
				for _, h := range v.HTTPExpect.HeadersInclude {
					if capturedHeaders[strings.ToLower(h)] == "" {
						t.Fatalf("missing header %s", h)
					}
				}
				for h, want := range v.HTTPExpect.HeadersExact {
					if got := capturedHeaders[strings.ToLower(h)]; got != want {
						t.Fatalf("header %s got %q want %q", h, got, want)
					}
				}
			}
		})
	}
}

func dispatchVector(ctx context.Context, c *Client, v vector) (map[string]any, error) {
	switch v.Call.Symbol {
	case "sdk.apps.create":
		return c.Apps.Create(ctx, v.Call.Args)
	case "sdk.operation":
		return c.Operation(ctx, v.Call.OperationID, OperationParams{PathParams: v.Call.PathParams, Query: v.Call.Query, Body: v.Call.Body})
	case "sdk.redactedToken":
		return map[string]any{"redactedToken": c.RedactedToken()}, nil
	default:
		return nil, &AxHubError{Category: "validation", Code: "unknown_vector_symbol", Message: v.Call.Symbol}
	}
}

func conformanceFiles(t *testing.T) []string {
	dirs := []string{}
	if d := os.Getenv("AXHUB_CONFORMANCE_DIR"); d != "" {
		dirs = append(dirs, d)
	}
	dirs = append(dirs, filepath.Join("testdata", "conformance", "vectors"))
	dirs = append(dirs, filepath.Join("..", "conformance", "vectors"))
	if spec := os.Getenv("AXHUB_SPEC_DIR"); spec != "" {
		dirs = append(dirs, filepath.Join(spec, "conformance", "vectors"))
	} else {
		dirs = append(dirs, filepath.Join("..", "..", "axhub-sdk-spec", "conformance", "vectors"))
	}
	for _, dir := range dirs {
		files, err := filepath.Glob(filepath.Join(dir, "*.json"))
		if err == nil && len(files) > 0 {
			return files
		}
	}
	t.Fatalf("vectors missing in %v", dirs)
	return nil
}

func wantString(m map[string]any, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}
