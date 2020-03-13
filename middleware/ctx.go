package middleware

import (
	"encoding/binary"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnehcaij/zeus/constant"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

const UNKNOWN_IP_ADDR = "0000"

var (
	localIP           string
	fullLengthLocalIP []byte
)

func Ctx() gin.HandlerFunc {
	localIP = os.Getenv(constant.HOST_IP_ADDR)
	if localIP == "" {
		localIP = getLocalIp()
	}
	elements := strings.Split(localIP, ".")
	for i := 0; i < len(elements); i++ {
		elements[i] = fmt.Sprintf("%03s", elements[i])
	}
	fullLengthLocalIP = []byte(strings.Join(elements, ""))

	return func(c *gin.Context) {
		c.Set(constant.LOG_ID, genLogId())
		c.Set(constant.LOCAL_IP_KEY, localIP)
		c.Next()
	}
}

// genLogId generates a global unique log id for request
// format: %Y%m%d%H%M%S + ip + 5位随机数
func genLogId() string {
	buf := make([]byte, 0, 64)
	buf = time.Now().AppendFormat(buf, "20060102150405")
	buf = append(buf, fullLengthLocalIP...)

	uuidBuf := make([]byte, 4)
	_, err := rand.Read(uuidBuf)
	if err != nil {
		panic(err)
	}
	uuidNum := binary.BigEndian.Uint32(uuidBuf)
	buf = append(buf, fmt.Sprintf("%05d", uuidNum)[:5]...)
	return string(buf)
}

func getLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return UNKNOWN_IP_ADDR
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return UNKNOWN_IP_ADDR
}
