// Package useragent builds the User-Agent header the CLI sends on every Wistia
// API request. A branded, versioned value (rather than the generic Speakeasy
// default) is what lets CLI traffic be identified and version-tracked
// server-side.
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
