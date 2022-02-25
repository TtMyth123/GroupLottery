package models

import (
	"github.com/astaxie/beego/orm"
	g_mconst "ttmyth123/GroupLottery/GameServer/models/mconst"
	"ttmyth123/GroupLottery/UserInfoRpc/models/mconst"
	"ttmyth123/kit/ttLog"
)

type TtAreaRefGame struct {
	Id       int
	AreaId   int
	GameType int
}

func (this *TtAreaRefGame) TableName() string {
	return mconst.TableName_TtAreaRefGame
}

func InitTtAreaRefGame() error {
	o := orm.NewOrm()
	c, _ := o.QueryTable(mconst.TableName_TtAreaRefGame).Count()
	if c == 0 {
		arrData := make([]TtAreaRefGame, 0)
		//arrData = append(arrData, TtArea{Id: 11, Area: "MYGE", Des: "小强游戏"})
		//arrData = append(arrData, TtArea{Id: 12, Area: "ZGGame", Des: "ZG游戏"})
		//arrData = append(arrData, TtArea{Id: 13, Area: "瞳瞳", Des: "广州客户1"})
		//arrData = append(arrData, TtArea{Id: 14, Area: "吴XS", Des: "广州客户2 越南彩"})

		arrData = append(arrData, TtAreaRefGame{AreaId: 12, GameType: g_mconst.GameType_G28_041})
		arrData = append(arrData, TtAreaRefGame{AreaId: 12, GameType: g_mconst.GameType_G28_042})
		arrData = append(arrData, TtAreaRefGame{AreaId: 12, GameType: g_mconst.GameType_G28_043})
		arrData = append(arrData, TtAreaRefGame{AreaId: 12, GameType: g_mconst.GameType_G28_044})

		arrData = append(arrData, TtAreaRefGame{AreaId: 14, GameType: g_mconst.GameType_Wsx_201})
		arrData = append(arrData, TtAreaRefGame{AreaId: 14, GameType: g_mconst.GameType_Wsx_202})
		arrData = append(arrData, TtAreaRefGame{AreaId: 14, GameType: g_mconst.GameType_Wsx_203})

		_, e := o.InsertMulti(len(arrData), arrData)
		if e != nil {
			ttLog.LogError("aaaaaaaaaaaa:", e)
			return e
		}
	}
	return nil
}
func (this *TtAreaRefGame) Add() error {
	o := orm.NewOrm()
	id, e := o.Insert(&this)
	this.Id = int(id)
	return e
}
