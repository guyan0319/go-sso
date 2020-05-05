package request

import "github.com/gin-gonic/gin"

func GetClientIp(c *gin.Context) string{
	ip:=c.ClientIP()
	if ip=="::1" {
		ip="127.0.0.1"
	}
	return ip
}