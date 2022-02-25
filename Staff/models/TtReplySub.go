package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"ttmyth123/GroupLottery/Staff/models/mconst"
)

type TtReplySub struct {
	BaseInfo
	MainId       int64  `orm:"column(main_id)" description:"主表Id"`
	ReplyContent string `orm:"size(1000)" description:"回复的内容"`
}

func (a *TtReplySub) TableName() string {
	return mconst.TableName_TtReplySub
}

func (this *TtReplySub) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	this.CreatedAt = time.Now()
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *TtReplySub) Update(o orm.Ormer, cols ...string) error {
	this.UpdateAt = time.Now()
	_, e := o.Update(this, cols...)
	return e
}

func (this *TtReplySub) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(&this)
	return e
}
