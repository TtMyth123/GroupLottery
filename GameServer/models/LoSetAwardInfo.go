package models

import (
	"github.com/TtMyth123/GameServer/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

type LoSetAwardInfo struct {
	Id         int
	GameType   int
	LotteryNum int64  //期号
	LotteryStr string //期号
	ResultNums string `orm:"size(1024)"` //开奖结果数据

	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func (this *LoSetAwardInfo) TableName() string {
	return mconst.TableName_LoSetAwardInfo
}
func (this *LoSetAwardInfo) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	if e == nil {
		this.Id = int(id)
	}

	return e
}

func (this *LoSetAwardInfo) Update(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	this.UpdatedAt = time.Now()
	id, e := o.Update(this, cols...)
	if e == nil {
		this.Id = int(id)
	}

	return e
}
