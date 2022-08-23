package ReportBll

import (
	"encoding/json"
	"fmt"
	"github.com/TtMyth123/Admin/CacheData"
	"github.com/TtMyth123/Admin/controllers/ReportBll/Chart"
	"github.com/TtMyth123/GameServer/models/mconst"
	userConst "github.com/TtMyth123/UserInfoRpc/models/mconst"
	"github.com/TtMyth123/kit/sqlKit"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"

	gameCacheData "github.com/TtMyth123/GameServer/CacheData"
)

type GroupBetInfo struct {
	Id           int
	UserId       int
	UserName     string
	GameName     string
	StrPeriod    string
	BetM         int
	Win          float64
	BetSn        string
	BetStr       string
	Status       int
	CreatedAt    time.Time
	ArrBetDetail []string
}

func (d GroupBetInfo) MarshalJSON() ([]byte, error) {
	type Alias GroupBetInfo
	return json.Marshal(&struct {
		Alias
		CreatedAt string
	}{
		Alias:     (Alias)(d),
		CreatedAt: d.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

type SumGroupBetInfo struct {
	C       int
	BetM    int
	Win     float64
	FirstId int
}

func GetGroupBetList(curAgentId, userId, gameType, userType, state int, beginDay, endDay, name, period, GroupBetSn string, pageIndex, pageSize, FirstId int) (int, []GroupBetInfo, SumGroupBetInfo) {
	o := orm.NewOrm()
	aSumGroupBetInfo := SumGroupBetInfo{}
	arrData := make([]GroupBetInfo, 0)
	sqlArgs := make([]interface{}, 0)

	maxId := FirstId
	if FirstId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, mconst.TableName_LoBetGroupInfo)
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&maxId)
		if e != nil {
			ttLog.LogError(e)
		}
	}

	sqlWhere := ` and a.id<=? and b.user_type<>?`
	sqlArgs = append(sqlArgs, maxId, userConst.UserType_3)

	if curAgentId != 0 {
		sqlWhere += ` and b.agent_user_id=? `
		sqlArgs = append(sqlArgs, curAgentId)
	}

	if beginDay != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.created_at,'%[1]s') >= ? `, "%Y-%m-%d")
		sqlArgs = append(sqlArgs, beginDay)
	}
	if endDay != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.created_at,'%[1]s') <= ? `, "%Y-%m-%d")
		sqlArgs = append(sqlArgs, endDay)
	}

	if state != 0 {
		sqlWhere += ` and a.status=?`
		sqlArgs = append(sqlArgs, state)
	}

	if userId != 0 {
		sqlWhere += ` and a.user_id=?`
		sqlArgs = append(sqlArgs, userId)
	}

	if gameType != 0 {
		sqlWhere += ` and a.game_type=?`
		sqlArgs = append(sqlArgs, gameType)
	}

	if name != "" {
		sqlWhere += ` and locate(?,b.user_name)>0`
		sqlArgs = append(sqlArgs, name)
	}
	if period != "" {
		sqlWhere += ` and locate(?,a.str_period)>0`
		sqlArgs = append(sqlArgs, period)
	}
	if GroupBetSn != "" {
		sqlWhere += ` and a.bet_sn=?`
		sqlArgs = append(sqlArgs, GroupBetSn)
	}

	if userType != 0 {
		sqlWhere += ` and b.user_type=?`
		sqlArgs = append(sqlArgs, userType)
	}

	sqlCount := fmt.Sprintf(`select count(1) c, sum(a.bet_m) as bet_m, sum(a.win) as win from %s a, %s b where a.user_id=b.id %s`,
		mconst.TableName_LoBetGroupInfo, userConst.TableName_TtGameUser, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aSumGroupBetInfo)
	if e != nil {
		ttLog.LogError(e)
		return aSumGroupBetInfo.C, arrData, aSumGroupBetInfo
	}

	offset, _ := sqlKit.GetOffset(aSumGroupBetInfo.C, pageSize, pageIndex)
	sqlWhere = sqlWhere + ` order by a.created_at DESC, a.id DESC LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, pageSize)

	sql := fmt.Sprintf(`select a.id, a.user_id, b.user_name, a.game_name, a.str_period, a.bet_m, a.win
, a.bet_sn, a.bet_str, a.status, a.created_at  from %s a, %s b where a.user_id=b.id %s`,
		mconst.TableName_LoBetGroupInfo, userConst.TableName_TtGameUser, sqlWhere)
	_, e = o.Raw(sql, sqlArgs).QueryRows(&arrData)
	if e != nil {
		ttLog.LogError(e)
		return aSumGroupBetInfo.C, arrData, aSumGroupBetInfo
	}

	for i := 0; i < len(arrData); i++ {
		arrData[i].ArrBetDetail = getNewArrBetDetail(arrData[i].BetStr)
	}

	aSumGroupBetInfo.FirstId = maxId
	return aSumGroupBetInfo.C, arrData, aSumGroupBetInfo
}

//func GetGroupBetListFirstId(gameType, userType, state int, beginDay, endDay,name,period, GroupBetSn string, pageIndex, pageSize,FirstId int) (int, []GroupBetInfo, SumGroupBetInfo) {
//	o := orm.NewOrm()
//	aSumGroupBetInfo := SumGroupBetInfo{}
//	arrData := make([]GroupBetInfo, 0)
//	sqlArgs := make([]interface{}, 0)
//
//	maxId := FirstId
//	if FirstId==0 {
//		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, mconst.TableName_LoBetGroupInfo)
//		e := o.Raw(sqlCount, sqlArgs).QueryRow(&maxId)
//		if e!= nil {
//			ttLog.LogError(e)
//		}
//	}
//
//	sqlWhere := ` and a.id<=? and b.user_type<>?`
//	sqlArgs = append(sqlArgs, maxId, userConst.UserType_3)
//
//	if beginDay != "" {
//		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.created_at,'%[1]s') >= ? `, "%Y-%m-%d")
//		sqlArgs = append(sqlArgs, beginDay)
//	}
//	if endDay != "" {
//		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.created_at,'%[1]s') <= ? `, "%Y-%m-%d")
//		sqlArgs = append(sqlArgs, endDay)
//	}
//
//	if state != 0 {
//		sqlWhere += ` and a.status=?`
//		sqlArgs = append(sqlArgs, state)
//	}
//
//	if gameType != 0 {
//		sqlWhere += ` and a.game_type=?`
//		sqlArgs = append(sqlArgs, gameType)
//	}
//
//	if name != "" {
//		sqlWhere += ` and locate(?,b.user_name)>0`
//		sqlArgs = append(sqlArgs, name)
//	}
//	if period != "" {
//		sqlWhere += ` and locate(?,a.str_period)>0`
//		sqlArgs = append(sqlArgs, period)
//	}
//	if GroupBetSn != "" {
//		sqlWhere += ` and a.bet_sn=?`
//		sqlArgs = append(sqlArgs, GroupBetSn)
//	}
//
//	if userType != 0 {
//		sqlWhere += ` and b.user_type=?`
//		sqlArgs = append(sqlArgs, userType)
//	}
//
//	sqlCount := fmt.Sprintf(`select count(1) c, sum(a.bet_m) as bet_m, sum(a.win) as win from %s a, %s b where a.user_id=b.id %s`,
//		mconst.TableName_LoBetGroupInfo, userConst.TableName_TtGameUser, sqlWhere)
//	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aSumGroupBetInfo)
//	if e != nil {
//		ttLog.LogError(e)
//		return aSumGroupBetInfo.C, arrData, aSumGroupBetInfo
//	}
//
//	offset, _ := sqlKit.GetOffset(aSumGroupBetInfo.C, pageSize, pageIndex)
//	sqlWhere = sqlWhere + ` order by a.created_at DESC, a.id DESC LIMIT ?,?`
//	sqlArgs = append(sqlArgs, offset, pageSize)
//
//	sql := fmt.Sprintf(`select a.id, b.user_name, a.game_name, a.str_period, a.bet_m, a.win
//, a.bet_sn, a.bet_str, a.status, a.created_at  from %s a, %s b where a.user_id=b.id %s`,
//		mconst.TableName_LoBetGroupInfo, userConst.TableName_TtGameUser, sqlWhere)
//	_, e = o.Raw(sql, sqlArgs).QueryRows(&arrData)
//	if e != nil {
//		ttLog.LogError(e)
//		return aSumGroupBetInfo.C, arrData, aSumGroupBetInfo
//	}
//
//	for i := 0; i < len(arrData); i++ {
//		arrData[i].ArrBetDetail = getNewArrBetDetail(arrData[i].BetStr)
//	}
//
//	aSumGroupBetInfo.FirstId = maxId
//	return aSumGroupBetInfo.C, arrData, aSumGroupBetInfo
//}

func getNewArrBetDetail(BetStr string) []string {
	BetStr = strings.Replace(BetStr, " ", "_", -1)
	BetStr = strings.Replace(BetStr, "$", "", -1)
	arr := strings.Split(BetStr, ";")
	ArrBetDetail := make([]string, len(arr)-1)
	for i := 1; i < len(arr); i++ {
		ArrBetDetail[i-1] = arr[i] + "元"
	}
	return ArrBetDetail
}

type DetailBet struct {
	Id         int
	GroupId    int
	UserName   string
	OddsName   string
	OddsDes    string
	BetM       int
	Odds       float64
	Win        float64
	GroupBetSn string
	Status     int
	CreatedAt  time.Time
	StrPeriod  string
	ResultNums string
}

func (d DetailBet) MarshalJSON() ([]byte, error) {
	type Alias DetailBet
	return json.Marshal(&struct {
		Alias
		CreatedAt string
	}{
		Alias:     (Alias)(d),
		CreatedAt: d.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

func GetGroupDetailBetList(GroupBetId int) []DetailBet {
	arrData := make([]DetailBet, 0)
	o := orm.NewOrm()

	//	sql := fmt.Sprintf(`select a.id, a.group_id, b.user_name, a.game_name, a.str_period, a.bet_m, a.win, a.odds
	//, a.odds_name, a.status, a.created_at  from %s a, %s b where a.user_id=b.id and a.group_id=?`,
	//		mconst.TableName_LoBetInfo, userConst.TableName_TtGameUser)

	sql := fmt.Sprintf(`select a.id, a.group_id, b.user_name, a.game_name, a.str_period, a.bet_m, a.win, a.odds, c.odds_des,a.result_nums, a.group_bet_sn
		, a.odds_name, a.status, a.created_at  from %s a LEFT JOIN %s b on (a.user_id=b.id)
		LEFT JOIN %s c on (a.big_odds_type=c.big_type) where a.group_id=? and c.game_type=a.game_type`,
		mconst.TableName_LoBetInfo, userConst.TableName_TtGameUser, mconst.TableName_LoOddsInfo)
	_, e := o.Raw(sql, GroupBetId).QueryRows(&arrData)

	if e != nil {
		ttLog.LogError(e)
		return arrData
	}
	return arrData
}

func Get7DayBetYieldChart() Chart.OptionsSchart {
	n := 7
	arrLabels := make([]string, n)
	arrDatasets := make([]Chart.DatasetsSchart, 2)

	arrBet := Chart.DatasetsSchart{Label: "投注"}
	arrWin := Chart.DatasetsSchart{Label: "赢得"}
	arrBet.Data = make([]float64, n)
	arrWin.Data = make([]float64, n)

	t := time.Now()
	for i := n - 1; i >= 0; i-- {
		arrLabels[i] = t.Format("01-02")

		bet := 0
		betKey := gameCacheData.CurDayBetMKey(t)
		CacheData.GetBeegoCache().GetCache(betKey, &bet)

		win := 0.0
		winKey := gameCacheData.CurDayWinMKey(t)
		CacheData.GetBeegoCache().GetCache(winKey, &win)

		arrBet.Data[i] = float64(bet)
		arrWin.Data[i] = win
		t = t.AddDate(0, 0, -1)

	}
	arrDatasets[0] = arrBet
	arrDatasets[1] = arrWin

	aOptionsSchart := Chart.NewOptionsSchart(Chart.TypeSchart_Bar, "近7天投注收益", arrLabels, arrDatasets)
	return aOptionsSchart
}

type MainData struct {
	SaveMoney    float64
	SaveMoneying float64

	DrawMoney    float64
	DrawMoneying float64

	Gold   float64
	Profit float64
}

func GetMainData() MainData {
	aMainData := MainData{}

	type Tmp struct {
		Gold    float64
		Golding float64
	}
	aTmpSaveMoney := Tmp{}

	o := orm.NewOrm()
	sql := fmt.Sprintf(`select sum(if(a.state=?,0,a.gold)) as golding
,sum(if(a.state=?,0,a.gold)) as gold from %s a`, mconst.TableName_TtSaveMoney)
	e := o.Raw(sql, mconst.SaveMoneyState_5_OK, mconst.SaveMoneyState_1_Apply).QueryRow(&aTmpSaveMoney)
	if e != nil {
		ttLog.LogError(e)
	}

	aMainData.SaveMoney = aTmpSaveMoney.Gold
	aMainData.SaveMoneying = aTmpSaveMoney.Golding

	aTmpDrawMoney := Tmp{}

	sql = fmt.Sprintf(`select sum(if(a.state=?,0,a.gold)) as golding
,sum(if(a.state=?,0,a.gold)) as gold  from %s a`, mconst.TableName_TtDrawMoney)
	e = o.Raw(sql, mconst.DrawMoneyState_4, mconst.DrawMoneyState_1_Apply).QueryRow(&aTmpDrawMoney)
	if e != nil {
		ttLog.LogError(e)
	}

	aMainData.DrawMoney = aTmpDrawMoney.Gold
	aMainData.DrawMoneying = aTmpDrawMoney.Golding

	gold := 0.0
	sql = fmt.Sprintf(`select sum(a.gold) as gold from %s a where a.user_type=?`, userConst.TableName_TtGameUser)
	e = o.Raw(sql, userConst.UserType_1).QueryRow(&gold)
	aMainData.Gold = gold
	aMainData.Profit = aMainData.SaveMoney - aMainData.DrawMoney - aMainData.Gold

	return aMainData
}
