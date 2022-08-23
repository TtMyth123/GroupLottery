package base

import "github.com/TtMyth123/LotteryResultServer/controllers/base/enums"

type JsonResult struct {
	Code enums.JsonResultCode `json:"code"`
	Msg  string               `json:"msg"`
	Obj  interface{}          `json:"obj"`
}
