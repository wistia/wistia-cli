package hooks

import (
	"github.com/wistia/wistia-cli/internal/useragent"
	"net/http"
)

// userAgentHook replaces the generated SDK's default User-Agent with the CLI's
// branded, versioned value on every request, so CLI traffic is attributable and
// version-trackable server-side. Registered in registration.go.
type userAgentHook struct{}

var _ beforeRequestHook = (*userAgentHook)(nil)

func (userAgentHook) BeforeRequest(_ BeforeRequestContext, req *http.Request) (*http.Request, error) {
	req.Header.Set("User-Agent", useragent.String())
	return req, nil
}
