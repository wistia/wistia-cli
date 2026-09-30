package hooks

import (
	"net/http"
	"testing"

	"github.com/wistia/wistia-cli/internal/sdk/sdkinternal/config"
)

func TestAPIVersion(t *testing.T) {
	tests := []struct {
		name, ua, want string
	}{
		{"dated release", "speakeasy-sdk/go 0.2.0 2.943.0 2026.09.0 github.com/wistia/wistia-cli/internal/sdk", "2026-09"},
		{"patched spec", "speakeasy-sdk/go 0.2.0 2.943.0 2026.09.3 github.com/wistia/wistia-cli/internal/sdk", "2026-09"},
		{"edge", "speakeasy-sdk/go 0.2.0 2.943.0 edge-version github.com/wistia/wistia-cli/internal/sdk", ""},
		{"unpadded month", "speakeasy-sdk/go 0.2.0 2.943.0 2026.9.0 github.com/wistia/wistia-cli/internal/sdk", ""},
		{"too few fields", "speakeasy-sdk/go 0.2.0", ""},
		{"empty", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := apiVersion(tt.ua); got != tt.want {
				t.Errorf("apiVersion(%q) = %q, want %q", tt.ua, got, tt.want)
			}
		})
	}
}

func TestAPIVersionHook(t *testing.T) {
	tests := []struct {
		name, ua, want string
	}{
		{"dated release sets the header", "speakeasy-sdk/go 0.2.0 2.943.0 2026.09.0 github.com/wistia/wistia-cli/internal/sdk", "2026-09"},
		{"edge sends no header", "speakeasy-sdk/go 0.2.0 2.943.0 edge-version github.com/wistia/wistia-cli/internal/sdk", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "https://api.wistia.com/modern/medias", nil)
			if err != nil {
				t.Fatal(err)
			}
			hookCtx := BeforeRequestContext{HookContext{SDKConfiguration: config.SDKConfiguration{UserAgent: tt.ua}}}

			got, err := apiVersionHook{}.BeforeRequest(hookCtx, req)
			if err != nil {
				t.Fatalf("BeforeRequest returned error: %v", err)
			}
			if _, present := got.Header["X-Wistia-Api-Version"]; present != (tt.want != "") {
				t.Errorf("X-Wistia-API-Version present = %v, want %v", present, tt.want != "")
			}
			if version := got.Header.Get("X-Wistia-API-Version"); version != tt.want {
				t.Errorf("X-Wistia-API-Version = %q, want %q", version, tt.want)
			}
		})
	}
}
