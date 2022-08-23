package GData

import (
	"github.com/TtMyth123/Staff/GData/chat"
	"github.com/TtMyth123/UserInfoRpc/models"
	"github.com/TtMyth123/kit/stringKit"
	"github.com/TtMyth123/kit/ttLog"
	"sync"
)

var (
	ChatId      int64
	mChatManage *ChatManage
)

func init() {
	ChatId = 100
	mChatManage = newChatManage()
}

type ChatInfo struct {
	MinChatId int64
	MaxChatId int64

	userMsg sync.Map
}

type ChatManage struct {
	mapUserRoom sync.Map //user to
}

func newChatManage() *ChatManage {
	mChatManage = new(ChatManage)
	return mChatManage
}

func (c *ChatManage) AddGameChatMsg(aChatInfo models.TtChatInfo) {
	var aUserMsg *chat.StaffMsg
	if msgs, ok := c.mapUserRoom.Load(aChatInfo.GameId2); ok {
		aUserMsg = msgs.(*chat.StaffMsg)
	} else {
		aUserMsg = chat.NewStaffMsg(aChatInfo.GameId2)
		c.mapUserRoom.Store(aChatInfo.GameId2, aUserMsg)
	}

	aUserMsg.AddMsg(aChatInfo)
}
func (c *ChatManage) AddAdminChatMsg(aChatInfo models.TtChatInfo) {
	var aUserMsg *chat.StaffMsg
	if msgs, ok := c.mapUserRoom.Load(aChatInfo.GameId); ok {
		aUserMsg = msgs.(*chat.StaffMsg)
	} else {
		aUserMsg = chat.NewStaffMsg(aChatInfo.GameId)
		c.mapUserRoom.Store(aChatInfo.GameId, aUserMsg)
	}

	aUserMsg.AddMsg(aChatInfo)
}

func (c *ChatManage) GetAllChatMsg(staffId int64, Way, Min, Max int64) chat.GroupChatMsg {

	aGroupChatMsg := chat.GroupChatMsg{}

	var aUserMsg *chat.StaffMsg
	if msgs, ok := c.mapUserRoom.Load(staffId); ok {
		aUserMsg = msgs.(*chat.StaffMsg)
	} else {
		aUserMsg = chat.NewStaffMsg(staffId)
		c.mapUserRoom.Store(staffId, aUserMsg)
	}
	aGroupChatMsg = aUserMsg.GetAllMsg(Way, Min, Max)

	ttLog.LogDebug("aa", stringKit.GetJsonStr(aGroupChatMsg))

	return aGroupChatMsg
}

func (c *ChatManage) GetRoomChatMsg(staffId int64, roomId int, Way, Min, Max int64) chat.GroupChatMsg {

	aGroupChatMsg := chat.GroupChatMsg{}

	var aUserMsg *chat.StaffMsg
	if msgs, ok := c.mapUserRoom.Load(staffId); ok {
		aUserMsg = msgs.(*chat.StaffMsg)
	} else {
		aUserMsg = chat.NewStaffMsg(staffId)
		c.mapUserRoom.Store(staffId, aUserMsg)
	}
	aGroupChatMsg = aUserMsg.GetRoomChatMsg(roomId, Way, Min, Max)

	return aGroupChatMsg
}

func GetChatManage() *ChatManage {
	return mChatManage
}
