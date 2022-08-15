package server

import (
	"context"
	"net/http"
	"net/netip"
	"net/url"
	"regexp"
)

func URIRoute(w http.ResponseWriter, r *http.Request) {
	// NOTE:
	// cache regexp or it'd be compiled every time a URI match happens

	// NOTE2:
	// we can add method check here but why bother
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

type ctxKey struct{}

var routes = []route{
	newRoute(`/`, ToIPInfo),
	newRoute(`/ip`, ToIPOnly),
	newRoute(`/info`, ToIPInfo),
	newRoute(`/info/([^/]+)`, ToIPInfoPara), // /info/{ip}
	newRoute(`/qr`, ToQRCode),
	newRoute(`/qr/..`, ToQRCodePara),
}

func getField(r *http.Request) string {
	fields := r.Context().Value(ctxKey{}).(string)
	return fields
}

func newRoute(pattern string, handler http.HandlerFunc) route {
	return route{regexp.MustCompile("^" + pattern), handler}
}

func ToIPInfo(w http.ResponseWriter, r *http.Request) { // /ip/info(?(gbk=)(cli=))
	if ip, err := netip.ParseAddr(GetClientIP(r)); err == nil {
		if !ip.Is4() { // is it a valid IPV4 Addr?
			http.NotFound(w, r)
		} else {
			if u, err := url.ParseRequestURI(r.RequestURI); err == nil {
				q := u.Query()                        // get parameters from query
				ipinfo := GetIPInfo(ip, q.Get("gbk")) // get ipinfo
				// TODO: change to "encode" to support other encoding
				if q.Get("cli") == "true" {
					OutputCli(ipinfo)
				} else {
					Output(ipinfo)
				}

			}
		}
	}
}

func ToIPOnly(w http.ResponseWriter, r *http.Request) {

}

func ToIPInfoPara(w http.ResponseWriter, r *http.Request) {

}

func ToQRCode(w http.ResponseWriter, r *http.Request) {

}

func ToQRCodePara(w http.ResponseWriter, r *http.Request) {

}

// credit(reference)
// https://benhoyt.com/writings/go-routing/
// https://github.com/benhoyt/go-routing
