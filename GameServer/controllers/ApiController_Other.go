package controllers

import (
	"errors"
	"ttmyth123/GroupLottery/GameServer/GInstance"
	"ttmyth123/GroupLottery/GameServer/controllers/base"
	"ttmyth123/GroupLottery/GameServer/controllers/base/enums"
	"ttmyth123/GroupLottery/GameServer/controllers/otherBll"
	"ttmyth123/GroupLottery/UserInfoRpc/GInstance/AreaConfig"
)

/**
oddschange 用户基本信息
*/
func (this *ApiController) OddsChange() {
	type TmpArgs struct {
		base.ArgsBox
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	server := GInstance.GetLotteryServer(aTmpArgs.GameType, aTmpArgs.Area)
	if server == nil {
		this.ErrJsonResultEx(errors.New("服务未启动, GetRoomInfo"))
	}
	server.ReLoadOddsInfo()
	this.JsonResult(enums.JRCodeSucc, "", "", nil)
}

/**
getgamenames 用户基本信息
*/
func (this *ApiController) GetGameNames() {
	type TmpArgs struct {
		base.ArgsBox
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	data := otherBll.GetGameNames()
	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}

/**
getarticlelist
*/
func (this *ApiController) GetArticleList() {
	type TmpArgs struct {
		base.ArgsBox
		ArticleType int
		PageIndex   int
		PageSize    int
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	PageTotal, arrData := otherBll.GetArticleList(aTmpArgs.ArticleType, aTmpArgs.PageIndex, aTmpArgs.PageSize)
	this.JsonResultCPage(enums.JRCodeSucc, "", PageTotal, arrData)
}

/**
getserviceinfo
*/
func (this *ApiController) GetServiceInfo() {
	type TmpArgs struct {
		base.ArgsBox
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	data := otherBll.GetServiceInfo()
	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}

func (this *ApiController) ModifyPwd() {
	type TmpArgs struct {
		base.ArgsBox
		OldPwd string
		NewPwd string
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	e = otherBll.ModifyPwd(aTmpArgs.Area, aTmpArgs.UserId, aTmpArgs.OldPwd, aTmpArgs.NewPwd)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", "", nil)
}

func (this *ApiController) ReLoadRebateSet() {
	type TmpArgs struct {
		base.ArgsBox
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	AreaConfig.ReLoadRebateSet(0)
	this.JsonResult(enums.JRCodeSucc, "", "", nil)
}

func (this *ApiController) SetStopBetHint() {
	type TmpArgs struct {
		base.ArgsBox
		StopBetHint string
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	server := GInstance.GetLotteryServer(aTmpArgs.GameType, aTmpArgs.Area)
	if server == nil {
		this.ErrJsonResultEx(errors.New("服务未启动, PreBet"))
	}
	e = server.SetStopBetHint(aTmpArgs.StopBetHint)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", "", nil)
}
