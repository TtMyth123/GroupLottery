package LotteryBox

type BetDataInfo struct {
	OddsType int    //赔率ID
	M        int    //投注金额
	OddsName string //赔率名称
	Nums     string //号码
}

type BetInfo struct {
	UserId        int
	GroupId       int
	LotteryNum    int64
	StrLotteryNum string
	BetData       []BetDataInfo
}
