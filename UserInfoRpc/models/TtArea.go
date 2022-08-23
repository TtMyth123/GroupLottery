package models

import (
	"github.com/TtMyth123/UserInfoRpc/models/mconst"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"time"
)

type TtArea struct {
	Id        int
	Area      string
	Des       string
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func (this *TtArea) TableName() string {
	return mconst.TableName_TtArea
}

func InitTtArea() error {
	o := orm.NewOrm()
	c, _ := o.QueryTable(mconst.TableName_TtArea).Count()
	if c == 0 {
		arrData := make([]TtArea, 0)
		arrData = append(arrData, TtArea{Id: 11, Area: "MYGE", Des: "小强游戏"})
		arrData = append(arrData, TtArea{Id: 12, Area: "ZGGame", Des: "ZG游戏"})
		arrData = append(arrData, TtArea{Id: 13, Area: "瞳瞳", Des: "广州客户1"})
		arrData = append(arrData, TtArea{Id: 14, Area: "吴XS", Des: "广州客户2 越南彩"})
		_, e := o.InsertMulti(len(arrData), arrData)
		if e != nil {
			ttLog.LogError("aaaaaaaaaaaa:", e)
			return e
		}
	}
	return nil
}
func (this *TtArea) Add() error {
	o := orm.NewOrm()
	id, e := o.Insert(this)
	this.Id = int(id)
	return e
}

func GetAddArea(Area string) (TtArea, error) {
	o := orm.NewOrm()
	aTtArea := TtArea{}
	e := o.QueryTable(mconst.TableName_TtArea).Filter("Area", Area).One(&aTtArea)
	if e != nil {
		aTtArea.Area = Area
		e = aTtArea.Add()
		if e != nil {
			return aTtArea, e
		}
	}
	return aTtArea, e
}
