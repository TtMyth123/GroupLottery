package controllers

import (
	"ttmyth123/GroupLottery/Staff/OtherServer/httpGameServer"
	"ttmyth123/GroupLottery/Staff/controllers/base"
	"ttmyth123/GroupLottery/Staff/controllers/base/box"
	"ttmyth123/GroupLottery/Staff/controllers/base/enums"
	"ttmyth123/GroupLottery/Staff/controllers/bll"
)

type TestController struct {
	base.ABaseController
}

func (c *TestController) AdminMsg() {
	RoomId, _ := c.GetInt("RoomId")

	RoomName := c.GetString("RoomName")
	ChatType, _ := c.GetInt("ChatType")
	Content := c.GetString("Content")

	aStaff := box.BaseSysUser{Id: 10018, Name: "admin1"}

	r := bll.AddAdminChatMsg(aStaff, RoomId, RoomName, ChatType, Content)

	//r := httpGameServer.AdminMsg(aTtChatInfo)
	c.JsonResult(enums.JRCodeSucc, "", r)
}

func (c *TestController) GetAdminList() {
	r := httpGameServer.GetAdminList()
	c.JsonResult(enums.JRCodeSucc, "", r)
}

func (c *TestController) CreatePrivateTo() {
	GameId, _ := c.GetInt64("GameId")
	ToGameId, _ := c.GetInt64("ToGameId")

	r := httpGameServer.CreatePrivateTo(GameId, ToGameId)
	c.JsonResult(enums.JRCodeSucc, "", r)
}

func (c *TestController) CreateAdminAccount() {
	UserName := c.GetString("UserName")
	Password := c.GetString("Password")
	GameName := c.GetString("GameName")

	r := httpGameServer.CreateAdminAccount(UserName, Password, GameName)
	c.JsonResult(enums.JRCodeSucc, "", r)
}

func (c *TestController) SendChat() {
	Id, _ := c.GetInt64("Id")
	UserName := c.GetString("UserName")
	RoomName := c.GetString("RoomName")

	GameId, _ := c.GetInt("GameId")
	RoomId, _ := c.GetInt("RoomId")
	ChatType, _ := c.GetInt("ChatType")
	Content := c.GetString("Content")

	GameId2, _ := c.GetInt("GameId2")
	UserName2 := c.GetString("UserName2")
	CreatedAt := c.GetString("CreatedAt")

	bll.AddChatMsg(Id, RoomId, RoomName, ChatType, Content, CreatedAt,
		GameId, UserName, GameId2, UserName2)
	c.JsonResult(enums.JRCodeSucc, "", nil)
}
