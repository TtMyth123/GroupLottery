package controllers

import (
	"fmt"
	"ttmyth123/GroupLottery/Admin/OtherServer/GameClientHttp"
	"ttmyth123/GroupLottery/Admin/controllers/AwardBll"
	"ttmyth123/GroupLottery/Admin/controllers/base/enums"
)

func (c *ApiController) GetAwardList() {
	GameType, _ := c.GetInt("GameType", 0)
	pageIndex, _ := c.GetInt("pageIndex", 0)
	pageSize, _ := c.GetInt("pageSize", 0)
	FirstId, _ := c.GetInt("FirstId", 0)
	beginDay := c.GetString("beginDay")
	endDay := c.GetString("endDay")
	LotteryStr := c.GetString("LotteryStr")
	PageTotal, data, groupData := AwardBll.GetAwardList(LotteryStr, beginDay, endDay, GameType, pageIndex, pageSize, FirstId)
	c.JsonPageResultGroup(enums.JRCodeSucc, "", PageTotal, data, groupData)
}

//func (c *ApiController) GetAwardListLastId()  {
//	GameType,_:=c.GetInt("GameType",0)
//	pageIndex,_:=c.GetInt("pageIndex",0)
//	pageSize,_:=c.GetInt("pageSize",0)
//	Period:=c.GetString("Period")
//	FirstId,_:=c.GetInt("FirstId",0)
//	PageTotal, data, groupData := AwardBll.GetAwardListFirstId(GameType,Period,pageIndex, pageSize,FirstId)
//	c.JsonPageResultGroup(enums.JRCodeSucc, "", PageTotal, data, groupData)
//}

func (c *ApiController) GetAwardDetail() {
	id, _ := c.GetInt("id", 0)
	data, e := AwardBll.GetAwardDetail(id)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", data)
}

func (c *ApiController) AddTestAward() {
	GameType, _ := c.GetInt("GameType", 0)
	fmt.Println(GameType)
	e := GameClientHttp.RandNewAwardInfo(GameType)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *ApiController) SetAwardInfo() {
	GameType, _ := c.GetInt("GameType", 0)
	LotteryAward := c.GetString("LotteryAward")
	LotteryStr := c.GetString("LotteryStr")

	e := GameClientHttp.SetAwardInfo(GameType, LotteryAward, LotteryStr)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *ApiController) GetSetAwardInfoList() {
	GameType, _ := c.GetInt("GameType", 0)
	pageIndex, _ := c.GetInt("pageIndex", 1)
	pageSize, _ := c.GetInt("pageSize", 10)
	FirstId, _ := c.GetInt("FirstId", 0)
	LotteryStr := c.GetString("LotteryStr")

	PageTotal, data, groupData := AwardBll.GetSetAwardInfo(GameType, pageIndex, pageSize, FirstId, LotteryStr)
	c.JsonPageResultGroup(enums.JRCodeSucc, "", PageTotal, data, groupData)
}
