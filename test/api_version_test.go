package integration

import (
	"os"
	"regexp"
	"testing"
)

var docVersionLine = regexp.MustCompile(`(?m)^\s*docVersion:\s*(\S+)\s*$`)

// Derives the expectation from gen.lock so one test covers both kinds of branch:
// version branches must send their YYYY-MM, and main (edge) must send nothing.
func TestAPIVersion_Header(t *testing.T) {
	lock, err := os.ReadFile("../.speakeasy/gen.lock")
	if err != nil {
		t.Fatal(err)
	}
	match := docVersionLine.FindSubmatch(lock)
	if match == nil {
		t.Fatal("no docVersion in .speakeasy/gen.lock")
	}
	docVersion := string(match[1])

	srv, got := newMockAPI(t, 200, `{}`)
	runMock(t, srv, "media", "get", "--media-hashed-id", "abc123")
	version, present := got.headers["X-Wistia-Api-Version"]

	if docVersion == "edge-version" {
		if present {
			t.Errorf("X-Wistia-API-Version = %q, want no header on an edge build", version)
		}
		return
	}

	dated := regexp.MustCompile(`^(\d{4})\.(\d{2})\.\d+$`).FindStringSubmatch(docVersion)
	if dated == nil {
		t.Fatalf("docVersion %q is neither edge-version nor YYYY.MM.patch", docVersion)
	}
	want := dated[1] + "-" + dated[2]
	if got := got.headers.Get("X-Wistia-API-Version"); got != want {
		t.Errorf("X-Wistia-API-Version = %q, want %q (docVersion %s)", got, want, docVersion)
	}
}
