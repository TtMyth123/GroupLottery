package controllers

import (
	"github.com/TtMyth123/LotteryResultSite/controllers/base"
	"github.com/TtMyth123/LotteryResultSite/controllers/base/enums"
	"github.com/TtMyth123/LotteryResultSite/controllers/bll"
)

type ApiController struct {
	base.ABaseController
}

/**
lotteryresult
*/
func (c *ApiController) LotteryResult() {
	LotteryStr := c.GetString("LotteryStr")
	strDay := c.GetString("strDay")

	gameType, _ := c.GetInt("gameType")
	pageIndex, _ := c.GetInt("pageIndex")
	pageSize, _ := c.GetInt("pageSize")
	MaxId, _ := c.GetInt("MaxId")

	arr, groupData := bll.GetAwardList(LotteryStr, strDay, gameType, pageIndex, pageSize, MaxId)
	c.JsonPageResultGroup(enums.JRCodeSucc, "", groupData.C, arr, groupData)
}
