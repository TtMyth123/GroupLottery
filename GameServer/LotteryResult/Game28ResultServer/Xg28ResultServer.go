package Game28ResultServer

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"ttmyth123/GroupLottery/GameServer/GInstance/GTtHint"
	"ttmyth123/GroupLottery/GameServer/LotteryResult"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/ResultBox"
	"ttmyth123/GroupLottery/GameServer/models"
	"ttmyth123/kit/httpClientKit"
	"ttmyth123/kit/lotteryKit"
	"ttmyth123/kit/strconvEx"
	"ttmyth123/kit/stringKit"
	"ttmyth123/kit/timeKit"
	"ttmyth123/kit/ttLog"
)

type Xg28ResultServer struct {
	CurOpenCodeInfo       ResultBox.Game28Result
	resultUrl             string
	mHttpClient           *httpClientKit.HttpClient
	mHandlerAwardInfoFunc LotteryResult.HandlerAwardInfoFunc
	pStrTime              string
	arrResultTime         []string
}

func NewXg28ResultServer(resultUrl string, aAwardInfoFunc LotteryResult.HandlerAwardInfoFunc) *Xg28ResultServer {
	aResultServer := new(Xg28ResultServer)
	aResultServer.resultUrl = resultUrl
	aResultServer.mHttpClient = httpClientKit.GetHttpClient("")
	aResultServer.mHandlerAwardInfoFunc = aAwardInfoFunc
	aResultServer.ReLoadResultTime()
	aResultServer.run()
	return aResultServer
}

func (this *Xg28ResultServer) GetNextTime(openT time.Time) (time.Time, error) {
	t := time.Now()
	if t.Before(openT) {
		t = openT
	}
	if len(this.arrResultTime) == 0 {
		//return t, errors.New(GTtHint.GetTtHint().GetHint("无效的时间数据"))
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

func (this *Xg28ResultServer) ReLoadResultTime() {
	this.arrResultTime = this.getResultTime("02:30")
}
func (this *Xg28ResultServer) getResultTime(strT string) []string {
	arrTime := make([]string, 0)
	t, _ := timeKit.GetTime("2000-01-01 00:00:00")
	oldDay := t.Day()
	for {
		strT := t.Format("15:04:05")
		arrTime = append(arrTime, strT)
		t = t.Add(time.Second * 60 * 3)
		if t.Day() != oldDay {
			break
		}
	}

	fmt.Println(arrTime)
	return arrTime
}

func (this *Xg28ResultServer) run() {
	ticker1 := time.NewTicker(2 * time.Second)
	go func() {
		for {
			select {
			case <-ticker1.C:
				this.saveAwardInfo()
			}
		}
	}()
}

func (this *Xg28ResultServer) saveAwardInfo() {
	aAwardInfoData, e := this.getAwardInfo()
	if e != nil {
		return
	}
	aAwardInfo := aAwardInfoData.(ResultBox.Game28Result)
	if this.CurOpenCodeInfo.NextIssue != aAwardInfo.NextIssue {
		aLoAwardInfo := models.LoAwardInfo{}
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

		aLoAwardInfo.NextLotteryTime, e = timeKit.GetTime(aAwardInfo.NextTime)
		//aLoAwardInfo.NextLotteryTime,e = this.GetNextTime()
		if e != nil {
			ttLog.LogError(e, "Xg28ResultServer 开奖时间有问题:", stringKit.GetJsonStr(aAwardInfo))
			return
		}

		this.CurOpenCodeInfo = aAwardInfo
		ttLog.LogDebug("Xg28ResultServer 开奖信息：", stringKit.GetJsonStr(aAwardInfo))

		if this.mHandlerAwardInfoFunc != nil {
			go this.mHandlerAwardInfoFunc(aLoAwardInfo)
		}
	}
}

func (this *Xg28ResultServer) getAwardInfo() (interface{}, error) {
	aXG28Result := ResultBox.XG28Result{}
	//{"row":1,"code":"xjp28","data":[{"opentime":"2020-08-16 14:00:00","expect":"3530926","opencode":"6,8,1","NextTime":"2020-08-16 14:02:00","NextIssue":"3530927"}],"SleepTime":0,"SleepMinute":0,"LastDataStr":"","LastIssue":"","SetTimerIng":false,"LastMsg":{},"Lock":{}}
	resurl, e := this.mHttpClient.GetBytes(this.resultUrl)
	if e != nil {
		ttLog.LogError(e)
	}
	e = json.Unmarshal(resurl, &aXG28Result)
	if e != nil {
		ttLog.LogError(e, "strJson:", string(resurl))
		return nil, e
	}

	if len(aXG28Result.Data.Arr) == 0 {
		return nil, errors.New("有问题的数据")
	}

	Opencode := lotteryKit.GetStrNum2ArrStr(aXG28Result.Data.Arr[0].Result, 1)
	aGame28Result := ResultBox.Game28Result{
		NextIssue: aXG28Result.Data.NextIssue,
		NextTime:  aXG28Result.Data.NextTime,
		Expect:    aXG28Result.Data.Arr[0].Issue,
		Opentime:  aXG28Result.Data.Arr[0].Time,
		Opencode:  Opencode,
	}

	return aGame28Result, nil
}
