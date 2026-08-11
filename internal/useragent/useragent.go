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
// User-Agent carries, with the CI suffix attached as SemVer build metadata
// (e.g. "2026.5.1+ci") so internal traffic stays distinguishable while the
// client name stays stable. Source builds report "dev".
func ClientVersion() string {
	return clientVersion(Version, strings.TrimSpace(os.Getenv(suffixEnvVar)))
}

func clientVersion(version, suffix string) string {
	if suffix == "" {
		return version
	}
	return version + "+" + suffix
}
