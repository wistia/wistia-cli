package useragent

import (
	"regexp"
	"strings"
	"testing"
)

// Shared with the wistia/wistia resolver that validates X-Wistia-Client-Version,
// and with #37's Dependency section. Widening it on one side without the other
// silently drops version attribution, so both sides pin the same grammar.
var resolverContract = regexp.MustCompile(`^(?:\d+\.\d+\.\d+|dev)(\+[0-9A-Za-z.-]{1,16})?$`)

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
		{"future suffix needs no code change", "2026.5.0", "staging", "2026.5.0+staging"},
		{"space becomes a separator", "2026.5.0", "nightly build", "2026.5.0+nightly-build"},
		{"invalid characters collapse", "2026.5.0", "ci_run/7", "2026.5.0+ci-run-7"},
		{"dotted identifiers survive", "2026.5.0", "staging.2", "2026.5.0+staging.2"},
		{"empty identifiers dropped", "2026.5.0", ".staging..2.", "2026.5.0+staging.2"},
		{"leading and trailing junk trimmed", "2026.5.0", " %ci% ", "2026.5.0+ci"},
		{"nothing usable omits the plus", "2026.5.0", "!!!", "2026.5.0"},
		{"dev build with suffix", "dev", "ci", "dev+ci"},
		{"long suffix truncates", "2026.5.0", "nightly analytics pipeline", "2026.5.0+nightly-analytic"},
		{"exactly at the bound is kept", "2026.5.0", "abcdefghijklmnop", "2026.5.0+abcdefghijklmnop"},
		{"one over the bound truncates", "2026.5.0", "abcdefghijklmnopq", "2026.5.0+abcdefghijklmnop"},
		{"trailing dot from truncation dropped", "2026.5.0", "abcdefghijklmno.pqrs", "2026.5.0+abcdefghijklmno"},
		{"trailing hyphen from truncation kept", "2026.5.0", "abcdefghijklmno-pqrs", "2026.5.0+abcdefghijklmno-"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clientVersion(tt.version, tt.suffix); got != tt.want {
				t.Errorf("clientVersion() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestClientVersionMatchesResolverContract(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want bool
	}{
		{"source build", clientVersion("dev", ""), true},
		{"source build in CI", clientVersion("dev", "ci"), true},
		{"release in CI", clientVersion("2026.5.1", "ci"), true},
		{"release with truncated metadata", clientVersion("2026.5.1", "nightly analytics pipeline"), true},
		{"metadata one over the bound", "2026.5.1+" + strings.Repeat("a", maxBuildMetadata+1), false},
		{"prerelease version", "0.1.2-beta", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resolverContract.MatchString(tt.got); got != tt.want {
				t.Errorf("resolverContract.MatchString(%q) = %v, want %v", tt.got, got, tt.want)
			}
		})
	}
}

// Invalid metadata makes the server drop version attribution while keeping the client name.
func TestClientVersionEmitsValidBuildMetadata(t *testing.T) {
	suffixes := []string{
		"ci", "staging", "nightly build", "ci_run/7", "  ci  ", "a.b.c",
		"CI-2026", "héllo", "ci\nrun", "!!!", "..", "-", "",
		"nightly analytics pipeline", strings.Repeat("a.", 40), strings.Repeat("-", 40),
		"aaaaaaaaaaaaaaa.bbbbb", strings.Repeat("é", 40) + " tail",
	}
	for _, suffix := range suffixes {
		got := clientVersion("2026.5.0", suffix)
		if !resolverContract.MatchString(got) {
			t.Errorf("clientVersion(_, %q) = %q, which the resolver would reject", suffix, got)
		}
		// Anchored to the Go constant rather than the pattern, so raising one
		// without the other fails here instead of drifting silently.
		metadata, _ := strings.CutPrefix(got, "2026.5.0")
		if n := len(strings.TrimPrefix(metadata, "+")); n > maxBuildMetadata {
			t.Errorf("clientVersion(_, %q) = %q: metadata is %d chars, want at most %d",
				suffix, got, n, maxBuildMetadata)
		}
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
