package integration

import (
	"os"
	"strings"
	"testing"
)

// Release gate for how the CLI identifies itself.
//
// Server-side CLI usage tracking keys on the version the binary reports, so a
// release that ships a "dev" or Speakeasy-default identity is invisible in those
// metrics. That has nearly happened: the User-Agent branding was absent from a
// release line because the commit carrying it was not on the release branch, and
// nothing in the release path noticed.
//
// This asserts against the actual release artifact that `wistia version` and the
// outbound User-Agent both report the tag's version. Run from
// .github/workflows/release.yaml after GoReleaser builds and before the tag is
// pushed, so a mismatch stops the release. Reproduce locally:
//
//	goreleaser build --single-target --clean --skip=validate -o /tmp/wistia
//	WISTIA_RELEASE_VERSION=2026.5.1 WISTIA_RELEASE_BINARY=/tmp/wistia \
//	  go test ./test/... -run TestReleaseIdentity -v
//
// The binary is supplied rather than built here on purpose: .goreleaser.yaml
// stays the single source of truth for the release ldflags, and the check then
// covers the ldflags themselves, not a restatement of them.
const (
	releaseVersionEnv = "WISTIA_RELEASE_VERSION"
	releaseBinaryEnv  = "WISTIA_RELEASE_BINARY"
	// Set in the release workflow so a misconfigured gate fails loudly instead of
	// skipping (mirrors WISTIA_REQUIRE_LIVE for the live tests).
	requireReleaseEnv = "WISTIA_REQUIRE_RELEASE_CHECK"
)

func TestReleaseIdentity(t *testing.T) {
	version := strings.TrimPrefix(strings.TrimSpace(os.Getenv(releaseVersionEnv)), "v")
	bin := strings.TrimSpace(os.Getenv(releaseBinaryEnv))

	if version == "" || bin == "" {
		if os.Getenv(requireReleaseEnv) != "" {
			t.Fatalf("%s is set, so the release gate must run, but %s=%q and %s=%q",
				requireReleaseEnv, releaseVersionEnv, version, releaseBinaryEnv, bin)
		}
		t.Skipf("set %s and %s to check a release build", releaseVersionEnv, releaseBinaryEnv)
	}
	if version == "dev" {
		t.Fatalf("%s = %q: a release must carry a real version", releaseVersionEnv, version)
	}

	t.Run("version command", func(t *testing.T) {
		res := runBinEnv(t, bin, nil, "version")
		if res.exitCode != 0 {
			t.Fatalf("exit = %d, want 0; stderr:\n%s", res.exitCode, res.stderr)
		}
		// Trailing lines carry the optional build time.
		first, _, _ := strings.Cut(res.stdout, "\n")
		if got, want := strings.TrimSpace(first), "wistia "+version; got != want {
			t.Errorf("wistia version = %q, want %q\n"+
				"Missing -X main.version means the binary reports the generated SDK version instead of the release.",
				got, want)
		}
	})

	// A unit test on useragent.Version would not prove the hook that sets the
	// header is registered, so capture the header off a real request.
	t.Run("outbound user agent", func(t *testing.T) {
		srv, got := newMockAPI(t, 200, `{"hashed_id":"abc123"}`)
		res := runBinEnv(t, bin, nil, "media", "get", "--media-hashed-id", "abc123",
			"--server-url", srv.URL, "--bearer-auth", "test-token")
		if res.exitCode != 0 {
			t.Fatalf("exit = %d, want 0; stderr:\n%s", res.exitCode, res.stderr)
		}
		ua := got.headers.Get("User-Agent")
		if want := "wistia-cli/" + version; !strings.Contains(ua, want) {
			t.Errorf("outbound User-Agent = %q, want it to contain %q\n"+
				"An unregistered User-Agent hook regresses this to the Speakeasy default; "+
				"a missing -X useragent.Version regresses it to wistia-cli/dev.", ua, want)
		}
	})
}
