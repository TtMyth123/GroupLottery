package UscBox

import "github.com/TtMyth123/GameServer/models/mconst"

const FT48Count = 48
const (
	GameIndex_5_cqssc   = 3  //重庆时时彩
	GameIndex_5_jsssc   = 11 //极速时时彩
	GameIndex_5_ygcyc   = 14 //英国幸运彩
	GameIndex_5_gzxy5   = 16 //澳洲幸运5
	GameIndex_5_yxssc   = 18 //腾讯分分彩
	GameIndex_5_ygssc   = 20 //英国时时彩
	GameIndex_10_bjsc   = 4  //北京赛车
	GameIndex_10_xyft   = 8  //幸运飞艇
	GameIndex_10_jskc   = 9  //极速快车
	GameIndex_10_jssc   = 12 //极速赛车
	GameIndex_10_ESPsm  = 13 //ESP赛马
	GameIndex_10_ygxyft = 15 //英国幸运飞艇
	GameIndex_10_gzxy10 = 17 //澳洲幸运10
	GameIndex_10_ygsc   = 19 //英国赛车
	GameIndex_8_gdkl10f = 2  //广东快乐十分
	GameIndex_8_cqxync  = 7  //重庆幸运农场
)

func GameType2Index(gameType int) int {
	switch gameType {
	case mconst.GameType_USC_cqssc:
		return GameIndex_5_cqssc
	case mconst.GameType_USC_jsssc:
		return GameIndex_5_jsssc
	case mconst.GameType_USC_ygcyc:
		return GameIndex_5_ygcyc
	case mconst.GameType_USC_ygssc:
		return GameIndex_5_ygssc
	case mconst.GameType_USC_gzxy5:
		return GameIndex_5_gzxy5
	case mconst.GameType_USC_yxssc:
		return GameIndex_5_yxssc

	case mconst.GameType_USC_bjsc:
		return GameIndex_10_bjsc
	case mconst.GameType_USC_xyft:
		return GameIndex_10_xyft
	case mconst.GameType_USC_jskc:
		return GameIndex_10_jskc
	case mconst.GameType_USC_jssc:
		return GameIndex_10_jssc
	case mconst.GameType_USC_ESPsm:
		return GameIndex_10_ESPsm
	case mconst.GameType_USC_ygxyft:
		return GameIndex_10_ygxyft
	case mconst.GameType_USC_ygsc:
		return GameIndex_10_ygsc
	case mconst.GameType_USC_gzxy10:
		return GameIndex_10_gzxy10

	case mconst.GameType_USC_cqxync:
		return GameIndex_8_cqxync
	case mconst.GameType_USC_gdkl10f:
		return GameIndex_8_gdkl10f
	}
	return 0
}
