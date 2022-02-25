package lottery

import "time"

type NewAwardInfo struct {
	LotteryNum      int64     //期号 如：2736604
	LotteryStr      string    //期号 如：2736604
	ResultNums      string    // 如：0,4,9
	NextLotteryStr  string    //  如：2736604
	NextLotteryTime time.Time //下一期开奖时间 如：2021-07-17 00:05:30
	CurLotteryTime  time.Time //当前开期时间 如：2021-07-17 00:05:30
}

func aa() {
	//a := time.Now()
	//a.Format()
}
