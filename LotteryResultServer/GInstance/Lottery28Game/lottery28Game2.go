package Lottery28Game

import (
	"errors"
	"fmt"
	"github.com/TtMyth123/LotteryResultServer/GInstance/lottery"
	"github.com/TtMyth123/kit/httpClientKit"
	"github.com/TtMyth123/kit/strconvEx"
	"github.com/TtMyth123/kit/stringKit"
	"github.com/TtMyth123/kit/timeKit"
	"github.com/astaxie/beego"
	"sync"
	"time"
)

type Game2Jnd28 struct {
	timeI time.Duration
	t     int

	CurNums       string
	mHttpClient   *httpClientKit.HttpClient
	aNewAwardInfo lottery.NewAwardInfo
	lottery.ILotteryServer
	resultUrl string
	dataLock  sync.Mutex
}

func NewGame2Jnd28(timeI time.Duration, t int) *Game2Jnd28 {
	aGame28 := new(Game2Jnd28)
	aGame28.mHttpClient = httpClientKit.GetHttpClient("")
	aGame28.timeI = timeI
	aGame28.t = t
	aGame28.resultUrl = beego.AppConfig.String("Game28::ResultHttpUrl2Jnd")

	aGame28.aNewAwardInfo, _ = aGame28.getAwardInfo()
	aGame28.run()
	return aGame28
}

func (this *Game2Jnd28) GetNewAwardInfo() lottery.NewAwardInfo {
	return this.aNewAwardInfo
}

func (this *Game2Jnd28) run() {
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
func (this *Game2Jnd28) getAwardInfo() (lottery.NewAwardInfo, error) {
	aNewAwardInfo := lottery.NewAwardInfo{}

	//aJnd28Result := ResultBox.Game28Result{}
	//https://www.kjw6589.com/data/Current/jndpc28/CurIssue.json?NL8RN4R1KWP3KPSZQXAJFMXXHW9X6BXNNOXXAYRYWDZNC7QLT4
	this.dataLock.Lock()
	kjw6589Url := fmt.Sprintf(`%s`, this.resultUrl)
	this.dataLock.Unlock()
	//{"row":1,"code":"xjp28","data":[{"opentime":"2020-08-16 14:00:00","expect":"3530926","opencode":"6,8,1","NextTime":"2020-08-16 14:02:00","NextIssue":"3530927"}],"SleepTime":0,"SleepMinute":0,"LastDataStr":"","LastIssue":"","SetTimerIng":false,"LastMsg":{},"Lock":{}}
	resultHtml, e := this.mHttpClient.GetBytes(kjw6589Url)
	if e != nil {
		return aNewAwardInfo, e
	}

	aNewAwardInfo, e = this.aoweidbResult2NewAwardInfo(string(resultHtml))

	return aNewAwardInfo, e
}
func (this *Game2Jnd28) UpdateResultUrl(url string) error {
	this.dataLock.Lock()
	defer this.dataLock.Unlock()
	this.resultUrl = url
	return nil
}

/**
resultData={"rows":1,"code":"xjn28","message":"","data":[{"expect":"2698060","opencode":"5,7,6","opentime":"2021-04-10 01:09:00"}]}
*/
func (this *Game2Jnd28) aoweidbResult2NewAwardInfo(resultData string) (lottery.NewAwardInfo, error) {
	aNewAwardInfo := lottery.NewAwardInfo{}
	type TmpRData struct {
		Expect   string `json:"expect"`
		Opencode string `json:"opencode"`
		Opentime string `json:"opentime"`
	}
	type TmpR struct {
		Data []TmpRData `json:"data"`
	}

	aTmpR := TmpR{}
	e := stringKit.GetJsonObj(resultData, &aTmpR)
	if e != nil {
		return aNewAwardInfo, errors.New("有问题的开奖号码的:" + resultData)
	}
	if len(aTmpR.Data) == 0 {
		return aNewAwardInfo, errors.New("有问题的开奖号码的2:" + resultData)
	}

	aNewAwardInfo.LotteryStr = aTmpR.Data[0].Expect
	aNewAwardInfo.LotteryNum = strconvEx.StrTry2Int64(aTmpR.Data[0].Expect, 0)
	aNewAwardInfo.NextLotteryStr = fmt.Sprintf("%d", aNewAwardInfo.LotteryNum+1)
	aNewAwardInfo.ResultNums = aTmpR.Data[0].Opencode

	aNewAwardInfo.CurLotteryTime, e = timeKit.GetTime(aTmpR.Data[0].Opentime)
	if e != nil {
		return aNewAwardInfo, errors.New("有问题的开奖时间:" + resultData)
	}
	aNewAwardInfo.NextLotteryTime = aNewAwardInfo.CurLotteryTime.Add(time.Second * 210)

	return aNewAwardInfo, nil
}
