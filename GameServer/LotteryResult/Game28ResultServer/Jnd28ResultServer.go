package Game28ResultServer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"sync"
	"time"
	"ttmyth123/GroupLottery/GameServer/GInstance/GTtHint"
	"ttmyth123/GroupLottery/GameServer/LotteryResult"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/ResultBox"
	"ttmyth123/GroupLottery/GameServer/models"
	"ttmyth123/kit"
	"ttmyth123/kit/httpClientKit"
	"ttmyth123/kit/lotteryKit"
	"ttmyth123/kit/strconvEx"
	"ttmyth123/kit/stringKit"
	"ttmyth123/kit/timeKit"
	"ttmyth123/kit/ttLog"
)

type Jnd28ResultServer struct {
	CurOpenCodeInfo ResultBox.Game28Result
	//resultUrl             string
	mHttpClient           *httpClientKit.HttpClient
	mHandlerAwardInfoFunc LotteryResult.HandlerAwardInfoFunc

	saveAwardInfoLock sync.RWMutex
	arrResultTime     []string
	pStrTime          string

	resultArrUrl map[string]string
}

func NewJnd28ResultServer(resultUrl string, aAwardInfoFunc LotteryResult.HandlerAwardInfoFunc) *Jnd28ResultServer {
	aResultServer := new(Jnd28ResultServer)
	aResultServer.resultArrUrl = make(map[string]string)
	aResultServer.resultArrUrl[resultUrl] = resultUrl
	resultUrl = beego.AppConfig.String("Game28::ResultHttpUrlJnd1")
	if resultUrl != "" {
		aResultServer.resultArrUrl[resultUrl] = resultUrl
	}
	resultUrl = beego.AppConfig.String("Game28::ResultHttpUrlJnd2")
	if resultUrl != "" {
		aResultServer.resultArrUrl[resultUrl] = resultUrl
	}

	aResultServer.mHttpClient = httpClientKit.GetHttpClient("")
	aResultServer.mHandlerAwardInfoFunc = aAwardInfoFunc

	strT := beego.AppConfig.String("Game28::JndBeginTime")
	if strT == "" {
		strT = "02:30"
	}
	aResultServer.ReLoadResultTime(strT)
	aResultServer.run()
	return aResultServer
}

//func Int2Minute(t time.Time) (time.Time,error) {
//	m := t.Minute()
//	strT := t.Format("2006-01-02 15:04")
//	addM := time.Duration(1)
//	if m%2==0 {
//		addM = 2
//	}
//
//	newT, e := time.ParseInLocation("2006-01-02 15:04", strT, time.Local)
//	newT= newT.Add(addM*time.Minute)
//
//	return newT,e
//}

//func GetNextTime() (time.Time, error) {
//	t := time.Now()
//	if len(arrResultTime) == 0 {
//		return t, errors.New("无效的时间数据")
//	}
//
//	strT := t.Format("15:04:05")
//	for i := 0; i < len(arrResultTime); i++ {
//		if strT < arrResultTime[i] {
//
//			ss := fmt.Sprintf("%s %s", t.Format("2006-01-02"), arrResultTime[i])
//			newT, e := timeKit.GetTime(ss)
//			return newT, e
//		}
//	}
//	t = t.Add(time.Hour)
//	ss := fmt.Sprintf("%s %s", t.Format("2006-01-02"), arrResultTime[0])
//	newT, e := timeKit.GetTime(ss)
//
//	return newT, e
//}

func (this *Jnd28ResultServer) GetNextTime(openT time.Time, rt int) (time.Time, error) {
	t := time.Now()
	if rt == 1 {
		if t.Before(openT) {
			t = openT
		}
	} else {
		t = openT
	}
	if len(this.arrResultTime) == 0 {
		return t, errors.New(GTtHint.GetTtHint().GetHint("无效的时间数据"))
	}

	strT := t.Format("15:04:05")
	for i := 0; i < len(this.arrResultTime); i++ {
		if strT < this.arrResultTime[i] {
			ss := fmt.Sprintf("%s %s", t.Format("2006-01-02"), this.arrResultTime[i])
			if this.pStrTime == ss {
				return t, errors.New(GTtHint.GetTtHint().GetHint("无效的时间数据1"))
			}

			newT, e := timeKit.GetTime(ss)
			if e == nil {
				iSub := newT.Sub(t) / time.Second
				if iSub < 6 {
					return newT, errors.New(GTtHint.GetTtHint().GetHint("无效的时间数据2"))
				}
			}
			this.pStrTime = ss
			return newT, e
		}
	}
	t = t.Add(time.Hour)
	ss := fmt.Sprintf("%s %s", t.Format("2006-01-02"), this.arrResultTime[0])
	newT, e := timeKit.GetTime(ss)

	return newT, e
}

func (this *Jnd28ResultServer) ReLoadResultTime(strT string) {

	this.arrResultTime = this.getResultTime(strT)
}
func (this *Jnd28ResultServer) getResultTime(strT string) []string {
	arrTime := make([]string, 0)
	if strT == "02:30" {
		t, _ := timeKit.GetTime("2000-01-01 00:02:30")
		for {
			strT := t.Format("15:04:05")
			arrTime = append(arrTime, strT)

			if t.Hour() == 19 {
				break
			}
			t = t.Add(time.Second * 210)
		}

		t, _ = timeKit.GetTime("2000-01-01 19:01:30")
		oldDay := t.Day()
		for {
			strT := t.Format("15:04:05")
			arrTime = append(arrTime, strT)
			t = t.Add(time.Second * 210)
			if t.Day() != oldDay {
				break
			}
		}
	} else if strT == "03:00" {
		t, _ := timeKit.GetTime("2000-01-01 00:03:00")
		for {
			strT := t.Format("15:04:05")
			arrTime = append(arrTime, strT)

			if t.Hour() == 20 {
				break
			}
			t = t.Add(time.Second * 210)
		}

		t, _ = timeKit.GetTime("2000-01-01 20:01:30")
		oldDay := t.Day()
		for {
			strT := t.Format("15:04:05")
			arrTime = append(arrTime, strT)
			t = t.Add(time.Second * 210)
			if t.Day() != oldDay {
				break
			}
		}
	}

	fmt.Println(arrTime)
	return arrTime
}

func (this *Jnd28ResultServer) run() {
	//ticker1 := time.NewTicker(2 * time.Second)
	ticker2 := time.NewTicker(6 * time.Second)
	go func() {
		for {
			select {
			//case <-ticker1.C:
			//	{
			//		aAwardInfoData, e := this.getAwardInfo()
			//		if e == nil {
			//			this.saveAwardInfo(aAwardInfoData)
			//		}
			//	}
			case <-ticker2.C:
				{
					aArrAwardInfoData, _ := this.getAwardInfoBy_kjw6589()

					for _, aAwardInfoData := range aArrAwardInfoData {
						this.saveAwardInfo(aAwardInfoData)
					}
				}
			}
		}
	}()
}

func (this *Jnd28ResultServer) saveAwardInfo(aAwardInfoData interface{}) {
	this.saveAwardInfoLock.Lock()
	defer this.saveAwardInfoLock.Unlock()

	aAwardInfo := aAwardInfoData.(ResultBox.Game28Result)
	if this.CurOpenCodeInfo.NextIssue < aAwardInfo.NextIssue {
		ttLog.LogDebug("Jnd28ResultServer__saveAwardInfo:", aAwardInfo.Expect, aAwardInfo.Opentime, aAwardInfo.Opencode)

		aLoAwardInfo := models.LoAwardInfo{}
		var e error
		aLoAwardInfo.LotteryNum = strconvEx.StrTry2Int64(aAwardInfo.Expect, 0)
		aLoAwardInfo.LotteryStr = aAwardInfo.Expect
		aLoAwardInfo.ResultNums = fmt.Sprintf("%s,%s,%s", aAwardInfo.Opencode[0], aAwardInfo.Opencode[1], aAwardInfo.Opencode[2])
		aLoAwardInfo.OriginalResult = stringKit.GetJsonStr(aAwardInfo)
		aLoAwardInfo.NextLotteryStr = aAwardInfo.NextIssue
		aLoAwardInfo.CurLotteryTime, e = timeKit.GetTime(aAwardInfo.Opentime)

		if e != nil {
			ttLog.LogError(e, "开奖时间有问题:", aAwardInfo.Opentime)
			return
		}

		//aLoAwardInfo.NextLotteryTime,e = timeKit.GetTime(aAwardInfo.NextTime)
		aLoAwardInfo.NextLotteryTime, e = this.GetNextTime(aLoAwardInfo.CurLotteryTime, aAwardInfo.R)
		if e != nil {
			ttLog.LogError(e, "Jnd28ResultServer 开奖时间有问题:", stringKit.GetJsonStr(aAwardInfo))
			return
		}

		this.CurOpenCodeInfo = aAwardInfo
		ttLog.LogDebug("Jnd28ResultServer 开奖信息：", stringKit.GetJsonStr(aAwardInfo))
		if this.mHandlerAwardInfoFunc != nil {
			go this.mHandlerAwardInfoFunc(aLoAwardInfo)
		}
	}
}

//func (this *Jnd28ResultServer) getAwardInfo() (interface{}, error) {
//	aJnd28Result := ResultBox.Game28Result{}
//	//{"row":1,"code":"xjp28","data":[{"opentime":"2020-08-16 14:00:00","expect":"3530926","opencode":"6,8,1","NextTime":"2020-08-16 14:02:00","NextIssue":"3530927"}],"SleepTime":0,"SleepMinute":0,"LastDataStr":"","LastIssue":"","SetTimerIng":false,"LastMsg":{},"Lock":{}}
//	resurl, e := this.mHttpClient.GetBytes(this.resultUrl)
//	if e != nil {
//		ttLog.LogError(e)
//	}
//	e = json.Unmarshal(resurl, &aJnd28Result)
//	if e != nil {
//		ttLog.LogError(e, "strJson:", string(resurl))
//		return nil, e
//	}
//	if len(aJnd28Result.Opencode) < 3 {
//		ttLog.LogError(e, "strJson:", string(resurl))
//		return nil, e
//	}
//	aJnd28Result.R = 1
//
//	return aJnd28Result, nil
//}

/**
resurl:=`{
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
}`
*/
func kjw6589Result2Game28Result(resurl string) (ResultBox.Game28Result, error) {
	aJnd28Result := ResultBox.Game28Result{}

	kjw6589MapData := make(map[string]interface{})
	e := stringKit.GetJsonObj(resurl, &kjw6589MapData)
	if e != nil {
		ttLog.LogError(e)
	}
	aJnd28Result.NextIssue = kit.GetInterface2Str(kjw6589MapData["issue"], "")
	aJnd28Result.Expect = kit.GetInterface2Str(kjw6589MapData["preIssue"], "")
	if openNum, ok := kjw6589MapData["openNum"].([]interface{}); !ok {
		//ttLog.LogError("有问题的开奖号码:",resurl)
		aa := stringKit.GetJsonStr(kjw6589MapData["openNum"])
		fmt.Println(aa)
		return aJnd28Result, errors.New("有问题的开奖号码")
	} else {
		iLen := len(openNum)
		if iLen == 3 {
			aJnd28Result.Opencode = make([]string, iLen)
			for i := 0; i < iLen; i++ {
				aJnd28Result.Opencode[i] = kit.GetInterface2Str(openNum[i], "0")
			}
		} else {
			//ttLog.LogError("有问题的开奖号码:",resurl)
			return aJnd28Result, errors.New("有问题的开奖号码")
		}
	}

	//lopenDateTime:= kit.GetInterface2Int64(kjw6589MapData["openDateTime"],0)
	//openDateTime1 := timeKit.NewTimeByJavaTimeLong(lopenDateTime)

	lcurrentOpenDateTime := kit.GetInterface2Int64(kjw6589MapData["currentOpenDateTime"], 0)
	currentOpenDateTime := timeKit.NewTimeByJavaTimeLong(lcurrentOpenDateTime)

	//lserverTime:= kit.GetInterface2Int64(kjw6589MapData["serverTime"],0)
	//serverTime := timeKit.NewTimeByJavaTimeLong(serverTime)

	//fmt.Println("openDateTime:",openDateTime1,"currentOpenDateTime:",currentOpenDateTime,"serverTime:",serverTime)

	aJnd28Result.Opentime = currentOpenDateTime.Format(timeKit.DateTimeLayout)

	return aJnd28Result, nil
}

func (this *Jnd28ResultServer) getAwardInfoBy_kjw6589() ([]interface{}, error) {

	if len(this.resultArrUrl) == 0 {
		return nil, errors.New("")
	}
	arrR := make([]interface{}, 0)
	for _, resultUrl2 := range this.resultArrUrl {
		aJnd28Result := ResultBox.Game28Result{}
		type NewAwardInfo struct {
			LotteryNum      int64  //期号
			LotteryStr      string //期号
			ResultNums      string
			NextLotteryStr  string
			NextLotteryTime time.Time //下一期开奖时间
			CurLotteryTime  time.Time //当前开期时间
		}
		aNewAwardInfo := NewAwardInfo{}

		resurl, e := this.mHttpClient.GetBytes(resultUrl2)
		if e != nil {
			continue
		}
		e = json.Unmarshal(resurl, &aNewAwardInfo)
		if e != nil {
			continue
		}

		aJnd28Result.NextIssue = aNewAwardInfo.NextLotteryStr
		aJnd28Result.NextTime = aNewAwardInfo.NextLotteryTime.Format(timeKit.DateTimeLayout)
		aJnd28Result.Expect = aNewAwardInfo.LotteryStr
		aJnd28Result.Opencode = lotteryKit.GetStrNum2ArrStr(aNewAwardInfo.ResultNums, 1)
		aJnd28Result.Opentime = aNewAwardInfo.CurLotteryTime.Add(time.Second).Format(timeKit.DateTimeLayout)
		aJnd28Result.R = 2
		arrR = append(arrR, aJnd28Result)
		//ttLog.LogDebug("aJnd28Result:clu",stringKit.GetJsonStr(aJnd28Result))
	}

	return arrR, nil
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
