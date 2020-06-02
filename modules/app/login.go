package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-sso/conf"
	"go-sso/models"
	"strconv"
	"time"
)

func DoLogin(c *gin.Context,user models.Users)  error{
	maxAge:=60*60*24
	if conf.Cfg.OpenJwt { //返回jwt
		customClaims :=&CustomClaims{
			UserId:         user.Id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Duration(maxAge)*time.Second).Unix(), // 过期时间，必须设置
			},
		}
		accessToken, err :=customClaims.MakeToken()
		if err != nil {
			return err
		}
		c.Header(ACCESS_TOKEN,accessToken)
	}
	//claims,err:=ParseToken(accessToken)
	//if err!=nil {
	//	return err
	//}
	id:=strconv.Itoa(int(user.Id))
	secure:=IsHttps(c)
	c.SetCookie(COOKIE_TOKEN,id,maxAge,"/", "",	 secure,true)
	return nil
}
func IsHttps(c *gin.Context) bool {
	if c.GetHeader(HEADER_FORWARDED_PROTO) =="https" || c.Request.TLS!=nil{
		return true
	}
	return false
}