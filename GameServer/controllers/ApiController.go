package controllers

import (
	"errors"
	"fmt"
	"ttmyth123/GroupLottery/GameServer/GInstance"
	"ttmyth123/GroupLottery/GameServer/LotteryServer"
	"ttmyth123/GroupLottery/GameServer/LotteryServer/LotteryBox"
	"ttmyth123/GroupLottery/GameServer/controllers/base"
	"ttmyth123/GroupLottery/GameServer/controllers/base/enums"
	"ttmyth123/kit/strconvEx"
)

type ApiController struct {
	base.AuthorBaseController
}

func (this *ApiController) GetRoomInfo() {
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

	data := server.GetRoomInfo()

	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}

func (this *ApiController) Bet() {
	type TmpArgs struct {
		base.ArgsBox
		LotteryNum    int64
		StrLotteryNum string
		BetData       []LotteryBox.BetDataInfo
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	if aTmpArgs.StrLotteryNum == "" {
		aTmpArgs.StrLotteryNum = fmt.Sprintf("%d", aTmpArgs.LotteryNum)
	}
	if aTmpArgs.LotteryNum == 0 {
		aTmpArgs.LotteryNum = strconvEx.StrTry2Int64(aTmpArgs.StrLotteryNum, 0)
	}

	server := GInstance.GetLotteryServer(aTmpArgs.GameType, aTmpArgs.Area)
	if server == nil {
		this.ErrJsonResultEx(errors.New("服务未启动, Bet"))
	}

	betInfo := LotteryBox.BetInfo{
		StrLotteryNum: aTmpArgs.StrLotteryNum,
		LotteryNum:    aTmpArgs.LotteryNum,
		BetData:       aTmpArgs.BetData,
		UserId:        aTmpArgs.UserId,
	}

	data, e := server.Bet(betInfo)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}

func (this *ApiController) GetCurResult() {
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
		this.ErrJsonResultEx(errors.New("服务未启动, GetCurResult"))
	}

	data, e := server.GetCurResult()
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}

func (this *ApiController) GetHistoryResultList() {
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

	server := GInstance.GetLotteryServer(aTmpArgs.GameType, aTmpArgs.Area)
	if server == nil {
		this.ErrJsonResultEx(errors.New("服务未启动, GetHistoryResultList"))
	}
	data, e := server.GetHistoryResultList(aTmpArgs.PageIndex, aTmpArgs.PageSize, aTmpArgs.LastId)

	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}
func (this *ApiController) GetHistoryResultByPeriod() {
	type TmpArgs struct {
		base.ArgsBox
		Period int
		Count  int
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	server := GInstance.GetLotteryServer(aTmpArgs.GameType, aTmpArgs.Area)
	if server == nil {
		this.ErrJsonResultEx(errors.New("服务未启动, GetHistoryResultByPeriod"))
	}
	data, e := server.GetHistoryResultByPeriod(aTmpArgs.Period, aTmpArgs.Count)

	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}

//gethistorylotterybyday
func (this *ApiController) GetHistoryLotteryByDay() {
	type TmpArgs struct {
		base.ArgsBox
		StrDay string
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	server := GInstance.GetLotteryServer(aTmpArgs.GameType, aTmpArgs.Area)
	if server == nil {
		this.ErrJsonResultEx(errors.New("服务未启动, GetHistoryLotteryByDay"))
	}
	data, e := server.GetHistoryLotteryByDay(aTmpArgs.UserId, aTmpArgs.StrDay)

	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}

func (this *ApiController) GetBetRecordList() {
	type TmpArgs struct {
		base.ArgsBox
		StrBeginDay string
		StrEndDay   string
		Status      int
		PageIndex   int
		PageSize    int
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	data, e := LotteryServer.GetBetRecordList(aTmpArgs.GameType, aTmpArgs.UserId, aTmpArgs.Status,
		aTmpArgs.PageIndex, aTmpArgs.PageSize, aTmpArgs.LastId, aTmpArgs.StrBeginDay, aTmpArgs.StrEndDay)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}

func (this *ApiController) GetBetOrder() {
	type TmpArgs struct {
		base.ArgsBox
		BetOrder string
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	data, e := LotteryServer.GetBetOrder(aTmpArgs.BetOrder)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}

func (this *ApiController) RandNewAwardInfo() {
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
		this.ErrJsonResultEx(errors.New("服务未启动, GetHistoryResultList"))
	}
	e = server.RandNewAwardInfo()

	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", "", nil)
}

func (this *ApiController) GetHistoryFTNum() {
	type TmpArgs struct {
		base.ArgsBox
		Count   int
		AwardId int
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	server := GInstance.GetLotteryServer(aTmpArgs.GameType, aTmpArgs.Area)
	if server == nil {
		this.ErrJsonResultEx(errors.New("服务未启动, GetHistoryFTNum"))
	}
	data, e := server.GetHistoryFTNum(aTmpArgs.LastId, aTmpArgs.Count)

	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}

func (this *ApiController) GetHistoryFTNumBy48() {
	type TmpArgs struct {
		base.ArgsBox
		AwardId int
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	server := GInstance.GetLotteryServer(aTmpArgs.GameType, aTmpArgs.Area)
	if server == nil {
		this.ErrJsonResultEx(errors.New("服务未启动, GetHistoryFTNumBy48"))
	}
	data, e := server.GetHistoryFTNumBy48(aTmpArgs.LastId)

	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}

func (this *ApiController) CurDayLoseWin() {
	type TmpArgs struct {
		base.ArgsBox
		AwardId int
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	data, e := LotteryServer.CurDayLoseWin(aTmpArgs.GameType, aTmpArgs.UserId)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}

func (this *ApiController) SetAwardInfo() {
	type TmpArgs struct {
		LotteryAward string
		LotteryNum   string
		base.ArgsBox
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	server := GInstance.GetLotteryServer(aTmpArgs.GameType, aTmpArgs.Area)
	if server == nil {
		this.ErrJsonResultEx(errors.New("服务未启动, SetAwardInfo"))
	}
	data, e := server.SetAwardInfo(aTmpArgs.LotteryAward, aTmpArgs.LotteryNum)

	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}
