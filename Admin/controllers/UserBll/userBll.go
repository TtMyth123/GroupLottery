package UserBll

import (
	"encoding/json"
	"fmt"
	"github.com/TtMyth123/Admin/GInstance"
	"github.com/TtMyth123/Admin/OtherServer/httpGameServer"
	"github.com/TtMyth123/UserInfoRpc/GData/gBox"
	"github.com/TtMyth123/UserInfoRpc/models"
	"github.com/TtMyth123/UserInfoRpc/models/mconst"
	"github.com/TtMyth123/kit/pwdKit"
	"github.com/TtMyth123/kit/sqlKit"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"time"
)

type GameUser struct {
	Id            int
	UserName      string `orm:"size(512)"`
	Nickname      string `orm:"size(256)"`
	LowerCount    int    //直接下级会员数
	AllLowerCount int    //全部的下级有会员数
	ReferrerCode  int    //推荐码
	State         int
	Gold          float64 //当前余额(金币)
	Rebate        float64 //可提佣金
	SumSaveMoney  float64 //累计上分金币
	MaxSaveMoney  float64 //最大上分金币
	SumDrawMoney  float64 //累计下分金币
	MaxDrawMoney  float64 //最大下分金币
	SumBet        float64 //累计投注
	SumWin        float64 //累计赢得

	CreatedAt time.Time
	LoginTime time.Time
}

func (d GameUser) MarshalJSON() ([]byte, error) {
	type Alias GameUser
	return json.Marshal(&struct {
		Alias
		CreatedAt string
		LoginTime string
	}{
		Alias:     (Alias)(d),
		CreatedAt: d.CreatedAt.Format("2006-01-02 15:04:05"),
		LoginTime: d.LoginTime.Format("2006-01-02 15:04:05"),
	})
}

type GroupUser struct {
	C            int
	Gold         float64 //当前余额(金币)
	SumSaveMoney float64 //累计上分金币
	MaxSaveMoney float64 //最大上分金币
	SumDrawMoney float64 //累计下分金币
	MaxDrawMoney float64 //最大下分金币
	SumBet       float64 //累计投注
	SumWin       float64 //累计赢得
	FirstId      int
}

/**
获取用户列表
*/
func GetGameUserList(curAgentId, userId, userType, state int, name string, pageIndex, pageSize, FirstId int) (int, []GameUser, GroupUser) {
	o := orm.NewOrm()
	aGroupUser := GroupUser{}
	arrData := make([]GameUser, 0)
	sqlArgs := make([]interface{}, 0)

	maxId := FirstId
	if FirstId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, "tt_game_user")
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&maxId)
		if e != nil {
			ttLog.LogError(e)
		}
	}

	sqlWhere := ` where a.id<=? and a.user_type<>?`
	sqlArgs = append(sqlArgs, maxId, mconst.UserType_3)
	if state != 0 {
		sqlWhere += ` and a.state=?`
		sqlArgs = append(sqlArgs, state)
	}

	if curAgentId != 0 {
		sqlWhere += ` and a.agent_user_id=? `
		sqlArgs = append(sqlArgs, curAgentId)
	}

	if userType != 0 {
		sqlWhere += ` and a.user_type=?`
		sqlArgs = append(sqlArgs, userType)
	}

	if userId != 0 {
		sqlWhere += ` and a.id=?`
		sqlArgs = append(sqlArgs, userId)
	}

	if name != "" {
		sqlWhere += ` and locate(?,a.user_name)>0`
		sqlArgs = append(sqlArgs, name)
	}
	//------------------------------------
	sqlCount := fmt.Sprintf(`select count(1) as c
,sum(a.gold) as gold, sum(a.sum_save_money+a.sum_add_money) as sum_save_money, max(a.max_save_money) as max_save_money
,sum(a.sum_draw_money+a.sum_dec_money) as sum_draw_money, max(a.max_draw_money) as max_draw_money 
,sum(a.sum_bet) as sum_bet, sum(a.sum_win) as sum_win, sum(a.rebate) as rebate
from tt_game_user a  %s`, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupUser)
	if e != nil {
		return 0, arrData, aGroupUser
	}

	offset, _ := sqlKit.GetOffset(aGroupUser.C, pageSize, pageIndex)
	sqlWhere = sqlWhere + ` order by a.created_at,a.id DESC LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, pageSize)

	sql := fmt.Sprintf(`select 
a.id, a.user_name, a.nickname, a.lower_count, a.all_lower_count, a.referrer_code, a.state
, a.gold, (a.sum_save_money+a.sum_add_money) as sum_save_money, a.max_save_money, (a.sum_draw_money+a.sum_dec_money) as sum_draw_money, a.max_draw_money
, a.sum_bet, a.sum_win, a.created_at, a.login_time, a.rebate
from tt_game_user a %s `, sqlWhere)
	_, e = o.Raw(sql, sqlArgs).QueryRows(&arrData)
	if e != nil {
		return 0, arrData, aGroupUser
	}

	aGroupUser.FirstId = maxId

	return aGroupUser.C, arrData, aGroupUser
}

/**
获取下级用户列表
*/
func GetJuniorGameUserList(pid, userType, state int, name string, pageIndex, pageSize, FirstId int) (int, []GameUser, GroupUser) {
	o := orm.NewOrm()
	aGroupUser := GroupUser{}
	arrData := make([]GameUser, 0)
	sqlArgs := make([]interface{}, 0)

	maxId := FirstId
	if FirstId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, "tt_game_user")
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&maxId)
		if e != nil {
			ttLog.LogError(e)
		}
	}

	sqlWhere := ` where a.id<=? and a.pid=? `
	sqlArgs = append(sqlArgs, maxId, pid)

	if state != 0 {
		sqlWhere = sqlWhere + ` and a.state=? `
		sqlArgs = append(sqlArgs, state)
	}
	if name != "" {
		sqlWhere = sqlWhere + ` and locate(?,a.user_name)>0 `
		sqlArgs = append(sqlArgs, name)
	}

	//------------------------------------
	sqlCount := fmt.Sprintf(`select count(1) as c 
,sum(a.gold) as gold, sum(a.sum_save_money+a.sum_add_money) as sum_save_money, max(a.max_save_money) as max_save_money
,sum(a.sum_draw_money+a.sum_dec_money) as sum_draw_money, max(a.max_draw_money) as max_draw_money 
,sum(a.sum_bet) as sum_bet, sum(a.sum_win) as sum_win, sum(a.rebate) as rebate 
from tt_game_user a %s`, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupUser)
	if e != nil {
		return 0, arrData, aGroupUser
	}

	offset, _ := sqlKit.GetOffset(aGroupUser.C, pageSize, pageIndex)
	sqlWhere = sqlWhere + ` order by a.created_at DESC LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, pageSize)

	sql := fmt.Sprintf(`select 
a.id, a.user_name, a.nickname, a.lower_count, a.all_lower_count, a.referrer_code, a.state
, a.gold, (a.sum_save_money+a.sum_add_money) as sum_save_money, a.max_save_money,  (a.sum_draw_money+a.sum_dec_money) as sum_draw_money, a.max_draw_money
, a.sum_bet, a.sum_win, a.created_at, a.login_time, a.rebate
from tt_game_user a %s `, sqlWhere)
	_, e = o.Raw(sql, sqlArgs).QueryRows(&arrData)
	if e != nil {
		return 0, arrData, aGroupUser
	}
	aGroupUser.FirstId = maxId
	return aGroupUser.C, arrData, aGroupUser
}

/**
获取下级用户列表
*/
func GetJuniorGameUserListFirstId(pid, userType, state int, name string, pageIndex, pageSize, FirstId int) (int, []GameUser, GroupUser) {
	o := orm.NewOrm()
	aGroupUser := GroupUser{}
	arrData := make([]GameUser, 0)
	sqlArgs := make([]interface{}, 0)
	maxId := FirstId
	if FirstId == 0 {
		sqlCount := fmt.Sprintf(`select max(id) as maxid from tt_game_user a `)
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&maxId)
		if e != nil {
			ttLog.LogError(e)
		}
	}

	sqlWhere := ` where a.id<=? and a.pid=? `
	sqlArgs = append(sqlArgs, maxId, pid)

	if state != 0 {
		sqlWhere = sqlWhere + ` and a.state=? `
		sqlArgs = append(sqlArgs, state)
	}
	if name != "" {
		sqlWhere = sqlWhere + ` and locate(?,a.user_name)>0 `
		sqlArgs = append(sqlArgs, name)
	}

	//------------------------------------
	sqlCount := fmt.Sprintf(`select count(1) as c 
,sum(a.gold) as gold, sum(a.sum_save_money+a.sum_add_money) as sum_save_money, max(a.max_save_money) as max_save_money
,sum(a.sum_draw_money+a.sum_dec_money) as sum_draw_money, max(a.max_draw_money) as max_draw_money 
,sum(a.sum_bet) as sum_bet, sum(a.sum_win) as sum_win, sum(a.rebate) as rebate 
from tt_game_user a %s`, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupUser)
	if e != nil {
		return 0, arrData, aGroupUser
	}

	offset, _ := sqlKit.GetOffset(aGroupUser.C, pageSize, pageIndex)
	sqlWhere = sqlWhere + ` order by a.created_at,a.id DESC LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, pageSize)

	sql := fmt.Sprintf(`select 
a.id, a.user_name, a.nickname, a.lower_count, a.all_lower_count, a.referrer_code, a.state
, a.gold, (a.sum_save_money+a.sum_add_money) as sum_save_money, a.max_save_money, (a.sum_draw_money+a.sum_dec_money) as sum_draw_money, a.max_draw_money
, a.sum_bet, a.sum_win, a.created_at, a.login_time, a.rebate
from tt_game_user a %s `, sqlWhere)
	_, e = o.Raw(sql, sqlArgs).QueryRows(&arrData)
	if e != nil {
		return 0, arrData, aGroupUser
	}
	aGroupUser.FirstId = maxId
	return aGroupUser.C, arrData, aGroupUser
}

func UpdateUserPwd(UserId int, pwd string) error {
	e := httpGameServer.ChangePaw(UserId, pwd)
	if e != nil {
		return e
	}
	updateData := make([]gBox.UpdateDataInfo, 0)
	newPwd1 := pwdKit.Sha1ToStr(pwd)
	updateData = append(updateData, gBox.UpdateDataInfo{FieldName: "Pwd", Type: 0, Value: newPwd1}, gBox.UpdateDataInfo{FieldName: "Pwd2", Type: 0, Value: pwd})

	_, e = GInstance.GetUserRpcClient().UpdateUserInfo(UserId, updateData)

	return e
}

func SetUserState(UserId, State int) error {
	updateData := make([]gBox.UpdateDataInfo, 0)
	updateData = append(updateData, gBox.UpdateDataInfo{FieldName: "State", Type: 0, Value: State})

	_, e := GInstance.GetUserRpcClient().UpdateUserInfo(UserId, updateData)

	return e
}

func GetUserDetailInfo(UserId int) (models.TtGameUser, error) {
	return GInstance.GetUserRpcClient().GetUser(UserId)
}
func SaveUserDetailInfo(UserId int, info []gBox.UpdateDataInfo) error {
	_, e := GInstance.GetUserRpcClient().UpdateUserInfo(UserId, info)
	return e
}
