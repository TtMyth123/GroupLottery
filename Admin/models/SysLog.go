package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"ttmyth123/GroupLottery/Admin/models/mconst"
)

type SysLog struct {
	Id         int64
	UserName   string
	Content    string
	RemoteAddr string
	UpdatedAt  time.Time `orm:"auto_now;type(datetime)"`
}

func (a *SysLog) TableName() string {
	return mconst.TableName_SysLog
}

func SysLogAdd(_user string, _content string, _remote_add string, _login_time time.Time) error {
	m := SysLog{UserName: _user, Content: _content, RemoteAddr: _remote_add, UpdatedAt: _login_time}

	o := orm.NewOrm()
	if _, err := o.Insert(&m); err == nil {
		return nil
	} else {
		return err
	}
}
