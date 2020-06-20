package main

import (
	"fmt"
	"strconv"
)

func main() {

	var  i int64
	
	i=111
	strings:=strconv.FormatInt(i,10)
	fmt.Println(strings)
}


