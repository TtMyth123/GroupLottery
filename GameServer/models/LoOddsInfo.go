package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/Game28ResultKit/Game28Const"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/UscResultKit/Usc10Kit"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/UscResultKit/Usc5Kit"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/UscResultKit/Usc8for20Kit"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/WsxConst"
	"ttmyth123/GroupLottery/GameServer/models/mconst"
	"ttmyth123/kit/ttLog"
)

type LoOddsInfo struct {
	Id            int
	GameType      int
	OddsType      int
	Handicap      int    //盘口
	OddsDes       string `orm:"size(256)"`
	BigOddsDes    string `orm:"size(256)"`
	Odds          float64
	SNum          string
	BigType       int
	OneUserMaxBet int //一个用户最多
	OneUserMinBet int //一个用户最少
	AllUserMaxBet int //全部用户最多
	N1            int
	N2            int
	CreatedAt     time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt     time.Time `orm:"auto_now;type(datetime)"`
}

func (this *LoOddsInfo) TableName() string {
	return mconst.TableName_LoOddsInfo
}

func InitLoOddsInfo() {
	InitOneLoOddsInfo(mconst.GameType_Wsx_201)
	InitOneLoOddsInfo(mconst.GameType_Wsx_202)
	//InitOneLoOddsInfo(mconst.GameType_Wsx_203)
	IniZbcOddsInfo()

	InitOneGame28OddsInfo(mconst.GameType_G28_041)
	InitOneGame28OddsInfo(mconst.GameType_G28_042)
	InitOneGame28OddsInfo(mconst.GameType_G28_043)
	InitOneGame28OddsInfo(mconst.GameType_G28_044)

	InitOneUscOddsInfo()
}
func GetInitOneLoOddsInfo(GameType int) []LoOddsInfo {
	arr := make([]LoOddsInfo, 0)

	arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_01, SNum: "2", OddsType: 1, Odds: 1.9, OddsDes: fmt.Sprintf("头特 大")})
	arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_01, SNum: "1", OddsType: 2, Odds: 1.9, OddsDes: fmt.Sprintf("头特 小")})
	arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_03, SNum: "2", OddsType: 3, Odds: 1.9, OddsDes: fmt.Sprintf("一等奖 大")})
	arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_03, SNum: "1", OddsType: 4, Odds: 1.9, OddsDes: fmt.Sprintf("一等奖 小")})

	arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_02, SNum: "1", OddsType: 5, Odds: 1.9, OddsDes: fmt.Sprintf("头特 单")})
	arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_02, SNum: "2", OddsType: 6, Odds: 1.9, OddsDes: fmt.Sprintf("头特 双")})
	arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_04, SNum: "1", OddsType: 7, Odds: 1.9, OddsDes: fmt.Sprintf("一等奖 单")})
	arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_04, SNum: "2", OddsType: 8, Odds: 1.9, OddsDes: fmt.Sprintf("一等奖 双")})

	//头等-头特
	for i := 10; i <= 19; i++ {
		m := i % 10
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_05, SNum: fmt.Sprintf("%d", m), OddsType: i, Odds: 9.9, OddsDes: fmt.Sprintf("头等-头特 %d", m)})
	}
	//头等-头特 B面
	for i := 10 + 20000; i <= 19+20000; i++ {
		m := i % 10
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_05B, SNum: fmt.Sprintf("%d", m), OddsType: i, Odds: 9.9, OddsDes: fmt.Sprintf("头等-头特 %d", m)})
	}

	//头等-尾特
	for i := 20; i <= 29; i++ {
		m := i % 10
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_06, SNum: fmt.Sprintf("%d", m), OddsType: i, Odds: 9.9, OddsDes: fmt.Sprintf("头等-尾特 %d", m)})
	}
	//头等-尾特
	for i := 20 + 20000; i <= 29+20000; i++ {
		m := i % 10
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_06B, SNum: fmt.Sprintf("%d", m), OddsType: i, Odds: 9.9, OddsDes: fmt.Sprintf("头等-尾特 %d", m)})
	}

	//一等-头特
	for i := 610; i <= 619; i++ {
		m := i % 10
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_18, SNum: fmt.Sprintf("%d", m), OddsType: i, Odds: 9.9, OddsDes: fmt.Sprintf("一等-头特 %d", m)})
	}
	//一等-头特
	for i := 610 + 20000; i <= 619+20000; i++ {
		m := i % 10
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_18B, SNum: fmt.Sprintf("%d", m), OddsType: i, Odds: 9.9, OddsDes: fmt.Sprintf("一等-头特 %d", m)})
	}

	//一等-尾特
	for i := 620; i <= 629; i++ {
		m := i % 10
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_19, SNum: fmt.Sprintf("%d", m), OddsType: i, Odds: 9.9, OddsDes: fmt.Sprintf("一等-尾特 %d", m)})
	}
	//一等-尾特
	for i := 620 + 20000; i <= 629+20000; i++ {
		m := i % 10
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_19B, SNum: fmt.Sprintf("%d", m), OddsType: i, Odds: 9.9, OddsDes: fmt.Sprintf("一等-尾特 %d", m)})
	}

	//二等-头特
	for i := 630; i <= 639; i++ {
		m := i % 10
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_20, SNum: fmt.Sprintf("%d", m), OddsType: i, Odds: 9.9, OddsDes: fmt.Sprintf("二等-头特 %d", m)})
	}
	//二等-头特
	for i := 630 + 20000; i <= 639+20000; i++ {
		m := i % 10
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_20B, SNum: fmt.Sprintf("%d", m), OddsType: i, Odds: 9.9, OddsDes: fmt.Sprintf("二等-头特 %d", m)})
	}
	//二等-尾特
	for i := 640; i <= 649; i++ {
		m := i % 10
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_21, SNum: fmt.Sprintf("%d", m), OddsType: i, Odds: 9.9, OddsDes: fmt.Sprintf("二等-尾特 %d", m)})
	}
	//二等-尾特
	for i := 640 + 20000; i <= 649+20000; i++ {
		m := i % 10
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_21B, SNum: fmt.Sprintf("%d", m), OddsType: i, Odds: 9.9, OddsDes: fmt.Sprintf("二等-尾特 %d", m)})
	}

	////波色
	//bs := []string{"绿波","红波","紫波","黄波"}
	//
	//for i:=600;i<600+96;i++ {
	//	m := i% 8
	//	m2:= m%2
	//
	//	arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType:WsxConst.BigType_14, SNum:fmt.Sprintf("%02d",m), OddsType: i, Odds: 99.9, OddsDes: fmt.Sprintf("%s %02d",bs[m2],m)})
	//}
	//arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType:WsxConst.BigType_14, SNum:fmt.Sprintf("%02d",96), OddsType: 696, Odds: 99.9, OddsDes: fmt.Sprintf("%s %02d",bs[0],96)})
	//arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType:WsxConst.BigType_14, SNum:fmt.Sprintf("%02d",97), OddsType: 697, Odds: 99.9, OddsDes: fmt.Sprintf("%s %02d",bs[1],97)})
	//arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType:WsxConst.BigType_14, SNum:fmt.Sprintf("%02d",98), OddsType: 698, Odds: 99.9, OddsDes: fmt.Sprintf("%s %02d",bs[2],98)})
	//arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType:WsxConst.BigType_14, SNum:fmt.Sprintf("%02d",99), OddsType: 699, Odds: 99.9, OddsDes: fmt.Sprintf("%s %02d",bs[3],99)})

	bs := []string{"绿", "红", "紫", "黄"}
	//波色 头等特码
	for i := 30; i < 30+4; i++ {
		m := i % 10
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_15, SNum: fmt.Sprintf("%d", m), OddsType: i, Odds: 9.9, OddsDes: fmt.Sprintf("头等特码 %s", bs[m])})
	}

	//波色 一等特码
	for i := 40; i < 40+4; i++ {
		m := i % 10
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_16, SNum: fmt.Sprintf("%d", m), OddsType: i, Odds: 9.9, OddsDes: fmt.Sprintf("一等特码 %s", bs[m])})
	}

	//波色 二等特码
	for i := 50; i < 50+4; i++ {
		m := i % 10
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_17, SNum: fmt.Sprintf("%d", m), OddsType: i, Odds: 9.9, OddsDes: fmt.Sprintf("二等特码 %s", bs[m])})
	}

	//平码两位区
	arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_12,
		SNum: fmt.Sprintf("%d", 0), OddsType: 60, Odds: 99.9, OddsDes: fmt.Sprintf("平码两位区")})
	//平码三位区
	arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_13,
		SNum: fmt.Sprintf("%d", 0), OddsType: 61, Odds: 999.9, OddsDes: fmt.Sprintf("平码三位区")})

	//头等特码
	for i := 100; i < 100+100; i++ {
		m := i % 100
		arr = append(arr, LoOddsInfo{Handicap: 0, OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_07, SNum: fmt.Sprintf("%02d", m), OddsType: i, Odds: 99.9, OddsDes: fmt.Sprintf("头等特码 %02d", m)})
	}
	//头等特码 //B面
	for i := 100 + 20000; i < 100+100+20000; i++ {
		m := i % 100
		arr = append(arr, LoOddsInfo{Handicap: 1, OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_07B, SNum: fmt.Sprintf("%02d", m), OddsType: i, Odds: 99.9, OddsDes: fmt.Sprintf("头等特码 %02d", m)})
	}

	//一等特码
	for i := 200; i < 200+100; i++ {
		m := i % 100
		arr = append(arr, LoOddsInfo{Handicap: 0, OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_08, SNum: fmt.Sprintf("%02d", m), OddsType: i, Odds: 99.9, OddsDes: fmt.Sprintf("一等特码 %02d", m)})
	}
	//一等特码 //B面
	for i := 200 + 20000; i < 200+100+20000; i++ {
		m := i % 100
		arr = append(arr, LoOddsInfo{Handicap: 1, OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_08B, SNum: fmt.Sprintf("%02d", m), OddsType: i, Odds: 99.9, OddsDes: fmt.Sprintf("一等特码 %02d", m)})
	}

	//二等特码
	for i := 300; i < 300+100; i++ {
		m := i % 100
		arr = append(arr, LoOddsInfo{Handicap: 0, OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_09, SNum: fmt.Sprintf("%02d", m), OddsType: i, Odds: 99.9, OddsDes: fmt.Sprintf("二等特码 %02d", m)})
	}
	//二等特码
	for i := 300 + 20000; i < 300+100+20000; i++ {
		m := i % 100
		arr = append(arr, LoOddsInfo{Handicap: 1, OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_09B, SNum: fmt.Sprintf("%02d", m), OddsType: i, Odds: 99.9, OddsDes: fmt.Sprintf("二等特码 %02d", m)})
	}

	//二连位
	for i := 400; i < 400+100; i++ {
		m := i % 100
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_10, SNum: fmt.Sprintf("%02d", m), OddsType: i, Odds: 99.9, OddsDes: fmt.Sprintf("二连位 %02d", m)})
	}

	//三连位
	if GameType == mconst.GameType_Wsx_202 {
		//头等 三连位
		for i := 1000; i < 1000+1000; i++ {
			m := i % 1000
			arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_11, SNum: fmt.Sprintf("%03d", m), OddsType: i, Odds: 999.9, OddsDes: fmt.Sprintf("头等 三连位 %03d", m)})
		}
	} else {
		//三连位
		for i := 1000; i < 1000+1000; i++ {
			m := i % 1000
			arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_11, SNum: fmt.Sprintf("%03d", m), OddsType: i, Odds: 999.9, OddsDes: fmt.Sprintf("三连位 %03d", m)})
		}
	}

	if GameType == mconst.GameType_Wsx_202 {
		//一等 三连位
		for i := 2000; i < 2000+1000; i++ {
			m := i % 1000
			arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: WsxConst.BigType_22, SNum: fmt.Sprintf("%03d", m), OddsType: i, Odds: 999.9, OddsDes: fmt.Sprintf("一等 三连位 %03d", m)})
		}
	}

	////平码两位区
	//for i:=500;i<500+100;i++ {
	//	m := i% 100
	//	arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType:WsxConst.BigType_12, SNum:fmt.Sprintf("%02d",m), OddsType: i, Odds: 99.9, OddsDes: fmt.Sprintf("平码两位区 %02d",m)})
	//}
	//
	////平码三位区
	//for i:=2000;i<2000+1000;i++ {
	//	m := i% 1000
	//	arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType:WsxConst.BigType_13, SNum:fmt.Sprintf("%03d",m), OddsType: i, Odds: 999.9, OddsDes: fmt.Sprintf("平码三位区 %03d",m)})
	//}

	return arr
}
func IniZbcOddsInfo() {
	o := orm.NewOrm()
	GameType := mconst.GameType_Wsx_203
	c, _ := o.QueryTable(mconst.TableName_LoOddsInfo).Filter("GameType", GameType).Count()
	if c == 0 {
		arr := make([]LoOddsInfo, 0)
		/*
		   1-1
		   2-2
		   3-2，3-3
		   4-2，4-3，4-4
		   5-3，5-4，5-5
		   6-3，6-4，6-5，6-6
		   7-3，7-4，7-5，7-6，7-7
		   8-4，8-5，8-6，8-7，8-8
		   9-4，9-5，9-6，9-7，9-8，9-9
		   10-5，10-6，10-7，10-8，10-9，10-10
		*/

		OddsTypeI := WsxConst.TO_WsxZbc_001 - 1
		minN := 1
		maxN := 1
		for i := minN; i <= maxN; i++ {
			OddsTypeI++
			arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, OddsType: OddsTypeI, BigType: WsxConst.TO_BigType_ZB01, N1: maxN, N2: i, Odds: 1.9, BigOddsDes: fmt.Sprintf("%d个球", maxN), OddsDes: fmt.Sprintf("%d个球 %d-%d", maxN, maxN, i)})
		}

		minN = 2
		maxN = 2
		for i := minN; i <= maxN; i++ {
			OddsTypeI++
			arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, OddsType: OddsTypeI, BigType: WsxConst.TO_BigType_ZB02, N1: maxN, N2: i, Odds: 1.9, BigOddsDes: fmt.Sprintf("%d个球", maxN), OddsDes: fmt.Sprintf("%d个球 %d-%d", maxN, maxN, i)})
		}

		minN = 2
		maxN = 3
		for i := minN; i <= maxN; i++ {
			OddsTypeI++
			arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, OddsType: OddsTypeI, BigType: WsxConst.TO_BigType_ZB03, N1: maxN, N2: i, Odds: 1.9, BigOddsDes: fmt.Sprintf("%d个球", maxN), OddsDes: fmt.Sprintf("%d个球 %d-%d", maxN, maxN, i)})
		}

		minN = 2
		maxN = 4
		for i := minN; i <= maxN; i++ {
			OddsTypeI++
			arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, OddsType: OddsTypeI, BigType: WsxConst.TO_BigType_ZB04, N1: maxN, N2: i, Odds: 1.9, BigOddsDes: fmt.Sprintf("%d个球", maxN), OddsDes: fmt.Sprintf("%d个球 %d-%d", maxN, maxN, i)})
		}

		minN = 3
		maxN = 5
		for i := minN; i <= maxN; i++ {
			OddsTypeI++
			arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, OddsType: OddsTypeI, BigType: WsxConst.TO_BigType_ZB05, N1: maxN, N2: i, Odds: 1.9, BigOddsDes: fmt.Sprintf("%d个球", maxN), OddsDes: fmt.Sprintf("%d个球 %d-%d", maxN, maxN, i)})
		}

		minN = 3
		maxN = 6
		for i := minN; i <= maxN; i++ {
			OddsTypeI++
			arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, OddsType: OddsTypeI, BigType: WsxConst.TO_BigType_ZB06, N1: maxN, N2: i, Odds: 1.9, BigOddsDes: fmt.Sprintf("%d个球", maxN), OddsDes: fmt.Sprintf("%d个球 %d-%d", maxN, maxN, i)})
		}

		minN = 3
		maxN = 7
		for i := minN; i <= maxN; i++ {
			OddsTypeI++
			arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, OddsType: OddsTypeI, BigType: WsxConst.TO_BigType_ZB07, N1: maxN, N2: i, Odds: 1.9, BigOddsDes: fmt.Sprintf("%d个球", maxN), OddsDes: fmt.Sprintf("%d个球 %d-%d", maxN, maxN, i)})
		}

		minN = 4
		maxN = 8
		for i := minN; i <= maxN; i++ {
			OddsTypeI++
			arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, OddsType: OddsTypeI, BigType: WsxConst.TO_BigType_ZB08, N1: maxN, N2: i, Odds: 1.9, BigOddsDes: fmt.Sprintf("%d个球", maxN), OddsDes: fmt.Sprintf("%d个球 %d-%d", maxN, maxN, i)})
		}

		minN = 4
		maxN = 9
		for i := minN; i <= maxN; i++ {
			OddsTypeI++
			arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, OddsType: OddsTypeI, BigType: WsxConst.TO_BigType_ZB09, N1: maxN, N2: i, Odds: 1.9, BigOddsDes: fmt.Sprintf("%d个球", maxN), OddsDes: fmt.Sprintf("%d个球 %d-%d", maxN, maxN, i)})
		}

		minN = 5
		maxN = 10
		for i := minN; i <= maxN; i++ {
			OddsTypeI++
			arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, OddsType: OddsTypeI, BigType: WsxConst.TO_BigType_ZB10, N1: maxN, N2: i, Odds: 1.9, BigOddsDes: fmt.Sprintf("%d个球", maxN), OddsDes: fmt.Sprintf("%d个球 %d-%d", maxN, maxN, i)})
		}

		OddsTypeI = WsxConst.TO_WsxZbc_037
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, OddsType: OddsTypeI, BigType: WsxConst.TO_BigType_ZB11, N1: 1, Odds: 1.9, BigOddsDes: fmt.Sprintf("大小"), OddsDes: fmt.Sprintf("大小")})

		OddsTypeI = WsxConst.TO_WsxZbc_039
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, OddsType: OddsTypeI, BigType: WsxConst.TO_BigType_ZB12, N1: 1, Odds: 1.9, BigOddsDes: fmt.Sprintf("单双"), OddsDes: fmt.Sprintf("单双")})

		OddsTypeI = WsxConst.TO_WsxZbc_041
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, OddsType: OddsTypeI, BigType: WsxConst.TO_BigType_ZB13, N1: 1, Odds: 1.9, BigOddsDes: fmt.Sprintf("大小和"), OddsDes: fmt.Sprintf("大小和")})
		OddsTypeI = WsxConst.TO_WsxZbc_042
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, OddsType: OddsTypeI, BigType: WsxConst.TO_BigType_ZB14, N1: 1, Odds: 1.9, BigOddsDes: fmt.Sprintf("单双和"), OddsDes: fmt.Sprintf("单双和")})

		_, e := o.InsertMulti(len(arr), arr)
		if e != nil {
			ttLog.LogError(e)
		}
	}

}

/**
南部彩，北部彩，中部彩
*/
func InitOneLoOddsInfo(GameType int) error {
	o := orm.NewOrm()
	c, _ := o.QueryTable(mconst.TableName_LoOddsInfo).Filter("GameType", GameType).Count()
	if c == 0 {
		arr := GetInitOneLoOddsInfo(GameType)

		_, e := o.InsertMulti(len(arr), arr)
		return e
	}

	return nil
}

func InitOneGame28OddsInfo(GameType int) error {
	o := orm.NewOrm()
	c, _ := o.QueryTable(mconst.TableName_LoOddsInfo).Filter("GameType", GameType).Count()
	if c == 0 {
		arr := make([]LoOddsInfo, 0)
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_01, SNum: "", OddsType: Game28Const.TO_G28_011, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_011})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_01, SNum: "", OddsType: Game28Const.TO_G28_012, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_012})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_01, SNum: "", OddsType: Game28Const.TO_G28_013, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_013})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_01, SNum: "", OddsType: Game28Const.TO_G28_014, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_014})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_01, SNum: "", OddsType: Game28Const.TO_G28_015, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_015})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_01, SNum: "", OddsType: Game28Const.TO_G28_016, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_016})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_01, SNum: "", OddsType: Game28Const.TO_G28_017, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_017})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_01, SNum: "", OddsType: Game28Const.TO_G28_018, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_018})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_01, SNum: "", OddsType: Game28Const.TO_G28_019, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_019})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_01, SNum: "", OddsType: Game28Const.TO_G28_020, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_020})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_02, SNum: "", OddsType: Game28Const.TO_G28_021, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_021})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_02, SNum: "", OddsType: Game28Const.TO_G28_022, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_022})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_02, SNum: "", OddsType: Game28Const.TO_G28_023, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_023})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_03, SNum: "", OddsType: Game28Const.TO_G28_024, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_024})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_03, SNum: "", OddsType: Game28Const.TO_G28_025, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_025})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_03, SNum: "", OddsType: Game28Const.TO_G28_026, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_026})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_04, SNum: "", OddsType: Game28Const.TO_G28_027, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_027})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_04, SNum: "", OddsType: Game28Const.TO_G28_028, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_028})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_04, SNum: "", OddsType: Game28Const.TO_G28_029, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_029})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_05, SNum: "", OddsType: Game28Const.TO_G28_031, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_031})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_05, SNum: "", OddsType: Game28Const.TO_G28_032, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_032})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_05, SNum: "", OddsType: Game28Const.TO_G28_033, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_033})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_05, SNum: "", OddsType: Game28Const.TO_G28_034, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_034})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_100, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_100})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_101, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_101})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_102, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_102})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_103, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_103})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_104, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_104})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_105, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_105})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_106, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_106})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_107, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_107})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_108, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_108})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_109, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_109})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_110, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_110})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_111, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_111})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_112, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_112})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_113, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_113})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_114, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_114})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_115, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_115})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_116, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_116})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_117, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_117})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_118, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_118})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_119, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_119})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_120, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_120})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_121, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_121})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_122, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_122})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_123, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_123})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_124, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_124})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_125, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_125})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_126, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_126})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Game28Const.BigType_06, SNum: "", OddsType: Game28Const.TO_G28_127, Odds: 1.9, OddsDes: Game28Const.TO_N_G28_127})

		_, e := o.InsertMulti(len(arr), arr)
		return e
	}
	return nil
}

func InitOneUsc5OddsInfo(GameType int) error {
	o := orm.NewOrm()
	c, _ := o.QueryTable(mconst.TableName_LoOddsInfo).Filter("GameType", GameType).Count()
	if c == 0 {
		arr := make([]LoOddsInfo, 0)
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_01, SNum: "", OddsType: Usc5Kit.OT_001, Odds: 1.9, OddsDes: Usc5Kit.N_OT_001})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_01, SNum: "", OddsType: Usc5Kit.OT_002, Odds: 1.9, OddsDes: Usc5Kit.N_OT_002})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_01, SNum: "", OddsType: Usc5Kit.OT_003, Odds: 1.9, OddsDes: Usc5Kit.N_OT_003})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_01, SNum: "", OddsType: Usc5Kit.OT_004, Odds: 1.9, OddsDes: Usc5Kit.N_OT_004})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_01, SNum: "", OddsType: Usc5Kit.OT_005, Odds: 1.9, OddsDes: Usc5Kit.N_OT_005})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_01, SNum: "", OddsType: Usc5Kit.OT_006, Odds: 1.9, OddsDes: Usc5Kit.N_OT_006})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_01, SNum: "", OddsType: Usc5Kit.OT_007, Odds: 1.9, OddsDes: Usc5Kit.N_OT_007})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_01, SNum: "", OddsType: Usc5Kit.OT_008, Odds: 1.9, OddsDes: Usc5Kit.N_OT_008})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_01, SNum: "", OddsType: Usc5Kit.OT_009, Odds: 1.9, OddsDes: Usc5Kit.N_OT_009})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_01, SNum: "", OddsType: Usc5Kit.OT_010, Odds: 1.9, OddsDes: Usc5Kit.N_OT_010})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_01, SNum: "", OddsType: Usc5Kit.OT_011, Odds: 1.9, OddsDes: Usc5Kit.N_OT_011})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_01, SNum: "", OddsType: Usc5Kit.OT_012, Odds: 1.9, OddsDes: Usc5Kit.N_OT_012})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_01, SNum: "", OddsType: Usc5Kit.OT_013, Odds: 1.9, OddsDes: Usc5Kit.N_OT_013})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_01, SNum: "", OddsType: Usc5Kit.OT_014, Odds: 1.9, OddsDes: Usc5Kit.N_OT_014})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_02, SNum: "", OddsType: Usc5Kit.OT_015, Odds: 1.9, OddsDes: Usc5Kit.N_OT_015})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_02, SNum: "", OddsType: Usc5Kit.OT_016, Odds: 1.9, OddsDes: Usc5Kit.N_OT_016})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_02, SNum: "", OddsType: Usc5Kit.OT_017, Odds: 1.9, OddsDes: Usc5Kit.N_OT_017})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_02, SNum: "", OddsType: Usc5Kit.OT_018, Odds: 1.9, OddsDes: Usc5Kit.N_OT_018})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_02, SNum: "", OddsType: Usc5Kit.OT_019, Odds: 1.9, OddsDes: Usc5Kit.N_OT_019})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_02, SNum: "", OddsType: Usc5Kit.OT_020, Odds: 1.9, OddsDes: Usc5Kit.N_OT_020})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_02, SNum: "", OddsType: Usc5Kit.OT_021, Odds: 1.9, OddsDes: Usc5Kit.N_OT_021})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_02, SNum: "", OddsType: Usc5Kit.OT_022, Odds: 1.9, OddsDes: Usc5Kit.N_OT_022})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_02, SNum: "", OddsType: Usc5Kit.OT_023, Odds: 1.9, OddsDes: Usc5Kit.N_OT_023})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_02, SNum: "", OddsType: Usc5Kit.OT_024, Odds: 1.9, OddsDes: Usc5Kit.N_OT_024})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_02, SNum: "", OddsType: Usc5Kit.OT_025, Odds: 1.9, OddsDes: Usc5Kit.N_OT_025})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_02, SNum: "", OddsType: Usc5Kit.OT_026, Odds: 1.9, OddsDes: Usc5Kit.N_OT_026})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_02, SNum: "", OddsType: Usc5Kit.OT_027, Odds: 1.9, OddsDes: Usc5Kit.N_OT_027})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_02, SNum: "", OddsType: Usc5Kit.OT_028, Odds: 1.9, OddsDes: Usc5Kit.N_OT_028})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_03, SNum: "", OddsType: Usc5Kit.OT_029, Odds: 1.9, OddsDes: Usc5Kit.N_OT_029})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_03, SNum: "", OddsType: Usc5Kit.OT_030, Odds: 1.9, OddsDes: Usc5Kit.N_OT_030})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_03, SNum: "", OddsType: Usc5Kit.OT_031, Odds: 1.9, OddsDes: Usc5Kit.N_OT_031})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_03, SNum: "", OddsType: Usc5Kit.OT_032, Odds: 1.9, OddsDes: Usc5Kit.N_OT_032})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_03, SNum: "", OddsType: Usc5Kit.OT_033, Odds: 1.9, OddsDes: Usc5Kit.N_OT_033})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_03, SNum: "", OddsType: Usc5Kit.OT_034, Odds: 1.9, OddsDes: Usc5Kit.N_OT_034})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_03, SNum: "", OddsType: Usc5Kit.OT_035, Odds: 1.9, OddsDes: Usc5Kit.N_OT_035})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_03, SNum: "", OddsType: Usc5Kit.OT_036, Odds: 1.9, OddsDes: Usc5Kit.N_OT_036})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_03, SNum: "", OddsType: Usc5Kit.OT_037, Odds: 1.9, OddsDes: Usc5Kit.N_OT_037})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_03, SNum: "", OddsType: Usc5Kit.OT_038, Odds: 1.9, OddsDes: Usc5Kit.N_OT_038})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_03, SNum: "", OddsType: Usc5Kit.OT_039, Odds: 1.9, OddsDes: Usc5Kit.N_OT_039})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_03, SNum: "", OddsType: Usc5Kit.OT_040, Odds: 1.9, OddsDes: Usc5Kit.N_OT_040})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_03, SNum: "", OddsType: Usc5Kit.OT_041, Odds: 1.9, OddsDes: Usc5Kit.N_OT_041})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_03, SNum: "", OddsType: Usc5Kit.OT_042, Odds: 1.9, OddsDes: Usc5Kit.N_OT_042})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_04, SNum: "", OddsType: Usc5Kit.OT_043, Odds: 1.9, OddsDes: Usc5Kit.N_OT_043})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_04, SNum: "", OddsType: Usc5Kit.OT_044, Odds: 1.9, OddsDes: Usc5Kit.N_OT_044})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_04, SNum: "", OddsType: Usc5Kit.OT_045, Odds: 1.9, OddsDes: Usc5Kit.N_OT_045})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_04, SNum: "", OddsType: Usc5Kit.OT_046, Odds: 1.9, OddsDes: Usc5Kit.N_OT_046})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_04, SNum: "", OddsType: Usc5Kit.OT_047, Odds: 1.9, OddsDes: Usc5Kit.N_OT_047})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_04, SNum: "", OddsType: Usc5Kit.OT_048, Odds: 1.9, OddsDes: Usc5Kit.N_OT_048})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_04, SNum: "", OddsType: Usc5Kit.OT_049, Odds: 1.9, OddsDes: Usc5Kit.N_OT_049})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_04, SNum: "", OddsType: Usc5Kit.OT_050, Odds: 1.9, OddsDes: Usc5Kit.N_OT_050})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_04, SNum: "", OddsType: Usc5Kit.OT_051, Odds: 1.9, OddsDes: Usc5Kit.N_OT_051})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_04, SNum: "", OddsType: Usc5Kit.OT_052, Odds: 1.9, OddsDes: Usc5Kit.N_OT_052})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_04, SNum: "", OddsType: Usc5Kit.OT_053, Odds: 1.9, OddsDes: Usc5Kit.N_OT_053})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_04, SNum: "", OddsType: Usc5Kit.OT_054, Odds: 1.9, OddsDes: Usc5Kit.N_OT_054})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_04, SNum: "", OddsType: Usc5Kit.OT_055, Odds: 1.9, OddsDes: Usc5Kit.N_OT_055})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_04, SNum: "", OddsType: Usc5Kit.OT_056, Odds: 1.9, OddsDes: Usc5Kit.N_OT_056})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_05, SNum: "", OddsType: Usc5Kit.OT_057, Odds: 1.9, OddsDes: Usc5Kit.N_OT_057})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_05, SNum: "", OddsType: Usc5Kit.OT_058, Odds: 1.9, OddsDes: Usc5Kit.N_OT_058})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_05, SNum: "", OddsType: Usc5Kit.OT_059, Odds: 1.9, OddsDes: Usc5Kit.N_OT_059})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_05, SNum: "", OddsType: Usc5Kit.OT_060, Odds: 1.9, OddsDes: Usc5Kit.N_OT_060})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_05, SNum: "", OddsType: Usc5Kit.OT_061, Odds: 1.9, OddsDes: Usc5Kit.N_OT_061})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_05, SNum: "", OddsType: Usc5Kit.OT_062, Odds: 1.9, OddsDes: Usc5Kit.N_OT_062})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_05, SNum: "", OddsType: Usc5Kit.OT_063, Odds: 1.9, OddsDes: Usc5Kit.N_OT_063})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_05, SNum: "", OddsType: Usc5Kit.OT_064, Odds: 1.9, OddsDes: Usc5Kit.N_OT_064})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_05, SNum: "", OddsType: Usc5Kit.OT_065, Odds: 1.9, OddsDes: Usc5Kit.N_OT_065})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_05, SNum: "", OddsType: Usc5Kit.OT_066, Odds: 1.9, OddsDes: Usc5Kit.N_OT_066})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_05, SNum: "", OddsType: Usc5Kit.OT_067, Odds: 1.9, OddsDes: Usc5Kit.N_OT_067})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_05, SNum: "", OddsType: Usc5Kit.OT_068, Odds: 1.9, OddsDes: Usc5Kit.N_OT_068})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_05, SNum: "", OddsType: Usc5Kit.OT_069, Odds: 1.9, OddsDes: Usc5Kit.N_OT_069})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_05, SNum: "", OddsType: Usc5Kit.OT_070, Odds: 1.9, OddsDes: Usc5Kit.N_OT_070})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_06, SNum: "", OddsType: Usc5Kit.OT_071, Odds: 1.9, OddsDes: Usc5Kit.N_OT_071})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_06, SNum: "", OddsType: Usc5Kit.OT_072, Odds: 1.9, OddsDes: Usc5Kit.N_OT_072})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_06, SNum: "", OddsType: Usc5Kit.OT_073, Odds: 1.9, OddsDes: Usc5Kit.N_OT_073})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_06, SNum: "", OddsType: Usc5Kit.OT_074, Odds: 1.9, OddsDes: Usc5Kit.N_OT_074})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_07, SNum: "", OddsType: Usc5Kit.OT_075, Odds: 1.9, OddsDes: Usc5Kit.N_OT_075})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_07, SNum: "", OddsType: Usc5Kit.OT_076, Odds: 1.9, OddsDes: Usc5Kit.N_OT_076})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_07, SNum: "", OddsType: Usc5Kit.OT_077, Odds: 1.9, OddsDes: Usc5Kit.N_OT_077})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_08, SNum: "", OddsType: Usc5Kit.OT_078, Odds: 1.9, OddsDes: Usc5Kit.N_OT_078})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_08, SNum: "", OddsType: Usc5Kit.OT_079, Odds: 1.9, OddsDes: Usc5Kit.N_OT_079})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_08, SNum: "", OddsType: Usc5Kit.OT_080, Odds: 1.9, OddsDes: Usc5Kit.N_OT_080})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_08, SNum: "", OddsType: Usc5Kit.OT_081, Odds: 1.9, OddsDes: Usc5Kit.N_OT_081})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_08, SNum: "", OddsType: Usc5Kit.OT_082, Odds: 1.9, OddsDes: Usc5Kit.N_OT_082})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_08, SNum: "", OddsType: Usc5Kit.OT_083, Odds: 1.9, OddsDes: Usc5Kit.N_OT_083})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_08, SNum: "", OddsType: Usc5Kit.OT_084, Odds: 1.9, OddsDes: Usc5Kit.N_OT_084})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_08, SNum: "", OddsType: Usc5Kit.OT_085, Odds: 1.9, OddsDes: Usc5Kit.N_OT_085})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_08, SNum: "", OddsType: Usc5Kit.OT_086, Odds: 1.9, OddsDes: Usc5Kit.N_OT_086})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_08, SNum: "", OddsType: Usc5Kit.OT_087, Odds: 1.9, OddsDes: Usc5Kit.N_OT_087})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_08, SNum: "", OddsType: Usc5Kit.OT_088, Odds: 1.9, OddsDes: Usc5Kit.N_OT_088})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_08, SNum: "", OddsType: Usc5Kit.OT_089, Odds: 1.9, OddsDes: Usc5Kit.N_OT_089})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_08, SNum: "", OddsType: Usc5Kit.OT_090, Odds: 1.9, OddsDes: Usc5Kit.N_OT_090})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_08, SNum: "", OddsType: Usc5Kit.OT_091, Odds: 1.9, OddsDes: Usc5Kit.N_OT_091})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_08, SNum: "", OddsType: Usc5Kit.OT_092, Odds: 1.9, OddsDes: Usc5Kit.N_OT_092})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_99, SNum: "", OddsType: Usc5Kit.OT_093, Odds: 1.9, OddsDes: Usc5Kit.N_OT_093})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_99, SNum: "", OddsType: Usc5Kit.OT_094, Odds: 1.9, OddsDes: Usc5Kit.N_OT_094})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_99, SNum: "", OddsType: Usc5Kit.OT_095, Odds: 1.9, OddsDes: Usc5Kit.N_OT_095})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_99, SNum: "", OddsType: Usc5Kit.OT_096, Odds: 1.9, OddsDes: Usc5Kit.N_OT_096})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_99, SNum: "", OddsType: Usc5Kit.OT_097, Odds: 1.9, OddsDes: Usc5Kit.N_OT_097})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_99, SNum: "", OddsType: Usc5Kit.OT_098, Odds: 1.9, OddsDes: Usc5Kit.N_OT_098})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_99, SNum: "", OddsType: Usc5Kit.OT_099, Odds: 1.9, OddsDes: Usc5Kit.N_OT_099})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_99, SNum: "", OddsType: Usc5Kit.OT_100, Odds: 1.9, OddsDes: Usc5Kit.N_OT_100})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_99, SNum: "", OddsType: Usc5Kit.OT_101, Odds: 1.9, OddsDes: Usc5Kit.N_OT_101})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_99, SNum: "", OddsType: Usc5Kit.OT_102, Odds: 1.9, OddsDes: Usc5Kit.N_OT_102})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_99, SNum: "", OddsType: Usc5Kit.OT_103, Odds: 1.9, OddsDes: Usc5Kit.N_OT_103})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_99, SNum: "", OddsType: Usc5Kit.OT_104, Odds: 1.9, OddsDes: Usc5Kit.N_OT_104})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_99, SNum: "", OddsType: Usc5Kit.OT_105, Odds: 1.9, OddsDes: Usc5Kit.N_OT_105})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_99, SNum: "", OddsType: Usc5Kit.OT_106, Odds: 1.9, OddsDes: Usc5Kit.N_OT_106})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_99, SNum: "", OddsType: Usc5Kit.OT_107, Odds: 1.9, OddsDes: Usc5Kit.N_OT_107})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_108, Odds: 1.9, OddsDes: Usc5Kit.N_OT_108})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_109, Odds: 1.9, OddsDes: Usc5Kit.N_OT_109})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_110, Odds: 1.9, OddsDes: Usc5Kit.N_OT_110})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_111, Odds: 1.9, OddsDes: Usc5Kit.N_OT_111})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_112, Odds: 1.9, OddsDes: Usc5Kit.N_OT_112})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_113, Odds: 1.9, OddsDes: Usc5Kit.N_OT_113})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_114, Odds: 1.9, OddsDes: Usc5Kit.N_OT_114})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_115, Odds: 1.9, OddsDes: Usc5Kit.N_OT_115})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_116, Odds: 1.9, OddsDes: Usc5Kit.N_OT_116})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_117, Odds: 1.9, OddsDes: Usc5Kit.N_OT_117})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_118, Odds: 1.9, OddsDes: Usc5Kit.N_OT_118})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_119, Odds: 1.9, OddsDes: Usc5Kit.N_OT_119})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_120, Odds: 1.9, OddsDes: Usc5Kit.N_OT_120})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_121, Odds: 1.9, OddsDes: Usc5Kit.N_OT_121})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_122, Odds: 1.9, OddsDes: Usc5Kit.N_OT_122})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_123, Odds: 1.9, OddsDes: Usc5Kit.N_OT_123})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_124, Odds: 1.9, OddsDes: Usc5Kit.N_OT_124})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_125, Odds: 1.9, OddsDes: Usc5Kit.N_OT_125})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_126, Odds: 1.9, OddsDes: Usc5Kit.N_OT_126})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_127, Odds: 1.9, OddsDes: Usc5Kit.N_OT_127})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_128, Odds: 1.9, OddsDes: Usc5Kit.N_OT_128})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_129, Odds: 1.9, OddsDes: Usc5Kit.N_OT_129})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_130, Odds: 1.9, OddsDes: Usc5Kit.N_OT_130})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_131, Odds: 1.9, OddsDes: Usc5Kit.N_OT_131})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_132, Odds: 1.9, OddsDes: Usc5Kit.N_OT_132})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc5Kit.GT_09, SNum: "", OddsType: Usc5Kit.OT_133, Odds: 1.9, OddsDes: Usc5Kit.N_OT_133})
		_, e := o.InsertMulti(len(arr), arr)
		return e
	}
	return nil
}

func InitOneUsc10OddsInfo(GameType int) error {
	o := orm.NewOrm()
	c, _ := o.QueryTable(mconst.TableName_LoOddsInfo).Filter("GameType", GameType).Count()
	if c == 0 {
		arr := make([]LoOddsInfo, 0)
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_001, Odds: 1.9, OddsDes: Usc10Kit.N_OT_001})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_002, Odds: 1.9, OddsDes: Usc10Kit.N_OT_002})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_003, Odds: 1.9, OddsDes: Usc10Kit.N_OT_003})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_004, Odds: 1.9, OddsDes: Usc10Kit.N_OT_004})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_005, Odds: 1.9, OddsDes: Usc10Kit.N_OT_005})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_006, Odds: 1.9, OddsDes: Usc10Kit.N_OT_006})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_007, Odds: 1.9, OddsDes: Usc10Kit.N_OT_007})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_008, Odds: 1.9, OddsDes: Usc10Kit.N_OT_008})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_009, Odds: 1.9, OddsDes: Usc10Kit.N_OT_009})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_010, Odds: 1.9, OddsDes: Usc10Kit.N_OT_010})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_011, Odds: 1.9, OddsDes: Usc10Kit.N_OT_011})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_012, Odds: 1.9, OddsDes: Usc10Kit.N_OT_012})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_013, Odds: 1.9, OddsDes: Usc10Kit.N_OT_013})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_014, Odds: 1.9, OddsDes: Usc10Kit.N_OT_014})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_015, Odds: 1.9, OddsDes: Usc10Kit.N_OT_015})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_016, Odds: 1.9, OddsDes: Usc10Kit.N_OT_016})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_017, Odds: 1.9, OddsDes: Usc10Kit.N_OT_017})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_018, Odds: 1.9, OddsDes: Usc10Kit.N_OT_018})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_019, Odds: 1.9, OddsDes: Usc10Kit.N_OT_019})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_020, Odds: 1.9, OddsDes: Usc10Kit.N_OT_020})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_01, SNum: "", OddsType: Usc10Kit.OT_021, Odds: 1.9, OddsDes: Usc10Kit.N_OT_021})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_02, SNum: "", OddsType: Usc10Kit.OT_022, Odds: 1.9, OddsDes: Usc10Kit.N_OT_022})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_02, SNum: "", OddsType: Usc10Kit.OT_023, Odds: 1.9, OddsDes: Usc10Kit.N_OT_023})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_02, SNum: "", OddsType: Usc10Kit.OT_024, Odds: 1.9, OddsDes: Usc10Kit.N_OT_024})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_02, SNum: "", OddsType: Usc10Kit.OT_025, Odds: 1.9, OddsDes: Usc10Kit.N_OT_025})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_02, SNum: "", OddsType: Usc10Kit.OT_026, Odds: 1.9, OddsDes: Usc10Kit.N_OT_026})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_02, SNum: "", OddsType: Usc10Kit.OT_027, Odds: 1.9, OddsDes: Usc10Kit.N_OT_027})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_02, SNum: "", OddsType: Usc10Kit.OT_028, Odds: 1.9, OddsDes: Usc10Kit.N_OT_028})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_02, SNum: "", OddsType: Usc10Kit.OT_029, Odds: 1.9, OddsDes: Usc10Kit.N_OT_029})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_02, SNum: "", OddsType: Usc10Kit.OT_030, Odds: 1.9, OddsDes: Usc10Kit.N_OT_030})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_02, SNum: "", OddsType: Usc10Kit.OT_031, Odds: 1.9, OddsDes: Usc10Kit.N_OT_031})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_02, SNum: "", OddsType: Usc10Kit.OT_032, Odds: 1.9, OddsDes: Usc10Kit.N_OT_032})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_02, SNum: "", OddsType: Usc10Kit.OT_033, Odds: 1.9, OddsDes: Usc10Kit.N_OT_033})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_02, SNum: "", OddsType: Usc10Kit.OT_034, Odds: 1.9, OddsDes: Usc10Kit.N_OT_034})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_02, SNum: "", OddsType: Usc10Kit.OT_035, Odds: 1.9, OddsDes: Usc10Kit.N_OT_035})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_02, SNum: "", OddsType: Usc10Kit.OT_036, Odds: 1.9, OddsDes: Usc10Kit.N_OT_036})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_02, SNum: "", OddsType: Usc10Kit.OT_037, Odds: 1.9, OddsDes: Usc10Kit.N_OT_037})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_03, SNum: "", OddsType: Usc10Kit.OT_038, Odds: 1.9, OddsDes: Usc10Kit.N_OT_038})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_03, SNum: "", OddsType: Usc10Kit.OT_039, Odds: 1.9, OddsDes: Usc10Kit.N_OT_039})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_03, SNum: "", OddsType: Usc10Kit.OT_040, Odds: 1.9, OddsDes: Usc10Kit.N_OT_040})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_03, SNum: "", OddsType: Usc10Kit.OT_041, Odds: 1.9, OddsDes: Usc10Kit.N_OT_041})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_03, SNum: "", OddsType: Usc10Kit.OT_042, Odds: 1.9, OddsDes: Usc10Kit.N_OT_042})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_03, SNum: "", OddsType: Usc10Kit.OT_043, Odds: 1.9, OddsDes: Usc10Kit.N_OT_043})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_03, SNum: "", OddsType: Usc10Kit.OT_044, Odds: 1.9, OddsDes: Usc10Kit.N_OT_044})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_03, SNum: "", OddsType: Usc10Kit.OT_045, Odds: 1.9, OddsDes: Usc10Kit.N_OT_045})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_03, SNum: "", OddsType: Usc10Kit.OT_046, Odds: 1.9, OddsDes: Usc10Kit.N_OT_046})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_03, SNum: "", OddsType: Usc10Kit.OT_047, Odds: 1.9, OddsDes: Usc10Kit.N_OT_047})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_03, SNum: "", OddsType: Usc10Kit.OT_048, Odds: 1.9, OddsDes: Usc10Kit.N_OT_048})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_03, SNum: "", OddsType: Usc10Kit.OT_049, Odds: 1.9, OddsDes: Usc10Kit.N_OT_049})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_03, SNum: "", OddsType: Usc10Kit.OT_050, Odds: 1.9, OddsDes: Usc10Kit.N_OT_050})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_03, SNum: "", OddsType: Usc10Kit.OT_051, Odds: 1.9, OddsDes: Usc10Kit.N_OT_051})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_03, SNum: "", OddsType: Usc10Kit.OT_052, Odds: 1.9, OddsDes: Usc10Kit.N_OT_052})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_03, SNum: "", OddsType: Usc10Kit.OT_053, Odds: 1.9, OddsDes: Usc10Kit.N_OT_053})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_04, SNum: "", OddsType: Usc10Kit.OT_054, Odds: 1.9, OddsDes: Usc10Kit.N_OT_054})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_04, SNum: "", OddsType: Usc10Kit.OT_055, Odds: 1.9, OddsDes: Usc10Kit.N_OT_055})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_04, SNum: "", OddsType: Usc10Kit.OT_056, Odds: 1.9, OddsDes: Usc10Kit.N_OT_056})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_04, SNum: "", OddsType: Usc10Kit.OT_057, Odds: 1.9, OddsDes: Usc10Kit.N_OT_057})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_04, SNum: "", OddsType: Usc10Kit.OT_058, Odds: 1.9, OddsDes: Usc10Kit.N_OT_058})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_04, SNum: "", OddsType: Usc10Kit.OT_059, Odds: 1.9, OddsDes: Usc10Kit.N_OT_059})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_04, SNum: "", OddsType: Usc10Kit.OT_060, Odds: 1.9, OddsDes: Usc10Kit.N_OT_060})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_04, SNum: "", OddsType: Usc10Kit.OT_061, Odds: 1.9, OddsDes: Usc10Kit.N_OT_061})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_04, SNum: "", OddsType: Usc10Kit.OT_062, Odds: 1.9, OddsDes: Usc10Kit.N_OT_062})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_04, SNum: "", OddsType: Usc10Kit.OT_063, Odds: 1.9, OddsDes: Usc10Kit.N_OT_063})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_04, SNum: "", OddsType: Usc10Kit.OT_064, Odds: 1.9, OddsDes: Usc10Kit.N_OT_064})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_04, SNum: "", OddsType: Usc10Kit.OT_065, Odds: 1.9, OddsDes: Usc10Kit.N_OT_065})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_04, SNum: "", OddsType: Usc10Kit.OT_066, Odds: 1.9, OddsDes: Usc10Kit.N_OT_066})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_04, SNum: "", OddsType: Usc10Kit.OT_067, Odds: 1.9, OddsDes: Usc10Kit.N_OT_067})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_04, SNum: "", OddsType: Usc10Kit.OT_068, Odds: 1.9, OddsDes: Usc10Kit.N_OT_068})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_04, SNum: "", OddsType: Usc10Kit.OT_069, Odds: 1.9, OddsDes: Usc10Kit.N_OT_069})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_05, SNum: "", OddsType: Usc10Kit.OT_070, Odds: 1.9, OddsDes: Usc10Kit.N_OT_070})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_05, SNum: "", OddsType: Usc10Kit.OT_071, Odds: 1.9, OddsDes: Usc10Kit.N_OT_071})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_05, SNum: "", OddsType: Usc10Kit.OT_072, Odds: 1.9, OddsDes: Usc10Kit.N_OT_072})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_05, SNum: "", OddsType: Usc10Kit.OT_073, Odds: 1.9, OddsDes: Usc10Kit.N_OT_073})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_05, SNum: "", OddsType: Usc10Kit.OT_074, Odds: 1.9, OddsDes: Usc10Kit.N_OT_074})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_05, SNum: "", OddsType: Usc10Kit.OT_075, Odds: 1.9, OddsDes: Usc10Kit.N_OT_075})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_05, SNum: "", OddsType: Usc10Kit.OT_076, Odds: 1.9, OddsDes: Usc10Kit.N_OT_076})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_05, SNum: "", OddsType: Usc10Kit.OT_077, Odds: 1.9, OddsDes: Usc10Kit.N_OT_077})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_05, SNum: "", OddsType: Usc10Kit.OT_078, Odds: 1.9, OddsDes: Usc10Kit.N_OT_078})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_05, SNum: "", OddsType: Usc10Kit.OT_079, Odds: 1.9, OddsDes: Usc10Kit.N_OT_079})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_05, SNum: "", OddsType: Usc10Kit.OT_080, Odds: 1.9, OddsDes: Usc10Kit.N_OT_080})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_05, SNum: "", OddsType: Usc10Kit.OT_081, Odds: 1.9, OddsDes: Usc10Kit.N_OT_081})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_05, SNum: "", OddsType: Usc10Kit.OT_082, Odds: 1.9, OddsDes: Usc10Kit.N_OT_082})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_05, SNum: "", OddsType: Usc10Kit.OT_083, Odds: 1.9, OddsDes: Usc10Kit.N_OT_083})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_05, SNum: "", OddsType: Usc10Kit.OT_084, Odds: 1.9, OddsDes: Usc10Kit.N_OT_084})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_05, SNum: "", OddsType: Usc10Kit.OT_085, Odds: 1.9, OddsDes: Usc10Kit.N_OT_085})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_06, SNum: "", OddsType: Usc10Kit.OT_086, Odds: 1.9, OddsDes: Usc10Kit.N_OT_086})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_06, SNum: "", OddsType: Usc10Kit.OT_087, Odds: 1.9, OddsDes: Usc10Kit.N_OT_087})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_06, SNum: "", OddsType: Usc10Kit.OT_088, Odds: 1.9, OddsDes: Usc10Kit.N_OT_088})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_06, SNum: "", OddsType: Usc10Kit.OT_089, Odds: 1.9, OddsDes: Usc10Kit.N_OT_089})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_06, SNum: "", OddsType: Usc10Kit.OT_090, Odds: 1.9, OddsDes: Usc10Kit.N_OT_090})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_06, SNum: "", OddsType: Usc10Kit.OT_091, Odds: 1.9, OddsDes: Usc10Kit.N_OT_091})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_06, SNum: "", OddsType: Usc10Kit.OT_092, Odds: 1.9, OddsDes: Usc10Kit.N_OT_092})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_06, SNum: "", OddsType: Usc10Kit.OT_093, Odds: 1.9, OddsDes: Usc10Kit.N_OT_093})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_06, SNum: "", OddsType: Usc10Kit.OT_094, Odds: 1.9, OddsDes: Usc10Kit.N_OT_094})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_06, SNum: "", OddsType: Usc10Kit.OT_095, Odds: 1.9, OddsDes: Usc10Kit.N_OT_095})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_06, SNum: "", OddsType: Usc10Kit.OT_096, Odds: 1.9, OddsDes: Usc10Kit.N_OT_096})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_06, SNum: "", OddsType: Usc10Kit.OT_097, Odds: 1.9, OddsDes: Usc10Kit.N_OT_097})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_06, SNum: "", OddsType: Usc10Kit.OT_098, Odds: 1.9, OddsDes: Usc10Kit.N_OT_098})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_06, SNum: "", OddsType: Usc10Kit.OT_099, Odds: 1.9, OddsDes: Usc10Kit.N_OT_099})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_06, SNum: "", OddsType: Usc10Kit.OT_100, Odds: 1.9, OddsDes: Usc10Kit.N_OT_100})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_06, SNum: "", OddsType: Usc10Kit.OT_101, Odds: 1.9, OddsDes: Usc10Kit.N_OT_101})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_07, SNum: "", OddsType: Usc10Kit.OT_102, Odds: 1.9, OddsDes: Usc10Kit.N_OT_102})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_07, SNum: "", OddsType: Usc10Kit.OT_103, Odds: 1.9, OddsDes: Usc10Kit.N_OT_103})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_07, SNum: "", OddsType: Usc10Kit.OT_104, Odds: 1.9, OddsDes: Usc10Kit.N_OT_104})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_07, SNum: "", OddsType: Usc10Kit.OT_105, Odds: 1.9, OddsDes: Usc10Kit.N_OT_105})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_07, SNum: "", OddsType: Usc10Kit.OT_106, Odds: 1.9, OddsDes: Usc10Kit.N_OT_106})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_07, SNum: "", OddsType: Usc10Kit.OT_107, Odds: 1.9, OddsDes: Usc10Kit.N_OT_107})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_07, SNum: "", OddsType: Usc10Kit.OT_108, Odds: 1.9, OddsDes: Usc10Kit.N_OT_108})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_07, SNum: "", OddsType: Usc10Kit.OT_109, Odds: 1.9, OddsDes: Usc10Kit.N_OT_109})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_07, SNum: "", OddsType: Usc10Kit.OT_110, Odds: 1.9, OddsDes: Usc10Kit.N_OT_110})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_07, SNum: "", OddsType: Usc10Kit.OT_111, Odds: 1.9, OddsDes: Usc10Kit.N_OT_111})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_07, SNum: "", OddsType: Usc10Kit.OT_112, Odds: 1.9, OddsDes: Usc10Kit.N_OT_112})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_07, SNum: "", OddsType: Usc10Kit.OT_113, Odds: 1.9, OddsDes: Usc10Kit.N_OT_113})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_07, SNum: "", OddsType: Usc10Kit.OT_114, Odds: 1.9, OddsDes: Usc10Kit.N_OT_114})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_07, SNum: "", OddsType: Usc10Kit.OT_115, Odds: 1.9, OddsDes: Usc10Kit.N_OT_115})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_08, SNum: "", OddsType: Usc10Kit.OT_118, Odds: 1.9, OddsDes: Usc10Kit.N_OT_118})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_08, SNum: "", OddsType: Usc10Kit.OT_119, Odds: 1.9, OddsDes: Usc10Kit.N_OT_119})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_08, SNum: "", OddsType: Usc10Kit.OT_120, Odds: 1.9, OddsDes: Usc10Kit.N_OT_120})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_08, SNum: "", OddsType: Usc10Kit.OT_121, Odds: 1.9, OddsDes: Usc10Kit.N_OT_121})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_08, SNum: "", OddsType: Usc10Kit.OT_122, Odds: 1.9, OddsDes: Usc10Kit.N_OT_122})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_08, SNum: "", OddsType: Usc10Kit.OT_123, Odds: 1.9, OddsDes: Usc10Kit.N_OT_123})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_08, SNum: "", OddsType: Usc10Kit.OT_124, Odds: 1.9, OddsDes: Usc10Kit.N_OT_124})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_08, SNum: "", OddsType: Usc10Kit.OT_125, Odds: 1.9, OddsDes: Usc10Kit.N_OT_125})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_08, SNum: "", OddsType: Usc10Kit.OT_126, Odds: 1.9, OddsDes: Usc10Kit.N_OT_126})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_08, SNum: "", OddsType: Usc10Kit.OT_127, Odds: 1.9, OddsDes: Usc10Kit.N_OT_127})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_08, SNum: "", OddsType: Usc10Kit.OT_128, Odds: 1.9, OddsDes: Usc10Kit.N_OT_128})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_08, SNum: "", OddsType: Usc10Kit.OT_129, Odds: 1.9, OddsDes: Usc10Kit.N_OT_129})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_08, SNum: "", OddsType: Usc10Kit.OT_130, Odds: 1.9, OddsDes: Usc10Kit.N_OT_130})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_08, SNum: "", OddsType: Usc10Kit.OT_131, Odds: 1.9, OddsDes: Usc10Kit.N_OT_131})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_09, SNum: "", OddsType: Usc10Kit.OT_134, Odds: 1.9, OddsDes: Usc10Kit.N_OT_134})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_09, SNum: "", OddsType: Usc10Kit.OT_135, Odds: 1.9, OddsDes: Usc10Kit.N_OT_135})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_09, SNum: "", OddsType: Usc10Kit.OT_136, Odds: 1.9, OddsDes: Usc10Kit.N_OT_136})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_09, SNum: "", OddsType: Usc10Kit.OT_137, Odds: 1.9, OddsDes: Usc10Kit.N_OT_137})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_09, SNum: "", OddsType: Usc10Kit.OT_138, Odds: 1.9, OddsDes: Usc10Kit.N_OT_138})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_09, SNum: "", OddsType: Usc10Kit.OT_139, Odds: 1.9, OddsDes: Usc10Kit.N_OT_139})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_09, SNum: "", OddsType: Usc10Kit.OT_140, Odds: 1.9, OddsDes: Usc10Kit.N_OT_140})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_09, SNum: "", OddsType: Usc10Kit.OT_141, Odds: 1.9, OddsDes: Usc10Kit.N_OT_141})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_09, SNum: "", OddsType: Usc10Kit.OT_142, Odds: 1.9, OddsDes: Usc10Kit.N_OT_142})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_09, SNum: "", OddsType: Usc10Kit.OT_143, Odds: 1.9, OddsDes: Usc10Kit.N_OT_143})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_09, SNum: "", OddsType: Usc10Kit.OT_144, Odds: 1.9, OddsDes: Usc10Kit.N_OT_144})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_09, SNum: "", OddsType: Usc10Kit.OT_145, Odds: 1.9, OddsDes: Usc10Kit.N_OT_145})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_09, SNum: "", OddsType: Usc10Kit.OT_146, Odds: 1.9, OddsDes: Usc10Kit.N_OT_146})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_09, SNum: "", OddsType: Usc10Kit.OT_147, Odds: 1.9, OddsDes: Usc10Kit.N_OT_147})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_10, SNum: "", OddsType: Usc10Kit.OT_150, Odds: 1.9, OddsDes: Usc10Kit.N_OT_150})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_10, SNum: "", OddsType: Usc10Kit.OT_151, Odds: 1.9, OddsDes: Usc10Kit.N_OT_151})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_10, SNum: "", OddsType: Usc10Kit.OT_152, Odds: 1.9, OddsDes: Usc10Kit.N_OT_152})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_10, SNum: "", OddsType: Usc10Kit.OT_153, Odds: 1.9, OddsDes: Usc10Kit.N_OT_153})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_10, SNum: "", OddsType: Usc10Kit.OT_154, Odds: 1.9, OddsDes: Usc10Kit.N_OT_154})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_10, SNum: "", OddsType: Usc10Kit.OT_155, Odds: 1.9, OddsDes: Usc10Kit.N_OT_155})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_10, SNum: "", OddsType: Usc10Kit.OT_156, Odds: 1.9, OddsDes: Usc10Kit.N_OT_156})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_10, SNum: "", OddsType: Usc10Kit.OT_157, Odds: 1.9, OddsDes: Usc10Kit.N_OT_157})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_10, SNum: "", OddsType: Usc10Kit.OT_158, Odds: 1.9, OddsDes: Usc10Kit.N_OT_158})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_10, SNum: "", OddsType: Usc10Kit.OT_159, Odds: 1.9, OddsDes: Usc10Kit.N_OT_159})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_10, SNum: "", OddsType: Usc10Kit.OT_160, Odds: 1.9, OddsDes: Usc10Kit.N_OT_160})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_10, SNum: "", OddsType: Usc10Kit.OT_161, Odds: 1.9, OddsDes: Usc10Kit.N_OT_161})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_10, SNum: "", OddsType: Usc10Kit.OT_162, Odds: 1.9, OddsDes: Usc10Kit.N_OT_162})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_10, SNum: "", OddsType: Usc10Kit.OT_163, Odds: 1.9, OddsDes: Usc10Kit.N_OT_163})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_11, SNum: "", OddsType: Usc10Kit.OT_166, Odds: 1.9, OddsDes: Usc10Kit.N_OT_166})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_11, SNum: "", OddsType: Usc10Kit.OT_167, Odds: 1.9, OddsDes: Usc10Kit.N_OT_167})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_11, SNum: "", OddsType: Usc10Kit.OT_168, Odds: 1.9, OddsDes: Usc10Kit.N_OT_168})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_11, SNum: "", OddsType: Usc10Kit.OT_169, Odds: 1.9, OddsDes: Usc10Kit.N_OT_169})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_11, SNum: "", OddsType: Usc10Kit.OT_170, Odds: 1.9, OddsDes: Usc10Kit.N_OT_170})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_11, SNum: "", OddsType: Usc10Kit.OT_171, Odds: 1.9, OddsDes: Usc10Kit.N_OT_171})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_11, SNum: "", OddsType: Usc10Kit.OT_172, Odds: 1.9, OddsDes: Usc10Kit.N_OT_172})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_11, SNum: "", OddsType: Usc10Kit.OT_173, Odds: 1.9, OddsDes: Usc10Kit.N_OT_173})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_11, SNum: "", OddsType: Usc10Kit.OT_174, Odds: 1.9, OddsDes: Usc10Kit.N_OT_174})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_11, SNum: "", OddsType: Usc10Kit.OT_175, Odds: 1.9, OddsDes: Usc10Kit.N_OT_175})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_11, SNum: "", OddsType: Usc10Kit.OT_176, Odds: 1.9, OddsDes: Usc10Kit.N_OT_176})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_11, SNum: "", OddsType: Usc10Kit.OT_177, Odds: 1.9, OddsDes: Usc10Kit.N_OT_177})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_11, SNum: "", OddsType: Usc10Kit.OT_178, Odds: 1.9, OddsDes: Usc10Kit.N_OT_178})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_11, SNum: "", OddsType: Usc10Kit.OT_179, Odds: 1.9, OddsDes: Usc10Kit.N_OT_179})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_99, SNum: "", OddsType: Usc10Kit.OT_180, Odds: 1.9, OddsDes: Usc10Kit.N_OT_180})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_99, SNum: "", OddsType: Usc10Kit.OT_181, Odds: 1.9, OddsDes: Usc10Kit.N_OT_181})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_182, Odds: 1.9, OddsDes: Usc10Kit.N_OT_182})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_183, Odds: 1.9, OddsDes: Usc10Kit.N_OT_183})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_184, Odds: 1.9, OddsDes: Usc10Kit.N_OT_184})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_185, Odds: 1.9, OddsDes: Usc10Kit.N_OT_185})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_186, Odds: 1.9, OddsDes: Usc10Kit.N_OT_186})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_187, Odds: 1.9, OddsDes: Usc10Kit.N_OT_187})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_188, Odds: 1.9, OddsDes: Usc10Kit.N_OT_188})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_189, Odds: 1.9, OddsDes: Usc10Kit.N_OT_189})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_190, Odds: 1.9, OddsDes: Usc10Kit.N_OT_190})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_191, Odds: 1.9, OddsDes: Usc10Kit.N_OT_191})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_192, Odds: 1.9, OddsDes: Usc10Kit.N_OT_192})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_193, Odds: 1.9, OddsDes: Usc10Kit.N_OT_193})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_194, Odds: 1.9, OddsDes: Usc10Kit.N_OT_194})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_195, Odds: 1.9, OddsDes: Usc10Kit.N_OT_195})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_196, Odds: 1.9, OddsDes: Usc10Kit.N_OT_196})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_197, Odds: 1.9, OddsDes: Usc10Kit.N_OT_197})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_198, Odds: 1.9, OddsDes: Usc10Kit.N_OT_198})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_199, Odds: 1.9, OddsDes: Usc10Kit.N_OT_199})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_200, Odds: 1.9, OddsDes: Usc10Kit.N_OT_200})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_201, Odds: 1.9, OddsDes: Usc10Kit.N_OT_201})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_202, Odds: 1.9, OddsDes: Usc10Kit.N_OT_202})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_203, Odds: 1.9, OddsDes: Usc10Kit.N_OT_203})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_204, Odds: 1.9, OddsDes: Usc10Kit.N_OT_204})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_205, Odds: 1.9, OddsDes: Usc10Kit.N_OT_205})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_206, Odds: 1.9, OddsDes: Usc10Kit.N_OT_206})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc10Kit.GT_12, SNum: "", OddsType: Usc10Kit.OT_207, Odds: 1.9, OddsDes: Usc10Kit.N_OT_207})
		_, e := o.InsertMulti(len(arr), arr)
		return e
	}
	return nil
}

func InitOneUsc8for20OddsInfo(GameType int) error {
	o := orm.NewOrm()
	c, _ := o.QueryTable(mconst.TableName_LoOddsInfo).Filter("GameType", GameType).Count()
	if c == 0 {
		arr := make([]LoOddsInfo, 0)
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_001, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_001})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_002, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_002})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_003, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_003})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_004, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_004})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_005, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_005})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_006, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_006})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_007, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_007})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_008, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_008})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_009, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_009})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_010, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_010})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_011, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_011})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_012, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_012})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_013, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_013})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_014, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_014})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_015, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_015})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_016, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_016})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_017, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_017})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_018, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_018})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_019, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_019})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_020, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_020})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_021, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_021})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_022, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_022})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_023, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_023})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_024, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_024})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_025, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_025})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_026, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_026})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_027, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_027})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_028, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_028})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_029, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_029})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_030, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_030})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_031, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_031})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_032, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_032})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_033, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_033})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_034, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_034})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_035, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_035})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_036, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_036})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_01, SNum: "", OddsType: Usc8for20Kit.OT_037, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_037})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_038, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_038})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_039, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_039})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_040, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_040})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_041, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_041})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_042, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_042})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_043, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_043})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_044, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_044})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_045, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_045})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_046, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_046})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_047, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_047})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_048, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_048})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_049, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_049})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_050, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_050})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_051, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_051})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_052, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_052})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_053, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_053})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_054, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_054})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_055, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_055})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_056, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_056})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_057, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_057})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_058, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_058})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_059, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_059})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_060, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_060})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_061, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_061})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_062, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_062})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_063, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_063})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_064, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_064})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_065, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_065})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_066, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_066})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_067, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_067})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_068, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_068})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_069, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_069})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_070, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_070})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_071, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_071})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_072, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_072})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_073, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_073})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_02, SNum: "", OddsType: Usc8for20Kit.OT_074, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_074})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_075, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_075})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_076, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_076})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_077, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_077})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_078, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_078})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_079, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_079})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_080, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_080})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_081, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_081})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_082, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_082})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_083, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_083})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_084, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_084})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_085, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_085})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_086, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_086})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_087, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_087})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_088, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_088})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_089, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_089})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_090, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_090})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_091, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_091})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_092, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_092})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_093, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_093})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_094, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_094})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_095, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_095})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_096, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_096})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_097, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_097})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_098, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_098})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_099, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_099})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_100, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_100})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_101, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_101})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_102, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_102})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_103, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_103})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_104, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_104})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_105, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_105})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_106, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_106})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_107, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_107})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_108, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_108})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_109, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_109})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_110, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_110})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_03, SNum: "", OddsType: Usc8for20Kit.OT_111, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_111})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_112, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_112})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_113, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_113})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_114, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_114})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_115, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_115})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_116, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_116})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_117, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_117})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_118, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_118})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_119, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_119})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_120, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_120})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_121, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_121})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_122, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_122})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_123, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_123})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_124, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_124})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_125, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_125})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_126, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_126})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_127, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_127})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_128, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_128})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_129, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_129})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_130, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_130})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_131, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_131})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_132, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_132})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_133, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_133})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_134, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_134})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_135, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_135})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_136, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_136})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_137, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_137})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_138, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_138})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_139, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_139})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_140, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_140})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_141, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_141})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_142, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_142})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_143, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_143})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_144, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_144})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_145, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_145})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_146, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_146})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_147, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_147})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_04, SNum: "", OddsType: Usc8for20Kit.OT_148, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_148})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_149, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_149})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_150, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_150})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_151, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_151})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_152, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_152})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_153, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_153})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_154, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_154})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_155, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_155})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_156, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_156})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_157, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_157})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_158, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_158})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_159, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_159})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_160, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_160})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_161, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_161})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_162, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_162})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_163, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_163})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_164, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_164})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_165, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_165})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_166, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_166})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_167, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_167})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_168, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_168})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_169, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_169})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_170, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_170})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_171, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_171})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_172, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_172})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_173, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_173})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_174, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_174})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_175, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_175})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_176, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_176})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_177, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_177})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_178, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_178})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_179, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_179})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_180, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_180})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_181, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_181})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_182, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_182})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_183, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_183})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_184, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_184})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_05, SNum: "", OddsType: Usc8for20Kit.OT_185, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_185})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_186, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_186})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_187, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_187})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_188, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_188})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_189, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_189})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_190, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_190})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_191, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_191})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_192, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_192})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_193, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_193})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_194, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_194})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_195, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_195})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_196, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_196})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_197, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_197})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_198, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_198})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_199, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_199})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_200, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_200})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_201, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_201})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_202, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_202})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_203, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_203})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_204, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_204})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_205, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_205})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_206, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_206})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_207, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_207})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_208, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_208})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_209, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_209})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_210, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_210})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_211, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_211})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_212, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_212})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_213, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_213})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_214, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_214})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_215, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_215})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_216, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_216})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_217, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_217})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_218, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_218})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_219, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_219})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_220, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_220})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_221, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_221})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_06, SNum: "", OddsType: Usc8for20Kit.OT_222, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_222})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_223, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_223})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_224, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_224})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_225, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_225})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_226, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_226})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_227, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_227})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_228, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_228})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_229, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_229})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_230, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_230})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_231, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_231})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_232, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_232})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_233, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_233})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_234, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_234})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_235, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_235})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_236, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_236})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_237, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_237})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_238, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_238})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_239, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_239})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_240, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_240})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_241, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_241})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_242, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_242})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_243, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_243})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_244, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_244})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_245, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_245})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_246, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_246})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_247, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_247})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_248, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_248})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_249, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_249})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_250, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_250})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_251, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_251})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_252, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_252})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_253, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_253})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_254, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_254})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_255, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_255})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_256, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_256})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_257, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_257})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_258, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_258})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_07, SNum: "", OddsType: Usc8for20Kit.OT_259, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_259})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_260, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_260})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_261, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_261})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_262, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_262})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_263, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_263})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_264, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_264})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_265, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_265})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_266, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_266})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_267, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_267})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_268, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_268})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_269, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_269})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_270, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_270})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_271, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_271})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_272, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_272})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_273, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_273})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_274, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_274})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_275, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_275})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_276, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_276})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_277, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_277})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_278, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_278})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_279, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_279})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_280, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_280})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_281, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_281})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_282, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_282})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_283, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_283})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_284, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_284})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_285, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_285})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_286, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_286})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_287, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_287})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_288, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_288})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_289, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_289})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_290, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_290})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_291, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_291})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_292, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_292})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_293, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_293})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_294, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_294})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_295, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_295})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_08, SNum: "", OddsType: Usc8for20Kit.OT_296, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_296})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_297, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_297})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_298, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_298})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_299, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_299})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_300, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_300})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_301, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_301})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_302, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_302})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_303, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_303})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_304, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_304})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_305, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_305})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_306, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_306})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_307, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_307})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_308, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_308})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_309, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_309})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_310, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_310})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_311, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_311})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_312, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_312})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_313, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_313})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_314, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_314})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_315, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_315})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_99, SNum: "", OddsType: Usc8for20Kit.OT_316, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_316})

		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_09, SNum: "", OddsType: Usc8for20Kit.OT_317, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_317})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_09, SNum: "", OddsType: Usc8for20Kit.OT_318, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_318})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_09, SNum: "", OddsType: Usc8for20Kit.OT_319, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_319})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_09, SNum: "", OddsType: Usc8for20Kit.OT_320, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_320})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_09, SNum: "", OddsType: Usc8for20Kit.OT_321, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_321})
		arr = append(arr, LoOddsInfo{OneUserMaxBet: 999999, AllUserMaxBet: 999999, GameType: GameType, BigType: Usc8for20Kit.GT_09, SNum: "", OddsType: Usc8for20Kit.OT_322, Odds: 1.9, OddsDes: Usc8for20Kit.N_OT_322})

		_, e := o.InsertMulti(len(arr), arr)
		return e
	}
	return nil
}

func InitOneUscOddsInfo() error {
	InitOneUsc5OddsInfo(mconst.GameType_USC_cqssc)
	InitOneUsc5OddsInfo(mconst.GameType_USC_jsssc)
	InitOneUsc5OddsInfo(mconst.GameType_USC_ygcyc)
	InitOneUsc5OddsInfo(mconst.GameType_USC_ygssc)
	InitOneUsc5OddsInfo(mconst.GameType_USC_gzxy5)
	InitOneUsc5OddsInfo(mconst.GameType_USC_yxssc)

	InitOneUsc10OddsInfo(mconst.GameType_USC_bjsc)
	InitOneUsc10OddsInfo(mconst.GameType_USC_xyft)
	InitOneUsc10OddsInfo(mconst.GameType_USC_jskc)
	InitOneUsc10OddsInfo(mconst.GameType_USC_jssc)
	InitOneUsc10OddsInfo(mconst.GameType_USC_ESPsm)
	InitOneUsc10OddsInfo(mconst.GameType_USC_ygxyft)
	InitOneUsc10OddsInfo(mconst.GameType_USC_ygsc)
	InitOneUsc10OddsInfo(mconst.GameType_USC_gzxy10)

	InitOneUsc8for20OddsInfo(mconst.GameType_USC_cqxync)
	InitOneUsc8for20OddsInfo(mconst.GameType_USC_gdkl10f)
	return nil
}

func GetAllLoOddsInfo(GameType int) ([]LoOddsInfo, error) {
	arrData := make([]LoOddsInfo, 0)
	o := orm.NewOrm()
	sql := fmt.Sprintf(`select a.* from %s a where a.game_type=?`, mconst.TableName_LoOddsInfo)
	_, e := o.Raw(sql, GameType).QueryRows(&arrData)
	return arrData, e
}
func (this *LoOddsInfo) Add(o orm.Ormer) error {
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

func (this *LoOddsInfo) Update(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	this.UpdatedAt = time.Now()
	arrC := []string{"UpdatedAt"}
	arrC = append(arrC, cols...)

	id, e := o.Update(this, arrC...)
	if e == nil {
		this.Id = int(id)
	}

	return e
}
