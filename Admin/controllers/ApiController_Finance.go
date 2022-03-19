package controllers

import (
	"ttmyth123/GroupLottery/Admin/CacheData"
	//"ttmyth123/DummyMarket/YhFilmAdmin/CacheData"
	"ttmyth123/GroupLottery/Admin/controllers/AFinanceBll"
	"ttmyth123/GroupLottery/Admin/controllers/base/enums"
	"ttmyth123/kit/httpKit/TmpFileKit"
)

func (this *ApiController) GetSaveDrawApplyInfo() {
	data := AFinanceBll.GetSaveDrawApplyInfo()

	this.JsonResult(enums.JRCodeSucc, "", data)
}

func (c *ApiController) GetDrawMoneyApplyList() {
	userId, _ := c.GetInt("userId", 0)
	beginDay := c.GetString("beginDay")
	endDay := c.GetString("endDay")
	userName := c.GetString("userName")
	pageIndex, _ := c.GetInt("pageIndex", 1)
	pageSize, _ := c.GetInt("pageSize", 10)
	FirstId, _ := c.GetInt("FirstId", 0)

	curAgentId := c.CurSysUserEx().GameId
	PageTotal, data, groupData := AFinanceBll.GetDrawMoneyApplyList(curAgentId, userName, beginDay, endDay, userId, pageIndex, pageSize, FirstId)
	c.JsonPageResultGroup(enums.JRCodeSucc, "", PageTotal, data, groupData)
}
func (c *ApiController) GetSaveMoneyApplyList() {
	userId, _ := c.GetInt("userId", 0)
	beginDay := c.GetString("beginDay")
	endDay := c.GetString("endDay")
	userName := c.GetString("userName")
	pageIndex, _ := c.GetInt("pageIndex", 1)
	pageSize, _ := c.GetInt("pageSize", 10)
	FirstId, _ := c.GetInt("FirstId", 0)
	curAgentId := c.CurSysUserEx().GameId

	PageTotal, data, groupData := AFinanceBll.GetSaveMoneyApplyList(curAgentId, userName, beginDay, endDay, userId, pageIndex, pageSize, FirstId)
	c.JsonPageResultGroup(enums.JRCodeSucc, "", PageTotal, data, groupData)
}

func (c *ApiController) SaveMoney() {
	UserId, _ := c.GetInt("UserId", 0)
	Money, _ := c.GetInt("Money", 0)
	GroupId, _ := c.GetInt("GroupId", 0)

	sysUser := c.CurSysUserEx()
	e := AFinanceBll.SaveMoney(GroupId, UserId, Money, sysUser.Id, sysUser.UserName)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}
func (c *ApiController) DrawMoney() {
	UserId, _ := c.GetInt("UserId", 0)
	Money, _ := c.GetInt("Money", 0)
	sysUser := c.CurSysUserEx()
	e := AFinanceBll.DrawMoney(UserId, Money, sysUser.Id, sysUser.UserName)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *ApiController) AgreeSaveMoney() {
	UserId, _ := c.GetInt("UserId", 0)
	SaveMoneyId, _ := c.GetInt("SaveMoneyId", 0)
	AuditorId := c.CurSysUserEx().Id
	AuditorName := c.CurSysUserEx().UserName
	e := AFinanceBll.AgreeSaveMoney(UserId, SaveMoneyId, AuditorId, AuditorName)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *ApiController) AgreeDrawMoney() {
	SaveMoneyId, _ := c.GetInt("id", 0)

	e := AFinanceBll.AgreeDrawMoney(SaveMoneyId, c.CurSysUserEx().Id, c.CurSysUserEx().UserName)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *ApiController) GetAccountList() {
	userName := c.GetString("userName")
	accountType, _ := c.GetInt("accountType", 0)
	userId, _ := c.GetInt("userId", 0)
	pageIndex, _ := c.GetInt("pageIndex", 0)
	pageSize, _ := c.GetInt("pageSize", 0)
	beginDay := c.GetString("beginDay")
	endDay := c.GetString("endDay")
	FirstId, _ := c.GetInt("FirstId", 0)

	curAgentId := c.CurSysUserEx().GameId
	if c.CurSysUserEx().IsSuper {
		curAgentId = 0
	}

	PageTotal, data, groupData := AFinanceBll.GetAccountList(curAgentId, beginDay, endDay, userName, userId, accountType, pageIndex, pageSize, FirstId)
	c.JsonPageResultGroup(enums.JRCodeSucc, "", PageTotal, data, groupData)
}

//func (c *ApiController) GetAccountListFirstId() {
//	userName := c.GetString("userName")
//	accountType, _ := c.GetInt("accountType", 0)
//	pageIndex, _ := c.GetInt("pageIndex", 0)
//	pageSize, _ := c.GetInt("pageSize", 0)
//	beginDay := c.GetString("beginDay")
//	endDay := c.GetString("endDay")
//	FirstId, _ := c.GetInt("FirstId", 0)
//
//	PageTotal, data,groupData := AFinanceBll.GetAccountListFirstId(beginDay, endDay, userName, accountType, pageIndex, pageSize,FirstId)
//	c.JsonPageResultGroup(enums.JRCodeSucc, "", PageTotal, data, groupData)
//}

func (c *ApiController) DrawMoneyList() {
	userName := c.GetString("userName")
	userId, _ := c.GetInt("userId", 0)
	pageIndex, _ := c.GetInt("pageIndex", 0)
	pageSize, _ := c.GetInt("pageSize", 0)
	beginDay := c.GetString("beginDay")
	endDay := c.GetString("endDay")
	FirstId, _ := c.GetInt("FirstId", 0)
	curAgentId := c.CurSysUserEx().GameId

	PageTotal, data, groupData := AFinanceBll.DrawMoneyList(curAgentId, beginDay, endDay, userName, userId, pageIndex, pageSize, FirstId)
	c.JsonPageResultGroup(enums.JRCodeSucc, "", PageTotal, data, groupData)
}

func (c *ApiController) SaveMoneyList() {
	userName := c.GetString("userName")
	userId, _ := c.GetInt("userId", 0)
	pageIndex, _ := c.GetInt("pageIndex", 0)
	pageSize, _ := c.GetInt("pageSize", 0)
	beginDay := c.GetString("beginDay")
	endDay := c.GetString("endDay")
	FirstId, _ := c.GetInt("FirstId", 0)
	curAgentId := c.CurSysUserEx().GameId

	PageTotal, data, groupData := AFinanceBll.SaveMoneyList(curAgentId, beginDay, endDay, userName, userId, pageIndex, pageSize, FirstId)
	c.JsonPageResultGroup(enums.JRCodeSucc, "", PageTotal, data, groupData)
}

/**
删除提现(下分)申请
*/
func (this *ApiController) DelSaveMoney() {
	id, _ := this.GetInt("Id")
	typeScore, e := AFinanceBll.DelSaveMoney(id, 0)
	if e != nil {
		this.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}

	CacheData.DelScoreRequisition(typeScore)

	this.JsonResult(enums.JRCodeSucc, "", "")
}

/**
删除提现(下分)申请
*/
func (this *ApiController) DelDrawMoney() {
	id, _ := this.GetInt("Id")
	Excuse := this.GetString("Excuse")
	e := AFinanceBll.DelDrawMoney(id, 0, Excuse)
	if e != nil {
		this.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}

	this.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *ApiController) GetRebateList() {
	userName := c.GetString("userName")
	RebateType, _ := c.GetInt("RebateType", 0)
	pageIndex, _ := c.GetInt("pageIndex", 0)
	pageSize, _ := c.GetInt("pageSize", 0)
	beginDay := c.GetString("beginDay")
	endDay := c.GetString("endDay")
	FirstId, _ := c.GetInt("FirstId", 0)
	userId, _ := c.GetInt("userId", 0)

	curAgentId := c.CurSysUserEx().GameId
	if c.CurSysUserEx().IsSuper {
		curAgentId = 0
	}

	PageTotal, data, groupData := AFinanceBll.GetRebateList(curAgentId, userId, RebateType, userName, beginDay, endDay, pageIndex, pageSize, FirstId)
	c.JsonPageResultGroup(enums.JRCodeSucc, "", PageTotal, data, groupData)
}

//func (c *ApiController) GetRebateListFirstId() {
//	userName := c.GetString("userName")
//	RebateType, _ := c.GetInt("RebateType", 0)
//	pageIndex, _ := c.GetInt("pageIndex", 0)
//	pageSize, _ := c.GetInt("pageSize", 0)
//	beginDay := c.GetString("beginDay")
//	endDay := c.GetString("endDay")
//
//	FirstId, _ := c.GetInt("FirstId", 0)
//	PageTotal, data, groupData := AFinanceBll.GetRebateListFirstId(RebateType, userName, beginDay, endDay, pageIndex, pageSize,FirstId)
//	c.JsonPageResultGroup(enums.JRCodeSucc, "", PageTotal, data, groupData)
//}

func (c *ApiController) GetFinanceAccount() {
	data, e := AFinanceBll.GetFinanceAccount()
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", data)
}

//func (c *ApiController) UpdateFinanceAccount() {
//	data, e := AFinanceBll.GetFinanceAccount()
//	if e != nil {
//		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
//	}
//	data.PayWay,_ = c.GetInt("PayWay",0)
//
//	data.OnlinePay = c.GetString("OnlinePay")
//	data.AlipayName = c.GetString("AlipayName")
//	data.AlipayUrl = c.GetString("AlipayUrl")
//
//	data.BankCard = c.GetString("BankCard")
//	data.BankName = c.GetString("BankName")
//	data.BankUser = c.GetString("BankUser")
//	data.UserMobile = c.GetString("UserMobile")
//
//	data.WXReceiptUrl = c.GetString("WXReceiptUrl")
//
//	imgFile, head, err := c.GetFile("AlipayQrcodeUrlFile")
//	if err == nil {
//		filePath, _ := httpKit.UploadFile(imgFile, head, "static/upload/FinanceAccount", "AlipayQrcodeUrlFile", "")
//		data.AlipayUrl = filePath
//	}
//
//	imgFile, head, err = c.GetFile("WXQrcodeUrlFile")
//	if err == nil {
//		filePath, _ := httpKit.UploadFile(imgFile, head, "static/upload/FinanceAccount", "WXQrcodeUrlFile", "")
//		data.WXReceiptUrl = filePath
//	}
//	e = data.Update()
//	if e != nil {
//		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
//	}
//	c.JsonResult(enums.JRCodeSucc, "", "")
//}
func (c *ApiController) UpdateFinanceAccount() {
	data, e := AFinanceBll.GetFinanceAccount()
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}

	AlipayQrcodePicGuid := c.GetString("AlipayQrcodePicGuid")
	if AlipayQrcodePicGuid != "" {
		filePath, e := TmpFileKit.ToNewPath(AlipayQrcodePicGuid, `static/upload/FinanceAccount`, "")
		if e != nil {
			c.JsonResult(enums.JRCodeFailed, "", "没有图片数据")
		}

		data.AlipayUrl = filePath
	}

	WXQrcodePicGuid := c.GetString("WXQrcodePicGuid")
	if WXQrcodePicGuid != "" {
		filePath, e := TmpFileKit.ToNewPath(WXQrcodePicGuid, `static/upload/FinanceAccount`, "")
		if e != nil {
			c.JsonResult(enums.JRCodeFailed, "", "没有图片数据")
		}

		data.WXReceiptUrl = filePath
	}

	data.PayWay, _ = c.GetInt("PayWay", 0)
	data.OnlinePay = c.GetString("OnlinePay")
	data.AlipayName = c.GetString("AlipayName")

	data.BankAddr = c.GetString("BankAddr")
	data.BankCard = c.GetString("BankCard")
	data.BankName = c.GetString("BankName")
	data.BankUser = c.GetString("BankUser")
	data.UserMobile = c.GetString("UserMobile")

	e = data.Update()
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

/**
删除提现(下分)申请
*/
func (this *ApiController) GetDrawMoneyAccountInfo() {
	id, _ := this.GetInt("DrawMoneyId")

	aDrawMoneyAccountInfo, e := AFinanceBll.GetDrawMoneyAccountInfo(id)
	if e != nil {
		this.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}

	this.JsonResult(enums.JRCodeSucc, "", aDrawMoneyAccountInfo)
}
