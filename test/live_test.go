//go:build integration

package integration

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"
	"time"
)

// Live tests run the CLI against the real Wistia API using the WISTIA_API_KEY
// env var. They are build-tagged ("integration") so a plain `go test` never hits
// the network, and they skip when the key is absent. They create and delete real
// resources on a dedicated test account, cleaning up after themselves.
//
// Run: WISTIA_API_KEY=... go test -tags integration -run TestLive ./test/...

// liveKey returns the API key, or skips the test when it is unset.
func liveKey(t *testing.T) string {
	t.Helper()
	key := os.Getenv("WISTIA_API_KEY")
	if key == "" {
		t.Skip("WISTIA_API_KEY not set; skipping live tests")
	}
	return key
}

// hashedID pulls a hashed id from a JSON response, tolerating either the API's
// snake_case (hashed_id) or a typed camelCase (hashedId) rendering.
func hashedID(t *testing.T, jsonStr string) string {
	t.Helper()
	var v map[string]any
	if err := json.Unmarshal([]byte(jsonStr), &v); err != nil {
		t.Fatalf("response is not JSON: %v\n%s", err, jsonStr)
	}
	for _, k := range []string{"hashed_id", "hashedId"} {
		if s, ok := v[k].(string); ok && s != "" {
			return s
		}
	}
	t.Fatalf("no hashed id in response: %s", jsonStr)
	return ""
}

// The key authenticates against the real API.
func TestLive_AccountGet(t *testing.T) {
	key := liveKey(t)
	res := runWistia(t, "account", "get", "--bearer-auth", key, "-o", "json")
	if res.exitCode != 0 {
		t.Fatalf("account get exit %d; stderr:\n%s", res.exitCode, res.stderr)
	}
	var v map[string]any
	if err := json.Unmarshal([]byte(res.stdout), &v); err != nil {
		t.Fatalf("account get output is not JSON: %v\n%s", err, res.stdout)
	}
}

// A folder can be created, fetched, and deleted end to end.
func TestLive_FolderRoundTrip(t *testing.T) {
	key := liveKey(t)
	name := "CLI-Integration-Test-" + strconv.FormatInt(time.Now().UnixNano(), 10)

	create := runWistia(t, "folders", "create", "--name", name, "--bearer-auth", key, "-o", "json")
	if create.exitCode != 0 {
		t.Fatalf("folders create exit %d; stderr:\n%s", create.exitCode, create.stderr)
	}
	id := hashedID(t, create.stdout)

	// Always clean up, even if a later assertion fails.
	t.Cleanup(func() {
		del := runWistia(t, "folders", "delete", "--id", id, "--bearer-auth", key)
		if del.exitCode != 0 {
			t.Logf("cleanup: folders delete %s exit %d: %s", id, del.exitCode, del.stderr)
		}
	})

	get := runWistia(t, "folders", "get", "--id", id, "--bearer-auth", key, "-o", "json")
	if get.exitCode != 0 {
		t.Fatalf("folders get exit %d; stderr:\n%s", get.exitCode, get.stderr)
	}
	var folder map[string]any
	if err := json.Unmarshal([]byte(get.stdout), &folder); err != nil {
		t.Fatalf("folders get output is not JSON: %v\n%s", err, get.stdout)
	}
	if folder["name"] != name {
		t.Errorf("folder name = %v, want %q", folder["name"], name)
	}
}
