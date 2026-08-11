package useragent

import (
	"strings"
	"testing"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		name                          string
		version, goos, goarch, suffix string
		want                          string
	}{
		{"release", "2026.5.0", "darwin", "arm64", "", "wistia-cli/2026.5.0 (darwin/arm64)"},
		{"dev default", "dev", "linux", "amd64", "", "wistia-cli/dev (linux/amd64)"},
		{"with suffix", "2026.5.0", "linux", "amd64", "ci", "wistia-cli/2026.5.0 (linux/amd64) ci"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := format(tt.version, tt.goos, tt.goarch, tt.suffix); got != tt.want {
				t.Errorf("format() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestClientVersion(t *testing.T) {
	tests := []struct {
		name            string
		version, suffix string
		want            string
	}{
		{"release", "2026.5.0", "", "2026.5.0"},
		{"dev default", "dev", "", "dev"},
		{"suffix as build metadata", "2026.5.0", "ci", "2026.5.0+ci"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clientVersion(tt.version, tt.suffix); got != tt.want {
				t.Errorf("clientVersion() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestClientVersionReadsAndTrimsEnvSuffix(t *testing.T) {
	t.Setenv(suffixEnvVar, "  ci  ")
	if got, want := ClientVersion(), Version+"+ci"; got != want {
		t.Errorf("ClientVersion() = %q, want %q", got, want)
	}
}

func TestStringReadsAndTrimsEnvSuffix(t *testing.T) {
	t.Setenv(suffixEnvVar, "  ci  ")
	got := String()
	if !strings.HasPrefix(got, "wistia-cli/") {
		t.Errorf("String() = %q, want prefix %q", got, "wistia-cli/")
	}
	if !strings.HasSuffix(got, " ci") {
		t.Errorf("String() = %q, want suffix %q (trimmed)", got, " ci")
	}
}
