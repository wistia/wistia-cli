package integration

import (
	"os"
	"strings"
	"testing"
)

// TestReleaseIdentity checks the GoReleaser-built artifact so the test covers
// the release ldflags from .goreleaser.yaml rather than a restatement of them.
const (
	releaseVersionEnv = "WISTIA_RELEASE_VERSION"
	releaseBinaryEnv  = "WISTIA_RELEASE_BINARY"
	// Prevent a misconfigured release workflow from silently skipping this gate.
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

	// A unit test on useragent.Version would not prove the hooks that set these
	// headers are registered, so capture them off a real request.
	t.Run("outbound identity", func(t *testing.T) {
		srv, got := newMockAPI(t, 200, `{"hashed_id":"abc123"}`)
		res := runBinEnv(t, bin, nil, "media", "get", "--media-hashed-id", "abc123",
			"--server-url", srv.URL, "--bearer-auth", "test-token")
		if res.exitCode != 0 {
			t.Fatalf("exit = %d, want 0; stderr:\n%s", res.exitCode, res.stderr)
		}

		ua := got.headers.Get("User-Agent")
		if reported := versionFromUserAgent(ua); reported != version {
			t.Errorf("outbound User-Agent reports version %q, want %q (User-Agent = %q)\n"+
				"An unregistered User-Agent hook regresses this to the Speakeasy default; "+
				"a missing -X useragent.Version regresses it to wistia-cli/dev.", reported, version, ua)
		}

		// runBinEnv sets no suffix, so no build metadata is expected here.
		if name := got.headers.Get("X-Wistia-Client-Name"); name != "wistia-cli" {
			t.Errorf("X-Wistia-Client-Name = %q, want %q", name, "wistia-cli")
		}
		if reported := got.headers.Get("X-Wistia-Client-Version"); reported != version {
			t.Errorf("X-Wistia-Client-Version = %q, want %q", reported, version)
		}
	})
}

// Prefix matching rejects Speakeasy's User-Agent, whose module path also
// contains "wistia-cli"; returning the whole token avoids version collisions.
func versionFromUserAgent(ua string) string {
	rest, ok := strings.CutPrefix(ua, "wistia-cli/")
	if !ok {
		return ""
	}
	version, _, _ := strings.Cut(rest, " ")
	return version
}
