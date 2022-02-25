package controllers

import (
	"ttmyth123/GroupLottery/GameServer/GInstance"
	"ttmyth123/GroupLottery/GameServer/GInstance/GTtHint"
	"ttmyth123/GroupLottery/GameServer/controllers/base"
	"ttmyth123/GroupLottery/GameServer/controllers/base/enums"
	"ttmyth123/GroupLottery/GameServer/controllers/financeBll"
	"ttmyth123/GroupLottery/GameServer/controllers/otherBll"
)

/**
（上分）充值申请
*/
func (this *ApiController) ApplySaveMoney() {
	type TmpArgs struct {
		base.ArgsBox
		Score float64
	}

	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	id, e := financeBll.SaveMoneyApply(aTmpArgs.UserId, aTmpArgs.Score)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	this.JsonResult(enums.JRCodeSucc, "", id, nil)
}

/**
上传支付凭证
*/
func (this *ApiController) UploadSaveMoneyVoucher() {
	type TmpArgs struct {
		base.ArgsBox
		SaveMoneyId int
		State       int
		VoucherUrl  string
	}

	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	e = financeBll.UploadSaveMoneyVoucher(aTmpArgs.UserId, aTmpArgs.SaveMoneyId, aTmpArgs.State, aTmpArgs.VoucherUrl)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", "", nil)
}

/**
（下分）提现申请 applydrawmoney
*/
func (this *ApiController) ApplyDrawMoney() {
	type TmpArgs struct {
		base.ArgsBox
		Score float64
		Pwd   string
	}

	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	e = financeBll.DrawMoneyApply(aTmpArgs.GroupId, aTmpArgs.UserId, aTmpArgs.Score, aTmpArgs.Pwd)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", "", nil)
}

/**
删除充值(上分)申请
*/
func (this *ApiController) DelSaveMoney() {
	type TmpArgs struct {
		base.ArgsBox
		SaveMoneyId int
	}

	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	_, e = financeBll.DelSaveMoney(aTmpArgs.SaveMoneyId, aTmpArgs.UserId)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	//CacheData.DelScoreRequisition(typeScore)

	this.JsonResult(enums.JRCodeSucc, "", "", nil)
}

/**
删除提现(下分)申请
*/
func (this *ApiController) DelDrawMoney() {
	type TmpArgs struct {
		base.ArgsBox
		DrawMoneyId int
	}

	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	e = financeBll.DelDrawMoney(aTmpArgs.DrawMoneyId, aTmpArgs.UserId)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	this.JsonResult(enums.JRCodeSucc, "", "", nil)
}

/**
获取上下分列表
*/
func (this *ApiController) GetDrawSaveMoneyList() {
	type TmpArgs struct {
		base.ArgsBox
		PageIndex int
		PageSize  int
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	if aTmpArgs.PageSize <= 0 {
		aTmpArgs.PageSize = 10
	}
	if aTmpArgs.PageIndex <= 0 {
		aTmpArgs.PageIndex = 1
	}

	PageTotal, arrData := financeBll.GetDrawSaveMoneyList(aTmpArgs.UserId, aTmpArgs.PageIndex, aTmpArgs.PageSize)
	this.JsonResultCPage(enums.JRCodeSucc, "", PageTotal, arrData)
}

/**
设置提现密码
*/
func (this *ApiController) SetDrawMoneyPwd() {
	type TmpArgs struct {
		base.ArgsBox
		NewPwd string
		OldPwd string
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	e = financeBll.SetDrawMoneyPwd(aTmpArgs.UserId, aTmpArgs.OldPwd, aTmpArgs.NewPwd)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	this.JsonResult(enums.JRCodeSucc, "", "", nil)
}

func (this *ApiController) GetReport() {
	type TmpArgs struct {
		base.ArgsBox
		BeginDay string
		EndDay   string
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	data := financeBll.GetReport(aTmpArgs.GameType, aTmpArgs.UserId, aTmpArgs.BeginDay, aTmpArgs.EndDay)
	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}

func (this *ApiController) Rebate2Gold() {
	type TmpArgs struct {
		base.ArgsBox
		Score float64
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	u, e := GInstance.GetUserRpcClient().Rebate2Gold(aTmpArgs.UserId, aTmpArgs.Score, "佣金转换",
		GTtHint.GetTtHint().GetHint("佣金转换"), GTtHint.GetTtHint().GetMpString())
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	type TmpResult struct {
		UserId  int
		TjMoney float64
		Money   float64
	}

	aTmpResult := TmpResult{
		UserId:  u.Id,
		TjMoney: u.Rebate,
		Money:   u.Gold,
	}

	this.JsonResult(enums.JRCodeSucc, "", aTmpResult, nil)
}

func (this *ApiController) Rebate2GoldRecord() {
	type TmpArgs struct {
		base.ArgsBox
		PageIndex int
		PageSize  int

		MaxId int
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	dataList, _, PageTotal, e := GInstance.GetUserRpcClient().Rebate2GoldRecord(aTmpArgs.UserId, aTmpArgs.PageIndex, aTmpArgs.PageSize)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	aUser, _ := GInstance.GetUserRpcClient().GetUser(aTmpArgs.UserId)

	this.JsonResultCPageEx(enums.JRCodeSucc, "", PageTotal, dataList, aUser.Rebate)
}

//getagentpayinfo
func (this *ApiController) GetAgentPayInfo() {
	type TmpArgs struct {
		base.ArgsBox
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	data, e := financeBll.GetAgentPayInfo(aTmpArgs.Area)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}
func (this *ApiController) GetAgentPayInfoEx() {
	type TmpArgs struct {
		base.ArgsBox
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	data, e := financeBll.GetAgentPayInfoEx(aTmpArgs.Area)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}

func (this *ApiController) ChangePayInfo() {
	type TmpArgs struct {
		base.ArgsBox
	}

	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.JsonResult(enums.JRCodeFailed, e.Error(), "", nil)
	}
	e = otherBll.ChangePayInfo(aTmpArgs.UserId)
	if e != nil {
		this.JsonResult(enums.JRCodeFailed, e.Error(), "", nil)
	}
	this.JsonResult(enums.JRCodeSucc, "", "", nil)
}

/**
上分
*/
func (this *ApiController) SaveMoney() {
	type TmpArgs struct {
		base.ArgsBox
		Score float64
	}

	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	e = financeBll.SaveMoney(aTmpArgs.UserId, aTmpArgs.Score, 0, "")
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	this.JsonResult(enums.JRCodeSucc, "", "", nil)
}

/**
下分
*/
func (this *ApiController) DrawMoney() {
	type TmpArgs struct {
		base.ArgsBox
		Score float64
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	e = financeBll.DrawMoney(aTmpArgs.GroupId, aTmpArgs.UserId, aTmpArgs.Score, 0, "")

	this.JsonResult(enums.JRCodeSucc, "", "", nil)
}

/**
下分待审核
State = 1，下分，审核ing
State = 2，下分审核OK，
State = 3，下分取消，
*/
func (this *ApiController) DrawMoneying() {
	type TmpArgs struct {
		base.ArgsBox
		Score   float64
		Money   float64
		State   int
		OrderId string `json:"orderId"`
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	e = financeBll.DrawMoneying(aTmpArgs.GroupId, aTmpArgs.UserId, aTmpArgs.Score, aTmpArgs.Money, aTmpArgs.State, aTmpArgs.OrderId)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	this.JsonResult(enums.JRCodeSucc, "", "", nil)
}

func AgentAccountInfo() {

}
