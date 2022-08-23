package LotteryResult

import (
	"encoding/json"
	"github.com/TtMyth123/GameServer/LotteryResult/WsxBbcResultKit"
	"github.com/TtMyth123/GameServer/LotteryResult/WsxBox"
	"github.com/TtMyth123/GameServer/models"
	"github.com/TtMyth123/kit/httpClientKit"
	"github.com/TtMyth123/kit/strconvEx"
	"github.com/TtMyth123/kit/stringKit"
	"github.com/TtMyth123/kit/timeKit"
	"github.com/TtMyth123/kit/ttLog"
	"time"
)

type WsxBbcResultServer struct {
	CurOpenCodeInfo WsxBox.WsxResultData

	resultUrl             string
	mHttpClient           *httpClientKit.HttpClient
	mHandlerAwardInfoFunc HandlerAwardInfoFunc
}

func NewWsxBbcResultServer(resultUrl string, aAwardInfoFunc HandlerAwardInfoFunc) *WsxBbcResultServer {
	aWsxResultServer := new(WsxBbcResultServer)
	aWsxResultServer.resultUrl = resultUrl
	aWsxResultServer.mHttpClient = httpClientKit.GetHttpClient("")
	aWsxResultServer.mHandlerAwardInfoFunc = aAwardInfoFunc
	aWsxResultServer.run()
	return aWsxResultServer
}

func (this *WsxBbcResultServer) run() {
	ticker1 := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker1.C:
				this.saveAwardInfo()
			}
		}
	}()
}

func (this *WsxBbcResultServer) saveAwardInfo() {
	aAwardInfo, e := this.getAwardInfo()
	if e != nil {
		return
	}

	if this.CurOpenCodeInfo.NextIssue != aAwardInfo.NextIssue {
		aOpenCodeInfo := WsxBox.HttpOpenCodeBbc{}
		e = json.Unmarshal([]byte(aAwardInfo.Opencode), &aOpenCodeInfo)
		if e != nil {
			ttLog.LogError(e, "strJson:", aAwardInfo.Opencode)
			return
		}

		aResultNums := WsxBbcResultKit.GetResultNums(aOpenCodeInfo)
		//aResultNums := make( map[string]interface{})
		//aResultNums["jackpots"] = aOpenCodeInfo.Jackpots
		//aResultNums["firstNum"] = aOpenCodeInfo.FirstNum
		//aResultNums["secondNum"] = aOpenCodeInfo.SecondNum
		//aResultNums["thirdNum"] = aOpenCodeInfo.ThirdNum
		//aResultNums["forthNum"] = aOpenCodeInfo.ForthNum
		//aResultNums["fifthNum"] = aOpenCodeInfo.FifthNum
		//aResultNums["sixthNum"] = aOpenCodeInfo.SixthNum
		//aResultNums["seventhNum"] = aOpenCodeInfo.SeventhNum
		//aResultNums["eighthNum"] = aOpenCodeInfo.EighthNum

		aLoAwardInfo := models.LoAwardInfo{}

		aLoAwardInfo.LotteryNum = strconvEx.StrTry2Int64(aAwardInfo.Expect, 0)
		aLoAwardInfo.LotteryStr = aAwardInfo.Expect
		aLoAwardInfo.ResultNums = stringKit.GetJsonStr(aResultNums)
		aLoAwardInfo.OriginalResult = stringKit.GetJsonStr(aAwardInfo)
		aLoAwardInfo.NextLotteryStr = aAwardInfo.NextIssue

		//aLoAwardInfo.CurLotteryTime,e = timeKit.GetTime(aAwardInfo.Opentime)
		//if e!= nil {
		//	ttLog.LogError(e,"开奖时间有问题:",aAwardInfo.Opentime)
		//	return
		//}
		//aLoAwardInfo.NextLotteryTime,e = timeKit.GetTime(aAwardInfo.NextTime)
		//if e!= nil {
		//	ttLog.LogError(e,"开奖时间有问题:",aAwardInfo.NextTime)
		//	return
		//}

		curTime, e := timeKit.GetTime(aAwardInfo.ServerTime)
		if e != nil {
			ttLog.LogError(e, "开奖时间有问题:", aAwardInfo.Opentime)
			return
		}
		aLoAwardInfo.CreatedAt = curTime
		NextTime := time.Duration(strconvEx.StrTry2Int64(aAwardInfo.NextTime, 0)) * time.Second
		aLoAwardInfo.NextTime = aAwardInfo.NextTime
		aLoAwardInfo.NextLotteryTime = curTime.Add(NextTime)
		aLoAwardInfo.CurLotteryTime = aLoAwardInfo.NextLotteryTime.AddDate(0, 0, -1)

		this.CurOpenCodeInfo = *aAwardInfo

		if this.mHandlerAwardInfoFunc != nil {
			go this.mHandlerAwardInfoFunc(aLoAwardInfo)
		}
	}
}

func (this *WsxBbcResultServer) getAwardInfo() (*WsxBox.WsxResultData, error) {
	aWsxResultData := WsxBox.WsxResultData{}
	aWsxResultInfo := WsxBox.WsxResultInfo{}
	resurl, e := this.mHttpClient.GetBytes(this.resultUrl)
	if e != nil {
		ttLog.LogError(e)
	}
	e = json.Unmarshal(resurl, &aWsxResultInfo)
	if e != nil {
		ttLog.LogError(e, "strJson:", string(resurl))
		return nil, e
	}

	//if len(	aWsxResultInfo.Data)>0 {
	//	e = json.Unmarshal( []byte(aWsxResultInfo.Data[0].Opencode), &aOpenCodeInfo)
	//	if e!= nil {
	//		ttLog.LogError(e,"strJson:",string(aWsxResultInfo.Data[0].Opencode))
	//		return nil,e
	//	}
	//}
	if len(aWsxResultInfo.Data) > 0 {
		aWsxResultData = aWsxResultInfo.Data[0]
		aWsxResultData.ServerTime = aWsxResultInfo.ServerTime
	}

	return &aWsxResultData, nil
}
