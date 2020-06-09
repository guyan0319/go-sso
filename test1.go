package main

import (
	"fmt"
	"go-sso/modules/app"
)

func main() {

	tokenStr:="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTc4MTY2OX0.xFcbRM1NxohURT0PZ7NCnxoqoZY1-TdgFNnisd_jyMQ"
	val,_:=app.ParseToken(tokenStr)
	fmt.Println(val)
	fmt.Println(val.Id)
}


