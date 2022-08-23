package base

import (
	"encoding/json"
	"github.com/TtMyth123/LotteryResultSite/controllers/base/enums"
	"github.com/TtMyth123/kit/stringKit"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego"
)

type ABaseController struct {
	beego.Controller
}
type BaseSysUser struct {
	UserId int
	Name   string
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
