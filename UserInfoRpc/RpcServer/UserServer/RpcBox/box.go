package RpcBox

import (
	"encoding/json"
	"github.com/TtMyth123/UserInfoRpc/GConfig"
	"github.com/TtMyth123/UserInfoRpc/GData/gBox"
	"github.com/TtMyth123/UserInfoRpc/GInstance/AreaConfig"
	"github.com/TtMyth123/UserInfoRpc/models"
	"github.com/TtMyth123/UserInfoRpc/models/mconst"
	"github.com/TtMyth123/kit/timeKit"
	"time"
)

type ArgsGetUser struct {
	ArgsBox
}

type ReplyGetUser struct {
	ReplyBox
	Data models.TtGameUser
}

type ArgsAddUser struct {
	ArgsBox
	Area     string
	UserName string
	Pwd      string
	MoneyPwd string
	Code     int
	UserType int

	InfoEx map[string]string
}

type ReplyAddUser struct {
	ReplyBox
	Data models.TtGameUser
}

type ArgsAddGold struct {
	ArgsBox
	Gold float64
	T    mconst.Account_Type
	Des  string
}
type ReplyAddGold struct {
	ReplyBox
	Data models.TtGameUser
}

type ArgsGetUserByNamePwd struct {
	ArgsBox
	UserName string
	Pwd      string
}
type ReplyGetUserByNamePwd struct {
	ReplyBox
	Data models.TtGameUser
}

type ArgsUpdateUserInfo struct {
	ArgsBox
	Datas []gBox.UpdateDataInfo
}
type ReplyUpdateUserInfo struct {
	ReplyBox
	Data models.TtGameUser
}

type ArgsRebate2Gold struct {
	ArgsBox
	Rebate float64
	Des    string
	Des2   string
	DesMp  string
}
type ReplyRebate2Gold struct {
	ReplyBox
	Data models.TtGameUser
}

type Rebate2GoldRecord struct {
	Id        int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	Gold      float64
}

func (d Rebate2GoldRecord) MarshalJSON() ([]byte, error) {
	type Alias Rebate2GoldRecord

	StrTime := ""
	TimeLag := GConfig.GetGConfig().TimeLag
	if TimeLag == 0 {
		StrTime = d.CreatedAt.Format(timeKit.DateTimeLayout)
	} else {
		StrTime = d.CreatedAt.Add(TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
	}

	return json.Marshal(&struct {
		Alias
		CreatedAt string
	}{
		Alias:     (Alias)(d),
		CreatedAt: StrTime,
	})
}

type ArgsRebate2GoldRecord struct {
	ArgsBox
	PageIndex int
	PageSize  int
	MaxId     int
}
type ReplyRebate2GoldRecord struct {
	ReplyBox
	PageTotal int
	DataList  []Rebate2GoldRecord
	GroupObj  interface{}
}

type ArgsAddRebate struct {
	ArgsBox
	UserSid     int
	BetUserName string

	GameType   int
	LotteryStr string
	Level      int
	OddsName   string
	BetM       int
	T          mconst.Rebate_Type
}
type ReplyAddRebate struct {
	ReplyBox
	Data models.TtGameUser
}

type ArgsGetRebateSetConfig struct {
	ArgsBox
	Area int
}
type ReplyGetRebateSetConfig struct {
	ReplyBox
	Data AreaConfig.RebateSetConfig
}

type ArgsReRebateSetConfig struct {
	ArgsBox
}
type ReplyReRebateSetConfig struct {
	ReplyBox
}

type ArgsGetGameUserList struct {
	ArgsBox
	UserName  string
	MaxId     int
	PageIndex int
	PageSize  int
}
type ReplyGetGameUserList struct {
	ReplyBox
	DataList interface{}
	GroupObj interface{}
}

type ArgsAddTtChatInfo struct {
	ChatInfo models.TtChatInfo
	ArgsBox
}
type ReplyAddTtChatInfo struct {
	ReplyBox
	DataInfo models.TtChatInfo
}

type ArgsGetLatelyChatUser struct {
	ArgsBox
	UserName  string
	MaxId     int
	PageIndex int
	PageSize  int
}
type ReplyGetLatelyChatUser struct {
	ReplyBox

	DataList interface{}
	GroupObj interface{}
}

type ArgsGetChatRoomMsgList struct {
	ArgsBox
	RoomId int
	MaxId  int64
	MinId  int64
	Count  int
	Way    int
}

type GetChatRoomMsgListGroup struct {
	C     int
	MaxId int64
	MinId int64
}
type ReplyGetChatRoomMsgList struct {
	ReplyBox

	//DataList interface{}
	//GroupObj interface{}
	DataList []models.TtChatInfo
	GroupObj GetChatRoomMsgListGroup
}

type ArgsGetChatStaffMsgList struct {
	ArgsBox
	RoomId int
	MaxId  int64
	MinId  int64
	Count  int
	Way    int
}

type GetChatStaffMsgListGroup struct {
	C     int
	MaxId int64
	MinId int64
}
type ReplyGetChatStaffMsgList struct {
	ReplyBox

	//DataList interface{}
	//GroupObj interface{}
	DataList []models.TtChatInfo
	GroupObj GetChatStaffMsgListGroup
}
