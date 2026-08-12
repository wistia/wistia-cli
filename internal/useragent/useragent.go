// Package useragent builds the CLI identity sent with Wistia API requests.
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

// ClientName remains stable because the server allowlists client names.
const ClientName = "wistia-cli"

// ClientVersion adds the optional traffic suffix as SemVer build metadata.
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

// Matches the bound the server validates against.
const maxBuildMetadata = 16

// Keep this generic so new traffic suffixes remain valid without a CLI release.
// The User-Agent retains the original suffix for existing log filters.
func buildMetadata(suffix string) string {
	var b strings.Builder
	skipped := false
	for _, r := range suffix {
		if !isBuildMetadataRune(r) {
			skipped = true
			continue
		}
		// Avoid a leading separator when the suffix starts with invalid characters.
		if skipped && b.Len() > 0 {
			b.WriteByte('-')
		}
		skipped = false
		b.WriteRune(r)
	}

	// SemVer forbids empty build identifiers.
	identifiers := strings.Split(b.String(), ".")
	kept := identifiers[:0]
	for _, id := range identifiers {
		if id != "" {
			kept = append(kept, id)
		}
	}
	metadata := strings.Join(kept, ".")

	// Truncate rather than drop: an over-long suffix still marks the traffic as
	// internal, where a bare version would pass for a release build.
	if len(metadata) > maxBuildMetadata {
		// A trailing "-" is a legal identifier character; a trailing "." is not.
		metadata = strings.TrimRight(metadata[:maxBuildMetadata], ".")
	}
	return metadata
}

func isBuildMetadataRune(r rune) bool {
	return r == '.' || r == '-' ||
		(r >= '0' && r <= '9') ||
		(r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z')
}
