package controllers

import (
	"github.com/astaxie/beego"
	"ttmyth123/GroupLottery/UserInfoRpc/OtherServer/httpGameServer"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//getaddmoney
func (c *MainController) GetAddMoney() {
	RoodId, _ := c.GetInt("RoodId", 0)
	GameId, _ := c.GetInt("GameId", 0)
	Money, _ := c.GetFloat("Money", 0)
	f, e := httpGameServer.AddMoney(RoodId, GameId, Money)
	mp := make(map[string]interface{})
	mp["Data"] = f
	mp["Err"] = e
	c.Data["json"] = mp
	c.ServeJSON()
	c.StopRun()
}
