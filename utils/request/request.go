package request

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func GetClientIp(c *gin.Context) string{
	ip:=c.ClientIP()
	if ip=="::1" {
		ip="127.0.0.1"
	}
	return ip
}
func GetJson(c *gin.Context) (map[string]interface{},error ){
	jsonstr, _ := ioutil.ReadAll(c.Request.Body)
	var data map[string]interface{}
	err := json.Unmarshal(jsonstr, &data)
	return data,err
}
func GetParam(c *gin.Context,key string)(string,bool){
	val:=c.GetHeader(key)
	if val!=""{
		return val,true
	}
	val,err :=c.Cookie(key)
	if err!=nil{
		return "",false
	}
	return val,true
}

