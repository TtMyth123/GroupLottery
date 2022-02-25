package Game28ResultServer

import (
	"encoding/json"
	"time"
	"ttmyth123/GroupLottery/GameServer/LotteryResult"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/ResultBox"
	"ttmyth123/GroupLottery/GameServer/models"
	"ttmyth123/kit/httpClientKit"
	"ttmyth123/kit/strconvEx"
	"ttmyth123/kit/stringKit"
	"ttmyth123/kit/timeKit"
	"ttmyth123/kit/ttLog"
)

type Xjp28ResultServer struct {
	CurOpenCodeInfo       ResultBox.Xjp28Data
	resultUrl             string
	mHttpClient           *httpClientKit.HttpClient
	mHandlerAwardInfoFunc LotteryResult.HandlerAwardInfoFunc
}

func NewXjp28ResultServer(resultUrl string, aAwardInfoFunc LotteryResult.HandlerAwardInfoFunc) *Xjp28ResultServer {
	aResultServer := new(Xjp28ResultServer)
	aResultServer.resultUrl = resultUrl
	aResultServer.mHttpClient = httpClientKit.GetHttpClient("")
	aResultServer.mHandlerAwardInfoFunc = aAwardInfoFunc
	aResultServer.run()
	return aResultServer
}

func (this *Xjp28ResultServer) run() {
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

func (this *Xjp28ResultServer) saveAwardInfo() {
	aAwardInfoData, e := this.getAwardInfo()
	if e != nil {
		return
	}
	aAwardInfo := aAwardInfoData.(ResultBox.Xjp28Data)
	if this.CurOpenCodeInfo.NextIssue != aAwardInfo.NextIssue {

		aLoAwardInfo := models.LoAwardInfo{}

		aLoAwardInfo.LotteryNum = strconvEx.StrTry2Int64(aAwardInfo.Expect, 0)
		aLoAwardInfo.LotteryStr = aAwardInfo.Expect
		aLoAwardInfo.ResultNums = aAwardInfo.OpenCode
		aLoAwardInfo.OriginalResult = stringKit.GetJsonStr(aAwardInfo)
		aLoAwardInfo.NextLotteryStr = aAwardInfo.NextIssue
		aLoAwardInfo.CurLotteryTime, e = timeKit.GetTime(aAwardInfo.Opentime)
		if e != nil {
			ttLog.LogError(e, "开奖时间有问题:", aAwardInfo.Opentime)
			return
		}

		aLoAwardInfo.NextLotteryTime, e = timeKit.GetTime(aAwardInfo.NextTime)
		if e != nil {
			ttLog.LogError(e, "开奖时间有问题:", aAwardInfo.NextTime)
			return
		}

		this.CurOpenCodeInfo = aAwardInfo
		if this.mHandlerAwardInfoFunc != nil {
			go this.mHandlerAwardInfoFunc(aLoAwardInfo)
		}
	}
}

func (this *Xjp28ResultServer) getAwardInfo() (interface{}, error) {
	aXjp28Result := ResultBox.Xjp28Result{}
	aXjp28Data := ResultBox.Xjp28Data{}
	//{"row":1,"code":"xjp28","data":[{"opentime":"2020-08-16 14:00:00","expect":"3530926","opencode":"6,8,1","NextTime":"2020-08-16 14:02:00","NextIssue":"3530927"}],"SleepTime":0,"SleepMinute":0,"LastDataStr":"","LastIssue":"","SetTimerIng":false,"LastMsg":{},"Lock":{}}
	resurl, e := this.mHttpClient.GetBytes(this.resultUrl)
	if e != nil {
		ttLog.LogError(e)
	}
	e = json.Unmarshal(resurl, &aXjp28Result)
	if e != nil {
		ttLog.LogError(e, "strJson:", string(resurl))
		return nil, e
	}

	if len(aXjp28Result.Data) > 0 {
		aXjp28Data = aXjp28Result.Data[0]
	}
	return aXjp28Data, nil
}
