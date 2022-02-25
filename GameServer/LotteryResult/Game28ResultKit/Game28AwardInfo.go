package Game28ResultKit

type Game28AwardInfo struct {
	LotteryStr     string //期号
	StrLotteryTime string //当前开期时间

	NextLotteryStr     string //下一期 期号
	StrNextLotteryTime string //下一期开期时间

	ResultNums string //当前期开奖结果 0,3,2,3,5,5,4,3
	ResultDX   string //大小： 小,小,小,小,大
	ResultDS   string //单双：单,双,单,单,双
	ResultLH   string //龙虎：龙,龙,龙,龙虎
	ResultGZH  string //冠亚和：10，大,小

	Countdown     int //倒计时（秒）
	StopCountdown int //封 倒计时（秒）
	GameType      int //游戏类型
}

type Game28HistoryResult struct {
	LotteryStr     string //期号
	StrLotteryTime string //当前开期时间

	NextLotteryStr     string
	StrNextLotteryTime string //下一期开期时间
	ResultNums         string
	HS                 int    //和数
	BS                 string //波色
	LH                 string //龙虎和
	JS                 string //极数
	BDS                string //豹对顺
	DS                 string //单双
	DX                 string //大小
}

type Game28HistoryResult2 struct {
	Period    int64  //期号
	StrPeriod string //期号
	Nums      []int  //号码
	SumNum    int    //和值
	DX        string //大小
	DXDS      string //大双，小双，小单，大单
	BDS       string //豹对顺
	JZ        string //极大，极小
	BS        string //红，绿，蓝
	LH        string //龙虎
}
