package controllers

import (
	"github.com/TtMyth123/Admin/controllers/base"
	"github.com/TtMyth123/Admin/controllers/base/enums"
	"github.com/TtMyth123/kit/ttLog"
)

type ChatController struct {
	base.ABaseController
}

/**
sendchat
*/
func (c *ChatController) SendChat() {
	//Id int64
	//UserName  string    `orm:"size(256);description"用户名"` //用户名
	//GameId    int64     `orm:"description(游戏ID)"`           //游戏ID
	//Way       int       `orm:"description(方向,0:客户端消息)"`
	//ChatType  int       `orm:"description(聊天类型.0:一般聊天内容)"`   //聊天类型.0:一般聊天内容
	//Content   string    `orm:"size(4000);description(聊天内容)"` //聊天内容
	//CreatedAt time.Time `orm:"auto_now_add;type(datetime);description(时间)"`
	Id, _ := c.GetInt("Id")
	UserName := c.GetString("UserName")
	Way, _ := c.GetInt("Way")
	GameId, _ := c.GetInt("GameId")
	ChatType, _ := c.GetInt("ChatType")
	Content := c.GetString("Content")
	CreatedAt := c.GetString("CreatedAt")

	ttLog.LogDebug("SendChat:", Id, UserName, Way, GameId, ChatType, Content, CreatedAt)
	c.JsonResult(enums.JRCodeSucc, "", nil)
}
