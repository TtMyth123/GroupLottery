package UscServer

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"sync"
	"time"
	"ttmyth123/GroupLottery/GameServer/GInstance/GTtHint"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/UscResultServer/UscBox"
	"ttmyth123/GroupLottery/GameServer/LotteryServer/LotteryBox"
	"ttmyth123/GroupLottery/GameServer/controllers/base/TtError"
	"ttmyth123/GroupLottery/GameServer/models"
	"ttmyth123/GroupLottery/GameServer/models/mconst"
	"ttmyth123/GroupLottery/UserInfoRpc/UserRpcClient"
	"ttmyth123/kit/stringKit"
	"ttmyth123/kit/ttLog"
)

type BaseUscServer struct {
	Name        string
	GameType    int
	GameIndex   int
	PreOpenTime time.Time
	PreIssue    int64

	StopBetTime int

	mpOdds map[int]UscBox.OddsInfo
	//curAwardInfo models.LoAwardInfo
	curAwardInfo UscBox.UscAwardInfo

	mUserRpcClient       *UserRpcClient.RpcClient
	RandNewAwardInfoLock sync.RWMutex
}

func (this *BaseUscServer) DelAfData(lottery_str string) error {
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

func (this *BaseUscServer) PreBet(betInfo LotteryBox.BetInfo) error {
	if this.curAwardInfo.LotteryNum == 0 {
		return errors.New(GTtHint.GetTtHint().GetHint("未开始，不能投注"))
	}
	if betInfo.LotteryNum != this.curAwardInfo.LotteryNum {
		return TtError.New("期号不正确，不能投注。当前期号%d", this.curAwardInfo.LotteryNum)
	}

	if this.curAwardInfo.Countdown < this.StopBetTime {
		return TtError.New("正在开奖,暂时无法下注.")
	}

	return nil
}

func (this *BaseUscServer) PreNewAwardInfo(newLoAwardInfo models.LoAwardInfo, aUscAwardInfo UscBox.UscAwardInfo) error {
	ttLog.LogDebug(stringKit.GetJsonStr(newLoAwardInfo))
	this.AddAwardInfo(&newLoAwardInfo)
	this.PreIssue = this.curAwardInfo.NewNum
	this.curAwardInfo = aUscAwardInfo
	this.PreOpenTime = time.Now()

	return nil
}

func (this *BaseUscServer) AddAwardInfo(newLoAwardInfo *models.LoAwardInfo) {
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

func (this *BaseUscServer) getBetGroupInfos(StrLotteryNum string) []*models.LoBetGroupInfo {
	o := orm.NewOrm()
	sql := fmt.Sprintf(`select a.* from %s a where a.str_period=? and a.status=? and a.game_type=?`, mconst.TableName_LoBetGroupInfo)

	arrLoBetGroupInfo := make([]*models.LoBetGroupInfo, 0)
	_, e := o.Raw(sql, StrLotteryNum, mconst.Bet_Status_1, this.GameType).QueryRows(&arrLoBetGroupInfo)
	if e != nil {
		ttLog.LogError(e)
	}

	return arrLoBetGroupInfo
}

func (this *BaseUscServer) getNotOpenBetGroupInfos(StrLotteryNum string) []*models.LoBetGroupInfo {
	o := orm.NewOrm()
	sql := fmt.Sprintf(`select a.* from %s a where a.str_period<? and a.status=? and a.game_type=?`, mconst.TableName_LoBetGroupInfo)

	arrLoBetGroupInfo := make([]*models.LoBetGroupInfo, 0)
	_, e := o.Raw(sql, StrLotteryNum, mconst.Bet_Status_1, this.GameType).QueryRows(&arrLoBetGroupInfo)
	if e != nil {
		ttLog.LogError(e)
	}

	return arrLoBetGroupInfo
}

func (this *BaseUscServer) getBetInfos(StrLotteryNum string) map[int][]*models.LoBetInfo {
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

func (this *BaseUscServer) getNotOpenBetInfos(StrLotteryNum string) map[int][]*models.LoBetInfo {
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
