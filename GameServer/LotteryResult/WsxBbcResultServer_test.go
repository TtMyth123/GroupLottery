package LotteryResult

import (
	"fmt"
	"testing"
	"ttmyth123/GroupLottery/GameServer/controllers/base/TtError"
	"ttmyth123/kit"
	"ttmyth123/kit/stringKit"
)

func TestAA(t *testing.T) {
	a := aaa()
	b := a.(*TtError.TtError)
	fmt.Print(a, "b:", stringKit.GetJsonStr(b))

}
func aaa() error {
	return TtError.New("aaa", nil)
}

func TestAA1(t *testing.T) {
	aaa := kit.GetGuid()
	aaa = aaa + aaa
	fmt.Print("bu:", aaa[:50], "___")
}
