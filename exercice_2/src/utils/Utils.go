package utils

import (
	"log"
	"net"
	"net/http"
)

// IPUnspecified Net complement
var IPUnspecified net.IP = net.IPv4(0, 0, 0, 0)

// GetIP https://blog.golang.org/context/userip/userip.go
func GetIP(req *http.Request) (net.IP, string) {
	ip, port, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		log.Panic("Error getIP: ", err)
	}
	return net.ParseIP(ip), port
}
