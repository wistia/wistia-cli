// Package useragent builds the headers that identify the CLI on every Wistia
// API request: a branded, versioned User-Agent (rather than the generic
// Speakeasy default), and the X-Wistia-Client-Name / X-Wistia-Client-Version
// pair. Together they let CLI traffic be identified and version-tracked
// server-side without parsing a vendor-generated string.
package useragent

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Version is the CLI release version, embedded at build time via
// -ldflags "-X github.com/wistia/wistia-cli/internal/useragent.Version=<version>".
// It stays "dev" for local and source builds.
var Version = "dev"

// suffixEnvVar, when set, is appended to the User-Agent. Wistia's own CI sets it
// so internal test traffic can be excluded from CLI adoption metrics.
const suffixEnvVar = "WISTIA_CLI_USER_AGENT_SUFFIX"

// String returns the User-Agent header value for outbound API requests.
func String() string {
	return format(Version, runtime.GOOS, runtime.GOARCH, strings.TrimSpace(os.Getenv(suffixEnvVar)))
}

func format(version, goos, goarch, suffix string) string {
	ua := fmt.Sprintf("wistia-cli/%s (%s/%s)", version, goos, goarch)
	if suffix != "" {
		ua += " " + suffix
	}
	return ua
}

// ClientName is the X-Wistia-Client-Name value. It is deliberately invariant:
// the server allowlists client names, so this must not vary by platform, build
// or environment.
const ClientName = "wistia-cli"

// ClientVersion returns the X-Wistia-Client-Version value: the same version the
// User-Agent carries, with the suffix attached as SemVer build metadata
// (e.g. "2026.5.1+ci") so internal traffic stays distinguishable while the
// client name stays stable. Source builds report "dev".
func ClientVersion() string {
	return clientVersion(Version, strings.TrimSpace(os.Getenv(suffixEnvVar)))
}

func clientVersion(version, suffix string) string {
	metadata := buildMetadata(suffix)
	if metadata == "" {
		return version
	}
	return version + "+" + metadata
}

// buildMetadata reduces a free-form suffix to valid SemVer build metadata:
// dot-separated identifiers drawn from [0-9A-Za-z-]. Any run of characters
// outside that set collapses to a single "-", so a suffix like "nightly build"
// becomes "nightly-build" rather than producing "2026.5.1+nightly build",
// which the server rejects — dropping version attribution silently while the
// client name still resolves.
//
// The rule is deliberately general, not an allowlist: "staging", "nightly" and
// anything else legitimate work without a code change here.
//
// Only the header needs this. The User-Agent carries the suffix verbatim,
// which is fine — that string is presentation, not a parsed identifier.
func buildMetadata(suffix string) string {
	var b strings.Builder
	skipped := false
	for _, r := range suffix {
		if !isBuildMetadataRune(r) {
			skipped = true
			continue
		}
		// Only separate what was actually joined; a leading invalid run must not
		// produce a leading "-".
		if skipped && b.Len() > 0 {
			b.WriteByte('-')
		}
		skipped = false
		b.WriteRune(r)
	}

	// Empty identifiers ("a..b", ".a", "a.") are not valid build metadata.
	identifiers := strings.Split(b.String(), ".")
	kept := identifiers[:0]
	for _, id := range identifiers {
		if id != "" {
			kept = append(kept, id)
		}
	}
	return strings.Join(kept, ".")
}

func isBuildMetadataRune(r rune) bool {
	return r == '.' || r == '-' ||
		(r >= '0' && r <= '9') ||
		(r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z')
}
