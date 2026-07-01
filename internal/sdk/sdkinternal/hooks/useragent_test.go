package hooks

import (
	"net/http"
	"strings"
	"testing"
)

func TestUserAgentHookOverwritesHeader(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "https://api.wistia.com/modern/medias", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("User-Agent", "speakeasy-sdk/go 0.0.1")

	got, err := userAgentHook{}.BeforeRequest(BeforeRequestContext{}, req)
	if err != nil {
		t.Fatalf("BeforeRequest returned error: %v", err)
	}
	ua := got.Header.Get("User-Agent")
	if !strings.HasPrefix(ua, "wistia-cli/") {
		t.Errorf("User-Agent = %q, want prefix %q", ua, "wistia-cli/")
	}
	if strings.Contains(ua, "speakeasy-sdk") {
		t.Errorf("User-Agent = %q still contains the default speakeasy value", ua)
	}
}
