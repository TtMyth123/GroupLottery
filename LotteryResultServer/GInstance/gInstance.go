package GInstance

import (
	"github.com/TtMyth123/GameServer/LotteryResult/UscResultServer/UscBox"
	"github.com/TtMyth123/GameServer/models/mconst"
	"github.com/TtMyth123/LotteryResultServer/GInstance/Lottery28Game"
	"github.com/TtMyth123/LotteryResultServer/GInstance/LotteryUscGame"
	"github.com/TtMyth123/LotteryResultServer/GInstance/lottery"
	"github.com/astaxie/beego"
	"sync"
)

var (
	GUscG     map[int]lottery.ILotteryServer
	b         lottery.BLotteryServer
	lockGUscG sync.Mutex
)

func Init() {
	GUscG = make(map[int]lottery.ILotteryServer)
	if isRun, _ := beego.AppConfig.Bool("Game28::JndIsRun"); isRun {
		GetServer(mconst.GameType_G28_041)
	}
}

func GetServerByIndex(GameIndex int) lottery.ILotteryServer {
	lockGUscG.Lock()
	defer lockGUscG.Unlock()
	if server, ok := GUscG[GameIndex+100]; ok {
		return server
	} else {
		switch GameIndex {
		case UscBox.GameIndex_5_cqssc,
			UscBox.GameIndex_5_jsssc,
			UscBox.GameIndex_5_ygcyc,
			UscBox.GameIndex_5_gzxy5,
			UscBox.GameIndex_5_yxssc,
			UscBox.GameIndex_5_ygssc:
			GUscG[GameIndex+100] = LotteryUscGame.NewUscGameR(5, 5)
		case UscBox.GameIndex_10_bjsc,
			UscBox.GameIndex_10_xyft,
			UscBox.GameIndex_10_jskc,
			UscBox.GameIndex_10_jssc,
			UscBox.GameIndex_10_ESPsm,
			UscBox.GameIndex_10_ygxyft,
			UscBox.GameIndex_10_gzxy10,
			UscBox.GameIndex_10_ygsc:
			GUscG[GameIndex+100] = LotteryUscGame.NewUscGameR(5, 10)
		case UscBox.GameIndex_8_gdkl10f,
			UscBox.GameIndex_8_cqxync:
			GUscG[GameIndex+100] = LotteryUscGame.NewUscGameR(5, 20)
		default:
			return b
		}

		return GUscG[GameIndex+100]
	}
}

func GetServer(GameType int) lottery.ILotteryServer {
	lockGUscG.Lock()
	defer lockGUscG.Unlock()
	if server, ok := GUscG[GameType]; ok {
		return server
	} else {
		switch GameType {
		case UscBox.GameIndex_5_cqssc,
			UscBox.GameIndex_5_jsssc,
			UscBox.GameIndex_5_ygcyc,
			UscBox.GameIndex_5_gzxy5,
			UscBox.GameIndex_5_yxssc,
			UscBox.GameIndex_5_ygssc:
			GUscG[GameType] = LotteryUscGame.NewUscGameR(5, 5)
		case UscBox.GameIndex_10_bjsc,
			UscBox.GameIndex_10_xyft,
			UscBox.GameIndex_10_jskc,
			UscBox.GameIndex_10_jssc,
			UscBox.GameIndex_10_ESPsm,
			UscBox.GameIndex_10_ygxyft,
			UscBox.GameIndex_10_gzxy10,
			UscBox.GameIndex_10_ygsc:
			GUscG[GameType] = LotteryUscGame.NewUscGameR(5, 10)
		case UscBox.GameIndex_8_gdkl10f,
			UscBox.GameIndex_8_cqxync:
			GUscG[GameType] = LotteryUscGame.NewUscGameR(5, 20)
		case mconst.GameType_G28_041:
			GUscG[GameType] = Lottery28Game.NewGameJnd28(5, GameType)
		case mconst.GameType_G28_2000:
			GUscG[GameType] = Lottery28Game.NewGame2Jnd28(5, GameType)
		case mconst.GameType_G28_7710:
			GUscG[GameType] = Lottery28Game.NewGame7710Jnd28(5, GameType)
		case mconst.GameType_G28_7711:
			GUscG[GameType] = Lottery28Game.NewGameKenoJnd28(5, GameType)
		default:
			return b
		}

		return GUscG[GameType]
	}
}
