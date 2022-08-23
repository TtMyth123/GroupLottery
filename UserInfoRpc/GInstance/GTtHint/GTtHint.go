package GTtHint

import (
	"github.com/TtMyth123/UserInfoRpc/GConfig/TtHint"
)

var (
	mTtHint *TtHint.TtHint
)

func Init() {
	mTtHint = TtHint.NewTtHint("conf/i18n", "dict")
}

func GetTtHint() *TtHint.TtHint {
	return mTtHint
}
