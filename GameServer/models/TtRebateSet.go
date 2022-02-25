package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"ttmyth123/GroupLottery/GameServer/models/mconst"
)

type TtRebateSet struct {
	Id    int
	Level int //返佣层级

	BetRebateRatio  float64   //投注佣金比例
	BetRebateRatio1 float64   //投注佣金比例
	BetRebateRatio2 float64   //投注佣金比例
	BetRebateRatio3 float64   //投注佣金比例
	CreatedAt       time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateAt        time.Time `orm:"auto_now_add;type(datetime)"`
}

func (this *TtRebateSet) TableName() string {
	return mconst.TableName_TtRebateSet
}

func (this *TtRebateSet) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	id, e := o.Insert(this)
	this.CreatedAt = time.Now()
	this.UpdateAt = this.CreatedAt
	this.Id = int(id)
	return e
}

func (this *TtRebateSet) Update(o orm.Ormer, cols ...string) error {
	_, e := o.Update(this, cols...)
	this.UpdateAt = time.Now()
	return e
}

func (this *TtRebateSet) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(&this)
	return e
}

func GetTtRebateSet() TtRebateSet {
	o := orm.NewOrm()
	aTtDrawSaveSet := TtRebateSet{Id: 1}
	o.Read(&aTtDrawSaveSet)
	return aTtDrawSaveSet
}

func InitTtRebateSet() error {
	o := orm.NewOrm()
	c, _ := o.QueryTable(mconst.TableName_TtRebateSet).Count()
	if c == 0 {
		aTtDrawSaveSet := TtRebateSet{Id: 1, BetRebateRatio: 0.01, BetRebateRatio1: 0.0, BetRebateRatio2: 0.0, BetRebateRatio3: 0.0, Level: 1}
		e := aTtDrawSaveSet.Add(o)
		return e
	}
	return nil
}
