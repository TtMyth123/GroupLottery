package base

import "ttmyth123/GroupLottery/Staff/controllers/base/enums"

type JsonResult struct {
	Code    enums.JsonResultCode `json:"code"`
	ErrCode int                  `json:"ErrCode"`
	Msg     string               `json:"msg"`
	Obj     interface{}          `json:"obj"`
}

type ListResult struct {
	LastId   int
	ListData interface{}
}

type JsonListResult struct {
	Code enums.JsonResultCode `json:"code"`
	Msg  string               `json:"msg"`
	Obj  ListResult           `json:"obj"`
}

type PageResult struct {
	PageTotal int
	ListData  interface{}
	GroupData interface{}
}
type JsonPageResult struct {
	Code enums.JsonResultCode `json:"code"`
	Msg  string               `json:"msg"`
	Obj  PageResult           `json:"obj"`
}
