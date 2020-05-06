package models

type Trace struct {
	Ctime int    `json:"ctime" xorm:"not null default 0 comment('注册时间') INT(11)"`
	Ext   string `json:"ext" xorm:"not null comment('扩展字段') VARCHAR(1000)"`
	Id    int64  `json:"id" xorm:"pk autoincr comment('主键') BIGINT(20)"`
	Ip    int    `json:"ip" xorm:"not null comment('ip') INT(10)"`
	Type  int    `json:"type" xorm:"not null default 0 comment('类型(0:注册1::登录2:退出3:修改4:删除)') index(UT) TINYINT(4)"`
	Uid   int64  `json:"uid" xorm:"not null default 0 comment('用户主键') index(UT) BIGINT(20)"`
}

var TraceTypeReg = 0
var TraceTypeLogin = 1
var TraceTypeOut = 2
var TraceTypeEdit = 3
var TraceTypeDel = 4
