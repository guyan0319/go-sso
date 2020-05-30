package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)
const (
	SECRETKEY = "243223ffslsfsldfl412fdsfsdf"//私钥
)
//自定义Claims
type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}
func main() {
	//生成token
	maxAge:=60*60*24

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(maxAge)*time.Second).Unix(), // 过期时间，必须设置,
		Issuer:    "jerry",// 非必须，也可以填充用户名，
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SECRETKEY)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Printf("token: %v\n", tokenString)

	//解析token
	ret,err :=ParseToken(tokenString)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Printf("userinfo: %v\n", ret)
}

//解析token
func ParseToken(tokenString string)(*CustomClaims,error)  {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims,nil
	} else {
		return nil,err
	}
}
