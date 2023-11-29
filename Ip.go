package tool

import (
	"io"
	"net"
	"net/http"
)

// ClientIp 客户端ip
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

// ServerInternalIp 服务端内网ip
func ServerInternalIp() string {
	faces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, face := range faces {
		if face.Flags&net.FlagUp == 0 {
			continue
		}
		if face.Flags&net.FlagLoopback != 0 {
			continue
		}
		adds, _ := face.Addrs()
		for _, addr := range adds {
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

// ServerPublicIp 服务端公网ip
func ServerPublicIp() string {
	resp, err := http.Get("http://ipinfo.io/ip")
	if err != nil {
		return ""
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}
