package WsxServer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"sync"
	"time"
	"ttmyth123/GroupLottery/GameServer/CacheData"
	"ttmyth123/GroupLottery/GameServer/GConfig"
	"ttmyth123/GroupLottery/GameServer/GInstance/GTtHint"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/WsxBbcResultKit"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/WsxBox"
	"ttmyth123/GroupLottery/GameServer/LotteryServer"
	"ttmyth123/GroupLottery/GameServer/LotteryServer/LotteryBox"
	"ttmyth123/GroupLottery/GameServer/controllers/base/TtError"
	"ttmyth123/GroupLottery/GameServer/models"
	"ttmyth123/GroupLottery/GameServer/models/mconst"
	"ttmyth123/GroupLottery/UserInfoRpc/GData/gBox"
	"ttmyth123/GroupLottery/UserInfoRpc/UserRpcClient"
	userModels "ttmyth123/GroupLottery/UserInfoRpc/models"
	userConst "ttmyth123/GroupLottery/UserInfoRpc/models/mconst"
	"ttmyth123/kit"
	"ttmyth123/kit/sqlKit"
	"ttmyth123/kit/strconvEx"
	"ttmyth123/kit/stringKit"
	"ttmyth123/kit/timeKit"
	"ttmyth123/kit/ttLog"
)

type WsxBbcServer struct {
	LotteryServer.BaseLotteryServer
	OddsLock sync.RWMutex
	BaseWsxServer
	mpOddsInfo map[int]models.LoOddsInfo
}

func NewWsxBbcServer(gameType int, UserRpcClient *UserRpcClient.RpcClient) *WsxBbcServer {
	aWsxServer := new(WsxBbcServer)
	aWsxServer.TimeLag = GConfig.GetGConfig().TimeLag
	aWsxServer.GameType = gameType
	aWsxServer.ReLoadOddsInfo()
	aWsxServer.mUserRpcClient = UserRpcClient
	return aWsxServer
}

func (this *WsxBbcServer) ReLoadOddsInfo() {
	this.OddsLock.Lock()
	defer this.OddsLock.Unlock()
	mpOddsInfo := make(map[int]models.LoOddsInfo)
	mpOdds := make(map[int]float64)
	arrOdds, e := models.GetAllLoOddsInfo(this.GameType)
	if e != nil {
		ttLog.LogError(e)
		return
	}

	for _, v := range arrOdds {
		mpOddsInfo[v.OddsType] = v
		mpOdds[v.OddsType] = v.Odds
	}
	this.mpOddsInfo = mpOddsInfo
	this.mpOdds = mpOdds
}

func (this *WsxBbcServer) GetRoomInfo() map[string]interface{} {
	mpInfo := make(map[string]interface{})

	mpInfo["1"] = this.StopBetTime
	mpInfo["2"] = this.mpOdds
	if this.TimeLag == 0 {
		mpInfo["3"] = this.curAwardInfo.NextLotteryTime.Format(timeKit.DateTimeLayout)
	} else {
		mpInfo["3"] = this.curAwardInfo.NextLotteryTime.Add(this.TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
	}

	mpInfo["4"] = this.curAwardInfo.LotteryStr
	mpInfo["5"] = this.curAwardInfo.NextLotteryStr
	t := time.Now()
	iSecond := int64(this.curAwardInfo.NextLotteryTime.Sub(t) / time.Second)
	if iSecond <= 0 {
		iSecond = 0
	}
	mpInfo["6"] = iSecond

	return mpInfo
}
func (this *WsxBbcServer) RandNewAwardInfo() error {
	this.RandNewAwardInfoLock.Lock()
	defer this.RandNewAwardInfoLock.Unlock()

	aOpenCodeInfo := WsxBox.OpenCodeInfo{}

	//85435
	aOpenCodeInfo.Jackpots = fmt.Sprintf("%05d", timeKit.GlobaRand.Intn(99999))
	//01738
	aOpenCodeInfo.FirstNum = fmt.Sprintf("%05d", timeKit.GlobaRand.Intn(99999))
	//41658
	aOpenCodeInfo.SecondNum = []string{
		fmt.Sprintf("%05d", timeKit.GlobaRand.Intn(99999)),
		fmt.Sprintf("%05d", timeKit.GlobaRand.Intn(99999)),
	}
	//90148
	aOpenCodeInfo.ThirdNum = []string{
		fmt.Sprintf("%05d", timeKit.GlobaRand.Intn(99999)),
		fmt.Sprintf("%05d", timeKit.GlobaRand.Intn(99999)),
		fmt.Sprintf("%05d", timeKit.GlobaRand.Intn(99999)),
		fmt.Sprintf("%05d", timeKit.GlobaRand.Intn(99999)),
		fmt.Sprintf("%05d", timeKit.GlobaRand.Intn(99999)),
		fmt.Sprintf("%05d", timeKit.GlobaRand.Intn(99999)),
	}

	//4413
	aOpenCodeInfo.ForthNum = []string{
		fmt.Sprintf("%04d", timeKit.GlobaRand.Intn(9999)),
		fmt.Sprintf("%04d", timeKit.GlobaRand.Intn(9999)),
		fmt.Sprintf("%04d", timeKit.GlobaRand.Intn(9999)),
		fmt.Sprintf("%04d", timeKit.GlobaRand.Intn(9999)),
	}
	//3328
	aOpenCodeInfo.FifthNum = []string{
		fmt.Sprintf("%04d", timeKit.GlobaRand.Intn(9999)),
		fmt.Sprintf("%04d", timeKit.GlobaRand.Intn(9999)),
		fmt.Sprintf("%04d", timeKit.GlobaRand.Intn(9999)),
		fmt.Sprintf("%04d", timeKit.GlobaRand.Intn(9999)),
		fmt.Sprintf("%04d", timeKit.GlobaRand.Intn(9999)),
		fmt.Sprintf("%04d", timeKit.GlobaRand.Intn(9999)),
	}
	//057
	aOpenCodeInfo.SixthNum = []string{
		fmt.Sprintf("%03d", timeKit.GlobaRand.Intn(999)),
		fmt.Sprintf("%03d", timeKit.GlobaRand.Intn(999)),
		fmt.Sprintf("%03d", timeKit.GlobaRand.Intn(999)),
	}
	//55
	aOpenCodeInfo.SeventhNum = []string{
		fmt.Sprintf("%02d", timeKit.GlobaRand.Intn(99)),
		fmt.Sprintf("%02d", timeKit.GlobaRand.Intn(99)),
		fmt.Sprintf("%02d", timeKit.GlobaRand.Intn(99)),
		fmt.Sprintf("%02d", timeKit.GlobaRand.Intn(99)),
	}

	aResultNums := make(map[string]interface{})
	aResultNums["jackpots"] = aOpenCodeInfo.Jackpots
	aResultNums["firstNum"] = aOpenCodeInfo.FirstNum
	aResultNums["secondNum"] = aOpenCodeInfo.SecondNum
	aResultNums["thirdNum"] = aOpenCodeInfo.ThirdNum
	aResultNums["forthNum"] = aOpenCodeInfo.ForthNum
	aResultNums["fifthNum"] = aOpenCodeInfo.FifthNum
	aResultNums["sixthNum"] = aOpenCodeInfo.SixthNum
	aResultNums["seventhNum"] = aOpenCodeInfo.SeventhNum
	//aResultNums["eighthNum"] = aOpenCodeInfo.EighthNum

	aAwardInfo := WsxBox.WsxResultData{}

	Expect := this.curAwardInfo.NextLotteryStr
	if Expect == "" {
		return errors.New("请等待...")
	}
	NextIssue := strconvEx.StrTry2Int64(Expect, 0)
	if NextIssue == 0 {
		return errors.New("请等待...")
	}

	aAwardInfo.Expect = Expect
	aAwardInfo.NextIssue = fmt.Sprintf("%d", NextIssue)
	aAwardInfo.NextTime = this.curAwardInfo.NextLotteryTime.Add(time.Hour * 24 * 2).Format("2006-01-02 15:04:05")
	aAwardInfo.Opentime = this.curAwardInfo.NextLotteryTime.Format("2006-01-02 15:04:05")
	aAwardInfo.Opencode = stringKit.GetJsonStr(aOpenCodeInfo)

	aLoAwardInfo := models.LoAwardInfo{}
	aLoAwardInfo.LotteryNum = strconvEx.StrTry2Int64(aAwardInfo.Expect, 0)
	aLoAwardInfo.LotteryStr = aAwardInfo.Expect
	aLoAwardInfo.ResultNums = stringKit.GetJsonStr(aResultNums)
	aLoAwardInfo.OriginalResult = stringKit.GetJsonStr(aAwardInfo)
	aLoAwardInfo.NextLotteryStr = aAwardInfo.NextIssue

	aLoAwardInfo.CurLotteryTime = time.Now()
	aLoAwardInfo.NextLotteryTime = time.Now().AddDate(0, 0, 2)

	e := this.DelAfData(aLoAwardInfo.LotteryStr)
	if e != nil {
		return e
	}

	this.NewAwardInfo(aLoAwardInfo)
	return nil
}
func (this *WsxBbcServer) NewAwardInfo(newLoAwardInfo models.LoAwardInfo) {
	this.BaseWsxServer.PreNewAwardInfo(newLoAwardInfo)
	o := orm.NewOrm()
	//ttLog.LogDebug(stringKit.GetJsonStr(newLoAwardInfo))
	//this.AddAwardInfo(&newLoAwardInfo)
	//this.curAwardInfo = newLoAwardInfo

	mpBetInfo := this.getBetInfos(newLoAwardInfo.LotteryStr)
	aOpenCodeInfo := WsxBox.OpenCodeInfo{}
	e := json.Unmarshal([]byte(newLoAwardInfo.ResultNums), &aOpenCodeInfo)
	if e != nil {
		ttLog.LogError(e, newLoAwardInfo.ResultNums)
		return
	}

	curT := time.Now()
	mpUserWin := make(map[int]float64)
	arrBetGroupInfo := this.getBetGroupInfos(newLoAwardInfo.LotteryStr)
	mpUserName := make(map[int]userModels.TtGameUser)
	mpUserXM := make(map[int]float64)
	mpSum2Rebate := make(map[int]float64)
	for _, groupBet := range arrBetGroupInfo {

		allWin := 0.0
		for _, bet := range mpBetInfo[groupBet.Id] {
			if _, ok := mpUserName[bet.UserId]; !ok {
				aU, _ := this.mUserRpcClient.GetUser(bet.UserId)
				mpUserName[bet.UserId] = aU
			}

			win := WsxBbcResultKit.ComputeLoseWin(*bet, aOpenCodeInfo, this.mpOddsInfo)

			bet.Status = mconst.Bet_Status_2
			bet.ResultNums = newLoAwardInfo.ResultNums
			bet.Win = win
			allWin += win

			if float64(bet.BetM) != win {
				rebateUser := mpUserName[bet.UserId]
				//betRebate := float64(bet.BetM) * BetRebateRatio
				mpUserXM[bet.UserId] += float64(bet.BetM)
				//mpSum2Rebate[bet.UserId] += betRebate

				//Des := fmt.Sprintf(`%s[%s][%s],投注：%s %d元.返利比：%g`, mconst.GetGameName(this.GameType), newLoAwardInfo.LotteryStr, mpUserName[bet.UserId], bet.OddsName, bet.BetM, BetRebateRatio)
				//Des2 := GTtHint.GetTtHint().GetHint(`%s[%s][%s],投注：%s %d元.返利比：%g`)
				//DesMp := GTtHint.GetTtHint().GetMpString(mconst.GetGameName(this.GameType), newLoAwardInfo.LotteryStr, mpUserName[bet.UserId], bet.OddsName, bet.BetM, BetRebateRatio)

				//rebateInfo := gBox.AddRebateInfo{UserId: bet.UserId, OldGold: float64(bet.BetM),
				//	Ratio: BetRebateRatio, Rebate: betRebate, RefId: bet.Id, T: userConst.Rebate_01_Guess,
				//	Des: Des, Des2: Des2, DesMp: DesMp}

				rebateInfo := gBox.AddRebateInfo{UserId: rebateUser.Pid, UserSid: rebateUser.Id, GameType: this.GameType,
					LotteryStr: newLoAwardInfo.LotteryStr, BetUserName: rebateUser.UserName, Level: 1,
					OddsName: bet.OddsName, BetM: bet.BetM, T: userConst.Rebate_01_Guess}
				this.mUserRpcClient.AddRebate(rebateInfo)
			}

			CacheData.AddCurPeriodUserWin(curT, this.GameType, bet.UserId, bet.OddsType, bet.BetM, newLoAwardInfo.LotteryStr, bet.Win)
		}

		groupBet.Win = allWin
		groupBet.Status = mconst.Bet_Status_2
		mpUserWin[groupBet.UserId] += allWin
		groupBet.ResultNums = newLoAwardInfo.ResultNums
	}

	for _, groupBet := range arrBetGroupInfo {
		for _, bet := range mpBetInfo[groupBet.Id] {
			bet.Update(o, "Status", "ResultNums", "Win")
		}
		groupBet.Update(o, "Status", "Win", "ResultNums")
		if groupBet.Win == 0 {
			continue
		}
		goldInfo := gBox.AddGoldInfo{GroupId: groupBet.GroupUserId, UserId: groupBet.UserId, Gold: groupBet.Win,
			T:    userConst.Account_02_Win,
			Des:  fmt.Sprintf("%s第%s期，[%s]投注%d元, 赢得%g", mconst.GetGameName(this.GameType), newLoAwardInfo.LotteryStr, groupBet.BetSn, groupBet.BetM, groupBet.Win),
			Des2: GTtHint.GetTtHint().GetHint("%s第%s期，[%s]投注%d元, 赢得%g"), DesMp: GTtHint.GetTtHint().GetMpString(mconst.GetGameName(this.GameType), newLoAwardInfo.LotteryStr, groupBet.BetSn, groupBet.BetM, groupBet.Win)}

		_, e := this.mUserRpcClient.AddGold(goldInfo)
		if e != nil {
			ttLog.LogError(e)
		}
	}

	infos := make([]gBox.UpdateDataInfo, 2)
	infos[0].FieldName = "SumXmBet"
	infos[0].Type = 1

	infos[1].FieldName = "Sum2Rebate"
	infos[1].Type = 1
	for userId, xm := range mpUserXM {
		infos[0].Value = xm
		infos[1].Value = mpSum2Rebate[userId]
		this.mUserRpcClient.UpdateUserInfo(userId, infos)
	}
}

func (this *WsxBbcServer) Bet(betInfo LotteryBox.BetInfo) (map[string]interface{}, error) {
	this.OddsLock.RLock()
	defer this.OddsLock.RUnlock()
	mpInfo := make(map[string]interface{})
	e := this.BaseWsxServer.PreBet(betInfo)

	if e != nil {
		return mpInfo, e
	}
	if this.PreIssue != "" {
		curTime := time.Now()
		subMinute := curTime.Sub(this.PreOpenTime) / time.Minute
		if subMinute <= 25 {
			return mpInfo, TtError.New("正在开奖,暂时无法下注.")
		}

		subMinute = this.curAwardInfo.NextLotteryTime.Sub(curTime) / time.Minute
		if subMinute <= 2 {
			return mpInfo, TtError.New("正在开奖,暂时无法下注.")
		}
	}

	mapOddTypeBet := make(map[int]int)
	mapOddTypeName := make(map[int]string)
	o := orm.NewOrm()
	allBetStr := ""
	allBetM := 0
	arrBetInfo := make([]models.LoBetInfo, 0)
	for _, v := range betInfo.BetData {
		mapOddTypeName[v.OddsType] = v.OddsName
		mapOddTypeBet[v.OddsType] += v.M
		if this.mpOddsInfo[v.OddsType].OneUserMinBet > v.M {
			return mpInfo, TtError.New("[%s]单注金额少于%d", v.OddsName, this.mpOddsInfo[v.OddsType].OneUserMinBet)
		}
		allBetM += v.M
		OddsName := v.OddsName
		if OddsName == "" {
			OddsName = this.mpOddsInfo[v.OddsType].OddsDes
		}
		allBetStr += fmt.Sprintf(";%s:$%d", OddsName, v.M)

		aBetInfo := models.LoBetInfo{
			BetSn:       kit.GetGuid(),
			BetStr:      fmt.Sprintf("%s,$%d", OddsName, v.M),
			Status:      mconst.Bet_Status_1,
			Period:      betInfo.LotteryNum,
			StrPeriod:   betInfo.StrLotteryNum,
			Odds:        this.mpOddsInfo[v.OddsType].Odds,
			OddsType:    v.OddsType,
			BigOddsType: this.mpOddsInfo[v.OddsType].BigType,
			Nums:        v.Nums,
			BetM:        v.M,
			Win:         0,
			GameType:    this.GameType,
			UserId:      betInfo.UserId,
			OddsName:    OddsName,
		}

		arrBetInfo = append(arrBetInfo, aBetInfo)
	}
	if len(allBetStr) > 4000 {
		return mpInfo, TtError.New("一次提交项目太多,无法投注.")
	}

	for oddsType, typeBet := range mapOddTypeBet {
		betUserGamePeriodOddsType, betGamePeriodOddsType := CacheData.GetCacheSumBet(betInfo.StrLotteryNum, this.GameType, oddsType, betInfo.UserId)

		if this.mpOddsInfo[oddsType].AllUserMaxBet < betGamePeriodOddsType+typeBet {
			return mpInfo, TtError.New("[%s]超出玩法投注限额", mapOddTypeName[oddsType])
			//return mpInfo, errors.New(GTtHint.GetTtHint().GetHint("[%s]超出玩法投注限额"))
		}

		if this.mpOddsInfo[oddsType].OneUserMaxBet < betUserGamePeriodOddsType+typeBet {
			return mpInfo, TtError.New("[%s]超出用户投注限额", mapOddTypeName[oddsType])
			//return mpInfo, errors.New(GTtHint.GetTtHint().GetHint("超出用户投注限额"))
		}
	}

	aLoBetGroupInfo := &models.LoBetGroupInfo{
		BetSn:     kit.GetGuid(),
		BetStr:    allBetStr,
		Status:    mconst.Bet_Status_1,
		Period:    betInfo.LotteryNum,
		StrPeriod: betInfo.StrLotteryNum,
		BetM:      allBetM,
		Win:       0,
		GameType:  this.GameType,
		UserId:    betInfo.UserId,
	}
	if len(allBetStr) > 0 {
		allBetStr = allBetStr[1:]
	}

	goldInfo := gBox.AddGoldInfo{GroupId: betInfo.GroupId, UserId: betInfo.UserId, T: userConst.Account_01_Guess, Gold: float64(allBetM),
		Des:  fmt.Sprintf("第%s期 投注:%s 共花费:%d", betInfo.StrLotteryNum, allBetStr, allBetM),
		Des2: GTtHint.GetTtHint().GetHint("第%s期 投注:%s 共花费:%d"), DesMp: GTtHint.GetTtHint().GetMpString(betInfo.StrLotteryNum, allBetStr, allBetM)}

	aUser, e := this.mUserRpcClient.AddGold(goldInfo)
	if e != nil {
		return mpInfo, e
	}

	id, _ := aLoBetGroupInfo.Add(o)
	for i := 0; i < len(arrBetInfo); i++ {
		arrBetInfo[i].GroupBetSn = aLoBetGroupInfo.BetSn
		arrBetInfo[i].GroupId = id
	}

	e = models.InsertMultiBetInfo(o, arrBetInfo)
	mpInfo["Money"] = aUser.Gold
	return mpInfo, e
}

type WsxAwardInfo struct {
	LotteryStr     string //期号
	StrLotteryTime string //当前开期时间

	NextLotteryStr     string
	StrNextLotteryTime string //下一期开期时间
	StrCurTime         string //当前时间
	NextTime           string //倒计时间
	WsxBox.OpenCodeInfo
}

func (this *WsxBbcServer) GetHistoryResultList(PageIndex, PageSize, LastId int) (interface{}, error) {
	sqlWhereLastId := ""
	mpData := make(map[string]interface{})
	PageTotal := 0
	arrData := make([]models.LoAwardInfo, 0)
	o := orm.NewOrm()
	sqlArgs := make([]interface{}, 0)
	if LastId == 0 {
		sql := fmt.Sprintf(`select max(id) from %s`, mconst.TableName_LoAwardInfo)
		o.Raw(sql).QueryRow(&LastId)
	} else {
		sqlWhereLastId = " a.id<? and "
		sqlArgs = append(sqlArgs, LastId)
	}

	sqlCount := fmt.Sprintf(`select count(1) c from %s a where %s a.game_type=? `,
		mconst.TableName_LoAwardInfo, sqlWhereLastId)
	sqlArgs = append(sqlArgs, this.GameType)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&PageTotal)
	if e != nil {
		return mpData, e
	}
	offset, _ := sqlKit.GetOffset(PageTotal, PageSize, PageIndex)

	sql := fmt.Sprintf(`select a.* from %s a where %s a.game_type=? order by a.created_at DESC  LIMIT ?,? `,
		mconst.TableName_LoAwardInfo, sqlWhereLastId)

	sqlArgs = append(sqlArgs, offset, PageSize)

	_, e = o.Raw(sql, sqlArgs).QueryRows(&arrData)

	arrWsxAwardInfo := make([]WsxAwardInfo, len(arrData))
	for i := 0; i < len(arrData); i++ {
		arrWsxAwardInfo[i].LotteryStr = arrData[i].LotteryStr

		if this.TimeLag == 0 {
			arrWsxAwardInfo[i].StrLotteryTime = arrData[i].CurLotteryTime.Format(timeKit.DateTimeLayout)
		} else {
			arrWsxAwardInfo[i].StrLotteryTime = arrData[i].CurLotteryTime.Add(this.TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
		}

		aOpenCodeInfo := WsxBox.OpenCodeInfo{}
		e = json.Unmarshal([]byte(arrData[i].ResultNums), &aOpenCodeInfo)
		if e != nil {
			ttLog.LogError(e)
		}
		arrWsxAwardInfo[i].OpenCodeInfo = aOpenCodeInfo
	}

	if e != nil {
		return mpData, e
	}

	mpData["LastId"] = LastId
	mpData["PageSize"] = PageSize
	mpData["Data"] = arrWsxAwardInfo
	return mpData, nil
}

func (this *WsxBbcServer) GetCurResult() (interface{}, error) {
	aWsxAwardInfo := WsxAwardInfo{}
	if this.curAwardInfo.LotteryStr == "" {
		return nil, errors.New(GTtHint.GetTtHint().GetHint("服务准备中"))
	}

	mpData := make(map[string]interface{})
	aWsxAwardInfo.NextLotteryStr = this.curAwardInfo.NextLotteryStr
	aWsxAwardInfo.LotteryStr = this.curAwardInfo.LotteryStr

	if this.TimeLag == 0 {
		aWsxAwardInfo.StrLotteryTime = this.curAwardInfo.CurLotteryTime.Format(timeKit.DateTimeLayout)
		aWsxAwardInfo.StrNextLotteryTime = this.curAwardInfo.NextLotteryTime.Format(timeKit.DateTimeLayout)
		aWsxAwardInfo.StrCurTime = this.curAwardInfo.CreatedAt.Format(timeKit.DateTimeLayout)
	} else {
		aWsxAwardInfo.StrLotteryTime = this.curAwardInfo.CurLotteryTime.Add(this.TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
		aWsxAwardInfo.StrNextLotteryTime = this.curAwardInfo.NextLotteryTime.Add(this.TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
		aWsxAwardInfo.StrCurTime = this.curAwardInfo.CreatedAt.Add(this.TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
	}
	aWsxAwardInfo.NextTime = this.curAwardInfo.NextTime

	aOpenCodeInfo := WsxBox.OpenCodeInfo{}
	e := json.Unmarshal([]byte(this.curAwardInfo.ResultNums), &aOpenCodeInfo)
	if e != nil {
		ttLog.LogError(e)
	}
	aWsxAwardInfo.OpenCodeInfo = aOpenCodeInfo

	if e != nil {
		return mpData, e
	}
	mpData["Data"] = aWsxAwardInfo
	return mpData, nil
}
