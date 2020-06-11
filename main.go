package main

import (
	"github.com/gin-gonic/gin"
	"go-sso/api"
	"go-sso/api/user"
	"go-sso/conf"
	"go-sso/modules/app"
	"go-sso/utils/common"
	"go-sso/utils/handle"
	"go-sso/utils/request"
	"go-sso/utils/response"
	"net/url"
	"strconv"
)

func main() {

	//初始化数据
	Load()
	//gin.SetMode(gin.DebugMode)//开发环境
	gin.SetMode(gin.ReleaseMode) //线上环境
	r := gin.Default()
	r.Use(Auth)
	r.POST("/renewal", user.Renewal)
	r.POST("/logout", user.Logout)
	r.POST("/login", user.Login)
	r.POST("/login/mobile", user.LoginByMobileCode)
	r.POST("/sendsms", user.SendSms)
	r.POST("/signup/mobile", user.SignupByMobile)
	r.POST("/signup/mobile/exist", user.MobileIsExists)
	r.GET("/", api.Index)
	r.GET("/my/info", user.Info)//用户信息
	r.GET("/pong", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8282") // listen and serve on 0.0.0.0:8080
}
func Load() {
	c := conf.Config{}
	c.Routes=[]string{"/ping","/renewal","/login","/login/mobile","/sendsms","/signup/mobile","/signup/mobile/exist"}
	c.OpenJwt=true//开启jwt
	conf.Set(c)
	//初始化数据验证
	handle.InitValidate()
}

func Auth(c *gin.Context){
	u,err:= url.Parse(c.Request.RequestURI)
	if err != nil {
		panic(err)
	}
	if common.InArrayString(u.Path,&conf.Cfg.Routes) {
		c.Next()
		return
	}
	//开启jwt
	if conf.Cfg.OpenJwt{
		accessToken,has:=request.GetParam(c,app.ACCESS_TOKEN)
		if !has {
			c.Abort()//组织调起其他函数
			response.ShowError(c,"nologin")
			return
		}
		ret,err:= app.ParseToken(accessToken)
		if err!=nil {
			c.Abort()
			response.ShowValidatorError(c,err.Error())
			return
		}
		uid := strconv.FormatInt(ret.UserId,10)
		has=app.CheckBlack(uid,accessToken)
		if has {
			c.Abort()//组织调起其他函数
			response.ShowError(c,"nologin")
			return
		}
		c.Set("uid",ret.UserId)
		c.Next()
		return
	}
	//cookie
	_,err=c.Cookie(app.COOKIE_TOKEN)
	if err!=nil {
		c.Abort()//组织调起其他函数
		response.ShowError(c,"nologin")
		return
	}
	c.Next()
	return
}
