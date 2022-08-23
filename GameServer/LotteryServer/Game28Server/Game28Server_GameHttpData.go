package Game28Server

import (
	"github.com/TtMyth123/GameServer/LotteryResult/Game28ResultKit"
	"github.com/TtMyth123/GameServer/OtherServer/httpGameServer"
	"github.com/TtMyth123/GameServer/OtherServer/httpGameServer/httpBox"
	"github.com/TtMyth123/GameServer/OtherServer/httpGameServer/httpConst"
	"github.com/TtMyth123/GameServer/models/mconst"
	"time"
)

type GameHttpData struct {
	Command string
	Data    interface{}
}

func (this *Game28Server) runGameHttpChan() {
	for {
		select {
		case c := <-this.gameHttpChan:
			this.dispGameHttpChan(c)
		}
	}
}

func (this *Game28Server) dispGameHttpChan(aGameHttpData GameHttpData) {
	switch aGameHttpData.Command {
	case httpConst.NewAwardInfo:
		{
			aData := aGameHttpData.Data.(Game28ResultKit.Game28AwardInfo)

			aAwardResultBox := httpBox.AwardResultBox{}
			aAwardResultBox.AwardInfo = aData
			aAwardResultBox.Game = mconst.GetGameName(this.GameType)
			httpGameServer.AwardResult(this.GameType, aAwardResultBox)

			time.Sleep(time.Second * 6)
			aSetGameTimerBox := httpBox.SetGameTimerBox{
				Timer:     aData.Countdown - 6,
				StopBet:   this.StopBetTime,
				NextIssue: aData.NextLotteryStr,
				Issue:     aData.LotteryStr}
			aSetGameTimerBox.Game = mconst.GetGameName(this.GameType)
			//httpGameServer.SetGameTimer(this.GameType,aSetGameTimerBox)
			aGameHttpData := GameHttpData{Command: httpConst.SetGameTimerBox, Data: aSetGameTimerBox}
			this.gameHttpChan <- aGameHttpData
		}
	case httpConst.SetGameTimerBox:
		{
			aSetGameTimerBox := aGameHttpData.Data.(httpBox.SetGameTimerBox)
			httpGameServer.SetGameTimer(this.GameType, aSetGameTimerBox)
		}
	case httpConst.PlayerBet:
		{
			aData := aGameHttpData.Data.(httpBox.PlayerBetBox)
			aData.Game = mconst.GetGameName(this.GameType)
			httpGameServer.PlayerBet(this.GameType, aData)
		}
	case httpConst.StopBetState:
		{
			aData := aGameHttpData.Data.(httpBox.StopBetStateBox)
			aData.Game = mconst.GetGameName(this.GameType)
			httpGameServer.StopBetState(this.GameType, aData)
		}
	}
}
