package UscServer

import (
	"errors"
	"fmt"
	"github.com/TtMyth123/GameServer/CacheData"
	"github.com/TtMyth123/GameServer/GInstance/GTtHint"
	"github.com/TtMyth123/GameServer/LotteryResult/UscResultKit/Usc10Kit"
	"github.com/TtMyth123/GameServer/LotteryResult/UscResultServer/UscBox"
	"github.com/TtMyth123/GameServer/LotteryServer"
	"github.com/TtMyth123/GameServer/LotteryServer/LotteryBox"
	"github.com/TtMyth123/GameServer/controllers/base/TtError"
	"github.com/TtMyth123/GameServer/models"
	"github.com/TtMyth123/GameServer/models/mconst"
	"github.com/TtMyth123/UserInfoRpc/GData/gBox"
	"github.com/TtMyth123/UserInfoRpc/UserRpcClient"
	userModels "github.com/TtMyth123/UserInfoRpc/models"
	userConst "github.com/TtMyth123/UserInfoRpc/models/mconst"
	"github.com/TtMyth123/kit"
	"github.com/TtMyth123/kit/lotteryKit"
	"github.com/TtMyth123/kit/sqlKit"
	"github.com/TtMyth123/kit/strconvEx"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"sync"
	"time"
)

type Usc10Server struct {
	LotteryServer.BaseLotteryServer
	OddsLock sync.RWMutex
	BaseUscServer
	mpOddsInfo map[int]models.LoOddsInfo
}

func NewUsc10Server(gameType int, UserRpcClient *UserRpcClient.RpcClient, StopBetTime int) *Usc10Server {
	aWsxServer := new(Usc10Server)
	aWsxServer.GameType = gameType
	aWsxServer.ReLoadOddsInfo()
	aWsxServer.mUserRpcClient = UserRpcClient
	aWsxServer.StopBetTime = StopBetTime
	aWsxServer.run()
	return aWsxServer
}

func (this *Usc10Server) run() {
	ticker1 := time.NewTicker(1000 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker1.C:
				this.downTime()
			}
		}
	}()
}
func (this *Usc10Server) downTime() {
	t := time.Now()
	iSecond := int(this.curAwardInfo.LotteryTime.Sub(t) / time.Second)
	if iSecond <= 0 {
		iSecond = 0
	}

	this.curAwardInfo.Countdown = iSecond
	this.curAwardInfo.StopCountdown = this.StopBetTime
	fmt.Println("GameType:", this.GameType, this.curAwardInfo.Countdown, this.curAwardInfo.Countdown, this.curAwardInfo.StopCountdown)
}

func (this *Usc10Server) ReLoadOddsInfo() {
	this.OddsLock.Lock()
	defer this.OddsLock.Unlock()
	mpOddsInfo := make(map[int]models.LoOddsInfo)
	mpOdds := make(map[int]UscBox.OddsInfo)
	arrOdds, e := models.GetAllLoOddsInfo(this.GameType)
	if e != nil {
		ttLog.LogError(e)
		return
	}

	for _, v := range arrOdds {
		mpOddsInfo[v.OddsType] = v
		mpOdds[v.OddsType] = UscBox.OddsInfo{v.OddsType, v.Odds}
	}
	this.mpOddsInfo = mpOddsInfo
	this.mpOdds = mpOdds
}

func (this *Usc10Server) GetRoomInfo() map[string]interface{} {
	mpInfo := make(map[string]interface{})

	mpInfo["Result"] = true
	mpInfo["MapOdds"] = this.mpOdds
	mpInfo["PageIndex"] = 0
	mpInfo["StopBetTime"] = this.StopBetTime
	mpInfo["Countdown"] = this.curAwardInfo.Countdown
	mpInfo["CurLotteryNum"] = this.curAwardInfo.NewNum
	mpInfo["NextLotteryNum"] = this.curAwardInfo.LotteryNum
	return mpInfo
}
func (this *Usc10Server) RandNewAwardInfo() error {
	this.RandNewAwardInfoLock.Lock()
	defer this.RandNewAwardInfoLock.Unlock()

	aLoAwardInfo := models.LoAwardInfo{}

	e := this.DelAfData(aLoAwardInfo.LotteryStr)
	if e != nil {
		return e
	}

	this.NewAwardInfo(aLoAwardInfo)
	return nil
}

func (this *Usc10Server) LoAwardInfo2UscAwardInfo(newLoAwardInfo models.LoAwardInfo) UscBox.UscAwardInfo {
	numList := lotteryKit.GetStrNum2Arr(newLoAwardInfo.ResultNums)
	ResultDX := Usc10Kit.GetResultDXs(numList)
	ResultDS := Usc10Kit.GetResultDSs(numList)
	ResultGZH := Usc10Kit.GetResultGZH(numList)
	ResultLH := Usc10Kit.GetResultLHs(numList)
	ResultFtH := Usc10Kit.GetFTHS(numList)
	ResultFtFS := Usc10Kit.GetFS(ResultFtH)
	ResultFtDS := Usc10Kit.GetResultFtDS(ResultFtFS)
	aUscAwardInfo := UscBox.UscAwardInfo{
		NewNum:      newLoAwardInfo.LotteryNum,
		NewNumTime:  newLoAwardInfo.CurLotteryTime,
		ResultNums:  newLoAwardInfo.ResultNums,
		ResultDX:    ResultDX,
		ResultDS:    ResultDS,
		ResultLH:    ResultLH,
		ResultGZH:   ResultGZH,
		ResultFtFS:  ResultFtFS,
		ResultFtH:   ResultFtH,
		ResultFtDS:  ResultFtDS,
		LotteryNum:  strconvEx.StrTry2Int64(newLoAwardInfo.NextLotteryStr, 0),
		LotteryTime: newLoAwardInfo.NextLotteryTime,
		ServerTime:  newLoAwardInfo.CreatedAt,
		GameIndex:   this.GameIndex,
	}

	return aUscAwardInfo
}

func (this *Usc10Server) NewAwardInfo(newLoAwardInfo models.LoAwardInfo) {
	aUscAwardInfo := this.LoAwardInfo2UscAwardInfo(newLoAwardInfo)

	this.BaseUscServer.PreNewAwardInfo(newLoAwardInfo, aUscAwardInfo)
	o := orm.NewOrm()

	mpBetInfo := this.getBetInfos(newLoAwardInfo.LotteryStr)
	mpUserWin := make(map[int]float64)
	//mpUserLossWin:= make(map[int]float64)

	//BetRebateRatio := AreaConfig.GetRebateSet(0).BetRebateRatio
	curT := time.Now()
	arrBetGroupInfo := this.getBetGroupInfos(newLoAwardInfo.LotteryStr)

	mpUserName := make(map[int]userModels.TtGameUser)
	mpUserXM := make(map[int]float64)
	mpSum2Rebate := make(map[int]float64)

	aAward, _ := Usc10Kit.GetAward(newLoAwardInfo.ResultNums)

	for _, groupBet := range arrBetGroupInfo {

		allWin := 0.0
		for _, bet := range mpBetInfo[groupBet.Id] {
			if _, ok := mpUserName[bet.UserId]; !ok {
				aU, _ := this.mUserRpcClient.GetUser(bet.UserId)
				mpUserName[bet.UserId] = aU
			}

			win := Usc10Kit.GetWinMoney(aAward, bet.OddsType, bet.BetM, bet.Odds)

			bet.Status = mconst.Bet_Status_2
			bet.ResultNums = newLoAwardInfo.ResultNums
			bet.Win = win
			allWin += win

			if float64(bet.BetM) != win {
				rebateUser := mpUserName[bet.UserId]
				mpUserXM[bet.UserId] += float64(bet.BetM)
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

func (this *Usc10Server) Bet(betInfo LotteryBox.BetInfo) (map[string]interface{}, error) {
	this.OddsLock.RLock()
	defer this.OddsLock.RUnlock()
	mpInfo := make(map[string]interface{})

	//if this.curAwardInfo.NextLotteryStr == "" {
	//	return mpInfo, errors.New(GTtHint.GetTtHint().GetHint("未开始，不能投注"))
	//}
	//if betInfo.StrLotteryNum != this.curAwardInfo.NextLotteryStr {
	//	return mpInfo, TtError.New("期号不正确，不能投注。当前期号%s",this.curAwardInfo.NextLotteryStr)
	//}
	e := this.BaseUscServer.PreBet(betInfo)
	if e != nil {
		return mpInfo, e
	}

	//if this.PreIssue !=0 {
	//	curTime:=time.Now()
	//	subMinute := curTime.Sub(this.PreOpenTime) / time.Minute
	//	if subMinute <= 25 {
	//		return mpInfo,  TtError.New("正在开奖,暂时无法下注.")
	//	}
	//
	//	subMinute = this.curAwardInfo.NextLotteryTime.Sub(curTime) / time.Minute
	//	if subMinute <= 5 {
	//		return mpInfo, TtError.New("正在开奖,暂时无法下注.")
	//	}
	//}
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

func (this *Usc10Server) GetHistoryResultList(PageIndex, PageSize, LastId int) (interface{}, error) {
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

	sql := fmt.Sprintf(`select a.* from %s a where %s a.game_type=? order by a.lottery_str DESC  LIMIT ?,? `,
		mconst.TableName_LoAwardInfo, sqlWhereLastId)

	sqlArgs = append(sqlArgs, offset, PageSize)

	_, e = o.Raw(sql, sqlArgs).QueryRows(&arrData)

	arrWsxAwardInfo := make([]UscBox.UscAwardInfo, len(arrData))
	for i := 0; i < len(arrData); i++ {
		arrWsxAwardInfo[i] = this.LoAwardInfo2UscAwardInfo(arrData[i])
	}

	if e != nil {
		return mpData, e
	}

	mpData["LastId"] = LastId
	mpData["PageSize"] = PageSize
	mpData["Data"] = arrWsxAwardInfo
	return mpData, nil
}

func (this *Usc10Server) GetCurResult() (interface{}, error) {
	if this.curAwardInfo.NewNum == 0 {
		return nil, errors.New(GTtHint.GetTtHint().GetHint("服务准备中"))
	}

	mpData := make(map[string]interface{})
	mpData["AwardInfo"] = this.curAwardInfo
	return mpData, nil
}
func (this *Usc10Server) getHistoryFTNum(LastId, PageSize int) []UscBox.UscAwardFTInfo {
	PageIndex := 1
	sqlWhereLastId := ""
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
		return make([]UscBox.UscAwardFTInfo, 0)
	}
	offset, _ := sqlKit.GetOffset(PageTotal, PageSize, PageIndex)

	sql := fmt.Sprintf(`select a.* from %s a where %s a.game_type=? order by a.lottery_str DESC  LIMIT ?,? `,
		mconst.TableName_LoAwardInfo, sqlWhereLastId)

	sqlArgs = append(sqlArgs, offset, PageSize)

	_, e = o.Raw(sql, sqlArgs).QueryRows(&arrData)

	arrWsxAwardInfo := make([]UscBox.UscAwardFTInfo, len(arrData))
	for i := 0; i < len(arrData); i++ {
		newLoAwardInfo := arrData[i]
		numList := lotteryKit.GetStrNum2Arr(newLoAwardInfo.ResultNums)
		FTHS := Usc10Kit.GetFTHS(numList)
		arrWsxAwardInfo[i].Id = arrData[i].Id
		arrWsxAwardInfo[i].ResultFtFS = Usc10Kit.GetFS(FTHS)
		arrWsxAwardInfo[i].ResultFtH = FTHS
		arrWsxAwardInfo[i].ResultFtDS = Usc10Kit.GetResultFtDS(arrWsxAwardInfo[i].ResultFtFS)
	}
	return arrWsxAwardInfo
}
func (this *Usc10Server) GetHistoryFTNum(PageSize int, LastId int) (interface{}, error) {
	mpData := make(map[string]interface{})
	datas := this.getHistoryFTNum(PageSize, LastId)
	mpData["AwardInfoList"] = datas
	mpData["Result"] = true
	mpData["Game"] = mconst.GetGameName(this.GameType)
	return mpData, nil
}

func (this *Usc10Server) GetHistoryFTNumBy48(LastId int) (interface{}, error) {
	mpData := make(map[string]interface{})
	datas := this.getHistoryFTNum(LastId, UscBox.FT48Count)
	mpData["AwardInfoList"] = datas
	mpData["Result"] = true
	mpData["Game"] = mconst.GetGameName(this.GameType)
	return mpData, nil
}

/**
 */
func (this *Usc10Server) GetHistoryLotteryByDay(UserId int, StrDay string) (interface{}, error) {
	arrData := make([]models.LoAwardInfo, 0)
	o := orm.NewOrm()
	sqlWhere := `a.game_type=?`
	sqlArgs := make([]interface{}, 0)
	sqlArgs = append(sqlArgs, this.GameType)

	if StrDay != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.cur_lottery_time,'%[1]s') >= ? `, "%Y-%m-%d")
		sqlArgs = append(sqlArgs, StrDay)
	}
	sql := fmt.Sprintf(`select a.* from %s a where %s order by a.cur_lottery_time DESC `,
		mconst.TableName_LoAwardInfo, sqlWhere)

	o.Raw(sql, sqlArgs).QueryRows(&arrData)
	iLen := len(arrData)
	arrR := make([]UscBox.UscAwardInfo, iLen)
	for i := 0; i < iLen; i++ {
		arrR[i] = this.LoAwardInfo2UscAwardInfo(arrData[i])

	}

	type TmpResult struct {
		Game          string
		AwardInfoList []UscBox.UscAwardInfo
	}

	aTmpResult := TmpResult{
		Game:          mconst.GetGameName(this.GameType),
		AwardInfoList: arrR,
	}

	return aTmpResult, nil
}
