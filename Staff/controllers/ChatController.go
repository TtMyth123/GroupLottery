package controllers

import (
	"fmt"
	"github.com/TtMyth123/Staff/controllers/base"
	"github.com/TtMyth123/Staff/controllers/base/enums"
	"github.com/TtMyth123/Staff/controllers/bll"
	"github.com/TtMyth123/kit/stringKit"
	"io/ioutil"
)

type ChatController struct {
	base.ABaseController
}

/**

SendChat 客户端主动调用 发信息 方法
*/
func (c *ChatController) SendChat() {
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

/**

SendChat 客户端主动调用 发信息 方法
*/
func (c *ChatController) SendChat1() {

	if v, err := ioutil.ReadAll(c.Ctx.Request.Body); err == nil {
		bbb := make(map[string]interface{})
		e := stringKit.GetJsonObj(string(v), &bbb)
		fmt.Println(e, "s:", string(v))
	} else {
		fmt.Println(err)
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}
