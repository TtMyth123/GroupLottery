package controllers

import (
	"ttmyth123/GroupLottery/Admin/controllers/ReportBll"
	"ttmyth123/GroupLottery/Admin/controllers/base/enums"
)

func (c *ApiController) GetGroupBetList() {
	state, _ := c.GetInt("state", 0)
	userName := c.GetString("userName")
	pageIndex, _ := c.GetInt("pageIndex", 0)
	pageSize, _ := c.GetInt("pageSize", 0)

	userId, _ := c.GetInt("userId", 0)
	userType, _ := c.GetInt("userType", 0)
	gameType, _ := c.GetInt("gameType", 0)
	FirstId, _ := c.GetInt("FirstId", 0)
	Period := c.GetString("Period")
	GroupBetSn := c.GetString("GroupBetSn")

	beginDay := c.GetString("beginDay")
	endDay := c.GetString("endDay")

	curAgentId := c.CurSysUserEx().GameId
	if c.CurSysUserEx().IsSuper {
		curAgentId = 0
	}

	PageTotal, data, groupData := ReportBll.GetGroupBetList(curAgentId, userId, gameType, userType, state, beginDay, endDay, userName, Period, GroupBetSn, pageIndex, pageSize, FirstId)
	c.JsonPageResultGroup(enums.JRCodeSucc, "", PageTotal, data, groupData)
}

//func (c *ApiController) GetGroupBetListFirstId() {
//	state, _ := c.GetInt("state", 0)
//	userName := c.GetString("userName")
//	pageIndex, _ := c.GetInt("pageIndex", 0)
//	pageSize, _ := c.GetInt("pageSize", 0)
//
//	userType, _ := c.GetInt("userType", 0)
//	gameType, _ := c.GetInt("gameType", 0)
//	FirstId, _ := c.GetInt("FirstId", 0)
//	Period := c.GetString("Period")
//	GroupBetSn := c.GetString("GroupBetSn")
//
//	beginDay := c.GetString("beginDay")
//	endDay := c.GetString("endDay")
//
//	PageTotal, data, groupData := ReportBll.GetGroupBetListFirstId(gameType, userType, state, beginDay, endDay, userName,Period,GroupBetSn, pageIndex, pageSize,FirstId)
//	c.JsonPageResultGroup(enums.JRCodeSucc, "", PageTotal, data, groupData)
//}

func (c *ApiController) GetGroupDetailBetList() {
	GroupBetId, _ := c.GetInt("GroupBetId", 0)

	data := ReportBll.GetGroupDetailBetList(GroupBetId)
	c.JsonResult(enums.JRCodeSucc, "", data)
}

func (c *ApiController) Get7DayBetYieldChart() {
	data := ReportBll.Get7DayBetYieldChart()
	c.JsonResult(enums.JRCodeSucc, "", data)
}

func (c *ApiController) GetMainData() {
	data := ReportBll.GetMainData()
	c.JsonResult(enums.JRCodeSucc, "", data)
}
