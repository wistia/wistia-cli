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

// ClientVersion returns the X-Wistia-Client-Version value:
// "\d+\.\d+\.\d+" or "dev", optionally followed by "+ci".
//
// The marker is a boolean — "this is internal, non-production traffic" — and is
// spelled "ci" whatever the suffix actually says, so a future "staging" caller
// also reads as ci in the warehouse. Its real label stays in the User-Agent.
func ClientVersion() string {
	return clientVersion(Version, strings.TrimSpace(os.Getenv(suffixEnvVar)))
}

func clientVersion(version, suffix string) string {
	if suffix == "" {
		return version
	}
	return version + "+ci"
}
