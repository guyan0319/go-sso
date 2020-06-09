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
	secure:=IsHttps(c)
	if conf.Cfg.OpenJwt { //返回jwt
		customClaims :=&CustomClaims{
			UserId:         user.Id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Duration(MAXAGE)*time.Second).Unix(), // 过期时间，必须设置
			},
		}
		accessToken, err :=customClaims.MakeToken()
		if err != nil {
			return err
		}
		refreshClaims :=&CustomClaims{
			UserId:         user.Id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Duration(MAXAGE+1800)*time.Second).Unix(), // 过期时间，必须设置
			},
		}
		refreshToken, err :=refreshClaims.MakeToken()
		if err != nil {
			return err
		}
		c.Header(ACCESS_TOKEN,accessToken)
		c.Header(REFRESH_TOKEN,refreshToken)
		c.SetCookie(ACCESS_TOKEN,accessToken,MAXAGE,"/", "",	 secure,true)
		c.SetCookie(REFRESH_TOKEN,refreshToken,MAXAGE,"/", "",	 secure,true)
	}
	//claims,err:=ParseToken(accessToken)
	//if err!=nil {
	//	return err
	//}
	id:=strconv.Itoa(int(user.Id))
	c.SetCookie(COOKIE_TOKEN,id,MAXAGE,"/", "",	 secure,true)

	return nil
}
//判断是否https
func IsHttps(c *gin.Context) bool {
	if c.GetHeader(HEADER_FORWARDED_PROTO) =="https" || c.Request.TLS!=nil{
		return true
	}
	return false
}