package UscResultServer

import (
	"encoding/json"
	"fmt"
	"github.com/TtMyth123/GameServer/LotteryResult"
	"github.com/TtMyth123/GameServer/LotteryResult/UscResultServer/UscBox"
	"github.com/TtMyth123/GameServer/models"
	"github.com/TtMyth123/kit/httpClientKit"
	"github.com/TtMyth123/kit/stringKit"
	"github.com/TtMyth123/kit/ttLog"
	"time"
)

type UscResultServer struct {
	CurOpenCodeInfo       UscBox.AwardInfo
	GameIndex             int
	resultUrl             string
	mHttpClient           *httpClientKit.HttpClient
	mHandlerAwardInfoFunc LotteryResult.HandlerAwardInfoFunc
}

/**
gameType:
3  //重庆时时彩
11 //极速时时彩
14 //英国幸运彩
16 //澳洲幸运5
18 //腾讯分分彩
20 //英国时时彩
 4  //北京赛车
 8  //幸运飞艇
 9  //极速快车
 12 //极速赛车
 13 //ESP赛马
 15 //英国幸运飞艇
 17 //澳洲幸运10
 19 //英国赛车
 2 //广东快乐十分
 = 7 //重庆幸运农场

resultUrl:  http://127.0.0.1:8000/api/lotteryresult?gameindex=?
*/
func NewUscResultServer(GameType int, resultUrl string, aAwardInfoFunc LotteryResult.HandlerAwardInfoFunc) *UscResultServer {
	aWsxResultServer := new(UscResultServer)
	aWsxResultServer.GameIndex = UscBox.GameType2Index(GameType)
	aWsxResultServer.resultUrl = fmt.Sprintf(resultUrl, aWsxResultServer.GameIndex)
	aWsxResultServer.mHttpClient = httpClientKit.GetHttpClient("")
	aWsxResultServer.mHandlerAwardInfoFunc = aAwardInfoFunc
	aWsxResultServer.run()
	return aWsxResultServer
}

func (this *UscResultServer) run() {
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

func (this *UscResultServer) saveAwardInfo() {
	aUscAwardInfo, e := this.getAwardInfo()
	if e != nil {
		return
	}

	if this.CurOpenCodeInfo.NextLotteryStr != aUscAwardInfo.NextLotteryStr {
		aLoAwardInfo := models.LoAwardInfo{}

		aLoAwardInfo.LotteryNum = aUscAwardInfo.LotteryNum
		aLoAwardInfo.LotteryStr = aUscAwardInfo.LotteryStr
		aLoAwardInfo.ResultNums = aUscAwardInfo.ResultNums
		aLoAwardInfo.OriginalResult = stringKit.GetJsonStr(aUscAwardInfo)
		aLoAwardInfo.NextLotteryStr = aUscAwardInfo.NextLotteryStr

		aLoAwardInfo.CreatedAt = aUscAwardInfo.ServerTime

		NextTime := time.Duration(aUscAwardInfo.NextLotteryTime.Sub(aUscAwardInfo.ServerTime)) / time.Second
		aLoAwardInfo.NextTime = fmt.Sprintf("%d", NextTime)
		aLoAwardInfo.CurLotteryTime = aUscAwardInfo.CurLotteryTime
		aLoAwardInfo.NextLotteryTime = aUscAwardInfo.NextLotteryTime
		aLoAwardInfo.CurLotteryTime = aUscAwardInfo.CurLotteryTime

		this.CurOpenCodeInfo = *aUscAwardInfo

		if this.mHandlerAwardInfoFunc != nil {
			go this.mHandlerAwardInfoFunc(aLoAwardInfo)
		}
	}
}

func (this *UscResultServer) getAwardInfo() (*UscBox.AwardInfo, error) {
	aAwardInfo := UscBox.AwardInfo{}

	result, e := this.mHttpClient.GetBytes(this.resultUrl)
	if e != nil {
		ttLog.LogError(e)
	}
	e = json.Unmarshal(result, &aAwardInfo)
	if e != nil {
		ttLog.LogError(e, "strJson:", string(result))
		return nil, e
	}

	return &aAwardInfo, nil
}
