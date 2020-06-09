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
$ curl -X POST "http://127.0.0.1:8282/signup/mobile"  -i -d '{"mobile":"1352227564","passwd":"123456","code": "111111"}'
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
$ curl -X POST "http://127.0.0.1:8282/login"  -i -d '{"mobile":"1352227564","passwd":"123456"}'
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    69  100    28  100    41  28000  41000 --:--:-- --:--:-- --:--:-- 69000HTTP/1.1 200 OK
Access_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTc4MTE3OX0.sBDYL0QKwpF9T-3X7NbwZsqBzTakhqYK54D09xyVgso
Content-Type: application/json; charset=utf-8
Rfresh_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTc4Mjk3OX0.NvdyjcsZHpaQQr0hHTWUlo9zTAzOCBvHQUEupsJtje4
Set-Cookie: Access_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTc4MTE3OX0.sBDYL0QKwpF9T-3X7NbwZsqBzTakhqYK54D09xyVgso; Path=/; Max-Age=86400; HttpOnly
Set-Cookie: Rfresh_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTc4Mjk3OX0.NvdyjcsZHpaQQr0hHTWUlo9zTAzOCBvHQUEupsJtje4; Path=/; Max-Age=86400; HttpOnly
Set-Cookie: UserId=1; Path=/; Max-Age=86400; HttpOnly
Date: Tue, 09 Jun 2020 09:26:19 GMT
Content-Length: 28

{"code":200,"msg":"success"}

```

手机号+验证码登录

```
$ curl -X POST "http://127.0.0.1:8282/login/mobile"  -i -d '{"mobile":"1352227564","code":"123456"}'
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    67  100    28  100    39   1750   2437 --:--:-- --:--:-- --:--:--  4187HTTP/1.1 200 OK
Access_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTc4MTY2OX0.xFcbRM1NxohURT0PZ7NCnxoqoZY1-TdgFNnisd_jyMQ
Content-Type: application/json; charset=utf-8
Rfresh_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTc4MzQ2OX0.syOp4W45YH2o4IGOBjvyolo6LpNfjbsNaKtN1PIa_gg
Set-Cookie: Access_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTc4MTY2OX0.xFcbRM1NxohURT0PZ7NCnxoqoZY1-TdgFNnisd_jyMQ; Path=/; Max-Age=86400; HttpOnly
Set-Cookie: Rfresh_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTc4MzQ2OX0.syOp4W45YH2o4IGOBjvyolo6LpNfjbsNaKtN1PIa_gg; Path=/; Max-Age=86400; HttpOnly
Set-Cookie: UserId=1; Path=/; Max-Age=86400; HttpOnly
Date: Tue, 09 Jun 2020 09:34:29 GMT
Content-Length: 28

{"code":200,"msg":"success"}

```

获取用户信息

```
 curl "http://127.0.0.1:8282/my/info" -b "Access_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTc4MTY2OX0.xFcbRM1NxohURT0PZ7NCnxoqoZY1-TdgFNnisd_jyMQ"
```

