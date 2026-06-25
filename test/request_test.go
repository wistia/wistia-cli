package integration

import (
	"encoding/json"
	"strings"
	"testing"
)

// These tests assert how the CLI constructs the HTTP request from flags. The
// mock records the request regardless of how the client later handles the
// response, so they validate the flag → request mapping independent of response
// shape. (Response parsing, output formats, and exit codes are covered in
// output_test.go.)

// A GET list maps to the collection endpoint with flags serialized to query
// params. There is no --all flag; --page / --per-page / --cursor are the real
// pagination controls.
func TestRequest_MediaList_MethodPathQuery(t *testing.T) {
	srv, got := newMockAPI(t, 200, `[]`)
	runMock(t, srv, "media", "list",
		"--page", "2",
		"--per-page", "50",
		"--name", "My Vid",
		"--folder-id", "fld1",
		"--sort-by", "name",
		"--archived",
	)

	if got.method != "GET" {
		t.Errorf("method = %s, want GET", got.method)
	}
	if !strings.HasSuffix(got.path, "/medias") {
		t.Errorf("path = %s, want suffix /medias", got.path)
	}
	for k, want := range map[string]string{
		"page":      "2",
		"per_page":  "50",
		"name":      "My Vid",
		"folder_id": "fld1",
		"sort_by":   "name",
		"archived":  "true",
	} {
		if g := got.query.Get(k); g != want {
			t.Errorf("query %s = %q, want %q (raw: %s)", k, g, want, got.rawQuery)
		}
	}
}

// A path parameter (the hashed id) is interpolated into the URL path.
func TestRequest_MediaGet_PathParam(t *testing.T) {
	srv, got := newMockAPI(t, 200, `{"hashed_id":"abc123"}`)
	runMock(t, srv, "media", "get", "--media-hashed-id", "abc123")

	if got.method != "GET" {
		t.Errorf("method = %s, want GET", got.method)
	}
	if !strings.HasSuffix(got.path, "/medias/abc123") {
		t.Errorf("path = %s, want suffix /medias/abc123", got.path)
	}
}

// A create command issues a POST with a JSON body built from its flags.
func TestRequest_FolderCreate_MethodBody(t *testing.T) {
	srv, got := newMockAPI(t, 201, `{"id":1,"name":"Test Folder"}`)
	runMock(t, srv, "folders", "create", "--name", "Test Folder")

	if got.method != "POST" {
		t.Errorf("method = %s, want POST", got.method)
	}
	if !strings.HasSuffix(got.path, "/folders") {
		t.Errorf("path = %s, want suffix /folders", got.path)
	}
	var body map[string]any
	if err := json.Unmarshal([]byte(got.body), &body); err != nil {
		t.Fatalf("request body is not valid JSON: %v\nbody: %s", err, got.body)
	}
	if body["name"] != "Test Folder" {
		t.Errorf("request body name = %v, want %q (body: %s)", body["name"], "Test Folder", got.body)
	}
}

// A delete command issues a DELETE to the resource path.
func TestRequest_MediaDelete_Method(t *testing.T) {
	srv, got := newMockAPI(t, 200, `{}`)
	runMock(t, srv, "media", "delete", "--media-hashed-id", "xyz789")

	if got.method != "DELETE" {
		t.Errorf("method = %s, want DELETE", got.method)
	}
	if !strings.HasSuffix(got.path, "/medias/xyz789") {
		t.Errorf("path = %s, want suffix /medias/xyz789", got.path)
	}
}
