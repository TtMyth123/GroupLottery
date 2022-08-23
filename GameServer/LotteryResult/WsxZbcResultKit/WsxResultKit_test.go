package WsxZbcResultKit

import (
	"fmt"
	"github.com/TtMyth123/GameServer/LotteryResult/WsxBox"
	"github.com/TtMyth123/GameServer/models"
	"testing"
)

func TestComputeLoseWin(t *testing.T) {
	aOpenCodeInfo := WsxBox.OpenCodeZbcInfo{}
	Nums := []string{"02", "05", "07", "08", "14", "17", "21", "23", "28", "31", "36", "40", "43", "44", "46", "56", "63", "71", "74", "77"}
	aOpenCodeInfo.Nums = Nums

	bet := models.LoBetInfo{BigOddsType: 704, BetM: 10, Nums: "07,08,14,18", Odds: 0.5}
	bet.StrOdds = `{"2":1,"3":5,"4":40}`

	f, odds := ComputeLoseWin(bet, aOpenCodeInfo)
	fmt.Println("a1:", f, odds)
}
