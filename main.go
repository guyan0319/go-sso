package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-sso/api"
	"go-sso/api/user"
	"go-sso/conf"
	"go-sso/modules/app"
	"go-sso/utils/handle"
	"log"
)

func main() {

	//初始化数据验证
	handle.InitValidate()
	//gin.SetMode(gin.DebugMode)//开发环境
	gin.SetMode(gin.ReleaseMode) //线上环境
	r := gin.Default()
	r.Use(Auth)
	//r.POST("/logout", user.Logout)
	r.POST("/login/mobile", user.LoginByMobile)
	r.POST("/login/mobile", user.LoginByMobileCode)
	r.POST("/sendsms", user.SendSms)
	r.POST("/signup/mobile", user.SignupByMobile)
	r.GET("/", api.Index)
	r.GET("/pong", func(c *gin.Context) {
		fmt.Println(c.Request.TLS)
		fmt.Println(c.Request.Proto)
		fmt.Println(	c.GetHeader(app.HEADER_FORWARDED_PROTO))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8282") // listen and serve on 0.0.0.0:8080
}
func Load() {
	c := conf.Config{}
	c.Routes=[]string{"/ping","/login"}
	c.OpenJwt=true//开启jwt
	conf.Set(c)
}

func Auth(c *gin.Context){
	//u,err:= url.Parse(c.Request.RequestURI)
	//if err != nil {
	//	panic(err)
	//}
	//if common.InArrayString(u.Path,&conf.Cfg.Routes) {
	//	//c.Next()
	//	return
	//}
	//session := sessions.Default(c)
	//v := session.Get(conf.Cfg.Token)
	//if v==nil {
	//	c.Abort()
	//	response.ShowError(c,"nologin")
	//	return
	//}
	//uid:=session.Get(v)
	//users := models.SystemUser{Id:uid.(int),Status:1}
	//has:=users.GetRow()
	//if !has {
	//	c.Abort()
	//	response.ShowError(c,"user_error")
	//	return
	//}
	////特殊账号
	//if users.Name==conf.Cfg.Super {
		c.Next()
		return
	//}
	//menuModel:=models.SystemMenu{}
	//menuMap,err:=menuModel.GetRouteByUid(uid)
	//if err!=nil {
	//	c.Abort()
	//	response.ShowError(c,"unauthorized")
	//	return
	//}
	//if _,ok:=menuMap[u.Path] ;!ok{
	//	c.Abort()
	//	response.ShowError(c,"unauthorized")
	//	return
	//}
	// access the status we are sending
	status := c.Writer.Status()
	log.Println(status) //状态 200
}
