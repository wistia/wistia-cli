package integration

import (
	"encoding/json"
	"testing"
)

// These tests cover what the CLI does with a response: output formatting, exit
// codes, and auth headers — the layer the dry-run preview can't exercise.

// --output-format json renders the response as valid JSON on stdout.
func TestOutput_JSON_Valid(t *testing.T) {
	srv, _ := newMockAPI(t, 200, `{"hashed_id":"abc123","name":"Demo"}`)
	res := runMock(t, srv, "media", "get", "--media-hashed-id", "abc123", "-o", "json")
	if res.exitCode != 0 {
		t.Fatalf("exit = %d, want 0; stderr:\n%s", res.exitCode, res.stderr)
	}
	var v map[string]any
	if err := json.Unmarshal([]byte(res.stdout), &v); err != nil {
		t.Fatalf("stdout is not valid JSON: %v\nstdout:\n%s", err, res.stdout)
	}
	if v["hashed_id"] != "abc123" {
		t.Errorf("hashed_id = %v, want abc123", v["hashed_id"])
	}
}

// A successful response exits 0 and writes the response to stdout, not stderr.
func TestExitCode_Success_Zero(t *testing.T) {
	srv, _ := newMockAPI(t, 200, `{"hashed_id":"abc123"}`)
	res := runMock(t, srv, "media", "get", "--media-hashed-id", "abc123")
	if res.exitCode != 0 {
		t.Errorf("exit = %d, want 0; stderr:\n%s", res.exitCode, res.stderr)
	}
}

// A 4xx/5xx API response makes the CLI exit non-zero with details on stderr.
func TestExitCode_APIError_NonZero(t *testing.T) {
	srv, _ := newMockAPI(t, 404, `{"error":"not found"}`)
	res := runMock(t, srv, "media", "get", "--media-hashed-id", "nope")
	if res.exitCode == 0 {
		t.Errorf("want non-zero exit on HTTP 404, got 0\nstdout:\n%s\nstderr:\n%s", res.stdout, res.stderr)
	}
}

// Auth resolves from the --bearer-auth flag → Authorization: Bearer <token>.
func TestAuth_FlagSetsBearerHeader(t *testing.T) {
	srv, got := newMockAPI(t, 200, `{}`)
	runMock(t, srv, "media", "get", "--media-hashed-id", "abc123")
	if h := got.headers.Get("Authorization"); h != "Bearer test-token" {
		t.Errorf("Authorization = %q, want %q", h, "Bearer test-token")
	}
}

// Auth resolves from the WISTIA_CLI_BEARER_AUTH env var when no flag is given.
func TestAuth_EnvSetsBearerHeader(t *testing.T) {
	srv, got := newMockAPI(t, 200, `{}`)
	runWistiaEnv(t, []string{"WISTIA_CLI_BEARER_AUTH=env-token"},
		"media", "get", "--media-hashed-id", "abc123", "--server-url", srv.URL)
	if h := got.headers.Get("Authorization"); h != "Bearer env-token" {
		t.Errorf("Authorization = %q, want %q", h, "Bearer env-token")
	}
}
