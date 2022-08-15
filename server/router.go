package server

import (
	"context"
	"net/http"
	"regexp"
)

// NOTE:
// cache regexp or it'd be compiled every time a URI match happens
// NOTE2:
// we can add method check here but why bother
func URIRoute(w http.ResponseWriter, r *http.Request) {
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	http.NotFound(w, r)
}

type route struct {
	regex   *regexp.Regexp   // route URI
	handler http.HandlerFunc // URI handler
}

// a strange magic that creates a valid map key
type ctxKey struct{}

// da router, add new route here and corresponding handle in handler.go
var routes = []route{
	newRoute(`/`, toIPInfo),
	newRoute(`/ip`, toIPOnly),
	newRoute(`/info`, toIPInfo),
	newRoute(`/info/([^/]+)`, toIPInfoPara), // /info/{ip}
	newRoute(`/qr`, toQRCode),
	newRoute(`/qr/..`, toQRCodePara),
}

// beautification
func newRoute(pattern string, handler http.HandlerFunc) route {
	return route{regexp.MustCompile("^" + pattern), handler}
}

// credit(reference)
// https://benhoyt.com/writings/go-routing/
// https://github.com/benhoyt/go-routing
