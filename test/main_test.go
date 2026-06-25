// Package integration holds black-box tests that exercise the real, compiled
// wistia binary the way a user would invoke it.
//
// Two tiers:
//
//   - Offline (request_test.go, output_test.go): run the CLI against an
//     in-process mock HTTP server (net/http/httptest) pointed at via
//     --server-url. Asserts request construction, response parsing, output
//     formats, auth headers, and exit codes. No account, no network, no secret —
//     runs on every push/PR, including in CI (it is just `go test`).
//   - Live (integration_test.go, build tag "integration"): real CRUD against a
//     dedicated test account, gated on the WISTIA_API_KEY secret.
//
// Run offline:  go test ./test/...
// Run live:     go test -tags integration ./test/...
package integration

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// wistiaBin is the path to the CLI binary built once for the whole suite.
var wistiaBin string

// TestMain builds the wistia binary from source so the tests exercise the
// assembled CLI end to end. Building here keeps `go test ./test/...`
// self-contained — no separate build step is required locally or in CI.
func TestMain(m *testing.M) {
	code, err := setupAndRun(m)
	if err != nil {
		fmt.Fprintln(os.Stderr, "integration test setup failed:", err)
		os.Exit(1)
	}
	os.Exit(code)
}

func setupAndRun(m *testing.M) (int, error) {
	root, err := moduleRoot()
	if err != nil {
		return 1, err
	}

	binDir, err := os.MkdirTemp("", "wistia-bin-")
	if err != nil {
		return 1, err
	}
	defer os.RemoveAll(binDir)

	wistiaBin = filepath.Join(binDir, "wistia")
	build := exec.Command("go", "build", "-o", wistiaBin, "./cmd/wistia")
	build.Dir = root
	build.Stderr = os.Stderr
	if err := build.Run(); err != nil {
		return 1, fmt.Errorf("building CLI: %w", err)
	}

	return m.Run(), nil
}

// moduleRoot returns the repository root (the parent of this test/ directory).
func moduleRoot() (string, error) {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("cannot determine source path")
	}
	return filepath.Dir(filepath.Dir(thisFile)), nil
}

// result captures the outcome of one CLI invocation.
type result struct {
	stdout   string
	stderr   string
	exitCode int
}

// runWistia executes the built CLI with the given args in an isolated
// environment and returns its captured output.
func runWistia(t *testing.T, args ...string) result {
	t.Helper()
	return runWistiaEnv(t, nil, args...)
}

// runWistiaEnv is runWistia with extra environment variables appended to the
// isolated base environment (used to exercise env-var-based config such as
// WISTIA_CLI_BEARER_AUTH).
func runWistiaEnv(t *testing.T, extraEnv []string, args ...string) result {
	t.Helper()
	if wistiaBin == "" {
		t.Fatal("wistiaBin not built (did TestMain run?)")
	}

	// --no-interactive guards against any prompt; the tests never feed stdin.
	full := append([]string{"--no-interactive"}, args...)
	cmd := exec.Command(wistiaBin, full...)
	cmd.Env = append(isolatedEnv(t), extraEnv...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	exit := 0
	if err := cmd.Run(); err != nil {
		var ee *exec.ExitError
		if errors.As(err, &ee) {
			exit = ee.ExitCode()
		} else {
			t.Fatalf("running %q: %v", strings.Join(full, " "), err)
		}
	}
	return result{stdout: stdout.String(), stderr: stderr.String(), exitCode: exit}
}

// isolatedEnv returns a minimal environment for a CLI invocation:
//   - a throwaway HOME, so the developer's ~/.config/wistia config can't leak
//     defaults (server URL, output format, credentials) into assertions;
//   - no WISTIA_CLI_* vars, so each test controls auth/config explicitly;
//   - no agent-detection vars, which would auto-enable --agent-mode and change
//     default output/error formatting.
func isolatedEnv(t *testing.T) []string {
	t.Helper()
	return []string{
		"HOME=" + t.TempDir(),
		"PATH=" + os.Getenv("PATH"),
		"NO_COLOR=1",
	}
}

// capturedRequest records what the mock API received.
type capturedRequest struct {
	method   string
	path     string
	query    url.Values
	rawQuery string
	body     string
	headers  http.Header
}

// newMockAPI starts an in-process HTTP server that records the request it
// receives and replies with status + jsonBody. It is closed automatically when
// the test ends. Point the CLI at it with --server-url srv.URL.
//
// This is the standard Go pattern for black-box testing an HTTP client and
// needs no network, account, or secret — it is just `go test`.
func newMockAPI(t *testing.T, status int, jsonBody string) (*httptest.Server, *capturedRequest) {
	t.Helper()
	captured := &capturedRequest{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		captured.method = r.Method
		captured.path = r.URL.Path
		captured.query = r.URL.Query()
		captured.rawQuery = r.URL.RawQuery
		captured.body = string(body)
		captured.headers = r.Header.Clone()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = io.WriteString(w, jsonBody)
	}))
	t.Cleanup(srv.Close)
	return srv, captured
}

// runMock runs the CLI against the mock server with a bearer token supplied via
// the --bearer-auth flag.
func runMock(t *testing.T, srv *httptest.Server, args ...string) result {
	t.Helper()
	return runWistia(t, append(args, "--server-url", srv.URL, "--bearer-auth", "test-token")...)
}
