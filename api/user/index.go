package user

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-sso/models"
	"go-sso/modules/app"
	"go-sso/utils/common"
	"go-sso/utils/handle"
	"go-sso/utils/request"
	"go-sso/utils/response"
	"go-sso/utils/sms"
	"go-sso/utils/verify"
	"strconv"
	"strings"
	"time"
)

type UserMobile struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
	Passwd string `form:"passwd" json:"passwd" binding:"required,max=20,min=6"`
	Code   string `form:"code" json:"code" binding:"required,len=6"`
}
type UserMobileCode struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
	Code   string `form:"code" json:"code" binding:"required,len=6"`
}

type UserMobilePasswd struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
	Passwd string `form:"passwd" json:"passwd" binding:"required,max=20,min=6"`
}

type Mobile struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
}

var MobileTrans = map[string]string{"mobile": "手机号"}

var UserMobileTrans = map[string]string{"Mobile": "手机号", "Passwd": "密码", "Code": "验证码"}

//手机密码
func Login(c *gin.Context) {
	var userMobile UserMobilePasswd
	if err := c.BindJSON(&userMobile); err != nil {
		msg := handle.TransTagName(&UserMobileTrans, err)
		response.ShowValidatorError(c, msg)
		return
	}
	model := models.Users{Mobile: userMobile.Mobile}
	if has := model.GetRow(); !has {
		response.ShowError(c, "mobile_not_exists")
		return
	}
	if common.Sha1En(userMobile.Passwd+model.Salt) != model.Passwd {
		response.ShowError(c, "login_error")
		return
	}
	err := app.DoLogin(c, model)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	response.ShowSuccess(c, "success")
	return
}

//注销登录
func Logout(c *gin.Context) {
	secure := app.IsHttps(c)
	//access_token  refresh_token 加黑名单
	accessToken, has := request.GetParam(c, app.ACCESS_TOKEN)
	if has {
		uid := strconv.FormatInt(c.MustGet("uid").(int64),10)
		app.AddBlack(uid, accessToken)
	}
	c.SetCookie(app.COOKIE_TOKEN, "", -1, "/", "", secure, true)
	c.SetCookie(app.ACCESS_TOKEN, "", -1, "/", "", secure, true)
	c.SetCookie(app.REFRESH_TOKEN, "", -1, "/", "", secure, true)
	response.ShowSuccess(c, "success")
	return
}

//手机验证码登录
func LoginByMobileCode(c *gin.Context) {
	var userMobile UserMobileCode
	if err := c.BindJSON(&userMobile); err != nil {
		msg := handle.TransTagName(&UserMobileTrans, err)
		fmt.Println(msg)
		response.ShowValidatorError(c, msg)
		return
	}
	//验证code
	//if sms.SmsCheck("code"+userMobile.Mobile, userMobile.Code) {
	//	response.ShowError(c, "code_error")
	//	return
	//}
	model := models.Users{Mobile: userMobile.Mobile}
	if has := model.GetRow(); !has {
		response.ShowError(c, "mobile_not_exists")
		return
	}
	err := app.DoLogin(c, model)
	if err != nil {
		fmt.Println(err)
		response.ShowError(c, "fail")
		return
	}
	response.ShowSuccess(c, "success")
	return
}
func MobileIsExists(c *gin.Context) {
	var userMobile Mobile
	if err := c.BindJSON(&userMobile); err != nil {
		msg := handle.TransTagName(&UserMobileTrans, err)
		fmt.Println(msg)
		response.ShowValidatorError(c, msg)
		return
	}
	if !verify.CheckMobile(userMobile.Mobile) {
		response.ShowError(c, "mobile_error")
		return
	}
	model := models.Users{Mobile: userMobile.Mobile}
	var data = map[string]bool{"is_exist": true}
	if has := model.GetRow(); !has {
		data["is_exist"] = false
	}
	response.ShowData(c, data)
	return
}

//发送短信验证码
func SendSms(c *gin.Context) {
	var p Mobile
	if err := c.BindJSON(&p); err != nil {
		msg := handle.TransTagName(&MobileTrans, err)
		response.ShowValidatorError(c, msg)
		return
	}
	if !verify.CheckMobile(p.Mobile) {
		response.ShowError(c, "mobile_error")
		return
	}
	//生成随机数
	code := common.GetRandomNum(6)
	msg := strings.Replace(sms.SMSTPL, "[code]", code, 1)
	err := sms.SendSms(p.Mobile, msg)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	response.ShowError(c, "success")
	return

}

//手机号注册
func SignupByMobile(c *gin.Context) {
	var userMobile UserMobile
	if err := c.BindJSON(&userMobile); err != nil {
		msg := handle.TransTagName(&UserMobileTrans, err)
		fmt.Println(msg)
		response.ShowValidatorError(c, msg)
		return
	}
	model := models.Users{Mobile: userMobile.Mobile}
	if has := model.GetRow(); has {
		response.ShowError(c, "mobile_exists")
		return
	}
	//验证code
	//if sms.SmsCheck("code"+userMobile.Mobile,userMobile.Code) {
	//	response.ShowError(c, "code_error")
	//	return
	//}

	model.Salt = common.GetRandomBoth(4)
	model.Passwd = common.Sha1En(userMobile.Passwd + model.Salt)
	model.Ctime = int(time.Now().Unix())
	model.Status = models.UsersStatusOk
	model.Mtime = time.Now()

	traceModel := models.Trace{Ctime: model.Ctime}
	traceModel.Ip = common.IpStringToInt(request.GetClientIp(c))
	traceModel.Type = models.TraceTypeReg

	deviceModel := models.Device{Ctime: model.Ctime, Ip: traceModel.Ip, Client: c.GetHeader("User-Agent")}
	_, err := model.Add(&traceModel, &deviceModel)
	if err != nil {
		fmt.Println(err)
		response.ShowError(c, "fail")
		return
	}
	response.ShowSuccess(c, "success")
	return
}

//access token 续期
func Renewal(c *gin.Context) {
	accessToken, has := request.GetParam(c, app.ACCESS_TOKEN)
	if !has {
		response.ShowValidatorError(c, "access token not found")
		return
	}
	refreshToken, has := request.GetParam(c, app.REFRESH_TOKEN)
	if !has {
		response.ShowError(c, "refresh_token")
		return
	}
	ret, err := app.ParseToken(refreshToken)
	if err != nil {
		response.ShowError(c, "refresh_token")
		return
	}
	//uid := strconv.FormatInt(ret.UserId,10)
	//has=app.CheckBlack(uid,accessToken)
	//if has {
	//	c.Abort()//组织调起其他函数
	//	response.ShowError(c,"nologin")
	//	return
	//}
	//_, err= app.ParseToken(accessToken)
	//if err == nil {
	//	response.ShowError(c, "access_token_ok")
	//	return
	//}
	customClaims := &app.CustomClaims{
		UserId: ret.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(app.MAXAGE) * time.Second).Unix(), // 过期时间，必须设置
		},
	}
	accessToken, err = customClaims.MakeToken()
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	customClaims = &app.CustomClaims{
		UserId: ret.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(app.MAXAGE+1800) * time.Second).Unix(), // 过期时间，必须设置
		},
	}
	refreshToken, err = customClaims.MakeToken()
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	c.Header(app.ACCESS_TOKEN, accessToken)
	c.Header(app.REFRESH_TOKEN, refreshToken)
	secure := app.IsHttps(c)
	c.SetCookie(app.ACCESS_TOKEN, accessToken, app.MAXAGE, "/", "", secure, true)
	c.SetCookie(app.REFRESH_TOKEN, refreshToken, app.MAXAGE, "/", "", secure, true)
	fmt.Println("ok")
	response.ShowError(c, "success")
	return
}
func Info(c *gin.Context)  {
	uid:=c.MustGet("uid").(int64)
	fmt.Println(uid)
	model:=models.Users{}
	model.Id=uid
	row,err:=model.GetRowById()
	if err!=nil {
		fmt.Println(err)
		response.ShowValidatorError(c, err)
		return
	}
	fmt.Println(row)
	fmt.Println(row.Name)
	//隐藏手机号中间数字
	s :=row.Mobile
	row.Mobile =string([]byte(s)[0:3])+"****"+string([]byte(s)[6:])
	response.ShowData(c,row)
	return
}
