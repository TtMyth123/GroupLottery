package models

import (
	"github.com/TtMyth123/Staff/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

//TtQuickReply

type TtQuickReply struct {
	BaseInfo
	StaffId int64  `orm:"column(staff_id)" description:"客服Id"`
	ReplyC  string `orm:"size(1000)" description:"快捷回复内容"`
}

func (a *TtQuickReply) TableName() string {
	return mconst.TableName_TtQuickReply
}

func (this *TtQuickReply) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	this.CreatedAt = time.Now()
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *TtQuickReply) Update(o orm.Ormer, cols ...string) error {
	this.UpdateAt = time.Now()
	_, e := o.Update(this, cols...)
	return e
}

func (this *TtQuickReply) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(&this)
	return e
}
