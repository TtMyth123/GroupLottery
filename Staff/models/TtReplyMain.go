package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"ttmyth123/GroupLottery/Staff/models/mconst"
)

type TtReplyMain struct {
	BaseInfo
	StaffId int64  `orm:"column(staff_id)" description:"客服Id"`
	MainKey string `orm:"size(256)" description:"关键字"`
}

func (a *TtReplyMain) TableName() string {
	return mconst.TableName_TtReplyMain
}

func (this *TtReplyMain) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	this.CreatedAt = time.Now()
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *TtReplyMain) Update(o orm.Ormer, cols ...string) error {
	this.UpdateAt = time.Now()
	_, e := o.Update(this, cols...)
	return e
}

func (this *TtReplyMain) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(&this)
	return e
}
