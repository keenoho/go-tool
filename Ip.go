package tool

import (
	"io/ioutil"
	"net"
	"net/http"
)

func ClientIp(request *http.Request) string {
	aliCDNRealIp := request.Header.Get("ali-cdn-real-ip")
	xForwardedFor := request.Header.Get("x-forwarded-for")
	xRealIp := request.Header.Get("x-real-ip")
	xForwardFor := request.Header.Get("x-forward-for")

	if aliCDNRealIp != "" {
		return aliCDNRealIp
	} else if xForwardedFor != "" {
		return xForwardedFor
	} else if xRealIp != "" {
		return xRealIp
	} else if xForwardFor != "" {
		return xForwardFor
	} else {
		return ""
	}
}

func ServerInternalIp() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			return ip.String()
		}
	}
	return ""
}

func ServerPublicIp() string {
	resp, err := http.Get("http://ipinfo.io/ip")
	if err != nil {
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}
