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
	// Hooks run after --header values are applied; a version the user chose wins.
	if req.Header.Get("X-Wistia-API-Version") != "" {
		return req, nil
	}
	if version := apiVersion(hookCtx.SDKConfiguration.UserAgent); version != "" {
		req.Header.Set("X-Wistia-API-Version", version)
	}
	return req, nil
}

var docVersionPattern = regexp.MustCompile(`^(\d{4})\.(\d{2})\.\d+$`)

// apiVersion maps the doc version in the generated User-Agent
// ("speakeasy-sdk/go <sdk> <generator> <docVersion> <package>") to the
// header's form: YYYY-MM for a dated release, edge-version for edge builds.
func apiVersion(generatedUserAgent string) string {
	fields := strings.Fields(generatedUserAgent)
	if len(fields) < 4 {
		return ""
	}
	if fields[3] == "edge-version" {
		return "edge-version"
	}
	match := docVersionPattern.FindStringSubmatch(fields[3])
	if match == nil {
		return ""
	}
	return match[1] + "-" + match[2]
}
