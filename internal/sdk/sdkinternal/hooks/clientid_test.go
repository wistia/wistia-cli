package hooks

import (
	"net/http"
	"testing"

	"github.com/wistia/wistia-cli/internal/useragent"
)

func TestClientIdentityHookSetsHeaders(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "https://api.wistia.com/modern/medias", nil)
	if err != nil {
		t.Fatal(err)
	}

	got, err := clientIdentityHook{}.BeforeRequest(BeforeRequestContext{}, req)
	if err != nil {
		t.Fatalf("BeforeRequest returned error: %v", err)
	}
	if name := got.Header.Get("X-Wistia-Client-Name"); name != "wistia-cli" {
		t.Errorf("X-Wistia-Client-Name = %q, want %q", name, "wistia-cli")
	}
	if version := got.Header.Get("X-Wistia-Client-Version"); version != useragent.ClientVersion() {
		t.Errorf("X-Wistia-Client-Version = %q, want %q", version, useragent.ClientVersion())
	}
}
