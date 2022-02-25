package LotteryResult

import (
	"encoding/json"
	"fmt"
	"time"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/WsxBox"
	"ttmyth123/GroupLottery/GameServer/models"
	"ttmyth123/kit/httpClientKit"
	"ttmyth123/kit/strconvEx"
	"ttmyth123/kit/stringKit"
	"ttmyth123/kit/timeKit"
	"ttmyth123/kit/ttLog"
)

type WsxZbcResultServer struct {
	CurOpenCodeInfo WsxBox.WsxResultData

	resultUrl             string
	mHttpClient           *httpClientKit.HttpClient
	mHandlerAwardInfoFunc HandlerAwardInfoFunc
}

func NewWsxZbcResultServer(resultUrl string, aAwardInfoFunc HandlerAwardInfoFunc) *WsxZbcResultServer {
	aWsxResultServer := new(WsxZbcResultServer)
	aWsxResultServer.resultUrl = resultUrl
	aWsxResultServer.mHttpClient = httpClientKit.GetHttpClient("")
	aWsxResultServer.mHandlerAwardInfoFunc = aAwardInfoFunc
	aWsxResultServer.run()
	return aWsxResultServer
}

func (this *WsxZbcResultServer) run() {
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

func (this *WsxZbcResultServer) saveAwardInfo() {
	aAwardInfo, e := this.getAwardInfo()
	if e != nil {
		return
	}

	if this.CurOpenCodeInfo.NextIssue != aAwardInfo.NextIssue {
		aOpenCodeInfo := make(map[int]string)
		e = json.Unmarshal([]byte(aAwardInfo.Opencode), &aOpenCodeInfo)
		if e != nil {
			ttLog.LogError(e, "strJson:", aAwardInfo.Opencode)
			return
		}

		aResultNums := make(map[string]interface{})
		Nums := make([]string, 20)
		for i := 0; i < 20; i++ {
			Nums[i] = aOpenCodeInfo[i]
		}

		aResultNums["Nums"] = Nums
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
		aLoAwardInfo.CurLotteryTime = curTime

		this.CurOpenCodeInfo = *aAwardInfo

		if this.mHandlerAwardInfoFunc != nil {
			go this.mHandlerAwardInfoFunc(aLoAwardInfo)
		}
	}
}

func (this *WsxZbcResultServer) getAwardInfo() (*WsxBox.WsxResultData, error) {
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

	if len(aWsxResultInfo.Data) > 0 {
		aWsxResultData = aWsxResultInfo.Data[0]
		aWsxResultData.ServerTime = aWsxResultInfo.ServerTime
	}

	//-------------------------------------------
	//aWsxResultData = getTestWsxResultData()
	//-------------------------------------------

	return &aWsxResultData, nil
}
func getTestWsxResultData() WsxBox.WsxResultData {
	t := time.Now()

	aWsxResultData := WsxBox.WsxResultData{}
	aWsxResultData.Opencode = getTestOpencode()
	aWsxResultData.Expect = t.Add(-time.Hour).Format("20060102") //20200805
	aWsxResultData.Opentime = t.Add(-time.Hour).Format(timeKit.DateTimeLayout)

	aWsxResultData.NextTime = t.Add(time.Hour * 23).Format(timeKit.DateTimeLayout)
	aWsxResultData.NextIssue = t.Add(time.Hour * 23).Format("20060102") //20200806
	aWsxResultData.ServerTime = t.Format(timeKit.DateTimeLayout)

	return aWsxResultData
}

func getTestOpencode() string {
	aResultNums := make(map[string]interface{})
	Nums := make([]string, 20)
	tmpM := make(map[int]int)
	for i := 0; i < 20; i++ {
		num := timeKit.GlobaRand.Intn(99)
		if _, ok := tmpM[num]; !ok {
			tmpM[num] = 1
			Nums[i] = fmt.Sprintf("%02d", num)
		}
	}
	aResultNums["Nums"] = Nums
	return stringKit.GetJsonStr(aResultNums)
}
