package GTtHint

import (
	"ttmyth123/GroupLottery/GameServer/GConfig/TtHint"
)

var (
	mTtHint *TtHint.TtHint
)

func Init() {
	//mTtHint = TtHint.NewTtHint("GameServer/conf/i18n","dict")
	mTtHint = TtHint.NewTtHint("conf/i18n", "dict")
}

func GetTtHint() *TtHint.TtHint {
	return mTtHint
}
