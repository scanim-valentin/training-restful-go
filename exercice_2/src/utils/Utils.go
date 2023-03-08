package utils

import (
	"log"
	"math/rand"
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

var charset = []byte("azertyuiopqsdfghjklmwxcvbnAZERTYUIOPQSDFHJKLMWXCVBN0123456789")

func RandomString(nbMaxChar int) []byte {
	nbChar := rand.Intn(nbMaxChar)
	content := make([]byte, 0)
	for k := 0; k < nbChar; k++ {
		content = append(content, charset[rand.Intn(len(charset)-1)])
	}
	return content
}
