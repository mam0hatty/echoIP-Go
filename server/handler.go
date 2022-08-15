package server

import (
	"net/http"
	"net/netip"
	"net/url"
)

func toIPInfo(w http.ResponseWriter, r *http.Request) {
	if ip, err := netip.ParseAddr(getClientIP(r)); err == nil {
		if !ip.Is4() { // is it a valid IPV4 Addr?
			http.NotFound(w, r)
		} else {
			if u, err := url.ParseRequestURI(r.RequestURI); err == nil {
				q := u.Query()                        // get parameters from query
				ipinfo := getIPInfo(ip, q.Get("gbk")) // get ipinfo
				if q.Get("cli") == "true" {
					outputCli(ipinfo)
				} else {
					output(ipinfo)
				}
			}
		}
	}
}

func toIPOnly(w http.ResponseWriter, r *http.Request) {

}

func toIPInfoPara(w http.ResponseWriter, r *http.Request) {

}

func toQRCode(w http.ResponseWriter, r *http.Request) {

}

func toQRCodePara(w http.ResponseWriter, r *http.Request) {

}

func getField(r *http.Request) string {
	fields := r.Context().Value(ctxKey{}).(string)
	return fields
}

func getClientIP(r *http.Request) string {
	if xip := r.Header.Get("X-Real-IP"); xip != "" {
		return xip // get client's real IP from nginx's custom header
	} else {
		return r.RemoteAddr // no proxy involved
	}
}

type IPInfo struct {
	ip       string
	as       string
	city     string
	region   string
	country  string
	timezone string
	location string
	isp      string
	scope    string
	detail   string
	status   string
}

func getIPInfo(IP netip.Addr, encoding string) *IPInfo {
	// try get it from redis
	// if no, query for it
	// if gbk true, convert it
	return nil

}

func outputCli(info *IPInfo) {
}

func output(info *IPInfo) {

}

func queryIPInfo() {
	// if ip is ipv6
	// compress it
	// if ip is special
	// set some fields in IPInfo
	// else
	// get info from IPDB
	// get info from ipinfo
	// get detail from qqwry
	// set fields in IPInfo according to results above
	// ipinfo conflict solution
	// set fields
	// minor tweaks
	// return IPInfo
}

// fmt.Fprintln(w, "\033[5m字符串闪耀\033[0m")
