package models

import (
	"time"
)

type Users struct {
	Ctime  int       `json:"ctime" xorm:"not null default 0 comment('创建时间') index INT(10)"`
	Email  string    `json:"email" xorm:"not null default '' comment('邮箱') VARCHAR(100)"`
	Ext    string    `json:"ext" xorm:"not null comment('扩展字段') TEXT"`
	Id     int64     `json:"id" xorm:"pk autoincr comment('主键') BIGINT(20)"`
	Mtime  time.Time `json:"mtime" xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
	Name   string    `json:"name" xorm:"not null default '' comment('用户名') VARCHAR(50)"`
	Passwd string    `json:"passwd" xorm:"not null comment('密码') VARCHAR(20)"`
	Phone  string    `json:"phone" xorm:"not null default '' comment('手机号') VARCHAR(20)"`
	Salt   string    `json:"salt" xorm:"not null comment('盐值') CHAR(4)"`
	Status int       `json:"status" xorm:"not null default 0 comment('状态（0：未审核,1:通过 10删除）') TINYINT(4)"`
}
var UsersStatusOk =1
var UsersStatusDel =10
var UsersStatusDef =0
func(u *Users) GetRow() bool {
	has, err := mEngine.Get(u)
	if err==nil &&  has  {
		return true
	}
	return false
}
func (u *Users) GetAll()([]Users,error) {
	var users[]Users
	err:=mEngine.Find(&users)
	return users,err
}

func (u *Users) Add(roles []interface{}) (int64 ,error){
	session := mEngine.NewSession()
	defer session.Close()
	// add Begin() before any action
	if err := session.Begin(); err != nil {
		// if returned then will rodefer session.Close()llback automatically
		return 0,err
	}
	//var uid int64
	_,err:=session.Insert(u)
	if err!=nil {
		return 0,err
	}

	//
	//
	//for _,k:=range roles{
	//	roleModel:=SystemRole{Name:k.(string)}
	//	has:=roleModel.GetRow()
	//	if !has {
	//		continue
	//	}
	//	if	roleModel.Status==0{
	//		continue
	//	}
	//	userroleModel:=SystemUserRole{SystemRoleId:roleModel.Id,SystemUserId:u.Id}
	//	has,err:=session.Get(&userroleModel)
	//	if err!=nil {
	//		return 0,err
	//	}
	//	if has {
	//		continue
	//	}
	//	_,err=session.Insert(&userroleModel)
	//	if err!=nil {
	//		return 0,err
	//	}
	//}
	return u.Id,session.Commit()
}
