package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-sso/api"
	"go-sso/api/user"
	"go-sso/conf"
	"go-sso/utils/handle"
)

func main() {

	//初始化数据验证
	handle.InitValidate()

	//gin.SetMode(gin.DebugMode)//开发环境
	gin.SetMode(gin.ReleaseMode) //线上环境
	r := gin.Default()

	//r.POST("/logout", user.Logout)
	//r.POST("/login", user.Login)
	//r.POST("/reg", user.Reg)
	r.POST("/signup/phone", user.SignupByPhone)
	r.GET("/", api.Index)
	r.GET("/pong", func(c *gin.Context) {

		fmt.Println(c.ClientIP())

		//fmt.Println("header \r\n",c.Request.Header)


		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8282") // listen and serve on 0.0.0.0:8080
}
func Load() {
	c := conf.Config{}
	c.Routes=[]string{"/ping","/login"}
	conf.Set(c)
}

func GetClientIp(c *gin.Context) string{
	ip:=c.ClientIP()
	if ip=="::1" {
		ip="127.0.0.1"
	}
	return ip
}