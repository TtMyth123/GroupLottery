package Lottery28Game

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"sync"
	"time"
	"ttmyth123/GroupLottery/LotteryResultServer/GInstance/lottery"
	"ttmyth123/kit/httpClientKit"
	"ttmyth123/kit/stringKit"
	"ttmyth123/kit/timeKit"
)

type GameKenoJnd28 struct {
	timeI time.Duration
	t     int

	CurNums       string
	mHttpClient   *httpClientKit.HttpClient
	aNewAwardInfo lottery.NewAwardInfo
	lottery.ILotteryServer
	resultUrl string
	dataLock  sync.Mutex
}

func NewGameKenoJnd28(timeI time.Duration, t int) *GameKenoJnd28 {
	aGame28 := new(GameKenoJnd28)
	aGame28.mHttpClient = httpClientKit.GetHttpClient("")
	aGame28.timeI = timeI
	aGame28.t = t
	aGame28.resultUrl = beego.AppConfig.String("Game28::ResultHttpUrlKenoJnd")

	aGame28.aNewAwardInfo, _ = aGame28.getAwardInfo()
	aGame28.run()
	return aGame28
}

func (this *GameKenoJnd28) GetNewAwardInfo() lottery.NewAwardInfo {
	return this.aNewAwardInfo
}

func (this *GameKenoJnd28) run() {
	ticker1 := time.NewTicker(this.timeI * time.Second)
	go func() {
		for {
			select {
			case <-ticker1.C:
				aNewAwardInfo, e := this.getAwardInfo()
				if e == nil {
					this.aNewAwardInfo = aNewAwardInfo
				}
			}
		}
	}()
}

/**
{
    "gameCode": "jndpc28",
    "preIssue": "2697269",
    "openNum": [
        7,
        3,
        9
    ],
    "dragonTigerArr": [],
    "sumArr": [
        19,
        1,
        1,
        0
    ],
    "issue": "2697270",
    "currentOpenDateTime": 1617814500000,
    "openDateTime": 1617814710000,
    "serverTime": 1617814521950,
    "openedCount": 287,
    "dailyTotal": 396,
    "formArr": [],
    "mimcryArr": [],
    "zodiacArr": [],
    "compareArr": [],
    "sumType": null,
    "wuxing": null
}
*/
func (this *GameKenoJnd28) getAwardInfo() (lottery.NewAwardInfo, error) {
	aNewAwardInfo := lottery.NewAwardInfo{}

	//aJnd28Result := ResultBox.Game28Result{}
	//https://www.kjw6589.com/data/Current/jndpc28/CurIssue.json?NL8RN4R1KWP3KPSZQXAJFMXXHW9X6BXNNOXXAYRYWDZNC7QLT4
	this.dataLock.Lock()
	iTime := timeKit.GetJavaTimeLong(time.Now())
	strU := fmt.Sprintf(`%s/?_%d`, this.resultUrl, iTime)
	this.dataLock.Unlock()
	//{"row":1,"code":"jnd28","data":[{"opentime":"2021-05-30 16:26:00","expect":"2717909","opencode":"0,4,9","Timer":0,"NextTime":"2021-05-30 16:29:30","NextIssue":"2717910","NextTime2":"","NextIssue2":""}],"SleepTime":0,"SleepMinute":0,"LastDataStr":"","LastIssue":"","SetTimerIng":false,"ServerTime":"","Point":null,"LastMsg":{},"Lock":{}}
	paramsHeader := make(map[string]string)
	paramsHeader["Accept"] = "application/json"
	resultHtml, e := this.mHttpClient.GetHeader(strU, paramsHeader)
	if e != nil {
		return aNewAwardInfo, e
	}

	aNewAwardInfo, e = this.kenoResult2NewAwardInfo(string(resultHtml))
	return aNewAwardInfo, e
}
func (this *GameKenoJnd28) UpdateResultUrl(url string) error {
	this.dataLock.Lock()
	defer this.dataLock.Unlock()
	this.resultUrl = url
	return nil
}

/**
resultData={"drawDate":"2021-07-17T00:05:30.000","timeSinceDraw":20428,"draw":2736604,"bonus":7,"num":[27,43,16,73,36,41,15,72,45,25,5,38,50,6,44,66,23,35,11,77],"nextKenoDraw":"00:03:10","nextKenoDrawTime":190}



加拿大 2 8 玩法说明：
例如：加拿大BCLC第"1749110"期数据从小到大排序 7,8,14,16,17,22,26,34,39,41,42,48,54,58,63,64,69,72,73,79

第一区[第2/5/8/11/14/17位数字] 8,17,34,42,58,69

计算：8+17+34+42+58+69= 228

结果为：8

第二区[第3/6/9/12/15/18位数字] 14,22,39,48,63,72

计算：14+22+39+48+63+72= 258

结果为：8

第三区[第4/7/10/13/16/19位数字] 16,26,41,54,64,73

计算：16+26+41+54+64+73= 274

结果为：4

最终游戏结果为：8+8+4=20
*/
func (this *GameKenoJnd28) kenoResult2NewAwardInfo(resultData string) (lottery.NewAwardInfo, error) {
	aNewAwardInfo := lottery.NewAwardInfo{}

	type TmpR struct {
		DrawDate         string `json:"drawDate"`
		TimeSinceDraw    int    `json:"timeSinceDraw"`
		Draw             int64  `json:"draw"`
		Bonus            int    `json:"bonus"`
		Num              []int  `json:"num"`
		NextKenoDraw     string `json:"nextKenoDraw"`
		NextKenoDrawTime int    `json:"nextKenoDrawTime"`
	}
	aTmpR := TmpR{}
	e := stringKit.GetJsonObj(resultData, &aTmpR)
	if e != nil {
		return aNewAwardInfo, errors.New("有问题的开奖号码的:" + resultData)
	}
	iLen := len(aTmpR.Num)
	if iLen != 20 {
		return aNewAwardInfo, fmt.Errorf("数据结果不对")
	}

	if len(aTmpR.Num) > 0 {
		if aTmpR.Num[0] == 0 {
			return aNewAwardInfo, fmt.Errorf("")
		}
	}

	aNewAwardInfo.LotteryStr = fmt.Sprintf("%d", aTmpR.Draw)
	if aNewAwardInfo.LotteryStr == this.aNewAwardInfo.LotteryStr {
		return aNewAwardInfo, fmt.Errorf("")
	}

	aNewAwardInfo.LotteryNum = aTmpR.Draw
	aNewAwardInfo.NextLotteryStr = fmt.Sprintf("%d", aNewAwardInfo.LotteryNum+1)

	for i := 0; i < iLen; i++ {
		for j := 0; j < iLen-i-1; j++ {
			if aTmpR.Num[j] > aTmpR.Num[j+1] {
				tmp := aTmpR.Num[j]
				aTmpR.Num[j] = aTmpR.Num[j+1]
				aTmpR.Num[j+1] = tmp
			}
		}
	}
	num1 := getNum(1, aTmpR.Num)
	num2 := getNum(2, aTmpR.Num)
	num3 := getNum(3, aTmpR.Num)

	aNewAwardInfo.ResultNums = fmt.Sprintf("%d,%d,%d", num1, num2, num3)
	Local_Shanghai, e := time.LoadLocation(`Asia/Shanghai`)
	curT := time.Now()
	if e == nil {
		curT = curT.In(Local_Shanghai)
	}

	aNewAwardInfo.CurLotteryTime = curT
	aNewAwardInfo.NextLotteryTime = aNewAwardInfo.CurLotteryTime.Add(time.Second * time.Duration(aTmpR.NextKenoDrawTime))

	return aNewAwardInfo, nil
}

func getNum(beginI int, nums []int) int {
	s := 0
	iLen := len(nums)
	for i := beginI; i < iLen-1; i += 3 {
		s += nums[i]
	}
	s = s % 10
	return s
}
