package integration

import (
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
	ua := got.headers.Get("User-Agent")
	want := versionFromUserAgent(ua)
	if want == "" {
		t.Fatalf("User-Agent = %q, want a branded wistia-cli/<version> value", ua)
	}
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
	ua := got.headers.Get("User-Agent")
	base := versionFromUserAgent(ua)
	if base == "" {
		t.Fatalf("User-Agent = %q, want a branded wistia-cli/<version> value", ua)
	}
	if version, want := got.headers.Get("X-Wistia-Client-Version"), base+"+ci"; version != want {
		t.Errorf("X-Wistia-Client-Version = %q, want %q", version, want)
	}
}
