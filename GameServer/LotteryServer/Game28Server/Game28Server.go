package Game28Server

import (
	"errors"
	"fmt"
	"github.com/TtMyth123/GameServer/CacheData"
	"github.com/TtMyth123/GameServer/GInstance/GTtHint"
	"github.com/TtMyth123/GameServer/LotteryResult/Game28ResultKit"
	"github.com/TtMyth123/GameServer/LotteryServer"
	"github.com/TtMyth123/GameServer/LotteryServer/LotteryBox"
	"github.com/TtMyth123/GameServer/OtherServer/httpGameServer"
	"github.com/TtMyth123/GameServer/OtherServer/httpGameServer/httpBox"
	"github.com/TtMyth123/GameServer/OtherServer/httpGameServer/httpConst"
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
	"github.com/TtMyth123/kit/stringKit"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"sync"
	"time"
)

type OddsInfo struct {
	Id   int
	Odds float64 `json:"1"`
}
type Game28Server struct {
	LotteryServer.ILotteryServer
	OddsLock sync.RWMutex
	Name     string
	GameType int

	StopBetTime        int
	mpOddsInfo         map[int]models.LoOddsInfo
	mpOdds             map[int]float64
	gameMpOdds         map[int]OddsInfo
	curAwardInfo       models.LoAwardInfo
	curGame28AwardInfo Game28ResultKit.Game28AwardInfo

	mUserRpcClient *UserRpcClient.RpcClient
	gameHttpChan   chan GameHttpData

	sendStopBetStateP string
	StopBetHint       string
	MaxCountdown      int
}

func NewGame28Server(gameType, StopBetTime int, UserRpcClient *UserRpcClient.RpcClient) *Game28Server {
	aServer := new(Game28Server)
	aServer.GameType = gameType
	aServer.ReLoadOddsInfo()
	aServer.mUserRpcClient = UserRpcClient
	aServer.gameHttpChan = make(chan GameHttpData, 50)
	aServer.StopBetTime = StopBetTime

	go aServer.runGameHttpChan()
	go aServer.run()

	return aServer
}
func (this *Game28Server) ReLoadOddsInfo() {
	this.OddsLock.Lock()
	defer this.OddsLock.Unlock()
	mpOddsInfo := make(map[int]models.LoOddsInfo)
	mpOdds := make(map[int]float64)
	gameMpOdds := make(map[int]OddsInfo)
	arrOdds, e := models.GetAllLoOddsInfo(this.GameType)
	if e != nil {
		ttLog.LogError(e)
		return
	}

	for _, v := range arrOdds {
		gameMpOdds[v.OddsType] = OddsInfo{Id: v.OddsType, Odds: v.Odds}
		mpOddsInfo[v.OddsType] = v
		mpOdds[v.OddsType] = v.Odds
	}
	this.mpOddsInfo = mpOddsInfo
	this.mpOdds = mpOdds
	this.gameMpOdds = gameMpOdds
}

func (this *Game28Server) run() {
	ticker1 := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker1.C:
			{
				this.curGame28AwardInfo.Countdown--
				this.curGame28AwardInfo.StopCountdown = this.curGame28AwardInfo.Countdown - this.StopBetTime
				if this.sendStopBetStateP == "" && this.curGame28AwardInfo.StopCountdown < 0 {
					if this.curGame28AwardInfo.LotteryStr != "" {
						this.sendStopBetStateP = this.curGame28AwardInfo.LotteryStr

						aStopBetStateBox := httpBox.StopBetStateBox{}
						aGameHttpData := GameHttpData{Command: httpConst.StopBetState, Data: aStopBetStateBox}
						this.gameHttpChan <- aGameHttpData
					}
				}
			}
		}
	}
}

func (this *Game28Server) AddAwardInfo(newLoAwardInfo models.LoAwardInfo) {
	o := orm.NewOrm()
	c, e := o.QueryTable(mconst.TableName_LoAwardInfo).Filter("GameType", this.GameType).Filter("LotteryStr", newLoAwardInfo.LotteryStr).Count()
	if e != nil {
		ttLog.LogError(e)
	}
	if c == 0 {
		newLoAwardInfo.GameType = this.GameType
		newLoAwardInfo.Add(o)
	}
}

func (this *Game28Server) getBetGroupInfos(StrLotteryNum string) []*models.LoBetGroupInfo {
	o := orm.NewOrm()
	sql := fmt.Sprintf(`select a.* from %s a where a.str_period=? and a.status=? and a.game_type=?`, mconst.TableName_LoBetGroupInfo)

	arrLoBetGroupInfo := make([]*models.LoBetGroupInfo, 0)
	_, e := o.Raw(sql, StrLotteryNum, mconst.Bet_Status_1, this.GameType).QueryRows(&arrLoBetGroupInfo)
	if e != nil {
		ttLog.LogError(e)
	}

	return arrLoBetGroupInfo
}

func (this *Game28Server) getBetInfos(StrLotteryNum string) map[int][]*models.LoBetInfo {
	mpBetInfo := make(map[int][]*models.LoBetInfo)
	o := orm.NewOrm()
	sql := fmt.Sprintf(`select a.* from %s a where a.str_period=? and a.status=? and a.game_type=? `, mconst.TableName_LoBetInfo)

	arrBets := make([]*models.LoBetInfo, 0)
	_, e := o.Raw(sql, StrLotteryNum, mconst.Bet_Status_1, this.GameType).QueryRows(&arrBets)
	if e != nil {
		ttLog.LogError(e)
	}

	for _, b := range arrBets {
		if m, ok := mpBetInfo[b.GroupId]; ok {
			m = append(m, b)
			mpBetInfo[b.GroupId] = m
		} else {
			tmpBets := make([]*models.LoBetInfo, 0)
			tmpBets = append(tmpBets, b)
			mpBetInfo[b.GroupId] = tmpBets
		}
	}

	return mpBetInfo
}

func (this *Game28Server) GetRoomInfo() map[string]interface{} {
	mpInfo := make(map[string]interface{})

	mpInfo["1"] = this.StopBetTime
	mpInfo["2"] = this.gameMpOdds
	mpInfo["3"] = 0

	//mpInfo["3"] = this.curAwardInfo.NextLotteryTime.Format(timeKit.DateTimeLayout)
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

func (this *Game28Server) NewAwardInfo(newLoAwardInfo models.LoAwardInfo) {
	ttLog.LogDebug(stringKit.GetJsonStr(newLoAwardInfo))
	o := orm.NewOrm()
	this.AddAwardInfo(newLoAwardInfo)
	this.curAwardInfo = newLoAwardInfo
	this.sendStopBetStateP = ""

	this.curGame28AwardInfo = Game28ResultKit.GetGame28AwardInfo(newLoAwardInfo, this.StopBetTime, 1)
	if this.MaxCountdown > 0 && this.curGame28AwardInfo.Countdown > this.MaxCountdown {
		this.curGame28AwardInfo.Countdown = this.MaxCountdown
	}

	if this.curGame28AwardInfo.Countdown < 30 {
		ttLog.LogDebug("CountdownCountdownCountdown:", stringKit.GetJsonStr(this.curGame28AwardInfo))
	}

	aGameHttpData := GameHttpData{Command: httpConst.NewAwardInfo,
		Data: this.curGame28AwardInfo}
	this.gameHttpChan <- aGameHttpData

	mpBetInfo := this.getBetInfos(newLoAwardInfo.LotteryStr)
	aZg28Result := Game28ResultKit.GetZg28Result(newLoAwardInfo.ResultNums)
	mpUserWin := make(map[int]float64)
	mpUserLossWin := make(map[int]float64)
	arrBetGroupInfo := this.getBetGroupInfos(newLoAwardInfo.LotteryStr)
	curT := time.Now()
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

			betLostWin := Game28ResultKit.ComputeLoseWin(*bet, aZg28Result, this.mpOddsInfo)

			bet.Status = mconst.Bet_Status_2
			bet.ResultNums = newLoAwardInfo.ResultNums
			bet.Win = betLostWin.WinM
			allWin += betLostWin.WinM

			mpUserLossWin[groupBet.UserId] += betLostWin.WinM - float64(betLostWin.BetM)

			if float64(betLostWin.BetM) != betLostWin.WinM {
				rebateUser := mpUserName[bet.UserId]
				//betRebate := strconvEx.Decimal(float64(bet.BetM) * BetRebateRatio)
				mpUserXM[bet.UserId] += float64(betLostWin.BetM)
				//mpSum2Rebate[bet.UserId] += betRebate

				mconst.GetGameName(this.GameType)

				//Des := fmt.Sprintf(`%s[%s][%s],投注：%s %d元.返利比：%g`, mconst.GetGameName(this.GameType), newLoAwardInfo.LotteryStr, mpUserName[bet.UserId].UserName, bet.OddsName, bet.BetM, BetRebateRatio)
				//Des2 := GTtHint.GetTtHint().GetHint(`%s[%s][%s],投注：%s %d元.返利比：%g`)
				//DesMp := GTtHint.GetTtHint().GetMpString(mconst.GetGameName(this.GameType), newLoAwardInfo.LotteryStr, mpUserName[bet.UserId].UserName, bet.OddsName, bet.BetM, BetRebateRatio)

				//rebateInfo := gBox.AddRebateInfo{UserId: rebateUser.Pid, UserSid: bet.UserId, OldGold: float64(bet.BetM),
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

func (this *Game28Server) Bet(betInfo LotteryBox.BetInfo) (map[string]interface{}, error) {
	this.OddsLock.RLock()
	defer this.OddsLock.RUnlock()

	mpInfo := make(map[string]interface{})

	if this.StopBetHint != "" {
		return mpInfo, TtError.New(this.StopBetHint)
	}

	if this.curAwardInfo.NextLotteryStr == "" {
		return mpInfo, errors.New(GTtHint.GetTtHint().GetHint("未开始，不能投注"))
	}

	if this.curGame28AwardInfo.Countdown < this.StopBetTime {
		return mpInfo, TtError.New("%s期号已封盘", this.curAwardInfo.NextLotteryStr)
	}

	if betInfo.StrLotteryNum != this.curAwardInfo.NextLotteryStr {
		return mpInfo, TtError.New("期号不正确，不能投注。当前期号%s", this.curAwardInfo.NextLotteryStr)
	}

	mapOddTypeBet := make(map[int]int)
	mapOddTypeName := make(map[int]string)
	for _, v := range betInfo.BetData {
		mapOddTypeName[v.OddsType] = v.OddsName
		if this.mpOddsInfo[v.OddsType].OneUserMinBet > v.M {
			return mpInfo, TtError.New("[%s]单注金额少于%d", v.OddsName, this.mpOddsInfo[v.OddsType].OneUserMinBet)
		}

		mapOddTypeBet[v.OddsType] += v.M
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

	o := orm.NewOrm()
	allBetStr := ""
	allBetM := 0
	arrBetInfo := make([]models.LoBetInfo, 0)
	arrBetDataInfo := make([]httpBox.BetDataInfo, 0)
	for _, v := range betInfo.BetData {
		allBetM += v.M
		OddsName := v.OddsName
		if OddsName == "" {
			OddsName = this.mpOddsInfo[v.OddsType].OddsDes
		}
		arrBetDataInfo = append(arrBetDataInfo, httpBox.BetDataInfo{OddsType: v.OddsType,
			M: v.M, OddsName: OddsName})

		allBetStr += fmt.Sprintf(";%s:$%d", OddsName, v.M)

		aBetInfo := models.LoBetInfo{
			BetSn:     kit.GetGuid(),
			BetStr:    fmt.Sprintf("%s,$%d", OddsName, v.M),
			Status:    mconst.Bet_Status_1,
			Period:    betInfo.LotteryNum,
			StrPeriod: betInfo.StrLotteryNum,
			Odds:      this.mpOddsInfo[v.OddsType].Odds,
			OddsType:  v.OddsType,
			BetM:      v.M,
			Win:       0,
			GameType:  this.GameType,
			UserId:    betInfo.UserId,
			OddsName:  OddsName,
		}

		CacheData.AddCurPeriodUserBet(this.GameType, betInfo.UserId, v.M, v.OddsType, betInfo.StrLotteryNum)
		arrBetInfo = append(arrBetInfo, aBetInfo)
	}
	if len(allBetStr) > 4000 {
		return mpInfo, TtError.New("一次提交项目太多,无法投注.")
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

	aPlayerBetBox := httpBox.PlayerBetBox{UserId: betInfo.UserId,
		Money: float64(allBetM), BetData: arrBetDataInfo,
	}
	aGameHttpData := GameHttpData{Command: httpConst.PlayerBet, Data: aPlayerBetBox}
	this.gameHttpChan <- aGameHttpData

	id, e := aLoBetGroupInfo.Add(o)
	for i := 0; i < len(arrBetInfo); i++ {
		arrBetInfo[i].GroupBetSn = aLoBetGroupInfo.BetSn
		arrBetInfo[i].GroupId = id
	}

	e = models.InsertMultiBetInfo(o, arrBetInfo)
	mpInfo["Money"] = aUser.Gold
	mpInfo["BetM"] = aLoBetGroupInfo.BetM
	mpInfo["OrderNo"] = aLoBetGroupInfo.BetSn
	mpInfo["BetTime"] = aLoBetGroupInfo.CreatedAt.Format("2006-01-02 15:04:05")

	mpInfo["Game"] = mconst.GetGameName(this.GameType)
	mpInfo["GameTypeBet"] = CacheData.GetCurUserPeriodBet(betInfo.UserId, this.GameType, betInfo.StrLotteryNum)

	return mpInfo, e
}

func (this *Game28Server) GetHistoryResultList(PageIndex, PageSize, LastId int) (interface{}, error) {
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

	arrWsxAwardInfo := make([]Game28ResultKit.Game28HistoryResult, len(arrData))
	for i := 0; i < len(arrData); i++ {
		arrWsxAwardInfo[i] = Game28ResultKit.GetGame28HistoryResult(arrData[i])
	}

	if e != nil {
		return mpData, e
	}

	mpData["LastId"] = LastId
	mpData["PageSize"] = PageSize
	mpData["Data"] = arrWsxAwardInfo
	return mpData, nil
}

func (this *Game28Server) GetHistoryResultByPeriod(period, Count int) (interface{}, error) {
	sqlWhereLastId := ""

	mpData := make(map[string]interface{})
	arrData := make([]models.LoAwardInfo, 0)
	o := orm.NewOrm()
	sqlArgs := make([]interface{}, 0)

	sqlWhereLastId = `a.lottery_str>? and a.game_type=?`
	sqlArgs = append(sqlArgs, fmt.Sprintf("%d", period), this.GameType)

	sql := fmt.Sprintf(`select a.* from %s a where %s order by a.lottery_str DESC  LIMIT ? `,
		mconst.TableName_LoAwardInfo, sqlWhereLastId)

	sqlArgs = append(sqlArgs, Count)

	_, e := o.Raw(sql, sqlArgs).QueryRows(&arrData)

	arrWsxAwardInfo := make([]Game28ResultKit.Game28HistoryResult2, len(arrData))
	for i := 0; i < len(arrData); i++ {
		arrWsxAwardInfo[i] = Game28ResultKit.GetGame28HistoryResult2(arrData[i])
	}

	if e != nil {
		return mpData, e
	}

	iLen := len(arrWsxAwardInfo)
	if iLen > 0 {
		mpData["DataPeriod"] = arrWsxAwardInfo[0].Period
	} else {
		mpData["DataPeriod"] = 0
	}

	t := time.Now()
	iSecond := int64(this.curAwardInfo.NextLotteryTime.Sub(t) / time.Second)
	if iSecond <= 0 {
		iSecond = 0
	}

	mpData["CurPeriod"] = this.curGame28AwardInfo.LotteryStr
	mpData["NextPeriod"] = this.curGame28AwardInfo.NextLotteryStr
	mpData["StopBetTime"] = iSecond
	mpData["DataList"] = arrWsxAwardInfo
	mpData["Result"] = true
	return mpData, nil
}

/**
 */
func (this *Game28Server) GetHistoryLotteryByDay(UserId int, StrDay string) (interface{}, error) {

	type UscAwardInfo struct {
		Id         int
		NewNum     int64  //当前期号
		NewNumTime string //当前期开奖时间
		ResultNums string //当前期开奖结果 0,3,2,3,5,5,4,3
		ResultDX   string //大小： 小,小,小,小,大
		ResultDS   string //单双：单,双,单,单,双
		ResultLH   string //龙虎：龙,龙,龙,龙虎
		ResultGZH  string //冠亚和：10，大,小
		ResultFtFS int    //番数
		ResultFtH  int    //番数和
		ResultFtDS string //番数 单双

		LotteryNum    int64  //下一期期号
		CloseTime     string //封盘时间
		LotteryTime   string //下一期开奖时间
		ServerTime    string `json:"-"` //服务器时间
		TotalNum      int    `json:"-"`
		RestNum       int    `json:"-"`
		GameIndex     int
		Countdown     int //倒计时（秒）
		StopCountdown int //封 倒计时（秒）
	}

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
	arrR := make([]UscAwardInfo, iLen)
	for i := 0; i < iLen; i++ {
		aGame28AwardInfo := Game28ResultKit.GetGame28AwardInfo(arrData[i], 0, 0)

		arrR[i].Id = arrData[i].Id
		arrR[i].NewNumTime = arrData[i].CurLotteryTime.Format("2006-01-02 15:04:05")

		arrR[i].NewNum = arrData[i].LotteryNum
		arrR[i].ResultNums = arrData[i].ResultNums
		arrR[i].NewNumTime = arrData[i].CurLotteryTime.Format("2006-01-02 15:04:05")
		arrR[i].ResultDX = aGame28AwardInfo.ResultDX
		arrR[i].ResultDS = aGame28AwardInfo.ResultDS
		arrR[i].ResultLH = aGame28AwardInfo.ResultLH
		nums := Game28ResultKit.GetAwardNumbers(aGame28AwardInfo.ResultNums)
		h := Game28ResultKit.GetAwardResultHS(nums)
		arrR[i].ResultGZH = fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s", Game28ResultKit.GetAwardResultDX(h),
			Game28ResultKit.GetAwardResultDS(h),
			Game28ResultKit.GetAwardResultLH(nums),
			Game28ResultKit.GetAwardResultBS(h),
			Game28ResultKit.GetAwardResultDX(h)+Game28ResultKit.GetAwardResultDS(h),
			Game28ResultKit.GetAwardResultJS(h),
			Game28ResultKit.GetAwardResultBDS(nums),
		)

		arrR[i].LotteryNum = strconvEx.StrTry2Int64(arrData[i].NextLotteryStr, 0)
		arrR[i].CloseTime = arrData[i].NextLotteryTime.Format("2006-01-02 15:04:05")
		arrR[i].LotteryTime = arrData[i].NextLotteryTime.Format("2006-01-02 15:04:05")
	}

	type TmpResult struct {
		Game          string
		AwardInfoList []UscAwardInfo
	}

	aTmpResult := TmpResult{
		Game:          mconst.GetGameName(this.GameType),
		AwardInfoList: arrR,
	}

	return aTmpResult, nil
}
func (this *Game28Server) GetCurResult() (interface{}, error) {
	aGame28AwardInfo := Game28ResultKit.GetGame28AwardInfo(this.curAwardInfo, 0, 0)
	if this.curAwardInfo.LotteryStr == "" {
		return nil, errors.New(GTtHint.GetTtHint().GetHint("服务准备中"))
	}
	type TmpResult struct {
		Game      string
		Money     float64
		AwardInfo map[string]interface{}
	}
	aTmpResult := TmpResult{}
	aTmpResult.Game = mconst.GetGameName(this.GameType)
	aTmpResult.AwardInfo = map[string]interface{}{}
	aTmpResult.Money = 0
	aTmpResult.AwardInfo["LotteryStr"] = aGame28AwardInfo.LotteryStr
	aTmpResult.AwardInfo["NextLotteryStr"] = aGame28AwardInfo.NextLotteryStr
	aTmpResult.AwardInfo["ResultNums"] = aGame28AwardInfo.ResultNums
	aTmpResult.AwardInfo["ResultDX"] = aGame28AwardInfo.ResultDX
	aTmpResult.AwardInfo["ResultDS"] = aGame28AwardInfo.ResultDS
	aTmpResult.AwardInfo["ResultLH"] = aGame28AwardInfo.ResultLH
	aTmpResult.AwardInfo["ResultGZH"] = aGame28AwardInfo.ResultGZH

	aTmpResult.AwardInfo["ResultFtFS"] = 0
	aTmpResult.AwardInfo["ResultFtH"] = 0
	aTmpResult.AwardInfo["ResultFtDS"] = 0
	aTmpResult.AwardInfo["StopCountdown"] = aGame28AwardInfo.StopCountdown

	aTmpResult.AwardInfo["Countdown"] = aGame28AwardInfo.Countdown
	aTmpResult.AwardInfo["LotteryTime"] = aGame28AwardInfo.StrNextLotteryTime
	aTmpResult.AwardInfo["NewNumTime"] = aGame28AwardInfo.StrLotteryTime

	//mpData := make(map[string]interface{})
	//mpData["Data"] = aGame28AwardInfo
	return aTmpResult, nil
}

func (this *Game28Server) SetStopBetHint(StopBetHint string) error {
	this.StopBetHint = StopBetHint
	return nil
}

func (this *Game28Server) SetAwardInfo(LotteryAward, LotteryStr string) (interface{}, error) {
	if this.GameType != mconst.GameType_G28_044 {
		return nil, errors.New("目前只有香港28可以设置")
	}
	nums := lotteryKit.GetStrNum2Arr(LotteryAward)
	if len(nums) != 3 {
		return nil, errors.New("结果只能是3个数值")
	}
	for _, n := range nums {
		if n < 0 {
			return nil, errors.New(fmt.Sprintf("开奖号码【%d】有误，它不能小于0，", n))
		}
		if n > 28 {
			return nil, errors.New(fmt.Sprintf("开奖号码【%d】有误，它不能大于28，", n))
		}
	}
	if LotteryStr == "" {
		LotteryStr = this.curAwardInfo.NextLotteryStr
	}
	newLotteryAward := lotteryKit.GetArrNum2String(nums, 1)
	LotteryNum := strconvEx.StrTry2Int64(LotteryStr, 0)

	e := httpGameServer.SetAwardResult(this.GameType, newLotteryAward, LotteryStr)
	if e != nil {
		ttLog.LogError(e)
		return nil, e
	}
	o := orm.NewOrm()
	aLoSetAwardInfo := models.LoSetAwardInfo{}
	e = o.QueryTable(mconst.TableName_LoSetAwardInfo).Filter("GameType", this.GameType).Filter("LotteryStr", LotteryStr).One(&aLoSetAwardInfo)

	if e == nil {
		aLoSetAwardInfo.ResultNums = newLotteryAward
		e = aLoSetAwardInfo.Update(nil, "ResultNums", "UpdatedAt")
	} else {
		aLoSetAwardInfo = models.LoSetAwardInfo{
			GameType:   this.GameType,
			LotteryStr: LotteryStr,
			LotteryNum: LotteryNum,
			ResultNums: newLotteryAward,
		}
		e = aLoSetAwardInfo.Add(nil)
	}

	if e != nil {
		ttLog.LogError(e)
		return nil, e
	}

	return nil, nil
}
