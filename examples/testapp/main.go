package main

import (
	"context"
	"encoding/json"
	"fmt"
	axhub "github.com/axhub/axhub-sdk-go"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

func main() {
	if token := os.Getenv("AXHUB_TOKEN"); token != "" {
		runProd(token)
		return
	}
	runLocal()
}

func runProd(token string) {
	tokenType, err := tokenTypeFromEnv(os.Getenv("AXHUB_TOKEN_TYPE"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "go prod test app failed: %v\n", err)
		os.Exit(1)
	}
	baseURL := os.Getenv("AXHUB_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.axhub.ai"
	}
	c := axhub.NewClient(axhub.Config{BaseURL: baseURL, Token: token, TokenType: tokenType, DefaultTenantID: os.Getenv("AXHUB_TENANT_ID")})
	got, err := c.Identity().AuthGetApiV1Me(context.Background(), axhub.OperationParams{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "go prod test app failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("go prod test app ok %s keys=%d\n", c.BaseURL(), len(got))
}

func runLocal() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"id": "app_demo", "tenant_id": "tnt_demo", "slug": "demo", "schema_name": "app_demo"})
	}))
	defer srv.Close()
	c := axhub.NewClient(axhub.Config{BaseURL: srv.URL, Token: "pat_demo", TokenType: axhub.TokenTypePAT, DefaultTenantID: "tnt_demo"})
	got, err := c.Apps.Create(context.Background(), map[string]any{"slug": "demo", "name": "Demo"})
	if err != nil {
		fmt.Fprintf(os.Stderr, "go test app failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("go test app ok %s %s\n", got["id"], c.BaseURL())
}

func tokenTypeFromEnv(value string) (axhub.TokenType, error) {
	switch strings.ToLower(value) {
	case "pat":
		return axhub.TokenTypePAT, nil
	case "jwt":
		return axhub.TokenTypeJWT, nil
	default:
		return "", fmt.Errorf("AXHUB_TOKEN_TYPE must be pat or jwt")
	}
}
