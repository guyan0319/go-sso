go-sso

这是一个基于Go语言开发的单点登录系统，实现手机号注册、手机号+验证码登录、手机号+密码登录、账号登出等功能，用户认证采用cookie和jwt两种方式。收发短信相关方法已提供，仅需根据短信通道提供商提供的接口做相应的参数配置即可使用。

环境介绍

golang语言：go1.13.3+ 、

数据库：mysql5.7

缓存：redis3.0

依赖包：

```
github.com/dgrijalva/jwt-go
github.com/gin-gonic/gin
github.com/go-xorm/xorm
github.com/go-sql-driver/mysql
```

注意：项目代码依赖管理工具采用的[go-modules](https://github.com/guyan0319/golang_development_notes/blob/master/zh/1.10.md)，需要了解的移步https://github.com/guyan0319/golang_development_notes/blob/master/zh/1.10.md

数据请求类型

```
Content-Type: application/json
```

注：这个一定要注意，其他类型服务端不识别。

快速开始

1、数据库

sql文件在data目录下，新建数据库名，通过mysql管理工具或登录mysql数据库直接创建数据库名，并导入ssodb.sql文件到数据库，完成数据库结构创建。

2、配置mysql、redis

配置文件在conf目录下,修改 mysql.go、redis.go配置成你自己的实际环境。

3、启动

```
 go  run  main.go
```

4、测试

这里我们可以采用postman等工具，进行客户端浏览器模拟操作，也可以通过curl实现。

本文以curl工具为例

注册手机号

```
$ curl -X POST "http://127.0.0.1:8282/signup/mobile"  -i -d '{"mobile":"13522227564","passwd":"123456","code": "111111"}'
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    86  100    28  100    58   1750   3625 --:--:-- --:--:-- --:--:-- 86000HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 09 Jun 2020 08:59:36 GMT
Content-Length: 28

{"code":200,"msg":"success"}
```

注：code参数值只要随意6位数即可，示例中服务端并没有配置短信通道，验证code代码已注释。

手机号+密码 登录

```
$ curl -X POST "http://127.0.0.1:8282/login"  -i -d '{"mobile":"13522227564","passwd":"123456"}'
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    69  100    28  100    41    451    661 --:--:-- --:--:-- --:--:--  1112HTTP/1.1 200 OK
Access_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTg1ODIwMX0.OMpRhdX2wXm1BuRmeaZtSH7L3skXhNhYc0YYUTQf7WI
Content-Type: application/json; charset=utf-8
Rfresh_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTg2MDAwMX0.Hglo9i8pxFGcRcvMw8mUvXlT2JLrHSC6ocQqJseknq8
Set-Cookie: Access_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTg1ODIwMX0.OMpRhdX2wXm1BuRmeaZtSH7L3skXhNhYc0YYUTQf7WI; Path=/; Max-Age=86400; HttpOnly
Set-Cookie: Rfresh_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTg2MDAwMX0.Hglo9i8pxFGcRcvMw8mUvXlT2JLrHSC6ocQqJseknq8; Path=/; Max-Age=86400; HttpOnly
Set-Cookie: UserId=1; Path=/; Max-Age=86400; HttpOnly
Date: Wed, 10 Jun 2020 06:50:01 GMT
Content-Length: 28

{"code":200,"msg":"success"}

```

手机号+验证码登录

```
$ curl -X POST "http://127.0.0.1:8282/login/mobile"  -i -d '{"mobile":"13522227564","code":"123456"}'
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    67  100    28  100    39  28000  39000 --:--:-- --:--:-- --:--:-- 67000HTTP/1.1 200 OK
Access_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTg1ODI1M30.-Yf8I1W-0lEjgVx19IYPoEZhMZ-hChlcA51u3k56hQI
Content-Type: application/json; charset=utf-8
Rfresh_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTg2MDA1M30.X1liLjGyJkTAnwag6unTZpbtxXX_9JQWGdp7q4II8ek
Set-Cookie: Access_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTg1ODI1M30.-Yf8I1W-0lEjgVx19IYPoEZhMZ-hChlcA51u3k56hQI; Path=/; Max-Age=86400; HttpOnly
Set-Cookie: Rfresh_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTg2MDA1M30.X1liLjGyJkTAnwag6unTZpbtxXX_9JQWGdp7q4II8ek; Path=/; Max-Age=86400; HttpOnly
Set-Cookie: UserId=1; Path=/; Max-Age=86400; HttpOnly
Date: Wed, 10 Jun 2020 06:50:53 GMT
Content-Length: 28

{"code":200,"msg":"success"}


```

获取用户信息

```
$  curl "http://127.0.0.1:8282/my/info" -b "Access_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTc4MTY2OX0.xFcbRM1NxohURT0PZ7NCnxoqoZY1-TdgFNnisd_jyMQ"
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    72  100    72    0     0    774      0 --:--:-- --:--:-- --:--:--   774{"code":200,"data":{"id":1,"name":"","email":"","mobile":"135****7564"}}
```

查看手机号是否存在

```
$  curl -X POST "http://127.0.0.1:8282/signup/mobile/exist"  -i -d '{"mobile":"13522227564"}'
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    61  100    37  100    24  37000  24000 --:--:-- --:--:-- --:--:-- 61000HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Wed, 10 Jun 2020 07:40:50 GMT
Content-Length: 37

{"code":200,"data":{"is_exist":true}}

```

登出系统

```
curl -X POST "http://127.0.0.1:8282/logout" -b "Access_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTc4MTY2OX0.xFcbRM1NxohURT0PZ7NCnxoqoZY1-TdgFNnisd_jyMQ"
```

