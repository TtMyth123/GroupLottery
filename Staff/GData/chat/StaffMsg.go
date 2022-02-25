package chat

import (
	"sync"
	"ttmyth123/GroupLottery/Staff/GInstance"
	"ttmyth123/GroupLottery/UserInfoRpc/RpcServer/UserServer/RpcBox"
	"ttmyth123/GroupLottery/UserInfoRpc/models"
)

type StaffMsg struct {
	MinChatId int64
	MaxChatId int64
	StaffId   int64

	mapRoomChat sync.Map
}

func NewStaffMsg(staffId int64) *StaffMsg {
	aStaffMsg := new(StaffMsg)
	aStaffMsg.StaffId = staffId
	aStaffMsg.MinChatId = 9999999

	return aStaffMsg
}
func (c *StaffMsg) AddMsg(aChatInfo models.TtChatInfo) {
	var aUserMsg *RoomUserMsg
	if userMsg, ok := c.mapRoomChat.Load(aChatInfo.RoomId); ok {
		aUserMsg = userMsg.(*RoomUserMsg)
	} else {
		aUserMsg = NewRoomUserMsg(aChatInfo.RoomId)
		c.mapRoomChat.Store(aChatInfo.RoomId, aUserMsg)
	}

	aNewChatInfo := TtChatInfoEx{
		TtChatInfo: aChatInfo,
	}
	aUserMsg.AddChatMsg(aNewChatInfo)
}

func (this *StaffMsg) GetAllMsg(way, MinId, MaxId int64) GroupChatMsg {
	aGroupChatMsg := GroupChatMsg{}
	aGroupChatMsg.MapMsg = make(map[int]*RoomMsg)

	aArgsGetChatRoomMsgList := new(RpcBox.ArgsGetChatStaffMsgList)
	aArgsGetChatRoomMsgList.Way = int(way)
	aArgsGetChatRoomMsgList.MinId = MinId
	aArgsGetChatRoomMsgList.MaxId = MaxId

	aArgsGetChatRoomMsgList.Count = 100
	r := GInstance.GetUserRpcClient().GetChatStaffMsgList(aArgsGetChatRoomMsgList)
	if r.Result {
		for _, aChatInfo := range r.DataList {
			aRoomMsg := aGroupChatMsg.MapMsg[aChatInfo.RoomId]
			if aRoomMsg == nil {
				aRoomMsg = NewRoomMsg()
				aGroupChatMsg.MapMsg[aChatInfo.RoomId] = aRoomMsg
			}

			aNewChatInfo := TtChatInfoEx{
				TtChatInfo: aChatInfo,
				ReplyC:     GetReplyManage().GetQuickReply(this.StaffId, aChatInfo.Content),
			}
			aRoomMsg.ArrMsg = append(aRoomMsg.ArrMsg, aNewChatInfo)
		}

		aGroupChatMsg.MinChatId = r.GroupObj.MinId
		aGroupChatMsg.MaxChatId = r.GroupObj.MaxId
	}

	return aGroupChatMsg
}

func (this *StaffMsg) GetRoomChatMsg(roomId int, way, MinId, MaxId int64) GroupChatMsg {
	aGroupChatMsg := GroupChatMsg{}
	aGroupChatMsg.MapMsg = make(map[int]*RoomMsg)

	aArgsGetChatRoomMsgList := new(RpcBox.ArgsGetChatRoomMsgList)
	aArgsGetChatRoomMsgList.Way = int(way)
	aArgsGetChatRoomMsgList.MinId = MinId
	aArgsGetChatRoomMsgList.MaxId = MaxId
	aArgsGetChatRoomMsgList.RoomId = roomId

	aArgsGetChatRoomMsgList.Count = 100
	r := GInstance.GetUserRpcClient().GetChatRoomMsgList(aArgsGetChatRoomMsgList)
	if r.Result {
		for _, aChatInfo := range r.DataList {
			aRoomMsg := aGroupChatMsg.MapMsg[aChatInfo.RoomId]
			if aRoomMsg == nil {
				aRoomMsg = NewRoomMsg()
				aGroupChatMsg.MapMsg[aChatInfo.RoomId] = aRoomMsg
			}

			aNewChatInfo := TtChatInfoEx{
				TtChatInfo: aChatInfo,
				ReplyC:     GetReplyManage().GetQuickReply(0, aChatInfo.Content),
			}
			aRoomMsg.ArrMsg = append(aRoomMsg.ArrMsg, aNewChatInfo)
		}

		aGroupChatMsg.MinChatId = r.GroupObj.MinId
		aGroupChatMsg.MaxChatId = r.GroupObj.MaxId
	}

	return aGroupChatMsg
}
