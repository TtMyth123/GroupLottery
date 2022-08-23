package controllers

import (
	"github.com/TtMyth123/LotteryResultServer/controllers/base"
	"github.com/TtMyth123/LotteryResultServer/controllers/uscBll"
)

type ApiController struct {
	base.BaseController
}

func (c *ApiController) LotteryResult() {
	gameindex, _ := c.GetInt("gameindex")
	data := uscBll.SendGetResultByIndex(gameindex)
	c.Data["json"] = data
	c.ServeJSON()
	c.StopRun()
}

func (c *ApiController) LotteryResultByType() {
	gameType, _ := c.GetInt("gameType")
	data := uscBll.SendGetResultByType(gameType)
	c.Data["json"] = data
	c.ServeJSON()
	c.StopRun()
}

func (c *ApiController) UpdateResultUrl() {
	gameType, _ := c.GetInt("gameType")
	url := c.GetString("url")
	e := uscBll.UpdateResultUrl(gameType, url)
	c.JsonResultEx(e, "")
}
