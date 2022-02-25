package base

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"ttmyth123/GroupLottery/LotteryResultServer/controllers/base/cBo"
	"ttmyth123/GroupLottery/LotteryResultServer/controllers/base/enums"
	"ttmyth123/kit/stringKit"
	"ttmyth123/kit/ttLog"
)

var isDev = false

type BaseController struct {
	beego.Controller
}

func (this *BaseController) ErrJsonResultEx(err error) {
	if err == nil {
		this.JsonResult(enums.JRCodeSucc, "", "")
	} else {
		this.JsonResult(enums.JRCodeFailed, err.Error(), "")
	}
}

func (this *BaseController) JsonResultCPage(code enums.JsonResultCode, msg string, PageTotal int, arrData interface{}) {
	if isDev {
		if code != enums.JRCodeSucc {
			_, Action := this.GetControllerAndAction()
			msg = fmt.Sprintf("M:%s +\n %s", Action, msg)
		}
	}
	type TmpR struct {
		PageTotal int
		LastId    int
		DataList  interface{}
	}
	r := TmpR{PageTotal: PageTotal, DataList: arrData, LastId: 0}
	this.JsonResult(code, msg, r)
}

func (this *BaseController) JsonResultCPageEx(code enums.JsonResultCode, msg string, PageTotal int, arrData interface{}, OtherInfo interface{}) {
	if isDev {
		if code != enums.JRCodeSucc {
			_, Action := this.GetControllerAndAction()
			msg = fmt.Sprintf("M:%s +\n %s", Action, msg)
		}
	}
	type TmpR struct {
		PageTotal int
		LastId    int
		DataList  interface{}
		OtherInfo interface{}
	}
	r := TmpR{PageTotal: PageTotal, DataList: arrData, LastId: 0, OtherInfo: OtherInfo}
	this.JsonResult(code, msg, r)
}

func (this *BaseController) JsonResult(code enums.JsonResultCode, msg string, obj interface{}) {
	if isDev {
		if code != enums.JRCodeSucc {
			_, Action := this.GetControllerAndAction()
			msg = fmt.Sprintf("M:%s +\n %s", Action, msg)
		}
	}
	res := JsonResult{Code: code, Msg: msg, Obj: obj}
	aTmp := cBo.ReplyBox{Result: true, Data: obj}
	if code != enums.JRCodeSucc {
		aTmp.Result = false
		aTmp.ErrMsg = msg
	}

	res.Obj = aTmp
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}
func (this *BaseController) JsonResultEx(err error, obj interface{}) {
	if err == nil {
		this.JsonResult(enums.JRCodeSucc, "", obj)
	} else {
		this.JsonResult(enums.JRCodeFailed, err.Error(), obj)
	}
}

func (this *BaseController) JsonListResult(code enums.JsonResultCode, msg string, LastId int, obj interface{}) {
	if isDev {
		if code != enums.JRCodeSucc {
			_, Action := this.GetControllerAndAction()
			msg = fmt.Sprintf("M:%s +\n %s", Action, msg)
		}
	}

	a := cBo.ArgsBox{}
	this.GetJsonData(&a)
	res := JsonResult{Code: code, Msg: msg, Obj: obj}
	aTmp := cBo.ReplyListBox{Result: true, ListData: obj, LastId: LastId}
	if code != enums.JRCodeSucc {
		aTmp.Result = false
		aTmp.ErrMsg = msg
	}

	res.Obj = aTmp
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) GetJsonData(v interface{}) error {
	strJson := this.GetString("jsonData")
	controllerName, actionName := this.GetControllerAndAction()
	ttLog.LogDebug(controllerName, actionName, "jsonData:", strJson)
	e := json.Unmarshal([]byte(strJson), v)
	ttLog.LogDebug("Data:", stringKit.GetJsonStr(v))
	return e
}
