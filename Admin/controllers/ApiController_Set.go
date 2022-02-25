package controllers

import (
	"ttmyth123/GroupLottery/Admin/OtherServer/GameClientHttp"
	"ttmyth123/GroupLottery/Admin/controllers/SetBll"
	"ttmyth123/GroupLottery/Admin/controllers/base/enums"
)

func (this *ApiController) GetOddsList() {
	type TmpArge struct {
		GameType   int
		ArrBigType []int
	}
	aTmp := TmpArge{}
	e := this.GetJsonData(&aTmp)
	if e != nil {
		this.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	data, e := SetBll.GetOddsList(aTmp.GameType, aTmp.ArrBigType)
	if e != nil {
		this.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}

	this.JsonResult(enums.JRCodeSucc, "", data)
}

func (c *ApiController) GetOddsListFirstId() {
	GameType, _ := c.GetInt("GameType", 0)

	OddsType, _ := c.GetInt("OddsType", 0)
	BigType, _ := c.GetInt("BigType", 0)

	pageIndex, _ := c.GetInt("pageIndex", 0)
	pageSize, _ := c.GetInt("pageSize", 0)
	FirstId, _ := c.GetInt("FirstId", 0)
	OddsDes := c.GetString("OddsDes")

	PageTotal, data, groupData := SetBll.GetOddsListFirstId(OddsDes, GameType, OddsType, BigType, pageIndex, pageSize, FirstId)
	c.JsonPageResultGroup(enums.JRCodeSucc, "", PageTotal, data, groupData)
}

func (this *ApiController) SaveOddsInfo() {
	Id, _ := this.GetInt("Id")
	odds, _ := this.GetFloat("Odds")
	OddsDes := this.GetString("OddsDes")
	OneUserMaxBet, _ := this.GetInt("OneUserMaxBet")
	AllUserMaxBet, _ := this.GetInt("AllUserMaxBet")
	OneUserMinBet, _ := this.GetInt("OneUserMinBet")

	aOddsInfo, e := SetBll.GetOddsInfo(Id)
	if e != nil {
		this.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}

	aOddsInfo.Odds = odds
	aOddsInfo.OddsDes = OddsDes
	aOddsInfo.OneUserMaxBet = OneUserMaxBet
	aOddsInfo.AllUserMaxBet = AllUserMaxBet
	aOddsInfo.OneUserMinBet = OneUserMinBet

	e = aOddsInfo.Update(nil, "Odds", "OddsDes", "OneUserMaxBet", "AllUserMaxBet", "OneUserMinBet")

	if e != nil {
		this.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}

	e = GameClientHttp.OddsChange(aOddsInfo.GameType)
	if e != nil {
		this.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	this.JsonResult(enums.JRCodeSucc, "", "")
}

func (this *ApiController) SaveBatchOdds() {
	GameType, _ := this.GetInt("GameType")
	odds, _ := this.GetFloat("Odds")
	MinOddType, _ := this.GetInt("MinOddType")
	MaxOddType, _ := this.GetInt("MaxOddType")

	e := SetBll.SaveBatchOdds(GameType, MinOddType, MaxOddType, odds)
	if e != nil {
		this.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}

	e = GameClientHttp.OddsChange(GameType)
	if e != nil {
		this.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	this.JsonResult(enums.JRCodeSucc, "", "")
}

func (this *ApiController) GetTtRebateSet() {
	data := SetBll.GetTtRebateSet(0)

	this.JsonResult(enums.JRCodeSucc, "", data)
}

func (this *ApiController) SaveTtRebateSet() {
	//Id,_:= this.GetInt("Id")
	Level, _ := this.GetInt("Level", 0)
	BetRebateRatio, _ := this.GetFloat("BetRebateRatio")
	BetRebateRatio1, _ := this.GetFloat("BetRebateRatio1")
	BetRebateRatio2, _ := this.GetFloat("BetRebateRatio2")
	BetRebateRatio3, _ := this.GetFloat("BetRebateRatio3")
	data := SetBll.GetTtRebateSet(0)
	data.BetRebateRatio = BetRebateRatio
	data.BetRebateRatio1 = BetRebateRatio1
	data.BetRebateRatio2 = BetRebateRatio2
	data.BetRebateRatio3 = BetRebateRatio3
	data.Level = Level
	e := SetBll.SaveTtRebateSet(data)

	if e != nil {
		this.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	this.JsonResult(enums.JRCodeSucc, "", "")
}

func (this *ApiController) GetWsxZbcOddsList() {
	//type TmpArge struct {
	//}
	//aTmp := TmpArge{}
	//e := this.GetJsonData(&aTmp)
	//if e != nil {
	//	this.JsonResult(enums.JRCodeFailed, e.Error(), "")
	//}
	data, e := SetBll.GetWsxZbcOddsList()
	if e != nil {
		this.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}

	this.JsonResult(enums.JRCodeSucc, "", data)
}
func (this *ApiController) SaveWsxZbcOddsInfo() {
	Id, _ := this.GetInt("Id")
	odds, _ := this.GetFloat("Odds")
	//OddsDes := this.GetString("OddsDes")
	OneUserMaxBet, _ := this.GetInt("OneUserMaxBet")
	AllUserMaxBet, _ := this.GetInt("AllUserMaxBet")
	OneUserMinBet, _ := this.GetInt("OneUserMinBet")

	aOddsInfo, e := SetBll.GetOddsInfo(Id)
	if e != nil {
		this.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}

	aOddsInfo.Odds = odds
	//aOddsInfo.OddsDes = OddsDes
	aOddsInfo.OneUserMaxBet = OneUserMaxBet
	aOddsInfo.AllUserMaxBet = AllUserMaxBet
	aOddsInfo.OneUserMinBet = OneUserMinBet
	e = SetBll.SaveWsxZbcOddsInfo(aOddsInfo)
	if e != nil {
		this.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	this.JsonResult(enums.JRCodeSucc, "", "")
}
