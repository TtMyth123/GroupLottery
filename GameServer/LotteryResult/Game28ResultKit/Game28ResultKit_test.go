package Game28ResultKit

import (
	"fmt"
	"github.com/TtMyth123/GameServer/LotteryResult/Game28ResultKit/Game28Const"
	"github.com/TtMyth123/GameServer/models"
	"github.com/TtMyth123/kit/stringKit"
	"testing"
)

func TestComputeLoseWin(t *testing.T) {
	mpOddsInfo := make(map[int]models.LoOddsInfo)

	mpOddsInfo[Game28Const.TO_G28_021] = models.LoOddsInfo{OddsType: Game28Const.TO_G28_021}

	aZg28Result := GetZg28Result("9,5,5")

	bet := models.LoBetInfo{BetM: 10, OddsType: Game28Const.TO_G28_021, Odds: 3}

	betLostWin := ComputeLoseWin(bet, aZg28Result, mpOddsInfo)
	fmt.Println(stringKit.GetJsonStr(betLostWin))
}

func TestComputeLoseWin_34(t *testing.T) {
	mpOddsInfo := make(map[int]models.LoOddsInfo)

	mpOddsInfo[Game28Const.TO_G28_034] = models.LoOddsInfo{OddsType: Game28Const.TO_G28_034}

	aZg28Result := GetZg28Result("8,0,5") //13 win:1
	bet := models.LoBetInfo{BetM: 300, OddsType: Game28Const.TO_G28_034, Odds: 1.9}
	betLostWin := ComputeLoseWin(bet, aZg28Result, mpOddsInfo)
	fmt.Println(aZg28Result.AwardNumbers, stringKit.GetJsonStr(betLostWin), betLostWin.WinM-float64(betLostWin.BetM))

	aZg28Result = GetZg28Result("5,9,1") //15 win:2
	bet = models.LoBetInfo{BetM: 300, OddsType: Game28Const.TO_G28_034, Odds: 1.9}
	betLostWin = ComputeLoseWin(bet, aZg28Result, mpOddsInfo)
	fmt.Println(aZg28Result.AwardNumbers, stringKit.GetJsonStr(betLostWin), betLostWin.WinM-float64(betLostWin.BetM))

	aZg28Result = GetZg28Result("5,5,5") //15 win:3
	bet = models.LoBetInfo{BetM: 300, OddsType: Game28Const.TO_G28_034, Odds: 1.9}
	betLostWin = ComputeLoseWin(bet, aZg28Result, mpOddsInfo)
	fmt.Println(aZg28Result.AwardNumbers, stringKit.GetJsonStr(betLostWin), betLostWin.WinM-float64(betLostWin.BetM))
}
