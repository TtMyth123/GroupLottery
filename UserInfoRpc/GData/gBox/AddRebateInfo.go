package gBox

import "github.com/TtMyth123/UserInfoRpc/models/mconst"

type AddRebateInfo struct {
	UserId      int
	UserSid     int
	BetUserName string

	GameType   int
	LotteryStr string
	Level      int
	OddsName   string
	BetM       int
	T          mconst.Rebate_Type

	//OldGold float64 //原金额
	//Ratio   float64//比例
	//RefId   int //原Id
	//Rebate  float64 //金额

	//Des string
	//Des2         string
	//DesMp    string
}
