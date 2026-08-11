package hooks

import (
	"github.com/wistia/wistia-cli/internal/useragent"
	"net/http"
)

// clientIdentityHook declares who the caller is, in headers the server can read
// directly instead of pattern-matching the User-Agent. Registered in
// registration.go alongside userAgentHook.
//
// These are declared attribution — any authenticated caller could send them —
// so the server normalizes and bounds the values. The CLI's job is only to send
// them consistently.
type clientIdentityHook struct{}

var _ beforeRequestHook = (*clientIdentityHook)(nil)

func (clientIdentityHook) BeforeRequest(_ BeforeRequestContext, req *http.Request) (*http.Request, error) {
	req.Header.Set("X-Wistia-Client-Name", useragent.ClientName)
	req.Header.Set("X-Wistia-Client-Version", useragent.ClientVersion())
	return req, nil
}
