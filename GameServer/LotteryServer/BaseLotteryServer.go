package LotteryServer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TtMyth123/GameServer/GConfig"
	"github.com/TtMyth123/GameServer/GInstance/GTtHint"
	"github.com/TtMyth123/GameServer/models/mconst"
	"github.com/TtMyth123/kit/sqlKit"
	"github.com/TtMyth123/kit/timeKit"
	"github.com/astaxie/beego/orm"
	"time"
)

type BaseLotteryServer struct {
	ILotteryServer
}

type GroupUserBetDataInfo struct {
	Id         int               `json:"-"`
	BetSn      string            `json:"3"`
	BetM       int               `json:"4"`
	Win        float64           `json:"5"`
	LotteryNum string            `json:"2"`
	CreatedAt  time.Time         `json:"1"`
	Status     int               `json:"7"`
	GameType   int               `json:"-"`
	GameName   string            `json:"6"`
	BetList    []UserBetDataInfo `json:"8"`
}

func (d GroupUserBetDataInfo) MarshalJSON() ([]byte, error) {
	type Alias GroupUserBetDataInfo

	newGameName := d.GameName
	strStatus := mconst.GetBetStatusName(d.Status)
	if GConfig.GetGConfig().IsI18n {
		newGameName = GTtHint.GetTtHint().GetHint(d.GameName)
		strStatus = GTtHint.GetTtHint().GetHint(strStatus)
	}

	StrTime := ""

	TimeLag := GConfig.GetGConfig().TimeLag
	if TimeLag == 0 {
		StrTime = d.CreatedAt.Format(timeKit.DateTimeLayout)
	} else {
		StrTime = d.CreatedAt.Add(TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
	}

	newWin := d.Win - float64(d.BetM)
	if d.Status == mconst.Bet_Status_1 {
		newWin = 0
	}
	return json.Marshal(&struct {
		Alias
		Status    string  `json:"7"`
		GameName  string  `json:"6"`
		CreatedAt string  `json:"1"`
		Win       float64 `json:"5"`
	}{
		Alias:     (Alias)(d),
		Status:    strStatus,
		GameName:  newGameName,
		Win:       newWin,
		CreatedAt: StrTime,
	})
}

type SumGroupUserBetDataInfo struct {
	C    int
	BetM int
	Win  float64
}

func (d SumGroupUserBetDataInfo) MarshalJSON() ([]byte, error) {
	type Alias SumGroupUserBetDataInfo

	newWin := d.Win - float64(d.BetM)
	return json.Marshal(&struct {
		Alias
		Win float64
	}{
		Alias: (Alias)(d),
		Win:   newWin,
	})
}

type UserBetDataInfo struct {
	Id         int       `json:"-"`
	GroupId    int       `json:"-"` //主ID
	OddsType   int       `json:"-"`
	BetM       int       `json:"2"`
	OddsName   string    `json:"1"`
	Status     int       `json:"5"`
	Win        float64   `json:"4"`
	LotteryNum string    `json:"-"`
	Odds       float64   `json:"3"`
	CreatedAt  time.Time `json:"-"`
	BetSn      string    `json:"-"`
	GameType   int       `json:"-"`
	GameName   string    `json:"-"`
}

func (d UserBetDataInfo) MarshalJSON() ([]byte, error) {
	type Alias UserBetDataInfo
	newWin := d.Win
	switch d.GameType {
	case mconst.GameType_Wsx_201, mconst.GameType_Wsx_202, mconst.GameType_Wsx_203:
		if d.Win == 0 {
			newWin = -float64(d.BetM)
		} else {
			newWin = d.Win
		}
	default:

	}
	return json.Marshal(&struct {
		Alias
		Win float64 `json:"4"`
	}{
		Alias: (Alias)(d),
		Win:   newWin,
	})
}

func getBetRecordList(gameType int, UserId, Status, PageIndex, PageSize, LastId int, StrBeginDay, StrEndDay string) (int, []*GroupUserBetDataInfo, SumGroupUserBetDataInfo) {
	arrGroupUserBetDataInfo := make([]*GroupUserBetDataInfo, 0)
	aSumGroupUserBetDataInfo := SumGroupUserBetDataInfo{}
	if PageIndex < 1 {
		PageIndex = 1
	}
	if PageSize <= 0 {
		PageSize = 10
	}

	o := orm.NewOrm()
	NewLastId := LastId
	if LastId == 0 {
		sql := fmt.Sprintf(`select max(id) from %s`, mconst.TableName_LoBetGroupInfo)
		o.Raw(sql).QueryRow(&NewLastId)
	}

	sqlArgs := make([]interface{}, 0)
	sqlSubArgs := make([]interface{}, 0)
	sqlWhere := ` where a.user_id=?`
	sqlSubWhere := ` where a.user_id=?`

	sqlArgs = append(sqlArgs, UserId)
	sqlSubArgs = append(sqlSubArgs, UserId)

	if NewLastId != 0 {
		sqlWhere += " and a.id<=?"
		sqlArgs = append(sqlArgs, NewLastId)
	}

	if Status != 0 {
		sqlWhere += " and a.status=?"
		sqlSubWhere += " and a.status=?"

		sqlArgs = append(sqlArgs, Status)
		sqlSubArgs = append(sqlSubArgs, Status)
	}
	if gameType != 0 {
		sqlWhere += " and a.game_type=?"
		sqlSubWhere += " and a.game_type=?"

		sqlArgs = append(sqlArgs, gameType)
		sqlSubArgs = append(sqlSubArgs, gameType)
	}

	if StrBeginDay != "" {
		sqlWhere += fmt.Sprintf(` and date_format(a.created_at,'%[1]s') >= ? `, "%Y-%m-%d")
		sqlSubWhere += fmt.Sprintf(` and date_format(a.created_at,'%[1]s') >= ? `, "%Y-%m-%d")

		sqlArgs = append(sqlArgs, StrBeginDay)
		sqlSubArgs = append(sqlSubArgs, StrBeginDay)
	}

	if StrEndDay != "" {
		sqlWhere += fmt.Sprintf(` and date_format(a.created_at,'%[1]s') <= ? `, "%Y-%m-%d")
		sqlSubWhere += fmt.Sprintf(` and date_format(a.created_at,'%[1]s') <= ? `, "%Y-%m-%d")

		sqlArgs = append(sqlArgs, StrEndDay)
		sqlSubArgs = append(sqlSubArgs, StrEndDay)
	}
	sqlCount := fmt.Sprintf(`select count(1) c, sum(a.bet_m) as bet_m, sum(a.win) as win 
from %s a %s`, mconst.TableName_LoBetGroupInfo, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aSumGroupUserBetDataInfo)
	if e != nil {
		return NewLastId, arrGroupUserBetDataInfo, aSumGroupUserBetDataInfo
	}

	offset, _ := sqlKit.GetOffset(aSumGroupUserBetDataInfo.C, PageSize, PageIndex)
	sqlWhere = sqlWhere + ` order by a.created_at DESC LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, PageSize)
	sql := fmt.Sprintf(`select a.id, a.bet_sn, a.win, a.str_period as lottery_num
, a.bet_m, a.created_at, a.status, a.game_name from %s a %s`, mconst.TableName_LoBetGroupInfo, sqlWhere)
	_, e = o.Raw(sql, sqlArgs).QueryRows(&arrGroupUserBetDataInfo)
	if e != nil {
		return NewLastId, arrGroupUserBetDataInfo, aSumGroupUserBetDataInfo
	}

	arrUserBetDataInfo := make([]UserBetDataInfo, 0)
	sql = fmt.Sprintf(`select a.id, a.group_id,a.odds_type, a.bet_m,a.odds_name
,a.status,  a.win, a.str_period as lottery_num, a.odds, a.created_at, a.bet_sn,a.game_type, a.game_name 
from %s a %s`, mconst.TableName_LoBetInfo, sqlSubWhere)
	_, e = o.Raw(sql, sqlSubArgs).QueryRows(&arrUserBetDataInfo)
	if e != nil {
		return NewLastId, arrGroupUserBetDataInfo, aSumGroupUserBetDataInfo
	}
	mpBets := make(map[int][]UserBetDataInfo)

	for _, bet := range arrUserBetDataInfo {
		if arrBet, ok := mpBets[bet.GroupId]; ok {
			arrBet = append(arrBet, bet)
			mpBets[bet.GroupId] = arrBet
		} else {
			arrBet := []UserBetDataInfo{bet}
			mpBets[bet.GroupId] = arrBet
		}
	}

	for i := 0; i < len(arrGroupUserBetDataInfo); i++ {
		arrGroupUserBetDataInfo[i].BetList = mpBets[arrGroupUserBetDataInfo[i].Id]
	}
	return NewLastId, arrGroupUserBetDataInfo, aSumGroupUserBetDataInfo
}

func GetBetRecordList(gameType, UserId, Status, PageIndex, PageSize, LastId int,
	StrBeginDay, StrEndDay string) (map[string]interface{}, error) {
	mpData := make(map[string]interface{})

	NewLastId, arrGroupUserBetDataInfo, aSumGroupUserBetDataInfo := getBetRecordList(gameType, UserId, Status, PageIndex, PageSize, LastId, StrBeginDay, StrEndDay)
	mpData["LastId"] = NewLastId
	mpData["BetList"] = arrGroupUserBetDataInfo
	mpData["GroupData"] = aSumGroupUserBetDataInfo
	return mpData, nil
}

func GetBetOrder(BetOrder string) (interface{}, error) {
	aGroupUserBetDataInfo := getBetOrder(BetOrder)
	return aGroupUserBetDataInfo, nil
}

func getBetOrder(BetOrder string) GroupUserBetDataInfo {
	aGroupUserBetDataInfo := GroupUserBetDataInfo{}

	o := orm.NewOrm()
	sqlArgs := make([]interface{}, 0)
	sqlWhere := ` where a.bet_sn=?`
	sqlArgs = append(sqlArgs, BetOrder)

	sql := fmt.Sprintf(`select a.id, a.bet_sn, a.win, a.str_period as lottery_num
, a.bet_m, a.created_at, a.status, a.game_name from %s a %s`, mconst.TableName_LoBetGroupInfo, sqlWhere)
	e := o.Raw(sql, sqlArgs).QueryRow(&aGroupUserBetDataInfo)
	if e != nil {
		return aGroupUserBetDataInfo
	}

	arrUserBetDataInfo := make([]UserBetDataInfo, 0)
	sql = fmt.Sprintf(`select a.id, a.group_id,a.odds_type, a.bet_m,a.odds_name
,a.status,  a.win, a.str_period as lottery_num, a.odds, a.created_at, a.bet_sn,a.game_type, a.game_name 
from %s a where a.group_id=?`, mconst.TableName_LoBetInfo)
	_, e = o.Raw(sql, aGroupUserBetDataInfo.Id).QueryRows(&arrUserBetDataInfo)
	if e != nil {
		return aGroupUserBetDataInfo
	}
	aGroupUserBetDataInfo.BetList = arrUserBetDataInfo

	return aGroupUserBetDataInfo
}

func CurDayLoseWin(gameType, UserId int) (interface{}, error) {
	mpData := make(map[string]interface{})
	mpData["CurBetM"] = 0
	mpData["CurWinM"] = 0
	mpData["CurGameBetM"] = 0
	mpData["CurGameWinM"] = 0

	return mpData, nil
}

func (this *BaseLotteryServer) SetStopBetHint(StopBetHint string) error {
	return errors.New("SetStopBetHint 没有这个接口")
}

func (this *BaseLotteryServer) GetHistoryLotteryByDay(UserId int, StrDay string) (interface{}, error) {
	return nil, errors.New("GetHistoryLotteryByDay 没有这个接口")
}

func (this *BaseLotteryServer) GetHistoryFTNum(PageSize int, LastId int) (interface{}, error) {
	return nil, errors.New("GetHistoryFTNum 没有这个接口")
}

func (this *BaseLotteryServer) GetHistoryFTNumBy48(LastId int) (interface{}, error) {
	return nil, errors.New("GetHistoryFTNumBy48 没有这个接口")
}

func (this *BaseLotteryServer) GetHistoryResultByPeriod(period, Count int) (interface{}, error) {
	return nil, errors.New("GetHistoryResultByPeriod 没有这个接口")
}
func (this *BaseLotteryServer) SetAwardInfo(LotteryAward, LotteryNum string) (interface{}, error) {
	return nil, errors.New("SetAwardInfo 没有这个接口")
}
