package userBll

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"ttmyth123/GroupLottery/GameServer/CacheData"
	"ttmyth123/GroupLottery/GameServer/GConfig"
	"ttmyth123/GroupLottery/GameServer/GInstance"
	"ttmyth123/GroupLottery/UserInfoRpc/models/mconst"
	"ttmyth123/kit"
	"ttmyth123/kit/sqlKit"
	"ttmyth123/kit/timeKit"
	"ttmyth123/kit/ttLog"
)

type UserInfo struct {
	Id            int
	ReferrerCode  int       //推荐码
	HeadImgurl    string    `orm:"size(512)"` //头像
	UserName      string    `orm:"size(512)"` //用户名
	Nickname      string    `orm:"size(256)"`
	Gold          float64   //当前余额(金币)
	SumXmBet      float64   //全部洗码金额
	SumWin        float64   //全部赢得金额
	UserType      int       //用户类型，1：一般玩家，2：业务玩家，3：游客
	PUserName     string    //上级用户名
	IsReferrer    int       //是否推手
	SumTime       time.Time //统计时间
	SumDrawMoney  float64   //累计提现金币
	SumSaveMoney  float64   //累计充值金币
	MemberLevel   int
	YjGold        float64 //佣金
	CurDayLossWin float64 //当前输赢
}

func (d UserInfo) MarshalJSON() ([]byte, error) {
	type Alias UserInfo

	StrTime := ""
	TimeLag := GConfig.GetGConfig().TimeLag
	if TimeLag == 0 {
		StrTime = d.SumTime.Format(timeKit.DateTimeLayout)
	} else {
		StrTime = d.SumTime.Add(TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
	}

	return json.Marshal(&struct {
		Alias
		SumTime string
	}{
		Alias:   (Alias)(d),
		SumTime: StrTime,
	})
}

func GetUserInfo(UserId int) (UserInfo, error) {
	aUserInfo := UserInfo{}
	aUser, e := GInstance.GetUserRpcClient().GetUser(UserId)
	if e != nil {
		return aUserInfo, e
	}
	aUserInfo.ReferrerCode = aUser.ReferrerCode
	aUserInfo.UserName = aUser.UserName
	aUserInfo.HeadImgurl = aUser.HeadImgurl
	aUserInfo.Nickname = aUser.Nickname
	aUserInfo.Gold = kit.Decimal(aUser.Gold)
	aUserInfo.SumXmBet = aUser.SumXmBet
	aUserInfo.SumWin = aUser.SumWin
	aUserInfo.UserType = aUser.UserType
	aUserInfo.SumTime = aUser.SumTime
	aUserInfo.IsReferrer = aUser.IsReferrer
	aUserInfo.SumDrawMoney = aUser.SumDrawMoney
	aUserInfo.SumSaveMoney = aUser.SumSaveMoney
	aUserInfo.Id = aUser.Id
	aUserInfo.YjGold = kit.Decimal(aUser.Rebate)
	aUserInfo.MemberLevel = aUser.MemberLevel
	aUserInfo.CurDayLossWin = CacheData.GetOneDayLossWin(time.Now(), aUser.Id)

	aPUser, e := GInstance.GetUserRpcClient().GetUser(UserId)
	if e == nil {
		aUserInfo.PUserName = aPUser.UserName
	}

	return aUserInfo, e
}

type MyReferrerCode struct {
	ReferrerCode int `json:"1"` //推荐码
}

func GetMyReferrerCode(UserId int) (MyReferrerCode, error) {

	aUser, e := GInstance.GetUserRpcClient().GetUser(UserId)
	if e != nil {
		return MyReferrerCode{}, e
	}
	aMyReferrerCode := MyReferrerCode{ReferrerCode: aUser.ReferrerCode}
	return aMyReferrerCode, e
}

type SubUserInfo struct {
	Id       int     `json:"1"`
	UserName string  `json:"2"`
	XmBet    float64 `json:"3"` //洗码
	Rebate   float64 `json:"4"` //返佣
}
type JuniorUserInfo struct {
	LowerCount    int           `json:"1"` //直接下级会员数
	AllLowerCount int           `json:"2"` //全部的下级有会员数
	XmBet         float64       `json:"3"` //累计洗码
	Rebate        float64       `json:"4"` //累计返佣
	DataList      []SubUserInfo `json:"5"`
	Pid           int           //父ID
	PageTotal     int
	LastId        int
}
type GroupJuniorUserInfo struct {
}

func GetJuniorUserInfo(RootUserId, UserId, pageIndex, pageSize, LastId int) (JuniorUserInfo, error) {
	if pageIndex < 1 {
		pageIndex = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	arrData := make([]SubUserInfo, 0)
	aUserInfo := JuniorUserInfo{}
	aUser, e := GInstance.GetUserRpcClient().GetUser(UserId)
	if e != nil {
		return aUserInfo, e
	}
	aUserInfo.LowerCount = aUser.LowerCount
	aUserInfo.AllLowerCount = aUser.AllLowerCount

	aUserInfo.Pid = aUser.Pid
	if RootUserId == UserId {
		aUserInfo.Pid = 0
	}

	//aUserInfo.SumRebate = aUser.Sum2Rebate
	//aUserInfo.Rebate = aUser.Rebate

	o := orm.NewOrm()
	sqlArgs := make([]interface{}, 0)

	maxId := LastId
	if LastId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, mconst.TableName_TtGameUser)
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&maxId)
		if e != nil {
			ttLog.LogError(e)
		}
	}
	type Tmp struct {
		Xm     float64
		Rebate float64
	}
	aTmp := Tmp{}
	sqlCount := fmt.Sprintf(`select sum(b.gold) xm,sum(b.rebate) rebate from tt_rebate_info b where b.rebate_type=? and b.user_id=? and b.user_sid<=?`)
	e = o.Raw(sqlCount, mconst.Rebate_01_Guess, RootUserId, maxId).QueryRow(&aTmp)

	offset, PageTotal := sqlKit.GetOffset(aUserInfo.LowerCount, pageSize, pageIndex)
	sql := fmt.Sprintf(`select  a.id, a.user_name, bb.xm_bet, bb.rebate from tt_game_user a
LEFT JOIN
(select c.id, sum(b.gold) as xm_bet,sum(b.rebate) rebate from tt_rebate_info b, tt_game_user c where b.rebate_type=? and b.user_id = ? and  b.user_sid = c.id and c.pid=? group by c.id) bb
on a.id=bb.id where a.pid=? and a.id<=? order by a.created_at DESC LIMIT ?,?`)
	_, e = o.Raw(sql, mconst.Rebate_01_Guess, RootUserId, UserId, UserId, maxId, offset, pageSize).QueryRows(&arrData)
	if e != nil {
		ttLog.LogError(e)
		return aUserInfo, e
	}
	aUserInfo.Rebate = aTmp.Rebate
	aUserInfo.XmBet = aTmp.Xm
	aUserInfo.DataList = arrData
	aUserInfo.LastId = maxId
	aUserInfo.PageTotal = PageTotal
	return aUserInfo, e
}

type JuniorUserInfoWsx struct {
	LowerCount    int           `json:"1"` //直接下级会员数
	AllLowerCount int           `json:"2"` //全部的下级有会员数
	SumRebate     float64       `json:"3"` //累计返佣
	Rebate        float64       `json:"4"` //累计洗码
	DataList      []SubUserInfo `json:"5"`
	Pid           int           //父ID
	PageTotal     int
	LastId        int
}

/**
越南彩用
*/
func GetJuniorUserInfoWsx(RootUserId, UserId, pageIndex, pageSize, LastId int) (JuniorUserInfoWsx, error) {
	if pageIndex < 1 {
		pageIndex = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	arrData := make([]SubUserInfo, 0)
	aUserInfo := JuniorUserInfoWsx{}
	aUser, e := GInstance.GetUserRpcClient().GetUser(UserId)
	if e != nil {
		return aUserInfo, e
	}
	aUserInfo.LowerCount = aUser.LowerCount
	aUserInfo.AllLowerCount = aUser.AllLowerCount
	aUserInfo.Rebate = aUser.Rebate

	aUserInfo.Pid = aUser.Pid
	if RootUserId == UserId {
		aUserInfo.Pid = 0
	}

	o := orm.NewOrm()
	sqlArgs := make([]interface{}, 0)

	maxId := LastId
	if LastId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, mconst.TableName_TtGameUser)
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&maxId)
		if e != nil {
			ttLog.LogError(e)
		}
	}
	type Tmp struct {
		Rebate float64
	}
	aTmp := Tmp{}
	sqlCount := fmt.Sprintf(`select sum(b.rebate) rebate  from tt_rebate_info b where b.rebate_type=? and b.user_id=? and b.user_sid<=?`)
	e = o.Raw(sqlCount, mconst.Rebate_01_Guess, RootUserId, maxId).QueryRow(&aTmp)

	offset, PageTotal := sqlKit.GetOffset(aUserInfo.LowerCount, pageSize, pageIndex)
	sql := fmt.Sprintf(`select  a.id, a.user_name, bb.xm_bet, bb.rebate from tt_game_user a
LEFT JOIN
(select c.id, sum(b.gold) as xm_bet,sum(b.rebate) rebate from tt_rebate_info b, tt_game_user c where b.rebate_type=? and b.user_id = ? and  b.user_sid = c.id and c.pid=? group by c.id) bb
on a.id=bb.id where a.pid=? and a.id<=? order by a.created_at DESC LIMIT ?,?`)
	_, e = o.Raw(sql, mconst.Rebate_01_Guess, RootUserId, UserId, UserId, maxId, offset, pageSize).QueryRows(&arrData)
	if e != nil {
		ttLog.LogError(e)
		return aUserInfo, e
	}
	aUserInfo.Rebate = aTmp.Rebate
	aUserInfo.SumRebate = aTmp.Rebate
	aUserInfo.DataList = arrData
	aUserInfo.LastId = maxId
	aUserInfo.PageTotal = PageTotal
	return aUserInfo, e
}
