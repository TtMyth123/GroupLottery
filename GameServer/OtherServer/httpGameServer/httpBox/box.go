package httpBox

type BaseHttpBox struct {
	Game   string `json:"Game"` //游戏名称
	Region string
}
type AwardResultBox struct {
	BaseHttpBox
	AwardInfo interface{}
}

type SetGameTimerBox struct {
	BaseHttpBox
	Timer     int    //倒计时时间（秒）
	StopBet   int    //倒计时时间（秒）
	NextIssue string //下一期期号
	Issue     string //当期期号
}

type BetDataInfo struct {
	OddsType int    //赔率ID
	M        int    //投注金额
	OddsName string //赔率名称
}
type PlayerBetBox struct {
	BaseHttpBox
	UserId  int     //用户ID
	Money   float64 //用户当前金额
	BetData []BetDataInfo
}

type StopBetStateBox struct {
	BaseHttpBox
}
