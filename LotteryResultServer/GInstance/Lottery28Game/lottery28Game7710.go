package Lottery28Game

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"sync"
	"time"
	"ttmyth123/GroupLottery/LotteryResultServer/GInstance/lottery"
	"ttmyth123/kit/httpClientKit"
	"ttmyth123/kit/strconvEx"
	"ttmyth123/kit/stringKit"
	"ttmyth123/kit/timeKit"
)

type Game7710Jnd28 struct {
	timeI time.Duration
	t     int

	CurNums       string
	mHttpClient   *httpClientKit.HttpClient
	aNewAwardInfo lottery.NewAwardInfo
	lottery.ILotteryServer
	resultUrl string
	dataLock  sync.Mutex
}

func NewGame7710Jnd28(timeI time.Duration, t int) *Game7710Jnd28 {
	aGame28 := new(Game7710Jnd28)
	aGame28.mHttpClient = httpClientKit.GetHttpClient("")
	aGame28.timeI = timeI
	aGame28.t = t
	aGame28.resultUrl = beego.AppConfig.String("Game28::ResultHttpUrl7710Jnd")

	aGame28.aNewAwardInfo, _ = aGame28.getAwardInfo()
	aGame28.run()
	return aGame28
}

func (this *Game7710Jnd28) GetNewAwardInfo() lottery.NewAwardInfo {
	return this.aNewAwardInfo
}

func (this *Game7710Jnd28) run() {
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
func (this *Game7710Jnd28) getAwardInfo() (lottery.NewAwardInfo, error) {
	aNewAwardInfo := lottery.NewAwardInfo{}

	//aJnd28Result := ResultBox.Game28Result{}
	//https://www.kjw6589.com/data/Current/jndpc28/CurIssue.json?NL8RN4R1KWP3KPSZQXAJFMXXHW9X6BXNNOXXAYRYWDZNC7QLT4
	this.dataLock.Lock()
	kjw6589Url := fmt.Sprintf(`%s`, this.resultUrl)
	this.dataLock.Unlock()
	//{"row":1,"code":"jnd28","data":[{"opentime":"2021-05-30 16:26:00","expect":"2717909","opencode":"0,4,9","Timer":0,"NextTime":"2021-05-30 16:29:30","NextIssue":"2717910","NextTime2":"","NextIssue2":""}],"SleepTime":0,"SleepMinute":0,"LastDataStr":"","LastIssue":"","SetTimerIng":false,"ServerTime":"","Point":null,"LastMsg":{},"Lock":{}}
	resultHtml, e := this.mHttpClient.GetBytes(kjw6589Url)
	if e != nil {
		return aNewAwardInfo, e
	}

	aNewAwardInfo, e = this.aoweidbResult2NewAwardInfo(string(resultHtml))

	return aNewAwardInfo, e
}
func (this *Game7710Jnd28) UpdateResultUrl(url string) error {
	this.dataLock.Lock()
	defer this.dataLock.Unlock()
	this.resultUrl = url
	return nil
}

/**
resultData={"row":1,"code":"jnd28","data":[{"opentime":"2021-05-30 16:26:00","expect":"2717909","opencode":"0,4,9","Timer":0,"NextTime":"2021-05-30 16:29:30","NextIssue":"2717910","NextTime2":"","NextIssue2":""}],"SleepTime":0,"SleepMinute":0,"LastDataStr":"","LastIssue":"","SetTimerIng":false,"ServerTime":"","Point":null,"LastMsg":{},"Lock":{}}
*/
func (this *Game7710Jnd28) aoweidbResult2NewAwardInfo(resultData string) (lottery.NewAwardInfo, error) {
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
