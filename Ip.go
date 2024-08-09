package tool

import (
	"net"
	"net/http"
	"strings"
)

// 客户端ip
func IpClient(request *http.Request) string {
	result := ""
	aliCDNRealIp := request.Header.Get("ali-cdn-real-ip")
	xForwardedFor := request.Header.Get("x-forwarded-for")
	xRealIp := request.Header.Get("x-real-ip")
	xForwardFor := request.Header.Get("x-forward-for")

	if aliCDNRealIp != "" {
		result = aliCDNRealIp
	} else if xForwardedFor != "" {
		result = xForwardedFor
	} else if xRealIp != "" {
		result = xRealIp
	} else if xForwardFor != "" {
		result = xForwardFor
	}

	if len(result) > 0 {
		arr := strings.Split(result, ",")
		result = strings.Trim(arr[0], " ")
	}

	return result
}

// 服务端内网
func IpServerInternal() string {
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
