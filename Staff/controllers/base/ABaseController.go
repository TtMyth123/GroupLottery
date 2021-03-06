package base

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"ttmyth123/GroupLottery/Staff/CacheData"
	"ttmyth123/GroupLottery/Staff/controllers/base/box"
	"ttmyth123/GroupLottery/Staff/controllers/base/enums"
	"ttmyth123/kit"
	"ttmyth123/kit/TtErrors"
	"ttmyth123/kit/stringKit"
	"ttmyth123/kit/ttLog"
)

type ABaseController struct {
	beego.Controller
}

func (this *ABaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	ttLog.LogDebug(controllerName, actionName)
}

func (this *ABaseController) GetJsonData(v interface{}) error {
	strJson := this.GetString("jsonData")
	controllerName, actionName := this.GetControllerAndAction()
	ttLog.LogDebug(controllerName, actionName, "jsonData:", strJson)
	e := json.Unmarshal([]byte(strJson), v)
	ttLog.LogDebug("Data:", stringKit.GetJsonStr(v))
	return e
}

func (this *ABaseController) JsonResult(code enums.JsonResultCode, msg string, obj interface{}) {
	res := JsonResult{Code: code, Msg: msg, Obj: obj}
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *ABaseController) JsonResultEx(e TtErrors.TtError, obj interface{}) {
	if e == nil {
		res := JsonResult{Code: enums.JRCodeSucc, Obj: obj}
		this.Data["json"] = res
	} else {
		res := JsonResult{Code: enums.JRCodeFailed, Msg: e.Error(), ErrCode: e.Code(), Obj: obj}
		this.Data["json"] = res
	}
	this.ServeJSON()
	this.StopRun()
}

func (this *ABaseController) JsonListResult(code enums.JsonResultCode, msg string, lastId int, obj interface{}) {
	listR := ListResult{LastId: lastId, ListData: obj}
	res := JsonListResult{Code: code, Msg: msg, Obj: listR}
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *ABaseController) JsonPageResult(code enums.JsonResultCode, msg string, PageTotal int, obj interface{}) {
	listR := PageResult{PageTotal: PageTotal, ListData: obj}
	res := JsonPageResult{Code: code, Msg: msg, Obj: listR}
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *ABaseController) JsonPageResultGroup(code enums.JsonResultCode, msg string, PageTotal int, obj interface{}, groupObj interface{}) {
	listR := PageResult{PageTotal: PageTotal, ListData: obj, GroupData: groupObj}
	res := JsonPageResult{Code: code, Msg: msg, Obj: listR}
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *ABaseController) GetToken(userInfo box.BaseSysUser) (string, error) {
	key := kit.GetGuid()

	aUser := box.SysUserEx{}
	aUser.BaseSysUser = userInfo
	CacheData.SetToken(key, aUser)

	return key, nil
}
