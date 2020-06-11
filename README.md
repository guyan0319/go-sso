# golang 单点登录系统（go-sso）

这是一个基于Go语言开发的单点登录系统，实现手机号注册、手机号+验证码登录、手机号+密码登录、账号登出等功能，用户认证采用cookie和jwt两种方式。收发短信相关方法已提供，仅需根据短信通道提供商提供的接口做相应的参数配置即可使用。

## 环境介绍

golang语言：go1.13.3+ 、

数据库：mysql5.7

缓存：redis3.0

## 项目地址

https://github.com/guyan0319/go-sso

## 依赖包：

```
github.com/dgrijalva/jwt-go
github.com/gin-gonic/gin
github.com/go-xorm/xorm
github.com/go-sql-driver/mysql
```

注意：项目代码依赖管理工具采用的[go-modules](https://github.com/guyan0319/golang_development_notes/blob/master/zh/1.10.md)，需要了解的移步https://github.com/guyan0319/golang_development_notes/blob/master/zh/1.10.md

## 数据请求类型

```
Content-Type: application/json
```

注：这个一定要注意，其他类型服务端不识别。

## 快速开始

### 1、数据库

sql文件在data目录下，新建数据库名，通过mysql管理工具或登录mysql数据库直接创建数据库名，并导入ssodb.sql文件到数据库，完成数据库结构创建。具体方法这里不再赘述。

### 2、配置mysql、redis

配置文件在conf目录下,修改 mysql.go、redis.go配置成你自己的实际环境。

### 3、启动

```
 go  run  main.go
```

### 4、测试

我们可以采用postman等工具，进行客户端浏览器模拟操作，也可以通过curl实现。

这里以curl工具为例：

##### 注册手机号

```
$ curl -X POST "http://127.0.0.1:8282/signup/mobile"  -i -d '{"mobile":"13522227564","passwd":"123456","code": "111111"}'
```

结果

```
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    86  100    28  100    58   1750   3625 --:--:-- --:--:-- --:--:-- 86000HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 09 Jun 2020 08:59:36 GMT
Content-Length: 28

{"code":200,"msg":"success"}
```

注：code参数值只要随意6位数即可，示例中服务端并没有配置短信通道，验证code代码已注释。

##### 手机号+密码 登录

```
$ curl -X POST "http://127.0.0.1:8282/login"  -i -d '{"mobile":"13522227564","passwd":"123456"}'
```

结果

```
 % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    69  100    28  100    41    451    661 --:--:-- --:--:-- --:--:--  1112HTTP/1.1 200 OK
Access_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTg1ODIwMX0.OMpRhdX2wXm1BuRmeaZtSH7L3skXhNhYc0YYUTQf7WI
Content-Type: application/json; charset=utf-8
Refresh_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTg2MDAwMX0.Hglo9i8pxFGcRcvMw8mUvXlT2JLrHSC6ocQqJseknq8
Set-Cookie: Access_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTg1ODIwMX0.OMpRhdX2wXm1BuRmeaZtSH7L3skXhNhYc0YYUTQf7WI; Path=/; Max-Age=86400; HttpOnly
Set-Cookie: Refresh_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTg2MDAwMX0.Hglo9i8pxFGcRcvMw8mUvXlT2JLrHSC6ocQqJseknq8; Path=/; Max-Age=86400; HttpOnly
Set-Cookie: UserId=1; Path=/; Max-Age=86400; HttpOnly
Date: Wed, 10 Jun 2020 06:50:01 GMT
Content-Length: 28

{"code":200,"msg":"success"}
```

##### 手机号+验证码登录

```
$  curl -X POST "http://127.0.0.1:8282/login/mobile"  -i -d '{"mobile":"13522227564","code":"123456"}'
```

结果

```
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    68  100    28  100    40    163    233 --:--:-- --:--:-- --:--:--   397HTTP/1.1 200 OK
Access_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTk1NDIxOX0.FK-6ofW-ni8E7BcQ1tH9Z4vzQHDBbnyIcDZLEytRrfQ
Content-Type: application/json; charset=utf-8
Refresh_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTk1NjAxOX0.gJzzho2gP1nNrkUKdtvMh0R3jGFZpA-ku0dWDvLftu0
Set-Cookie: Access_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTk1NDIxOX0.FK-6ofW-ni8E7BcQ1tH9Z4vzQHDBbnyIcDZLEytRrfQ; Path=/; Max-Age=86400; HttpOnly
Set-Cookie: Refresh_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTk1NjAxOX0.gJzzho2gP1nNrkUKdtvMh0R3jGFZpA-ku0dWDvLftu0; Path=/; Max-Age=86400; HttpOnly
Set-Cookie: UserId=1; Path=/; Max-Age=86400; HttpOnly
Date: Thu, 11 Jun 2020 09:30:19 GMT
Content-Length: 28

{"code":200,"msg":"success"}
```

##### 获取用户信息

使用上面登录获取的Access_Token获取用户信息。

```
$ curl "http://127.0.0.1:8282/my/info" -b "Access_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTk1NDIxOX0.FK-6ofW-ni8E7BcQ1tH9Z4vzQHDBbnyIcDZLEytRrfQ"
```

结果

```
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    73  100    73    0     0  73000      0 --:--:-- --:--:-- --:--:-- 73000{"code":200,"data":{"id":1,"name":"","email":"","mobile":"135****27564"}}
```

##### 查看手机号是否存在

```
$  curl -X POST "http://127.0.0.1:8282/signup/mobile/exist"  -i -d '{"mobile":"13522227564"}'
```

结果

```
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    61  100    37  100    24  37000  24000 --:--:-- --:--:-- --:--:-- 61000HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Wed, 10 Jun 2020 07:40:50 GMT
Content-Length: 37

{"code":200,"data":{"is_exist":true}}
```

##### access token 续期

```
$ curl -X POST "http://127.0.0.1:8282/renewal" -i -b "Access_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTk1NDIxOX0.FK-6ofW-ni8E7BcQ1tH9Z4vzQHDBbnyIcDZLEytRrfQ;Rfresh_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTk1NjAxOX0.gJzzho2gP1nNrkUKdtvMh0R3jGFZpA-ku0dWDvLftu0"
```

结果：

```
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    28  100    28    0     0    595      0 --:--:-- --:--:-- --:--:--   595HTTP/1.1 200 OK
Access_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTk1NTEzMn0.KCNVQ39HoMZFG-Xl0xwDrVdhQO_w4-tDxWY0ebyhfyk
Content-Type: application/json; charset=utf-8
Refresh_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTk1NjkzMn0.XijN8ZjPbZkhjJB33igti3dSm1tfd2kp_iNRtPqpwyM
Set-Cookie: Access_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTk1NTEzMn0.KCNVQ39HoMZFG-Xl0xwDrVdhQO_w4-tDxWY0ebyhfyk; Path=/; Max-Age=86400; HttpOnly
Set-Cookie: Refresh_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTk1NjkzMn0.XijN8ZjPbZkhjJB33igti3dSm1tfd2kp_iNRtPqpwyM; Path=/; Max-Age=86400; HttpOnly
Date: Thu, 11 Jun 2020 09:45:32 GMT
Content-Length: 28

{"code":400,"msg":"success"}

```

##### 登出系统

```
$ curl -X POST "http://127.0.0.1:8282/logout" -b "Access_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU5MTk1NDIxOX0.FK-6ofW-ni8E7BcQ1tH9Z4vzQHDBbnyIcDZLEytRrfQ"
```

结果

```
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    28  100    28    0     0  28000      0 --:--:-- --:--:-- --:--:-- 28000{"code":200,"msg":"success"}
```

### 小结

仅使用jwt实现单点登录会遇到两个问题

- 用户无法主动登出，即服务端发出token后，无法主动销毁token，用户还可以用通过token访问系统，本项目增加了缓存登出用户token到黑名单的方式，变相实现登出。
- token续期问题，access_token携带有效期，有效期过了无法自动续期。本项目提供了续期接口（renewal），服务端在生成access_token同时还会生成refresh_token（有效期比access_token长），用户可以通过有效的refresh_token和access_token访问renewal接口重新获取新的refresh_token和access_token。

在使用过程如遇到任何问题，欢迎[issues](https://github.com/guyan0319/go-sso/issues)。