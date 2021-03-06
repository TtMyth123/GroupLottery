package WsxServer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"sync"
	"time"
	"ttmyth123/GroupLottery/GameServer/CacheData"
	"ttmyth123/GroupLottery/GameServer/GConfig"
	"ttmyth123/GroupLottery/GameServer/GInstance/GTtHint"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/WsxBox"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/WsxZbcResultKit"
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

type WsxZbcServer struct {
	LotteryServer.BaseLotteryServer
	OddsLock sync.RWMutex
	BaseWsxServer
	mpOddsInfo map[int]ZbcOdds

	lotteryTime LotteryTime
	iii         int
}

func NewWsxZbcServer(gameType int, UserRpcClient *UserRpcClient.RpcClient) *WsxZbcServer {
	aWsxServer := new(WsxZbcServer)
	aWsxServer.TimeLag = GConfig.GetGConfig().TimeLag
	aWsxServer.GameType = gameType
	aWsxServer.ReLoadOddsInfo()
	aWsxServer.mUserRpcClient = UserRpcClient
	aWsxServer.iii = 0
	aWsxServer.run()
	return aWsxServer
}

func (this *WsxZbcServer) run() {
	ticker1 := time.NewTicker(500 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker1.C:
				this.getNextIssueAndTime()
			}
		}
	}()
}

type ZbcOdds struct {
	OneUserMaxBet int
	OneUserMinBet int
	AllUserMaxBet int
	Odds          float64
	StrOdds       string
	BigOddsDes    string
}

func (this *WsxZbcServer) ReLoadOddsInfo() {
	this.OddsLock.Lock()
	defer this.OddsLock.Unlock()
	mpBigTypeOdds := make(map[int]ZbcOdds)
	mpOddsInfo := make(map[int]map[int]float64)
	mpOdds := make(map[int]float64)
	arrOdds, e := models.GetAllLoOddsInfo(this.GameType)
	if e != nil {
		ttLog.LogError(e)
		return
	}

	for _, v := range arrOdds {
		if bigTypeOdds, ok := mpOddsInfo[v.BigType]; !ok {
			bigTypeOdds = make(map[int]float64)
			bigTypeOdds[v.N2] = v.Odds
			mpOddsInfo[v.BigType] = bigTypeOdds
		} else {
			bigTypeOdds[v.N2] = v.Odds
			mpOddsInfo[v.BigType] = bigTypeOdds
		}

		if _, ok := mpBigTypeOdds[v.BigType]; !ok {
			aZbcOdds := ZbcOdds{
				OneUserMaxBet: v.OneUserMaxBet,
				OneUserMinBet: v.OneUserMinBet,
				AllUserMaxBet: v.AllUserMaxBet,
				Odds:          v.Odds,
				BigOddsDes:    v.BigOddsDes,
			}
			mpBigTypeOdds[v.BigType] = aZbcOdds
		}

		mpOdds[v.BigType] = v.Odds
	}
	for k, v := range mpBigTypeOdds {
		v.StrOdds = stringKit.GetJsonStr(mpOddsInfo[k])
		mpBigTypeOdds[k] = v
	}

	this.mpOddsInfo = mpBigTypeOdds
	this.mpOdds = mpOdds
}

func (this *WsxZbcServer) GetRoomInfo() map[string]interface{} {
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
func (this *WsxZbcServer) RandNewAwardInfo() error {
	aOpenCodeInfo := WsxBox.OpenCodeZbcInfo{}
	aOpenCodeInfo.Nums = make([]string, 20)
	tmpM := make(map[int]int)
	for i := 0; i < 20; i++ {
		num := timeKit.GlobaRand.Intn(99)
		if _, ok := tmpM[num]; !ok {
			tmpM[num] = 1
			aOpenCodeInfo.Nums[i] = fmt.Sprintf("%02d", num)
		}
	}

	aResultNums := make(map[string]interface{})
	aResultNums["Nums"] = aOpenCodeInfo.Nums

	aAwardInfo := WsxBox.WsxResultData{}

	Expect := this.curAwardInfo.NextLotteryStr
	if Expect == "" {
		return errors.New("?????????...")
	}
	NextIssue := strconvEx.StrTry2Int64(Expect, 0)
	if NextIssue == 0 {
		return errors.New("?????????...")
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

type LotteryTime struct {
	NextIssue string
	NextTime  time.Time

	NextIssue2 string
	NextTime2  time.Time
}

func (this *WsxZbcServer) GetLotteryTime(aWsxResultData WsxBox.WsxResultData) (LotteryTime, error) {
	aLotteryTime := LotteryTime{}
	NextTime, e := timeKit.GetTime(aWsxResultData.NextTime)
	if e != nil {
		return LotteryTime{}, e
	}
	aLotteryTime.NextTime = NextTime
	aLotteryTime.NextIssue = aWsxResultData.NextIssue

	NextTime2, e := timeKit.GetTime(aWsxResultData.NextTime2)
	if e != nil {
		return LotteryTime{}, e
	}
	aLotteryTime.NextTime2 = NextTime2
	aLotteryTime.NextIssue2 = aWsxResultData.NextIssue2

	return aLotteryTime, nil
}
func (this *WsxZbcServer) getNextIssueAndTime() {
	t := time.Now()
	if this.lotteryTime.NextIssue == "" {
		return
	}

	ii := int(t.Sub(this.lotteryTime.NextTime) / time.Second)
	//ii2 := int(this.lotteryTime.NextTime2.Sub(t) / time.Second)
	//
	//fmt.Println("        T:",t)
	//fmt.Println(" NextTime:",this.lotteryTime.NextTime,ii)
	//fmt.Println("NextTime2:",this.lotteryTime.NextTime2,ii2)
	if ii > 0 {
		this.curAwardInfo.NextLotteryStr = this.lotteryTime.NextIssue2
		this.curAwardInfo.NextLotteryTime = this.lotteryTime.NextTime2
		NextTime := int(this.lotteryTime.NextTime2.Sub(t) / time.Second)

		this.curAwardInfo.NextTime = strconv.Itoa(NextTime)
	} else {
		this.curAwardInfo.NextLotteryStr = this.lotteryTime.NextIssue
		this.curAwardInfo.NextLotteryTime = this.lotteryTime.NextTime
		NextTime := int(this.lotteryTime.NextTime.Sub(t) / time.Second)
		this.curAwardInfo.NextTime = strconv.Itoa(NextTime)
	}
}

/**
??????????????????????????????
*/
func (this *WsxZbcServer) processorNotOpenBet(LotteryStr string) {
	o := orm.NewOrm()
	mpBetInfo := this.getNotOpenBetInfos(LotteryStr)
	arrBetGroupInfo := this.getNotOpenBetGroupInfos(LotteryStr)
	for _, groupBet := range arrBetGroupInfo {
		for _, bet := range mpBetInfo[groupBet.Id] {
			bet.Status = mconst.Bet_Status_3
			bet.Win = float64(bet.BetM)
			bet.Update(o, "Status", "Win")
		}
		groupBet.Win = float64(groupBet.BetM)
		groupBet.Status = mconst.Bet_Status_3
		groupBet.Update(o, "Status", "Win")

		goldInfo := gBox.AddGoldInfo{GroupId: groupBet.GroupUserId,UserId: groupBet.UserId, Gold: float64(groupBet.BetM),
			T:    userConst.Account_14_NotOpen,
			Des:  fmt.Sprintf("%s???%s???????????????[%s]??????%d???, ??????%d", mconst.GetGameName(this.GameType), LotteryStr, groupBet.BetSn, groupBet.BetM, groupBet.BetM),
			Des2: GTtHint.GetTtHint().GetHint("%s???%s???????????????[%s]??????%d???, ??????%d"), DesMp: GTtHint.GetTtHint().GetMpString(mconst.GetGameName(this.GameType), LotteryStr, groupBet.BetSn, groupBet.BetM, groupBet.BetM)}
		_, e := this.mUserRpcClient.AddGold(goldInfo)
		if e != nil {
			ttLog.LogError(e)
		}
	}
}

func (this *WsxZbcServer) NewAwardInfo(newLoAwardInfo models.LoAwardInfo) {
	aWsxResultData := WsxBox.WsxResultData{}
	e := stringKit.GetJsonObj(newLoAwardInfo.OriginalResult, &aWsxResultData)
	if e != nil {
		ttLog.LogError(e, newLoAwardInfo)
		return
	}

	this.AddAwardInfo(&newLoAwardInfo)
	this.lotteryTime, e = this.GetLotteryTime(aWsxResultData)
	if e != nil {
		ttLog.LogError(e, newLoAwardInfo)
		return
	}

	this.curAwardInfo = newLoAwardInfo
	this.PreOpenTime = time.Now()

	o := orm.NewOrm()

	this.processorNotOpenBet(newLoAwardInfo.LotteryStr)

	mpBetInfo := this.getBetInfos(newLoAwardInfo.LotteryStr)
	aOpenCodeInfo := WsxBox.OpenCodeZbcInfo{}
	e = json.Unmarshal([]byte(newLoAwardInfo.ResultNums), &aOpenCodeInfo)
	if e != nil {
		ttLog.LogError(e, newLoAwardInfo.ResultNums)
		return
	}

	//BetRebateRatio := AreaConfig.GetRebateSet(0).BetRebateRatio
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

			win, Odds := WsxZbcResultKit.ComputeLoseWin(*bet, aOpenCodeInfo)

			bet.Status = mconst.Bet_Status_2
			if Odds != 0 {
				bet.Odds = Odds
			}
			//
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
			bet.Update(o, "Status", "ResultNums", "Win", "Odds")
		}
		groupBet.Update(o, "Status", "Win", "ResultNums")
		if groupBet.Win == 0 {
			continue
		}
		goldInfo := gBox.AddGoldInfo{GroupId: groupBet.GroupUserId, UserId: groupBet.UserId, Gold: groupBet.Win,
			T:    userConst.Account_02_Win,
			Des:  fmt.Sprintf("%s???%s??????[%s]??????%d???, ??????%g", mconst.GetGameName(this.GameType), newLoAwardInfo.LotteryStr, groupBet.BetSn, groupBet.BetM, groupBet.Win),
			Des2: GTtHint.GetTtHint().GetHint("%s???%s??????[%s]??????%d???, ??????%g"), DesMp: GTtHint.GetTtHint().GetMpString(mconst.GetGameName(this.GameType), newLoAwardInfo.LotteryStr, groupBet.BetSn, groupBet.BetM, groupBet.Win)}

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

func (this *WsxZbcServer) Bet(betInfo LotteryBox.BetInfo) (map[string]interface{}, error) {
	this.OddsLock.RLock()
	defer this.OddsLock.RUnlock()
	mpInfo := make(map[string]interface{})

	if this.curAwardInfo.NextLotteryStr == "" {
		return mpInfo, errors.New(GTtHint.GetTtHint().GetHint("????????????????????????"))
	}
	if betInfo.StrLotteryNum != this.curAwardInfo.NextLotteryStr {
		return mpInfo, TtError.New("?????????????????????????????????????????????%s", this.curAwardInfo.NextLotteryStr)
	}

	t := time.Now()
	iSecond := int64(this.curAwardInfo.NextLotteryTime.Sub(t) / time.Second)
	if iSecond <= 30 {
		return mpInfo, errors.New(GTtHint.GetTtHint().GetHint("????????????,??????????????????."))
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
			return mpInfo, TtError.New("[%s]??????????????????%d", v.OddsName, this.mpOddsInfo[v.OddsType].OneUserMinBet)
		}
		allBetM += v.M
		OddsName := v.OddsName
		if OddsName == "" {
			OddsName = this.mpOddsInfo[v.OddsType].BigOddsDes
		}
		allBetStr += fmt.Sprintf(";%s:$%d", OddsName, v.M)

		aBetInfo := models.LoBetInfo{
			BetSn:       kit.GetGuid(),
			BetStr:      fmt.Sprintf("%s,$%d", OddsName, v.M),
			Status:      mconst.Bet_Status_1,
			Period:      betInfo.LotteryNum,
			StrPeriod:   betInfo.StrLotteryNum,
			Odds:        this.mpOdds[v.OddsType],
			StrOdds:     this.mpOddsInfo[v.OddsType].StrOdds,
			OddsType:    0,
			BigOddsType: v.OddsType,
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
		return mpInfo, TtError.New("????????????????????????,????????????.")
	}
	for oddsType, typeBet := range mapOddTypeBet {
		betUserGamePeriodOddsType, betGamePeriodOddsType := CacheData.GetCacheSumBet(betInfo.StrLotteryNum, this.GameType, oddsType, betInfo.UserId)

		if this.mpOddsInfo[oddsType].AllUserMaxBet < betGamePeriodOddsType+typeBet {
			return mpInfo, TtError.New("[%s]????????????????????????", mapOddTypeName[oddsType])
			//return mpInfo, errors.New(GTtHint.GetTtHint().GetHint("[%s]????????????????????????"))
		}

		if this.mpOddsInfo[oddsType].OneUserMaxBet < betUserGamePeriodOddsType+typeBet {
			return mpInfo, TtError.New("[%s]????????????????????????", mapOddTypeName[oddsType])
			//return mpInfo, errors.New(GTtHint.GetTtHint().GetHint("????????????????????????"))
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
		Des:  fmt.Sprintf("???%s??? ??????:%s ?????????:%d", betInfo.StrLotteryNum, allBetStr, allBetM),
		Des2: GTtHint.GetTtHint().GetHint("???%s??? ??????:%s ?????????:%d"), DesMp: GTtHint.GetTtHint().GetMpString(betInfo.StrLotteryNum, allBetStr, allBetM)}

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

type WsxZbcAwardInfo struct {
	LotteryStr     string //??????
	StrLotteryTime string //??????????????????

	NextLotteryStr     string
	StrNextLotteryTime string //?????????????????????
	StrCurTime         string //????????????
	NextTime           string //????????????

	Nums []string
}

func (this *WsxZbcServer) GetHistoryResultList(PageIndex, PageSize, LastId int) (interface{}, error) {
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

	sql := fmt.Sprintf(`select a.* from %s a where %s a.game_type=? order by a.created_at DESC LIMIT ?,? `,
		mconst.TableName_LoAwardInfo, sqlWhereLastId)

	sqlArgs = append(sqlArgs, offset, PageSize)

	_, e = o.Raw(sql, sqlArgs).QueryRows(&arrData)

	arrWsxAwardInfo := make([]WsxZbcAwardInfo, len(arrData))
	for i := 0; i < len(arrData); i++ {
		arrWsxAwardInfo[i].LotteryStr = arrData[i].LotteryStr
		if this.TimeLag == 0 {
			arrWsxAwardInfo[i].StrLotteryTime = arrData[i].CurLotteryTime.Format(timeKit.DateTimeLayout)
		} else {
			arrWsxAwardInfo[i].StrLotteryTime = arrData[i].CurLotteryTime.Add(this.TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
		}
		aOpenCodeInfo := WsxBox.OpenCodeZbcInfo{}
		e = json.Unmarshal([]byte(arrData[i].ResultNums), &aOpenCodeInfo)
		if e != nil {
			ttLog.LogError(e)
		}
		arrWsxAwardInfo[i].Nums = aOpenCodeInfo.Nums
	}

	if e != nil {
		return mpData, e
	}
	mpData["LastId"] = LastId
	mpData["PageSize"] = PageSize
	mpData["Data"] = arrWsxAwardInfo
	return mpData, nil
}

func (this *WsxZbcServer) GetCurResult() (interface{}, error) {
	aWsxAwardInfo := WsxZbcAwardInfo{}
	if this.curAwardInfo.LotteryStr == "" {
		return nil, errors.New(GTtHint.GetTtHint().GetHint("???????????????"))
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

	aOpenCodeInfo := WsxBox.OpenCodeZbcInfo{}
	e := json.Unmarshal([]byte(this.curAwardInfo.ResultNums), &aOpenCodeInfo)
	if e != nil {
		ttLog.LogError(e)
	}
	aWsxAwardInfo.Nums = aOpenCodeInfo.Nums

	if e != nil {
		return mpData, e
	}
	mpData["Data"] = aWsxAwardInfo
	return mpData, nil
}
