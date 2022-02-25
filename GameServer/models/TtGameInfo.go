package models

import (
	"github.com/astaxie/beego/orm"
	"ttmyth123/GroupLottery/GameServer/models/mconst"
)

type TtGameInfo struct {
	Id       int
	GameType int
	GameName string
}

func (this *TtGameInfo) TableName() string {
	return mconst.TableName_TtGameInfo
}

func InitTtGameInfo() error {
	o := orm.NewOrm()
	c, _ := o.QueryTable("tt_game_info").Count()
	if c == 0 {
		arr := make([]TtGameInfo, 0)
		arr = append(arr, TtGameInfo{Id: mconst.GameType_Wsx_201, GameType: mconst.GameType_Wsx_201, GameName: mconst.GetGameName(mconst.GameType_Wsx_201)})
		arr = append(arr, TtGameInfo{Id: mconst.GameType_Wsx_202, GameType: mconst.GameType_Wsx_202, GameName: mconst.GetGameName(mconst.GameType_Wsx_202)})
		arr = append(arr, TtGameInfo{Id: mconst.GameType_Wsx_203, GameType: mconst.GameType_Wsx_203, GameName: mconst.GetGameName(mconst.GameType_Wsx_203)})

		arr = append(arr, TtGameInfo{Id: mconst.GameType_G28_041, GameType: mconst.GameType_G28_041, GameName: mconst.GetGameName(mconst.GameType_G28_041)})
		arr = append(arr, TtGameInfo{Id: mconst.GameType_G28_042, GameType: mconst.GameType_G28_042, GameName: mconst.GetGameName(mconst.GameType_G28_042)})
		arr = append(arr, TtGameInfo{Id: mconst.GameType_G28_043, GameType: mconst.GameType_G28_043, GameName: mconst.GetGameName(mconst.GameType_G28_043)})
		arr = append(arr, TtGameInfo{Id: mconst.GameType_G28_044, GameType: mconst.GameType_G28_044, GameName: mconst.GetGameName(mconst.GameType_G28_044)})

		arr = append(arr, TtGameInfo{Id: mconst.GameType_USC_cqssc, GameType: mconst.GameType_USC_cqssc, GameName: mconst.GetGameName(mconst.GameType_USC_cqssc)})
		arr = append(arr, TtGameInfo{Id: mconst.GameType_USC_jsssc, GameType: mconst.GameType_USC_jsssc, GameName: mconst.GetGameName(mconst.GameType_USC_jsssc)})
		arr = append(arr, TtGameInfo{Id: mconst.GameType_USC_ygcyc, GameType: mconst.GameType_USC_ygcyc, GameName: mconst.GetGameName(mconst.GameType_USC_ygcyc)})
		arr = append(arr, TtGameInfo{Id: mconst.GameType_USC_ygssc, GameType: mconst.GameType_USC_ygssc, GameName: mconst.GetGameName(mconst.GameType_USC_ygssc)})
		arr = append(arr, TtGameInfo{Id: mconst.GameType_USC_gzxy5, GameType: mconst.GameType_USC_gzxy5, GameName: mconst.GetGameName(mconst.GameType_USC_gzxy5)})
		arr = append(arr, TtGameInfo{Id: mconst.GameType_USC_yxssc, GameType: mconst.GameType_USC_yxssc, GameName: mconst.GetGameName(mconst.GameType_USC_yxssc)})

		arr = append(arr, TtGameInfo{Id: mconst.GameType_USC_bjsc, GameType: mconst.GameType_USC_bjsc, GameName: mconst.GetGameName(mconst.GameType_USC_bjsc)})
		arr = append(arr, TtGameInfo{Id: mconst.GameType_USC_xyft, GameType: mconst.GameType_USC_xyft, GameName: mconst.GetGameName(mconst.GameType_USC_xyft)})
		arr = append(arr, TtGameInfo{Id: mconst.GameType_USC_jskc, GameType: mconst.GameType_USC_jskc, GameName: mconst.GetGameName(mconst.GameType_USC_jskc)})
		arr = append(arr, TtGameInfo{Id: mconst.GameType_USC_jssc, GameType: mconst.GameType_USC_jssc, GameName: mconst.GetGameName(mconst.GameType_USC_jssc)})
		arr = append(arr, TtGameInfo{Id: mconst.GameType_USC_ESPsm, GameType: mconst.GameType_USC_ESPsm, GameName: mconst.GetGameName(mconst.GameType_USC_ESPsm)})
		arr = append(arr, TtGameInfo{Id: mconst.GameType_USC_ygxyft, GameType: mconst.GameType_USC_ygxyft, GameName: mconst.GetGameName(mconst.GameType_USC_ygxyft)})
		arr = append(arr, TtGameInfo{Id: mconst.GameType_USC_ygsc, GameType: mconst.GameType_USC_ygsc, GameName: mconst.GetGameName(mconst.GameType_USC_ygsc)})
		arr = append(arr, TtGameInfo{Id: mconst.GameType_USC_gzxy10, GameType: mconst.GameType_USC_gzxy10, GameName: mconst.GetGameName(mconst.GameType_USC_gzxy10)})

		arr = append(arr, TtGameInfo{Id: mconst.GameType_USC_cqxync, GameType: mconst.GameType_USC_cqxync, GameName: mconst.GetGameName(mconst.GameType_USC_cqxync)})
		arr = append(arr, TtGameInfo{Id: mconst.GameType_USC_gdkl10f, GameType: mconst.GameType_USC_gdkl10f, GameName: mconst.GetGameName(mconst.GameType_USC_gdkl10f)})

		_, e := o.InsertMulti(len(arr), arr)
		return e
	}
	return nil
}
