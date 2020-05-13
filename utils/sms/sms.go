package sms

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go-sso/utils/cache"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"unicode/utf8"
)

const SMSTPL  ="【xxxx】您正在申请手机注册，验证码为：[code]，若非本人操作请忽略！"
type SmsSeting struct {




}
func SmsCheck(key, code string) bool {
	key = cache.RedisSuf + key
	// 从池里获取连接
	rc := cache.RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	val, err := redis.String(rc.Do("GET", key))
	if err != nil || val != code {
		return false
	}
	return true
}
func SmsSet(key,val string)(err error)  {
	key = cache.RedisSuf + key
	// 从池里获取连接
	rc := cache.RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	_, err= rc.Do("Set", key, val, "EX", 600)
	if err != nil {
		return
	}
	return
}

func HttpPostForm(url string, data url.Values) (string, error) {

	resp, err := http.PostForm(url, data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

//发送短信
func SendSms(c map[string]string, mobile, msg string) error {
	if mobile == "" {
		return errors.New("mobile is not null")
	}
	reg := `^1\d{10}$`
	if c["internation"] == "1" {
		reg = `^(00){1}\d+`
	}
	rgx := regexp.MustCompile(reg)
	if !rgx.MatchString(mobile) {
		return errors.New("mobile is irregular")
	}
	if utf8.RuneCountInString(msg) < 10 {
		return errors.New("Character length is not enough.")
	}

	url_send := c["url"]
	data_send := url.Values{"sname": {c["sname"]}, "spwd": {c["spwd"]}, "scorpid": {c["scorpid"]}, "sprdid": {c["sprdid"]}, "sdst": {mobile}, "smsg": {msg}}
	re, err := HttpPostForm(url_send, data_send)
	fmt.Println(re, "zhe")
	return err
}
