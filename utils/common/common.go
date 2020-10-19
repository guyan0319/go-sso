package common

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/rand"
	crand "crypto/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

//获取随机数 纯文字
func GetRandomString(n int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//获取随机数  数字和文字
func GetRandomBoth(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//获取随机数  纯数字
func GetRandomNum(n int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}


//获取随机数  base32
func GetRandomBase32(n int) string {
	str := "234567abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//生成区间随机数
func RandInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}

//sha1加密
func Sha1En(data string) string {
	t := sha1.New()///产生一个散列值得方式
	_,_=io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//对字符串进行MD5哈希
func Md5En(data string) string {
	t := md5.New()
	_,_=io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//自定义唯一id
func GetUniqueId() string {
	cur := time.Now()
	timestamps := cur.UnixNano()
	uid := strconv.FormatInt(timestamps, 10) + GetRandomNum(5)
	return Md5En(uid)
}

//自定义唯一id
func OrderUniqueId() string {
	cur := time.Now()
	timestamps := cur.UnixNano() / 1000000 //获取毫秒
	return strconv.FormatInt(timestamps, 10) + GetRandomNum(5)
}
//查找某值是否在数组中
func InArrayString(v string, m *[]string) bool {
	for _, value := range *m {
		if value == v {
			return true
		}
	}
	return false
}
func IpStringToInt(ipstring string) int {
	if net.ParseIP(ipstring)==nil {
		return 0
	}
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}

func IpIntToString(ipInt int) string{
	var bytes [4]byte
	bytes[0] = byte(ipInt & 0xFF)
	bytes[1] = byte((ipInt >> 8) & 0xFF)
	bytes[2] = byte((ipInt >> 16) & 0xFF)
	bytes[3] = byte((ipInt >> 24) & 0xFF)
	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0]).String()
}

// 生成区间[-m, n]的安全随机数
func RangeRand(min, max int64) int64 {
	if min > max {
		panic("the min is greater than max!")
	}
	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := crand.Int(crand.Reader, big.NewInt(max+1+i64Min))

		return result.Int64() - i64Min
	} else {
		result, _ := crand.Int(crand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}
//整数区间的随机数
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}

	return rand.Int63n(max-min) + min
}
