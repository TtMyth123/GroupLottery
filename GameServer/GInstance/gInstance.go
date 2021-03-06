package GInstance

import (
	"github.com/astaxie/beego"
	"ttmyth123/GroupLottery/GameServer/LotteryResult"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/Game28ResultServer"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/UscResultServer"
	"ttmyth123/GroupLottery/GameServer/LotteryServer"
	"ttmyth123/GroupLottery/GameServer/LotteryServer/Game28Server"
	"ttmyth123/GroupLottery/GameServer/LotteryServer/UscServer"
	"ttmyth123/GroupLottery/GameServer/LotteryServer/WsxServer"
	"ttmyth123/GroupLottery/GameServer/models/mconst"
	"ttmyth123/GroupLottery/UserInfoRpc/UserRpcClient"
)

var (
	mUserRpcClient *UserRpcClient.RpcClient

	mpLotteryServer map[int]LotteryServer.ILotteryServer
)

func Init() {
	addr := beego.AppConfig.String("UserRpcClientAddr")
	mUserRpcClient = UserRpcClient.NewRpcClient(addr)
	initLotteryServer()
}

func initLotteryServer() {
	mpLotteryServer = make(map[int]LotteryServer.ILotteryServer)
	if isRun, _ := beego.AppConfig.Bool("Wsx::NbcIsRun"); isRun {
		WsxServer := WsxServer.NewWsxNbcServer(mconst.GameType_Wsx_201, mUserRpcClient)
		strUrl := beego.AppConfig.String("Wsx::ResultHttpUrl1")
		LotteryResult.NewWsxResultServer1(strUrl, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_Wsx_201] = WsxServer
	}

	if isRun, _ := beego.AppConfig.Bool("Wsx::BbcIsRun"); isRun {
		WsxServer := WsxServer.NewWsxBbcServer(mconst.GameType_Wsx_202, mUserRpcClient)

		aResultHttpUrl2 := beego.AppConfig.String("Wsx::ResultHttpUrl2")
		LotteryResult.NewWsxBbcResultServer(aResultHttpUrl2, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_Wsx_202] = WsxServer
	}

	if isRun, _ := beego.AppConfig.Bool("Wsx::ZbcIsRun"); isRun {
		WsxServer := WsxServer.NewWsxZbcServer(mconst.GameType_Wsx_203, mUserRpcClient)
		aResultHttpUrl3 := beego.AppConfig.String("Wsx::ResultHttpUrl3")
		LotteryResult.NewWsxZbcResultServer(aResultHttpUrl3, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_Wsx_203] = WsxServer
	}

	if isRun, _ := beego.AppConfig.Bool("Game28::JndIsRun"); isRun {
		Game28ServerJnd := Game28Server.NewGame28Server(mconst.GameType_G28_041, 10, mUserRpcClient)

		url := beego.AppConfig.String("Game28::ResultHttpUrlJnd")
		Game28ResultServer.NewJnd28ResultServer(url, Game28ServerJnd.NewAwardInfo)
		mpLotteryServer[mconst.GameType_G28_041] = Game28ServerJnd
	}

	if isRun, _ := beego.AppConfig.Bool("Game28::BjIsRun"); isRun {
		Game28ServerBj := Game28Server.NewGame28Server(mconst.GameType_G28_042, 40, mUserRpcClient)

		url := beego.AppConfig.String("Game28::ResultHttpUrlBj")
		Game28ResultServer.NewBj28ResultServer(url, Game28ServerBj.NewAwardInfo)
		mpLotteryServer[mconst.GameType_G28_042] = Game28ServerBj
	}

	if isRun, _ := beego.AppConfig.Bool("Game28::XjpIsRun"); isRun {
		Game28ServerXjp := Game28Server.NewGame28Server(mconst.GameType_G28_043, 10, mUserRpcClient)

		url := beego.AppConfig.String("Game28::ResultHttpUrlXjp")
		Game28ResultServer.NewXjp28ResultServer(url, Game28ServerXjp.NewAwardInfo)
		mpLotteryServer[mconst.GameType_G28_043] = Game28ServerXjp
	}

	if isRun, _ := beego.AppConfig.Bool("Game28::XgIsRun"); isRun {
		Game28ServerXg := Game28Server.NewGame28Server(mconst.GameType_G28_044, 10, mUserRpcClient)

		url := beego.AppConfig.String("Game28::ResultHttpUrlXg")
		Game28ResultServer.NewXg28ResultServer(url, Game28ServerXg.NewAwardInfo)
		mpLotteryServer[mconst.GameType_G28_044] = Game28ServerXg
	}

	//#???????????????
	if isRun, _ := beego.AppConfig.Bool("USC::cqssc_03_IsRun"); isRun {
		WsxServer := UscServer.NewUsc5Server(mconst.GameType_USC_cqssc, mUserRpcClient, 10)
		aResultHttpUrl := beego.AppConfig.String("USC::ResultHttpUrl")
		UscResultServer.NewUscResultServer(mconst.GameType_USC_cqssc, aResultHttpUrl, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_USC_cqssc] = WsxServer
	}
	//#???????????????
	if isRun, _ := beego.AppConfig.Bool("USC::jsssc_11_IsRun"); isRun {
		WsxServer := UscServer.NewUsc5Server(mconst.GameType_USC_jsssc, mUserRpcClient, 10)
		aResultHttpUrl := beego.AppConfig.String("USC::ResultHttpUrl")
		UscResultServer.NewUscResultServer(mconst.GameType_USC_jsssc, aResultHttpUrl, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_USC_jsssc] = WsxServer
	}
	//#???????????????
	if isRun, _ := beego.AppConfig.Bool("USC::ygcyc_14_IsRun"); isRun {
		WsxServer := UscServer.NewUsc5Server(mconst.GameType_USC_ygcyc, mUserRpcClient, 10)
		aResultHttpUrl := beego.AppConfig.String("USC::ResultHttpUrl")
		UscResultServer.NewUscResultServer(mconst.GameType_USC_ygcyc, aResultHttpUrl, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_USC_ygcyc] = WsxServer
	}
	//#????????????5
	if isRun, _ := beego.AppConfig.Bool("USC::gzxy5_16_IsRun"); isRun {
		WsxServer := UscServer.NewUsc5Server(mconst.GameType_USC_gzxy5, mUserRpcClient, 10)
		aResultHttpUrl := beego.AppConfig.String("USC::ResultHttpUrl")
		UscResultServer.NewUscResultServer(mconst.GameType_USC_gzxy5, aResultHttpUrl, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_USC_gzxy5] = WsxServer
	}
	//#???????????????
	if isRun, _ := beego.AppConfig.Bool("USC::yxssc_18_IsRun"); isRun {
		WsxServer := UscServer.NewUsc5Server(mconst.GameType_USC_yxssc, mUserRpcClient, 10)
		aResultHttpUrl := beego.AppConfig.String("USC::ResultHttpUrl")
		UscResultServer.NewUscResultServer(mconst.GameType_USC_yxssc, aResultHttpUrl, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_USC_yxssc] = WsxServer
	}
	//#???????????????
	if isRun, _ := beego.AppConfig.Bool("USC::ygssc_20_IsRun"); isRun {
		WsxServer := UscServer.NewUsc5Server(mconst.GameType_USC_ygssc, mUserRpcClient, 10)
		aResultHttpUrl := beego.AppConfig.String("USC::ResultHttpUrl")
		UscResultServer.NewUscResultServer(mconst.GameType_USC_ygssc, aResultHttpUrl, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_USC_ygssc] = WsxServer
	}

	//#????????????
	if isRun, _ := beego.AppConfig.Bool("USC::bjsc_04_IsRun"); isRun {
		WsxServer := UscServer.NewUsc10Server(mconst.GameType_USC_bjsc, mUserRpcClient, 10)
		aResultHttpUrl := beego.AppConfig.String("USC::ResultHttpUrl")
		UscResultServer.NewUscResultServer(mconst.GameType_USC_bjsc, aResultHttpUrl, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_USC_bjsc] = WsxServer
	}
	//#????????????
	if isRun, _ := beego.AppConfig.Bool("USC::xyft_08_IsRun"); isRun {
		WsxServer := UscServer.NewUsc10Server(mconst.GameType_USC_xyft, mUserRpcClient, 10)
		aResultHttpUrl := beego.AppConfig.String("USC::ResultHttpUrl")
		UscResultServer.NewUscResultServer(mconst.GameType_USC_xyft, aResultHttpUrl, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_USC_xyft] = WsxServer
	}
	//#????????????
	if isRun, _ := beego.AppConfig.Bool("USC::jskc_09_IsRun"); isRun {
		WsxServer := UscServer.NewUsc10Server(mconst.GameType_USC_jskc, mUserRpcClient, 10)
		aResultHttpUrl := beego.AppConfig.String("USC::ResultHttpUrl")
		UscResultServer.NewUscResultServer(mconst.GameType_USC_jskc, aResultHttpUrl, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_USC_jskc] = WsxServer
	}
	//#????????????
	if isRun, _ := beego.AppConfig.Bool("USC::jssc_12_IsRun"); isRun {
		WsxServer := UscServer.NewUsc10Server(mconst.GameType_USC_jssc, mUserRpcClient, 10)
		aResultHttpUrl := beego.AppConfig.String("USC::ResultHttpUrl")
		UscResultServer.NewUscResultServer(mconst.GameType_USC_jssc, aResultHttpUrl, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_USC_jssc] = WsxServer
	}
	//ESP??????
	if isRun, _ := beego.AppConfig.Bool("USC::ESPsm_13_IsRun"); isRun {
		WsxServer := UscServer.NewUsc10Server(mconst.GameType_USC_ESPsm, mUserRpcClient, 10)
		aResultHttpUrl := beego.AppConfig.String("USC::ResultHttpUrl")
		UscResultServer.NewUscResultServer(mconst.GameType_USC_ESPsm, aResultHttpUrl, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_USC_ESPsm] = WsxServer
	}
	//#??????????????????
	if isRun, _ := beego.AppConfig.Bool("USC::ygxyft_15_IsRun"); isRun {
		WsxServer := UscServer.NewUsc10Server(mconst.GameType_USC_ygxyft, mUserRpcClient, 10)
		aResultHttpUrl := beego.AppConfig.String("USC::ResultHttpUrl")
		UscResultServer.NewUscResultServer(mconst.GameType_USC_ygxyft, aResultHttpUrl, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_USC_ygxyft] = WsxServer
	}
	//#????????????10
	if isRun, _ := beego.AppConfig.Bool("USC::gzxy10_17_IsRun"); isRun {
		WsxServer := UscServer.NewUsc10Server(mconst.GameType_USC_gzxy10, mUserRpcClient, 10)
		aResultHttpUrl := beego.AppConfig.String("USC::ResultHttpUrl")
		UscResultServer.NewUscResultServer(mconst.GameType_USC_gzxy10, aResultHttpUrl, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_USC_gzxy10] = WsxServer
	}
	//#????????????
	if isRun, _ := beego.AppConfig.Bool("USC::ygsc_19_IsRun"); isRun {
		WsxServer := UscServer.NewUsc10Server(mconst.GameType_USC_ygsc, mUserRpcClient, 10)
		aResultHttpUrl := beego.AppConfig.String("USC::ResultHttpUrl")
		UscResultServer.NewUscResultServer(mconst.GameType_USC_ygsc, aResultHttpUrl, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_USC_ygsc] = WsxServer
	}

	//#??????????????????
	if isRun, _ := beego.AppConfig.Bool("USC::gdkl10f_02_IsRun"); isRun {
		WsxServer := UscServer.NewUsc8for20Server(mconst.GameType_USC_gdkl10f, mUserRpcClient, 10)
		aResultHttpUrl := beego.AppConfig.String("USC::ResultHttpUrl")
		UscResultServer.NewUscResultServer(mconst.GameType_USC_gdkl10f, aResultHttpUrl, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_USC_gdkl10f] = WsxServer
	}
	//#??????????????????
	if isRun, _ := beego.AppConfig.Bool("USC::cqxync_07_IsRun"); isRun {
		WsxServer := UscServer.NewUsc8for20Server(mconst.GameType_USC_cqxync, mUserRpcClient, 10)
		aResultHttpUrl := beego.AppConfig.String("USC::ResultHttpUrl")
		UscResultServer.NewUscResultServer(mconst.GameType_USC_cqxync, aResultHttpUrl, WsxServer.NewAwardInfo)
		mpLotteryServer[mconst.GameType_USC_cqxync] = WsxServer
	}
}

func GetUserRpcClient() *UserRpcClient.RpcClient {
	return mUserRpcClient
}

func GetLotteryServer(gameType int, Area string) LotteryServer.ILotteryServer {
	return mpLotteryServer[gameType]
}
