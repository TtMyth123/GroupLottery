package WsxBbcResultKit

import (
	"fmt"
	"testing"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/WsxBox"
	"ttmyth123/GroupLottery/GameServer/models"
	"ttmyth123/GroupLottery/GameServer/models/mconst"
	"ttmyth123/kit/stringKit"
)

func TestAA(t *testing.T) {
	jackpots := "123456"

	a1 := GetEndStr(jackpots, 3)
	fmt.Println("a1:", a1)

	a1 = GetBeginStr(jackpots, 7)
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
	aOpenCodeInfo := WsxBox.OpenCodeInfo{}

	//strC := `{"eighthNum":null,"fifthNum":["3018","6857","2805","6012","7802","2051"],"firstNum":"42840","forthNum":["5812","5760","0179","6915"],"jackpots":"62428","secondNum":["13867","38422"],"seventhNum":["33","79","50","26"],"sixthNum":["385","480","556"],"thirdNum":["87100","43529","98334","47034","95413","95247"]}`
	strC := `{"fifthNum":["1677","9524","7016","4746","3257","3405"],"firstNum":"71602","forthNum":["7045","7208","0361","9902"],"jackpots":"42050","secondNum":["46241","14507"],"seventhNum":["11","08","00","69"],"sixthNum":["030","710","607"],"thirdNum":["74322","38732","64804","78836","98669","54004"]}`
	e := stringKit.GetJsonObj(strC, &aOpenCodeInfo)
	if e != nil {

	}
	bet := models.LoBetInfo{OddsType: 60, BetM: 10, Nums: "50,05", Odds: 0.5}

	f := ComputeLoseWin(bet, aOpenCodeInfo, mpLoOdds)
	fmt.Println("a1:", f)
}
