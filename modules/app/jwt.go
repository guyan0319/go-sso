package app

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"go-sso/utils/cache"
)

const (
	SECRETKEY = "42wqTE23123wffLU94342wgldgFs"
	MAXAGE=3600*24
	CACHE_BLACK_TOKEN="black.token."
)

type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}
//产生token
func(cc *CustomClaims ) MakeToken() (string,error) {
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, cc)
	return token.SignedString([]byte(SECRETKEY))
}
//解析token
func ParseToken(tokenString string)(*CustomClaims,error)  {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims,nil
	} else {
		return nil,err
	}
}
//加入到黑名单
func AddBlack( key , token string ) (err error) {
	key = cache.RedisSuf + CACHE_BLACK_TOKEN+ key
	// 从池里获取连接
	rc := cache.RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	_, err= rc.Do("Set", key, token, "EX", MAXAGE)
	if err != nil {
		return
	}
	return
}
//检查token是否在黑名单
func CheckBlack(key,token string)bool  {
	key = cache.RedisSuf + CACHE_BLACK_TOKEN+ key
	// 从池里获取连接
	rc := cache.RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	val, err := redis.String(rc.Do("GET", key))
	if err != nil || val != token {
		return false
	}
	return true
}
