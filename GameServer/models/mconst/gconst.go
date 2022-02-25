package mconst

const (
	GameType_Wsx_201 = 201 //"南部彩"
	GameType_Wsx_202 = 202 //"北部彩"
	GameType_Wsx_203 = 203 //"中部彩"

	GameType_G28_2000 = 2000 //"加拿大28"
	GameType_G28_7710 = 7710 //"加拿大28"
	GameType_G28_7711 = 7711 //"加拿大28"

	GameType_G28_041 = 41 //"加拿大28"
	GameType_G28_042 = 42 //"北京28"
	GameType_G28_043 = 43 //"新加坡28"
	GameType_G28_044 = 44 //"香港28"

	GameType_USC_cqssc = 103 //重庆时时彩
	GameType_USC_jsssc = 111 //极速时时彩
	GameType_USC_ygcyc = 114 //英国幸运彩
	GameType_USC_ygssc = 120 //英国时时彩
	GameType_USC_gzxy5 = 116 //澳洲幸运5
	GameType_USC_yxssc = 118 //腾讯分分彩

	GameType_USC_bjsc   = 104 //北京赛车
	GameType_USC_xyft   = 108 //幸运飞艇
	GameType_USC_jskc   = 109 //极速快车
	GameType_USC_jssc   = 112 //极速赛车
	GameType_USC_ESPsm  = 113 //ESP赛马
	GameType_USC_ygxyft = 115 //英国幸运飞艇
	GameType_USC_ygsc   = 119 //英国赛车
	GameType_USC_gzxy10 = 117 //澳洲幸运10

	GameType_USC_cqxync  = 107 //重庆幸运农场
	GameType_USC_gdkl10f = 102 //广东快乐十分

)
const (
	GameType_N_Wsx_201 = "南部彩"
	GameType_N_Wsx_202 = "北部彩"
	GameType_N_Wsx_203 = "中部彩"

	GameType_N_G28_041 = "加拿大28"
	GameType_N_G28_042 = "北京28"
	GameType_N_G28_043 = "新加坡28"
	GameType_N_G28_044 = "香港28"

	GameType_N_USC_cqssc = "重庆时时彩"
	GameType_N_USC_jsssc = "极速时时彩"
	GameType_N_USC_ygcyc = "英国幸运彩"
	GameType_N_USC_ygssc = "英国时时彩"
	GameType_N_USC_gzxy5 = "澳洲幸运5"
	GameType_N_USC_yxssc = "腾讯分分彩"

	GameType_N_USC_bjsc   = "北京赛车"
	GameType_N_USC_xyft   = "幸运飞艇"
	GameType_N_USC_jskc   = "极速快车"
	GameType_N_USC_jssc   = "极速赛车"
	GameType_N_USC_ESPsm  = "ESP赛马"
	GameType_N_USC_ygxyft = "英国幸运飞艇"
	GameType_N_USC_ygsc   = "英国赛车"
	GameType_N_USC_gzxy10 = "澳洲幸运10"

	GameType_N_USC_cqxync  = "重庆幸运农场"
	GameType_N_USC_gdkl10f = "广东快乐十分"
)

var mpGameName map[int]string

func init() {
	mpGameName = make(map[int]string)
	mpGameName[GameType_Wsx_201] = GameType_N_Wsx_201
	mpGameName[GameType_Wsx_202] = GameType_N_Wsx_202
	mpGameName[GameType_Wsx_203] = GameType_N_Wsx_203

	mpGameName[GameType_G28_041] = GameType_N_G28_041
	mpGameName[GameType_G28_042] = GameType_N_G28_042
	mpGameName[GameType_G28_043] = GameType_N_G28_043
	mpGameName[GameType_G28_044] = GameType_N_G28_044

	mpGameName[GameType_USC_cqssc] = GameType_N_USC_cqssc
	mpGameName[GameType_USC_jsssc] = GameType_N_USC_jsssc
	mpGameName[GameType_USC_ygcyc] = GameType_N_USC_ygcyc
	mpGameName[GameType_USC_ygssc] = GameType_N_USC_ygssc
	mpGameName[GameType_USC_gzxy5] = GameType_N_USC_gzxy5
	mpGameName[GameType_USC_yxssc] = GameType_N_USC_yxssc

	mpGameName[GameType_USC_bjsc] = GameType_N_USC_bjsc
	mpGameName[GameType_USC_xyft] = GameType_N_USC_xyft
	mpGameName[GameType_USC_jskc] = GameType_N_USC_jskc
	mpGameName[GameType_USC_jssc] = GameType_N_USC_jssc
	mpGameName[GameType_USC_ESPsm] = GameType_N_USC_ESPsm
	mpGameName[GameType_USC_ygxyft] = GameType_N_USC_ygxyft
	mpGameName[GameType_USC_ygsc] = GameType_N_USC_ygsc
	mpGameName[GameType_USC_gzxy10] = GameType_N_USC_gzxy10

	mpGameName[GameType_USC_cqxync] = GameType_N_USC_cqxync
	mpGameName[GameType_USC_gdkl10f] = GameType_N_USC_gdkl10f
}

func GetGameName(gameType int) string {
	return mpGameName[gameType]
}
