package chat

import (
	"encoding/json"
	"github.com/TtMyth123/Staff/GConfig"
	"github.com/TtMyth123/UserInfoRpc/models"
	"github.com/TtMyth123/UserInfoRpc/models/mconst"
	"github.com/TtMyth123/kit/httpKit"
	"github.com/TtMyth123/kit/timeKit"
	"time"
)

type GroupChatMsg struct {
	MinChatId int64
	MaxChatId int64

	CurMinChatId int64
	CurMaxChatId int64

	MapMsg map[int]*RoomMsg
}
type TtChatInfoEx struct {
	models.TtChatInfo
	ReplyC []string
}

type RoomMsg struct {
	MinChatId int64
	MaxChatId int64
	RoomName  string
	ArrMsg    []TtChatInfoEx
}

func (d *TtChatInfoEx) MarshalJSON() ([]byte, error) {
	type Alias TtChatInfoEx

	aContent := d.Content
	if d.ChatType == mconst.ChatType_Pic_1 {
		aContent = httpKit.GetImgUrl(GConfig.GetGConfig().BasePicRootPath, d.Content)
	}

	StrTime1 := ""
	StrTime2 := ""
	TimeLag := GConfig.GetGConfig().TimeLag
	if TimeLag == 0 {
		StrTime1 = d.CreatedAt.Format(timeKit.DateTimeLayout)
		StrTime2 = d.UpdateAt.Format(timeKit.DateTimeLayout)
	} else {
		StrTime1 = d.CreatedAt.Add(TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
		StrTime2 = d.UpdateAt.Add(TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
	}

	return json.Marshal(&struct {
		Id        int64
		Id2       int64
		RoomId    int
		Letter    string
		RoomName  string
		GameId    int64
		UserName  string
		GameId2   int64
		UserName2 string
		Way       int
		ChatType  int
		Content   string
		CreatedAt string
		UpdateAt  string
		ReplyC    []string
	}{
		Id:        d.Id,
		Id2:       d.Id2,
		RoomId:    d.RoomId,
		Letter:    d.Letter,
		RoomName:  d.RoomName,
		GameId:    d.GameId,
		UserName:  d.UserName,
		GameId2:   d.GameId2,
		UserName2: d.UserName2,
		Way:       d.Way,
		ChatType:  d.ChatType,
		ReplyC:    d.ReplyC,
		CreatedAt: StrTime1,
		UpdateAt:  StrTime2,
		Content:   aContent,
	})
}

func NewRoomMsg() *RoomMsg {
	aRoomMsg := new(RoomMsg)
	aRoomMsg.ArrMsg = make([]TtChatInfoEx, 0)
	return aRoomMsg
}
