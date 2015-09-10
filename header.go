package luddite

import (
	"net/http"
	"net/url"
	"strconv"
)

const (
	HeaderAccept            = "Accept"
	HeaderAuthorization     = "Authorization"
	HeaderContentType       = "Content-Type"
	HeaderForwardedFor      = "X-Forwarded-For"
	HeaderLocation          = "Location"
	HeaderRequestId         = "X-Request-Id"
	HeaderSpirentApiVersion = "X-Spirent-Api-Version"
	HeaderSpirentNextLink   = "X-Spirent-Next-Link"
)

func RequestBearerToken(r *http.Request) (token string) {
	if authStr := r.Header.Get(HeaderAuthorization); authStr != "" && authStr[:7] == "Bearer " {
		token = authStr[7:]
	}
	return
}

func RequestApiVersion(r *http.Request, defaultVersion int) (version int) {
	version = defaultVersion
	if versionStr := r.Header.Get(HeaderSpirentApiVersion); versionStr != "" {
		if i, err := strconv.Atoi(versionStr); err == nil && i > 0 {
			version = i
		}
	}
	return
}

func RequestQueryCursor(r *http.Request) string {
	return r.URL.Query().Get("cursor")
}

func RequestNextLink(r *http.Request, cursor string) *url.URL {
	next := *r.URL
	v := next.Query()
	v.Set("cursor", cursor)
	next.RawQuery = v.Encode()
	return &next
}
