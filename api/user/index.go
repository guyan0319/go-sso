package user

import (
	"github.com/gin-gonic/gin"
	"go-sso/models"
	"go-sso/utils/common"
	"go-sso/utils/handle"
	"go-sso/utils/response"
	"time"
)
type UserPhone struct {
	Phone   string `form:"phone" json:"phone" binding:"required"`
	Passwd   string `form:"passwd" json:"passwd" binding:"required,max=20,min=6"`
	Code   string `form:"code" json:"code" binding:"required,len=6"`
}
var  UserPhoneTrans =map[string]string{"Phone":"手机号","Passwd":"密码","Code":"验证码"}
func Login(c *gin.Context) {

	//var u User
	//err :=c.BindJSON(&u)
	//if err!=nil	{
	//	response.ShowError(c, "fail")
	//	return
	//}
	//if u.Username == "" || u.Password == "" {
	//	response.ShowError(c, "fail")
	//	return
	//}
	//user := models.SystemUser{Name: u.Username}
	//has := user.GetRow()
	//if !has {
	//	response.ShowError(c, "fail")
	//	return
	//}
	//if common.Sha1En(u.Password+user.Salt) != user.Password {
	//	response.ShowError(c, "fail")
	//	return
	//}
	//session := sessions.Default(c)
	//var data = make(map[string]interface{}, 0)
	//v := session.Get(conf.Cfg.Token)
	//fmt.Println(v)
	//if v == nil {
	//	cur := time.Now()
	//	//纳秒
	//	timestamps := cur.UnixNano()
	//	times := strconv.FormatInt(timestamps, 10)
	//	v = common.Md5En(common.GetRandomString(16) + times)
	//	session.Set(conf.Cfg.Token, v)
	//	session.Set(v, user.Id)
	//	err=session.Save()
	//	fmt.Println("设置成功")
	//}
	//data[conf.Cfg.Token] = v
	//response.ShowData(c, data)
	return
}
func SignupByPhone(c *gin.Context)  {
	var userPhone UserPhone
	if err := c.BindJSON(&userPhone); err != nil {
		msg:=handle.TransTagName(&UserPhoneTrans,err)
		response.ShowValidatorError(c,msg)
		return
	}
	model:=models.Users{Phone:userPhone.Phone}
	if has:=model.GetRow(); has{
		response.ShowError(c, "phone_exists")
		return
	}

	model.Salt=common.GetRandomBoth(4)
	model.Passwd = common.Sha1En(userPhone.Passwd+model.Salt)
	model.Ctime=time.Now().Second()
	model.Status=models.UsersStatusOk

	//traceModel := models.Trace{Ctime:model.Ctime}
	//
	//c.Request.RemoteAddr


	response.ShowSuccess(c,"success")
	return
}
