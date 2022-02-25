package base

import "ttmyth123/GroupLottery/GameServer/controllers/base/enums"

type JsonResult struct {
	Code   enums.JsonResultCode `json:"code"`
	Msg    string               `json:"msg"`
	Obj    interface{}          `json:"obj"`
	MsgAdd int
	Mp     interface{} `json:"mp"`
}
