package Lottery28Game

import (
	"errors"
	"fmt"
	"github.com/TtMyth123/LotteryResultServer/GInstance/lottery"
	"github.com/TtMyth123/kit"
	"github.com/TtMyth123/kit/httpClientKit"
	"github.com/TtMyth123/kit/lotteryKit"
	"github.com/TtMyth123/kit/strconvEx"
	"github.com/TtMyth123/kit/stringKit"
	"github.com/TtMyth123/kit/timeKit"
	"github.com/astaxie/beego"
	"strings"
	"sync"
	"time"
	//"github.com/TtMyth123/LotteryResultServer/GInstance/lottery"
)

type GameJnd28 struct {
	timeI time.Duration
	t     int

	CurNums       string
	mHttpClient   *httpClientKit.HttpClient
	aNewAwardInfo lottery.NewAwardInfo
	lottery.ILotteryServer
	resultUrl string
	dataLock  sync.Mutex
}

func NewGameJnd28(timeI time.Duration, t int) *GameJnd28 {
	aGame28 := new(GameJnd28)
	aGame28.mHttpClient = httpClientKit.GetHttpClient("")
	aGame28.timeI = timeI
	aGame28.t = t
	aGame28.resultUrl = beego.AppConfig.String("Game28::ResultHttpUrlJnd")

	aGame28.aNewAwardInfo, _ = aGame28.getAwardInfo()
	aGame28.run()
	return aGame28
}

func (this *GameJnd28) GetNewAwardInfo() lottery.NewAwardInfo {
	return this.aNewAwardInfo
}

func (this *GameJnd28) run() {
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
func (this *GameJnd28) getAwardInfo() (lottery.NewAwardInfo, error) {
	aNewAwardInfo := lottery.NewAwardInfo{}

	//aJnd28Result := ResultBox.Game28Result{}
	//https://www.kjw6589.com/data/Current/jndpc28/CurIssue.json?NL8RN4R1KWP3KPSZQXAJFMXXHW9X6BXNNOXXAYRYWDZNC7QLT4
	rand := kit.GetGuid()
	rand = rand + rand
	rand = strings.ToUpper(rand[:50])
	this.dataLock.Lock()
	kjw6589Url := fmt.Sprintf(`%s?%s`, this.resultUrl, rand)
	this.dataLock.Unlock()
	//{"row":1,"code":"xjp28","data":[{"opentime":"2020-08-16 14:00:00","expect":"3530926","opencode":"6,8,1","NextTime":"2020-08-16 14:02:00","NextIssue":"3530927"}],"SleepTime":0,"SleepMinute":0,"LastDataStr":"","LastIssue":"","SetTimerIng":false,"LastMsg":{},"Lock":{}}
	resultHtml, e := this.mHttpClient.GetBytes(kjw6589Url)
	if e != nil {
		return aNewAwardInfo, e
	}

	aNewAwardInfo, e = kjw6589Result2NewAwardInfo(string(resultHtml))

	return aNewAwardInfo, e
}
func (this *GameJnd28) UpdateResultUrl(url string) error {
	this.dataLock.Lock()
	defer this.dataLock.Unlock()
	this.resultUrl = url
	return nil
}
func kjw6589Result2NewAwardInfo(resultData string) (lottery.NewAwardInfo, error) {
	aNewAwardInfo := lottery.NewAwardInfo{}

	kjw6589MapData := make(map[string]interface{})
	e := stringKit.GetJsonObj(resultData, &kjw6589MapData)
	if e != nil {
		return aNewAwardInfo, errors.New("有问题的开奖号码的:" + resultData)
	}
	aNewAwardInfo.NextLotteryStr = kit.GetInterface2Str(kjw6589MapData["issue"], "")
	aNewAwardInfo.LotteryStr = kit.GetInterface2Str(kjw6589MapData["preIssue"], "")
	aNewAwardInfo.LotteryNum = strconvEx.StrTry2Int64(aNewAwardInfo.LotteryStr, 0)

	if openNum, ok := kjw6589MapData["openNum"].([]interface{}); !ok {
		return aNewAwardInfo, errors.New("有问题的开奖号码2:" + resultData)
	} else {
		iLen := len(openNum)
		if iLen == 3 {
			arrNums := make([]int, iLen)
			for i := 0; i < iLen; i++ {
				arrNums[i] = kit.GetInterface2Int(openNum[i], 0)
			}
			aNewAwardInfo.ResultNums = lotteryKit.GetArrNum2String(arrNums, 1)
		} else {
			return aNewAwardInfo, errors.New("有问题的开奖号码3:" + resultData)
		}
	}

	lopenDateTime := kit.GetInterface2Int64(kjw6589MapData["openDateTime"], 0)
	openDateTime := timeKit.NewTimeByJavaTimeLong(lopenDateTime)
	aNewAwardInfo.NextLotteryTime = openDateTime

	lcurrentOpenDateTime := kit.GetInterface2Int64(kjw6589MapData["currentOpenDateTime"], 0)
	currentOpenDateTime := timeKit.NewTimeByJavaTimeLong(lcurrentOpenDateTime)
	aNewAwardInfo.CurLotteryTime = currentOpenDateTime

	return aNewAwardInfo, nil
}

//func (this *GameJnd28) getAwardInfoBy_kjw6589() (Game28Result, error) {
//	//aJnd28Result := ResultBox.Game28Result{}
//	//https://www.kjw6589.com/data/Current/jndpc28/CurIssue.json?NL8RN4R1KWP3KPSZQXAJFMXXHW9X6BXNNOXXAYRYWDZNC7QLT4
//	rand :=kit.GetGuid()
//	rand=rand+rand
//	rand = strings.ToUpper(rand[:50])
//	//kjw6589Url := fmt.Sprintf(`https://www.kjw6589.com/data/Current/jndpc28/CurIssue.json?%s`,rand)
//	kjw6589Url := fmt.Sprintf(`%s?%s`,this.resultUrl,rand)
//	//{"row":1,"code":"xjp28","data":[{"opentime":"2020-08-16 14:00:00","expect":"3530926","opencode":"6,8,1","NextTime":"2020-08-16 14:02:00","NextIssue":"3530927"}],"SleepTime":0,"SleepMinute":0,"LastDataStr":"","LastIssue":"","SetTimerIng":false,"LastMsg":{},"Lock":{}}
//	//resurl, e := this.mHttpClient.GetBytes(kjw6589Url)
//	paramsHeader := make(map[string]string)
//	paramsHeader["accept"] = `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3`
//	paramsHeader["Connection"] = ``
//	paramsHeader["accept-encoding"] = `gzip, deflate, br`
//	paramsHeader["accept-language"] = `zh-CN,zh;q=0.9`
//	paramsHeader["cache-control"] = `max-age=0`
//	paramsHeader["if-modified-since"] = `Wed, 07 Apr 2021 19:29:21 GMT`
//	paramsHeader["if-none-match"] = `"606e0811-14e"`
//	paramsHeader["upgrade-insecure-requests"] = `1`
//	//paramsHeader["Connection"] = ``
//	//paramsHeader["Connection"] = "aaaaa"
//	//paramsHeader["Connection"] = "aaaaa"
//	resurl, e :=this.mHttpClient.DoRequest("GET",kjw6589Url,paramsHeader,nil)
//
//	aJnd28Result,e := kjw6589Result2Game28Result(string(resurl))
//
//	return aJnd28Result, e
//}

//func kjw6589Result2Game28Result(resultData string) (Game28Result,error) {
//	aJnd28Result := Game28Result{}
//
//	kjw6589MapData := make(map[string]interface{})
//	e := stringKit.GetJsonObj(resultData, &kjw6589MapData)
//	if e != nil {
//		return aJnd28Result,errors.New("有问题的开奖号码1:"+ resultData)
//	}
//	aJnd28Result.NextIssue = kit.GetInterface2Str(kjw6589MapData["issue"],"")
//	aJnd28Result.Expect = kit.GetInterface2Str(kjw6589MapData["preIssue"],"")
//	if openNum,ok :=  kjw6589MapData["openNum"].([]interface{});!ok {
//		return aJnd28Result,errors.New("有问题的开奖号码2:"+ resultData)
//	} else {
//		iLen :=len(openNum)
//		if iLen==3 {
//			aJnd28Result.Opencode = make([]string,iLen)
//			for i:=0;i<iLen;i++{
//				aJnd28Result.Opencode[i] = kit.GetInterface2Str(openNum[i],"0")
//			}
//		} else {
//			return aJnd28Result,errors.New("有问题的开奖号码3:"+ resultData)
//		}
//	}
//	lcurrentOpenDateTime:= kit.GetInterface2Int64(kjw6589MapData["currentOpenDateTime"],0)
//	currentOpenDateTime := timeKit.NewTimeByJavaTimeLong(lcurrentOpenDateTime)
//
//	aJnd28Result.Opentime = currentOpenDateTime.Format(timeKit.DateTimeLayout)
//	return  aJnd28Result,nil
//}
