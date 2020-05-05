package main

import (
	"fmt"
	"go-sso/utils/common"
)

func main() {
	ip:="13.13.13.12"
	intv:=common.IpStringToInt(ip)
	fmt.Println(intv)
	fmt.Println(common.IpStringToInt(ip))


	fmt.Println(common.IpIntToString(intv))
}

