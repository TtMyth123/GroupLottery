package base

import "ttmyth123/GroupLottery/LotteryResultServer/controllers/base/enums"

type JsonResult struct {
	Code enums.JsonResultCode `json:"code"`
	Msg  string               `json:"msg"`
	Obj  interface{}          `json:"obj"`
}
