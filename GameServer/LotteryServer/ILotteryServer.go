package LotteryServer

import "github.com/TtMyth123/GameServer/LotteryServer/LotteryBox"

type ILotteryServer interface {
	ReLoadOddsInfo()
	GetRoomInfo() map[string]interface{}
	Bet(betInfo LotteryBox.BetInfo) (map[string]interface{}, error)

	GetHistoryResultList(PageIndex, PageSize, LastId int) (interface{}, error)
	GetHistoryLotteryByDay(UserId int, StrDay string) (interface{}, error)
	GetCurResult() (interface{}, error)

	RandNewAwardInfo() error

	SetStopBetHint(StopBetHint string) error
	//GetBetRecordList(UserId, Status, PageIndex, PageSize int, StrBeginDay, StrEndDay string)(map[string]interface{}, error)
	GetHistoryFTNum(int, int) (interface{}, error)
	GetHistoryFTNumBy48(int) (interface{}, error)
	GetHistoryResultByPeriod(int, int) (interface{}, error)
	SetAwardInfo(LotteryAward, LotteryNum string) (interface{}, error)
}
