package main

import (
	"fmt"
	"go-sso/utils/common"
)

func main() {
	ip:="13.13.13.12"
	fmt.Println(common.Sha1En(ip))
	fmt.Println(common.Sha1En1(ip))
	fmt.Println(common.Md5En(ip))
}

