package UscBox

import "time"

type AwardInfo struct {
	LotteryNum      int64  //期号
	LotteryStr      string //期号
	ResultNums      string
	NextLotteryStr  string
	NextLotteryTime time.Time //下一期开奖时间
	CurLotteryTime  time.Time //当前开期时间
	ServerTime      time.Time //服务的当前时间
}

type UscAwardInfo struct {
	Id         int
	NewNum     int64     //当前期号
	NewNumTime time.Time //当前期开奖时间
	ResultNums string    //当前期开奖结果 0,3,2,3,5,5,4,3
	ResultDX   string    //大小： 小,小,小,小,大
	ResultDS   string    //单双：单,双,单,单,双
	ResultLH   string    //龙虎：龙,龙,龙,龙虎
	ResultGZH  string    //冠亚和：10，大,小
	ResultFtFS int       //番数
	ResultFtH  int       //番数和
	ResultFtDS string    //番数 单双

	LotteryNum    int64     //下一期期号
	CloseTime     time.Time //封盘时间
	LotteryTime   time.Time //下一期开奖时间
	ServerTime    time.Time `json:"-"` //服务器时间
	TotalNum      int       `json:"-"`
	RestNum       int       `json:"-"`
	GameIndex     int
	Countdown     int //倒计时（秒）
	StopCountdown int //封 倒计时（秒）
}

type OddsInfo struct {
	Id   int
	Odds float64 `json:"1"`
}

type UscAwardFTInfo struct {
	Id         int
	ResultFtFS int
	ResultFtH  int
	ResultFtDS string
}
