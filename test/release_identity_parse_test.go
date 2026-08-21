package integration

import (
	"strings"
	"testing"
)

func TestVersionFromUserAgent(t *testing.T) {
	tests := []struct {
		name, ua, want string
	}{
		{"release", "wistia-cli/2026.5.1 (darwin/arm64)", "2026.5.1"},
		{"with CI suffix", "wistia-cli/2026.5.1 (linux/amd64) ci", "2026.5.1"},
		{"dev build", "wistia-cli/dev (linux/amd64)", "dev"},
		{"two-digit patch", "wistia-cli/2026.5.10 (darwin/arm64)", "2026.5.10"},
		{"speakeasy default", "speakeasy-sdk/go 0.0.1 2.911.0 edge-version github.com/wistia/wistia-cli/internal/sdk", ""},
		{"empty", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := versionFromUserAgent(tt.ua); got != tt.want {
				t.Errorf("versionFromUserAgent(%q) = %q, want %q", tt.ua, got, tt.want)
			}
		})
	}
}

func TestVersionFromUserAgentRejectsPrefixMatch(t *testing.T) {
	const ua = "wistia-cli/2026.5.10 (darwin/arm64)"
	const expected = "2026.5.1"

	if !strings.Contains(ua, "wistia-cli/"+expected) {
		t.Fatal("precondition: this User-Agent should contain the expected version as a substring")
	}
	if got := versionFromUserAgent(ua); got == expected {
		t.Errorf("versionFromUserAgent(%q) = %q, want a mismatch with %q", ua, got, expected)
	}
}
