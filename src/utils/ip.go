package utils

import (
	"net"
	"net/http"
	"strings"

	"orderbento/src/utils/zapLog"

	"go.uber.org/zap"
)

var maxCidrBlocks = [...]string{
	"127.0.0.1/8",    // localhost
	"10.0.0.0/8",     // 24-bit block
	"172.16.0.0/12",  // 20-bit block
	"192.168.0.0/16", // 16-bit block
	"169.254.0.0/16", // link local address
	"::1/128",        // localhost IPv6
	"fc00::/7",       // unique local address IPv6
	"fe80::/10",      // link local address IPv6
}

var ipnets []*net.IPNet

func init() {
	ipnets = make([]*net.IPNet, len(maxCidrBlocks))
	for i, v := range maxCidrBlocks {
		_, ipnet, _ := net.ParseCIDR(v)
		ipnets[i] = ipnet
	}
}

func IsPrivateAddress(address string) bool {
	ipAddress := net.ParseIP(address)
	if ipAddress == nil {
		return false
	}

	for _, value := range ipnets {
		if value.Contains(ipAddress) {
			return true
		}
	}

	return false
}

func GetRealIp(req *http.Request) string {
	xRealIP := req.Header.Get("X-Real-Ip")
	xForwardedFor := req.Header.Get("X-Forwarded-For")

	if xRealIP == "" && xForwardedFor == "" {
		var remoteIP string
		// var err error
		// If there are colon in remote address, remove the port number
		// otherwise, return remote address as is
		if strings.ContainsRune(req.RemoteAddr, ':') {
			var err error
			remoteIP, _, err = net.SplitHostPort(req.RemoteAddr)
			if err != nil {
				zapLog.WriteLogError("net SplitHostPort error!:", zap.Error(err))
				return "無法解析的IP"
			}
		} else {
			remoteIP = req.RemoteAddr
		}

		return remoteIP
	}

	for _, address := range strings.Split(xForwardedFor, ",") {
		address = strings.TrimSpace(address)
		isPrivate := IsPrivateAddress(address)
		if !isPrivate { //不為私人域名才認定為真實IP
			return address
		}
	}

	// If nothing succeed, return X-Real-IP
	return xRealIP
}
