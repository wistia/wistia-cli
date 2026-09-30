package hooks

import (
	"net/http"
	"regexp"
	"strings"
)

// Without this header the API serves its newest version, so a release would
// silently drift from the spec it was generated from as new versions ship.
type apiVersionHook struct{}

var _ beforeRequestHook = (*apiVersionHook)(nil)

func (apiVersionHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	if version := apiVersion(hookCtx.SDKConfiguration.UserAgent); version != "" {
		req.Header.Set("X-Wistia-API-Version", version)
	}
	return req, nil
}

var docVersionPattern = regexp.MustCompile(`^(\d{4})\.(\d{2})\.\d+$`)

// apiVersion maps the doc version in the generated User-Agent
// ("speakeasy-sdk/go <sdk> <generator> <docVersion> <package>") to the
// header's YYYY-MM form. Edge builds return "" so they send no header.
func apiVersion(generatedUserAgent string) string {
	fields := strings.Fields(generatedUserAgent)
	if len(fields) < 4 {
		return ""
	}
	match := docVersionPattern.FindStringSubmatch(fields[3])
	if match == nil {
		return ""
	}
	return match[1] + "-" + match[2]
}
