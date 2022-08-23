package models

import (
	"github.com/TtMyth123/GameServer/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

type TtDrawSaveSet struct {
	Id            int
	RegGiveType   int       //注册赠送数量
	RegGiveCount  int       //注册赠送数量
	MinSave       int       //最小充值
	MinDraw       int       //最小提现
	DrawHint      string    //充值提示
	SaveHint      string    //提现提示
	SaveGiveType  int       //充值赠送类型
	SaveGiveRatio float64   //充值赠送比例
	DrawBeginHr   int       //提现的开始时间（小时）
	DrawWorkHr    int       //提现的工作（小时）
	CreatedAt     time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateAt      time.Time `orm:"auto_now_add;type(datetime)"`
}

func (this *TtDrawSaveSet) TableName() string {
	return mconst.TableName_TtDrawSaveSet
}

func (this *TtDrawSaveSet) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	id, e := o.Insert(this)
	this.CreatedAt = time.Now()
	this.UpdateAt = this.CreatedAt
	this.Id = int(id)
	return e
}

func (this *TtDrawSaveSet) Update(o orm.Ormer, cols ...string) error {
	_, e := o.Update(this, cols...)
	this.UpdateAt = time.Now()
	return e
}

func (this *TtDrawSaveSet) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(&this)
	return e
}

func GetTtDrawSaveSet() TtDrawSaveSet {
	o := orm.NewOrm()
	aTtDrawSaveSet := TtDrawSaveSet{Id: 1}
	o.Read(&aTtDrawSaveSet)
	return aTtDrawSaveSet
}

func InitTtDrawSaveSet() error {
	o := orm.NewOrm()
	c, _ := o.QueryTable(mconst.TableName_TtDrawSaveSet).Count()
	if c == 0 {
		aTtDrawSaveSet := TtDrawSaveSet{Id: 1, RegGiveType: mconst.RegGiveType_1, RegGiveCount: 0, MinDraw: 100, MinSave: 100,
			SaveHint: "充值最少100.",
			DrawHint: "提现最少100,手续费10%", SaveGiveType: mconst.SaveGiveType_1, SaveGiveRatio: 0, DrawBeginHr: 0, DrawWorkHr: 24}
		e := aTtDrawSaveSet.Add(o)
		return e
	}
	return nil
}
