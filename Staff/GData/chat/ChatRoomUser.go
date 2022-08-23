package chat

import (
	"github.com/TtMyth123/Staff/GInstance"
	"github.com/TtMyth123/UserInfoRpc/RpcServer/UserServer/RpcBox"
)

const defLen = 100

type RoomUserMsg struct {
	MinChatId int64
	MaxChatId int64
	ArrMsg    []TtChatInfoEx
	MaxLen    int
	//mapUser sync.Map
	RoomId int
}

func NewRoomUserMsg(RoomId int) *RoomUserMsg {
	aRoomUserMsg := new(RoomUserMsg)
	aRoomUserMsg.ArrMsg = make([]TtChatInfoEx, 0)
	aRoomUserMsg.MaxLen = defLen
	return aRoomUserMsg
}

func (this *RoomUserMsg) AddChatMsg(aChatInfo TtChatInfoEx) {
	this.ArrMsg = append(this.ArrMsg, aChatInfo)
	if this.MaxChatId < aChatInfo.Id {
		this.MaxChatId = aChatInfo.Id
	}
	len := len(this.ArrMsg)

	if this.MaxLen < len {
		this.ArrMsg = this.ArrMsg[len-this.MaxLen : len]
		this.MinChatId = this.ArrMsg[0].Id
	}
}
func (c *RoomUserMsg) GetUnreadMsg(way, MinId, MaxId int64) RoomMsg {
	if true {
		aArgsGetChatRoomMsgList := new(RpcBox.ArgsGetChatRoomMsgList)
		aArgsGetChatRoomMsgList.MaxId = 0
		aArgsGetChatRoomMsgList.Way = int(way)
		aArgsGetChatRoomMsgList.MinId = MinId
		aArgsGetChatRoomMsgList.MaxId = MaxId
		aArgsGetChatRoomMsgList.Count = 20
		r := GInstance.GetUserRpcClient().GetChatRoomMsgList(aArgsGetChatRoomMsgList)

		if r.Result {
			newDataList := make([]TtChatInfoEx, len(r.DataList))
			for i, v := range r.DataList {
				newDataList[i] = TtChatInfoEx{
					TtChatInfo: v,
					ReplyC:     GetReplyManage().GetQuickReply(0, v.Content),
				}
			}

			aRoomMsg := RoomMsg{
				MinChatId: r.GroupObj.MinId,
				MaxChatId: r.GroupObj.MaxId,
				ArrMsg:    newDataList,
			}
			return aRoomMsg
		}
	}
	aRoomMsg := RoomMsg{MinChatId: c.MinChatId,
		MaxChatId: c.MaxChatId,
		ArrMsg:    c.ArrMsg,
	}

	return aRoomMsg
}
