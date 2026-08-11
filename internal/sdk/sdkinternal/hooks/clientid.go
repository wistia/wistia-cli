package hooks

import (
	"github.com/wistia/wistia-cli/internal/useragent"
	"net/http"
)

// These headers are declared attribution, so the server validates and bounds them.
type clientIdentityHook struct{}

var _ beforeRequestHook = (*clientIdentityHook)(nil)

func (clientIdentityHook) BeforeRequest(_ BeforeRequestContext, req *http.Request) (*http.Request, error) {
	req.Header.Set("X-Wistia-Client-Name", useragent.ClientName)
	req.Header.Set("X-Wistia-Client-Version", useragent.ClientVersion())
	return req, nil
}
