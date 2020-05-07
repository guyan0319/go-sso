package sms

import (
	"github.com/gomodule/redigo/redis"
	"go-sso/utils/cache"
)

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