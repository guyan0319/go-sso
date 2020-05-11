package app

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SecretKey = "42wqTE23123wffLU94342wgldgFs"
)

type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}

func(cc CustomClaims ) MakeToken() (string,error) {
	cc.StandardClaims.ExpiresAt=time.Now().Add(60*60*time.Second).Unix() // 过期时间，必须设置
	//cc.StandardClaims.Issuer=""//可自定义
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, cc)
	return token.SignedString([]byte(SecretKey))
}
func ParseToken(tokenString string)  {
	//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	//	// Don't forget to validate the alg is what you expect:
	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	//	}
	//	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	//	return SecretKey, nil
	//})
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		fmt.Println(claims)

	} else {
		fmt.Println("aaa")
		fmt.Println(err)
	}






}
