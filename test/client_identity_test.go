package integration

import (
	"strings"
	"testing"
)

// The CLI declares who it is in dedicated headers so server-side attribution
// doesn't have to parse the User-Agent. These capture the headers off a real
// request, which is what proves the hook is registered.

func TestClientIdentity_Headers(t *testing.T) {
	srv, got := newMockAPI(t, 200, `{}`)
	runMock(t, srv, "media", "get", "--media-hashed-id", "abc123")

	if name := got.headers.Get("X-Wistia-Client-Name"); name != "wistia-cli" {
		t.Errorf("X-Wistia-Client-Name = %q, want %q", name, "wistia-cli")
	}
	// The suite builds without the release ldflags, so the expected version is
	// whatever this build reports — "dev" here, the release version in a release
	// build. Reading it from the User-Agent also asserts the two agree.
	want := userAgentVersion(t, got.headers.Get("User-Agent"))
	if version := got.headers.Get("X-Wistia-Client-Version"); version != want {
		t.Errorf("X-Wistia-Client-Version = %q, want %q", version, want)
	}
}

// The CI suffix rides along as SemVer build metadata, so Wistia's own traffic
// stays distinguishable without the client name varying.
func TestClientIdentity_CISuffixRidesOnVersion(t *testing.T) {
	srv, got := newMockAPI(t, 200, `{}`)
	runWistiaEnv(t, []string{"WISTIA_CLI_USER_AGENT_SUFFIX=ci"},
		"media", "get", "--media-hashed-id", "abc123",
		"--server-url", srv.URL, "--bearer-auth", "test-token")

	if name := got.headers.Get("X-Wistia-Client-Name"); name != "wistia-cli" {
		t.Errorf("X-Wistia-Client-Name = %q, want %q — the name must not vary with the suffix",
			name, "wistia-cli")
	}
	want := userAgentVersion(t, got.headers.Get("User-Agent")) + "+ci"
	if version := got.headers.Get("X-Wistia-Client-Version"); version != want {
		t.Errorf("X-Wistia-Client-Version = %q, want %q", version, want)
	}
}

// userAgentVersion pulls the version token out of "wistia-cli/<version> (os/arch)".
func userAgentVersion(t *testing.T, ua string) string {
	t.Helper()
	_, rest, ok := strings.Cut(ua, "wistia-cli/")
	if !ok {
		t.Fatalf("User-Agent = %q, want it to contain %q", ua, "wistia-cli/")
	}
	version, _, _ := strings.Cut(rest, " ")
	return version
}
