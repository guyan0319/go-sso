package main

import (
	"github.com/gin-gonic/gin"
	"go-sso/api"
	"go-sso/api/user"
	"go-sso/conf"
	"go-sso/modules/app"
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

		tokenString :="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjcsImV4cCI6MTU4OTIxNTI4NH0.JAeukEGhUFIhEsFjJ12UGHsMsBGW1xYhqRuRlyMlmRc"
		app.ParseToken(tokenString)
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

