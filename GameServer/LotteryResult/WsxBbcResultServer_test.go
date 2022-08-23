package LotteryResult

import (
	"fmt"
	"github.com/TtMyth123/GameServer/controllers/base/TtError"
	"github.com/TtMyth123/kit"
	"github.com/TtMyth123/kit/stringKit"
	"testing"
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
