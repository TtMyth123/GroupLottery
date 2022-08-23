package WsxNbcResultKit

import (
	"fmt"
	"github.com/TtMyth123/GameServer/LotteryResult/WsxBox"
	"github.com/TtMyth123/GameServer/models"
	"github.com/TtMyth123/GameServer/models/mconst"
	"github.com/TtMyth123/kit/stringKit"
	"github.com/astaxie/beego/orm"
	"testing"
)

var (
	mpLoOddsInfo map[int]models.LoOddsInfo
)

//func TestMain(m *testing.M)  {
//	ttLog.InitLogs()
//	//models.Init()
//}
func iniDataMpLoOddsInfo() {
	mpLoOddsInfo := make(map[int]models.LoOddsInfo)
	arrData, _ := models.GetAllLoOddsInfo(mconst.GameType_Wsx_202)
	for i := 0; i < len(arrData); i++ {
		mpLoOddsInfo[arrData[i].OddsType] = arrData[i]
	}
}
func GetLoBetInfo(id int) (models.LoBetInfo, error) {
	bet := models.LoBetInfo{Id: id}
	o := orm.NewOrm()
	e := o.Read(&bet)
	return bet, e
}

func TestGetEndStr(t *testing.T) {
	jackpots := "123456"

	a1 := GetEndStr(jackpots, 3)
	fmt.Println("a1:", a1)

	a1 = GetBeginStr(jackpots, 7)
	fmt.Println("a1:", a1)
}

func TestGetBeginStr(t *testing.T) {
	jackpots := "123456"

	a1 := GetBeginStr(jackpots, 1)
	fmt.Println("a1:", a1)
}

func GetMpLoOddsInfo(GameType int) map[int]models.LoOddsInfo {
	mpLoOdds := make(map[int]models.LoOddsInfo)
	arr := models.GetInitOneLoOddsInfo(GameType)
	for _, v := range arr {
		mpLoOdds[v.OddsType] = v
	}

	return mpLoOdds
}

func TestComputeLoseWin(t *testing.T) {
	mpLoOdds := GetMpLoOddsInfo(mconst.GameType_Wsx_202)
	fmt.Println("a1:aaa")
	aOpenCodeInfo := WsxBox.OpenCodeInfoNbc{}

	strC := `{"eighthNum":null,"fifthNum":["3018","6857","2805","6012","7802","2051"],"firstNum":"42840","forthNum":["5812","5760","0179","6915"],"jackpots":"62428","secondNum":["13867","38422"],"seventhNum":["33","79","50","26"],"sixthNum":["385","480","556"],"thirdNum":["87100","43529","98334","47034","95413","95247"]}`
	e := stringKit.GetJsonObj(strC, &aOpenCodeInfo)
	if e != nil {

	}
	bet := models.LoBetInfo{OddsType: 128, BetM: 10}

	f := ComputeLoseWin(bet, aOpenCodeInfo, mpLoOdds)
	fmt.Println("a1:", f)
}
