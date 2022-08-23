package WsxServer

import (
	"errors"
	"fmt"
	"github.com/TtMyth123/GameServer/GInstance/GTtHint"
	"github.com/TtMyth123/GameServer/LotteryServer"
	"github.com/TtMyth123/GameServer/LotteryServer/LotteryBox"
	"github.com/TtMyth123/GameServer/controllers/base/TtError"
	"github.com/TtMyth123/GameServer/models"
	"github.com/TtMyth123/GameServer/models/mconst"
	"github.com/TtMyth123/UserInfoRpc/UserRpcClient"
	"github.com/TtMyth123/kit/stringKit"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"sync"
	"time"
)

type BaseWsxServer struct {
	TimeLag time.Duration
	LotteryServer.ILotteryServer
	Name        string
	GameType    int
	PreOpenTime time.Time
	PreIssue    string

	StopBetTime int

	mpOdds       map[int]float64
	curAwardInfo models.LoAwardInfo

	mUserRpcClient       *UserRpcClient.RpcClient
	RandNewAwardInfoLock sync.RWMutex
}

func (this *BaseWsxServer) DelAfData(lottery_str string) error {
	o := orm.NewOrm()
	sql := fmt.Sprintf(`delete from %s where game_type=? and str_period>?`, mconst.TableName_LoBetInfo)
	_, e := o.Raw(sql, this.GameType, lottery_str).Exec()
	if e != nil {
		return e
	}

	sql = fmt.Sprintf(`delete from %s where game_type=? and str_period>?`, mconst.TableName_LoBetGroupInfo)
	_, e = o.Raw(sql, this.GameType, lottery_str).Exec()
	if e != nil {
		return e
	}

	sql = fmt.Sprintf(`delete from %s where game_type=? and lottery_str>=?`, mconst.TableName_LoAwardInfo)
	_, e = o.Raw(sql, this.GameType, lottery_str).Exec()
	if e != nil {
		return e
	}

	return nil
}

func (this *BaseWsxServer) PreBet(betInfo LotteryBox.BetInfo) error {
	if this.curAwardInfo.NextLotteryStr == "" {
		return errors.New(GTtHint.GetTtHint().GetHint("未开始，不能投注"))
	}
	if betInfo.StrLotteryNum != this.curAwardInfo.NextLotteryStr {
		return TtError.New("期号不正确，不能投注。当前期号%s", this.curAwardInfo.NextLotteryStr)
	}

	return nil
}

func (this *BaseWsxServer) PreNewAwardInfo(newLoAwardInfo models.LoAwardInfo) error {
	ttLog.LogDebug(stringKit.GetJsonStr(newLoAwardInfo))
	this.AddAwardInfo(&newLoAwardInfo)
	this.PreIssue = this.curAwardInfo.LotteryStr
	this.curAwardInfo = newLoAwardInfo
	this.PreOpenTime = time.Now()

	return nil
}

func (this *BaseWsxServer) AddAwardInfo(newLoAwardInfo *models.LoAwardInfo) {
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

func (this *BaseWsxServer) getBetGroupInfos(StrLotteryNum string) []*models.LoBetGroupInfo {
	o := orm.NewOrm()
	sql := fmt.Sprintf(`select a.* from %s a where a.str_period=? and a.status=? and a.game_type=?`, mconst.TableName_LoBetGroupInfo)

	arrLoBetGroupInfo := make([]*models.LoBetGroupInfo, 0)
	_, e := o.Raw(sql, StrLotteryNum, mconst.Bet_Status_1, this.GameType).QueryRows(&arrLoBetGroupInfo)
	if e != nil {
		ttLog.LogError(e)
	}

	return arrLoBetGroupInfo
}

func (this *BaseWsxServer) getNotOpenBetGroupInfos(StrLotteryNum string) []*models.LoBetGroupInfo {
	o := orm.NewOrm()
	sql := fmt.Sprintf(`select a.* from %s a where a.str_period<? and a.status=? and a.game_type=?`, mconst.TableName_LoBetGroupInfo)

	arrLoBetGroupInfo := make([]*models.LoBetGroupInfo, 0)
	_, e := o.Raw(sql, StrLotteryNum, mconst.Bet_Status_1, this.GameType).QueryRows(&arrLoBetGroupInfo)
	if e != nil {
		ttLog.LogError(e)
	}

	return arrLoBetGroupInfo
}

func (this *BaseWsxServer) getBetInfos(StrLotteryNum string) map[int][]*models.LoBetInfo {
	mpBetInfo := make(map[int][]*models.LoBetInfo)
	o := orm.NewOrm()
	sql := fmt.Sprintf(`select a.* from %s a where a.str_period=? and a.status=? and a.game_type=?`, mconst.TableName_LoBetInfo)

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

func (this *BaseWsxServer) getNotOpenBetInfos(StrLotteryNum string) map[int][]*models.LoBetInfo {
	mpBetInfo := make(map[int][]*models.LoBetInfo)
	o := orm.NewOrm()
	sql := fmt.Sprintf(`select a.* from %s a where a.str_period<? and a.status=? and a.game_type=?`, mconst.TableName_LoBetInfo)

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
