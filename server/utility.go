package server

import (
	"net/http"
	"net/netip"
)

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

func GetClientIP(r *http.Request) string {
	if xip := r.Header.Get("X-Real-IP"); xip != "" {
		return xip // get client's real IP from nginx's custom header
	} else {
		return r.RemoteAddr // no proxy involved

	}
}

func GetIPInfo(IP netip.Addr, encoding string) *IPInfo {
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
	return nil

}

func OutputCli(info *IPInfo) {
}

func Output(info *IPInfo) {

}

func QueryIPInfo() {

}

// fmt.Fprintln(w, "\033[5m字符串闪耀\033[0m")
