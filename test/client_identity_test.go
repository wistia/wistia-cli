package integration

import (
	"strings"
	"testing"
)

// Black-box requests verify that the generated SDK registers the identity hook.
func TestClientIdentity_Headers(t *testing.T) {
	srv, got := newMockAPI(t, 200, `{}`)
	runMock(t, srv, "media", "get", "--media-hashed-id", "abc123")

	if name := got.headers.Get("X-Wistia-Client-Name"); name != "wistia-cli" {
		t.Errorf("X-Wistia-Client-Name = %q, want %q", name, "wistia-cli")
	}
	// Compare against the User-Agent so this assertion works in dev and release builds.
	want := userAgentVersion(t, got.headers.Get("User-Agent"))
	if version := got.headers.Get("X-Wistia-Client-Version"); version != want {
		t.Errorf("X-Wistia-Client-Version = %q, want %q", version, want)
	}
}

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

func TestClientIdentity_SuffixSanitizedForHeaderOnly(t *testing.T) {
	srv, got := newMockAPI(t, 200, `{}`)
	runWistiaEnv(t, []string{"WISTIA_CLI_USER_AGENT_SUFFIX=nightly build"},
		"media", "get", "--media-hashed-id", "abc123",
		"--server-url", srv.URL, "--bearer-auth", "test-token")

	ua := got.headers.Get("User-Agent")
	if !strings.HasSuffix(ua, " nightly build") {
		t.Errorf("User-Agent = %q, want it to end with %q verbatim", ua, " nightly build")
	}
	want := userAgentVersion(t, ua) + "+nightly-build"
	if version := got.headers.Get("X-Wistia-Client-Version"); version != want {
		t.Errorf("X-Wistia-Client-Version = %q, want %q", version, want)
	}
}

func userAgentVersion(t *testing.T, ua string) string {
	t.Helper()
	_, rest, ok := strings.Cut(ua, "wistia-cli/")
	if !ok {
		t.Fatalf("User-Agent = %q, want it to contain %q", ua, "wistia-cli/")
	}
	version, _, _ := strings.Cut(rest, " ")
	return version
}
