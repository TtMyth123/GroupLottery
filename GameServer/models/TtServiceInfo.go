package models

import (
	"github.com/TtMyth123/GameServer/models/mconst"
	"github.com/astaxie/beego/orm"
)

type TtServiceInfo struct {
	Id        int
	Nickname  string
	QQ        string
	Wechat    string
	WechatUrl string
}

func (a *TtServiceInfo) TableName() string {
	return mconst.TableName_TtServiceInfo
}

func InitTtServiceInfo() error {
	o := orm.NewOrm()
	c, _ := o.QueryTable(mconst.TableName_TtServiceInfo).Count()
	if c == 0 {
		arrData := make([]TtServiceInfo, 0)
		arrData = append(arrData, TtServiceInfo{Id: 1, Nickname: "小小"})
		_, e := o.InsertMulti(len(arrData), arrData)
		return e
	}

	return nil
}
