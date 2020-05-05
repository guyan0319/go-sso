package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	ip:=GetRemoteIp(r)
	fmt.Println(ip)

	fmt.Fprintf(w, "Hello world!\n") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayHelloHandler) //   设置访问路由
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetRemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}
	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}