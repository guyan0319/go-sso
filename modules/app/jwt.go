package app

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

const (
	SECRETKEY = "42wqTE23123wffLU94342wgldgFs"
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
