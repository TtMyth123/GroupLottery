package AFinanceBll

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"ttmyth123/GroupLottery/Admin/CacheData"
	"ttmyth123/GroupLottery/Admin/GConfig"
	"ttmyth123/GroupLottery/Admin/GInstance"
	"ttmyth123/GroupLottery/GameServer/models"
	"ttmyth123/GroupLottery/GameServer/models/mconst"
	"ttmyth123/GroupLottery/UserInfoRpc/GData/gBox"
	userConst "ttmyth123/GroupLottery/UserInfoRpc/models/mconst"
	"ttmyth123/kit/httpKit"
	"ttmyth123/kit/ttLog"

	"ttmyth123/kit/sqlKit"
)

func getSaveDrawApplyInfoKey() string {
	return fmt.Sprintf("SaveDrawApplyInfoKey")
}

type SaveDrawApplyInfo struct {
	SaveCount int
	DrawCount int
	SaveId    int
	DrawId    int
}

func GetSaveDrawApplyInfo() SaveDrawApplyInfo {
	key := getSaveDrawApplyInfoKey()
	aSaveDrawApplyInfo := SaveDrawApplyInfo{}
	e := CacheData.GetBeegoCache().GetCache(key, &aSaveDrawApplyInfo)
	if e != nil {
		aSaveDrawApplyInfo, e = GetSaveDrawApplyInfoAndSave()
	}

	return aSaveDrawApplyInfo
}

func GetSaveDrawApplyInfoAndSave() (SaveDrawApplyInfo, error) {
	aSaveDrawApplyInfo := SaveDrawApplyInfo{}
	type Tmp struct {
		C     int
		MaxId int
	}

	aTmp := Tmp{}
	o := orm.NewOrm()
	sqlCount := fmt.Sprintf(`select count(1) as c,max(a.id) as max_id from 
%s a, tt_game_user b where a.user_id=b.id and a.state <>?`, mconst.TableName_TtSaveMoney)

	e := o.Raw(sqlCount, mconst.SaveMoneyState_5_OK).QueryRow(&aTmp)
	if e != nil {
		return aSaveDrawApplyInfo, e
	}
	aSaveDrawApplyInfo.SaveCount = aTmp.C
	aSaveDrawApplyInfo.SaveId = aTmp.MaxId

	sqlCount = fmt.Sprintf(`select count(1) as c,max(a.id) as max_id from 
%s a, tt_game_user b where a.user_id=b.id and a.state = ?`, mconst.TableName_TtDrawMoney)
	e = o.Raw(sqlCount, mconst.DrawMoneyState_1_Apply).QueryRow(&aTmp)
	if e != nil {
		return aSaveDrawApplyInfo, e
	}
	aSaveDrawApplyInfo.DrawCount = aTmp.C
	aSaveDrawApplyInfo.DrawId = aTmp.MaxId

	key := getSaveDrawApplyInfoKey()
	CacheData.GetBeegoCache().SetCache(key, &aSaveDrawApplyInfo, 60*5)

	return aSaveDrawApplyInfo, nil
}

type DrawMoneyApply struct {
	Id          int
	UserId      int
	Gold        float64
	State       int
	SN          string `orm:"size(40)"`
	UserName    string
	FullName    string
	AuditorId   int    //审核人Id
	AuditorName string `orm:"size(256)"` //审核人名称
	UpdatedAt   time.Time
	CreatedAt   time.Time
	CurGold     float64

	UserType int
}

func (d DrawMoneyApply) MarshalJSON() ([]byte, error) {
	type Alias DrawMoneyApply
	return json.Marshal(&struct {
		Alias
		UpdatedAt string
		CreatedAt string
	}{
		Alias:     (Alias)(d),
		UpdatedAt: d.UpdatedAt.Format("2006-01-02 15:04:05"),
		CreatedAt: d.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

type GroupDrawMoneyApply struct {
	C       int
	Gold    float64
	FirstId int
}

/**
获取提现申请列表
*/
func GetDrawMoneyApplyList(curAgentId int, userName, beginDay, endDay string, userId, pageIndex, pageSize, FirstId int) (int, []DrawMoneyApply, GroupDrawMoneyApply) {
	aGroupDrawMoneyApply := GroupDrawMoneyApply{}
	o := orm.NewOrm()
	arrData := make([]DrawMoneyApply, 0)
	sqlArgs := make([]interface{}, 0)

	maxId := FirstId
	if FirstId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, mconst.TableName_TtDrawMoney)
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&maxId)
		if e != nil {
			ttLog.LogError(e)
		}
	}

	sqlWhere := "and b.agent_user_id=? and a.id<=? and a.state=? and b.user_type<>? "
	sqlArgs = append(sqlArgs, curAgentId, maxId, mconst.DrawMoneyState_1_Apply, userConst.UserType_3)
	if userId != 0 {
		sqlWhere += " and a.user_id=? "
		sqlArgs = append(sqlArgs, userId)
	}

	if userName != "" {
		sqlWhere = sqlWhere + ` and locate(?,b.user_name)>0`
		sqlArgs = append(sqlArgs, userName)
	}

	if beginDay != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.created_at,'%[1]s') >= ? `, "%Y-%m-%d")
		sqlArgs = append(sqlArgs, beginDay)
	}
	if endDay != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.created_at,'%[1]s') <= ? `, "%Y-%m-%d")
		sqlArgs = append(sqlArgs, endDay)
	}
	if userId != 0 {
		sqlWhere += " and a.user_id=?"
		sqlArgs = append(sqlArgs, userId)
	}

	sqlCount := fmt.Sprintf(`select count(1) c, sum(a.gold) as gold from 
tt_draw_money a, tt_game_user b where a.user_id=b.id %s`, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupDrawMoneyApply)
	if e != nil {
		return aGroupDrawMoneyApply.C, arrData, aGroupDrawMoneyApply
	}

	offset, _ := sqlKit.GetOffset(aGroupDrawMoneyApply.C, pageSize, pageIndex)
	sqlWhere = sqlWhere + ` LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, pageSize)

	sql := fmt.Sprintf(`select a.*, b.user_name, b.gold as cur_gold, b.user_type from
tt_draw_money a, tt_game_user b  where a.user_id=b.id %s `,
		sqlWhere)
	o.Raw(sql, sqlArgs).QueryRows(&arrData)

	aGroupDrawMoneyApply.FirstId = maxId
	return aGroupDrawMoneyApply.C, arrData, aGroupDrawMoneyApply
}

type SaveMoneyApply struct {
	Id          int
	UserId      int
	Gold        float64
	Money       float64
	State       int
	SN          string `orm:"size(40)"`
	UserName    string
	AuditorId   int    //审核人Id
	AuditorName string `orm:"size(256)"` //审核人名称
	UpdatedAt   time.Time
	CreatedAt   time.Time
	CurGold     float64

	FullName string
	UserType int
}

func (d SaveMoneyApply) MarshalJSON() ([]byte, error) {
	type Alias SaveMoneyApply
	return json.Marshal(&struct {
		Alias
		UpdatedAt string
		CreatedAt string
	}{
		Alias:     (Alias)(d),
		UpdatedAt: d.UpdatedAt.Format("2006-01-02 15:04:05"),
		CreatedAt: d.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

type GroupSaveMoneyApply struct {
	C       int
	Gold    float64
	Money   float64
	FirstId int
}

/**
获取 充值申请列表
*/
func GetSaveMoneyApplyList(curAgentId int, userName, beginDay, endDay string, userId, pageIndex, pageSize, FirstId int) (int, []SaveMoneyApply, GroupSaveMoneyApply) {
	aGroupSaveMoneyApply := GroupSaveMoneyApply{}
	o := orm.NewOrm()
	arrData := make([]SaveMoneyApply, 0)
	sqlArgs := make([]interface{}, 0)

	maxId := FirstId
	if FirstId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, mconst.TableName_TtSaveMoney)
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&maxId)
		if e != nil {
			ttLog.LogError(e)
		}
	}

	sqlWhere := "and b.agent_user_id=? and a.id<=? and a.state <> ? and b.user_type<>? "
	sqlArgs = append(sqlArgs, curAgentId, maxId, mconst.SaveMoneyState_5_OK, userConst.UserType_3)
	if userId != 0 {
		sqlWhere += " and a.user_id=? "
		sqlArgs = append(sqlArgs, userId)
	}

	if userName != "" {
		sqlWhere = sqlWhere + ` and locate(?,b.user_name)>0`
		sqlArgs = append(sqlArgs, userName)
	}

	if beginDay != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.created_at,'%[1]s') >= ? `, "%Y-%m-%d")
		sqlArgs = append(sqlArgs, beginDay)
	}
	if endDay != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.created_at,'%[1]s') <= ? `, "%Y-%m-%d")
		sqlArgs = append(sqlArgs, endDay)
	}
	if userId != 0 {
		sqlWhere += " and a.user_id=?"
		sqlArgs = append(sqlArgs, userId)
	}

	sqlCount := fmt.Sprintf(`select count(1) c ,sum(a.gold) as gold, sum(a.money) as money 
from  tt_save_money a, tt_game_user b where a.user_id=b.id %s`, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupSaveMoneyApply)
	if e != nil {
		return aGroupSaveMoneyApply.C, arrData, aGroupSaveMoneyApply
	}

	offset, _ := sqlKit.GetOffset(aGroupSaveMoneyApply.C, pageSize, pageIndex)
	sqlWhere = sqlWhere + ` LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, pageSize)

	sql := fmt.Sprintf(`select a.*, b.user_name, b.gold as cur_gold, b.user_type from
tt_save_money a, tt_game_user b where a.user_id=b.id %s `,
		sqlWhere)
	o.Raw(sql, sqlArgs).QueryRows(&arrData)

	aGroupSaveMoneyApply.FirstId = maxId
	return aGroupSaveMoneyApply.C, arrData, aGroupSaveMoneyApply
}

/**
上分
*/
func SaveMoney(GroupId, UserId, Money, AuditorId int, AuditorName string) error {
	o := orm.NewOrm()
	aTtSaveMoney := models.TtSaveMoney{GroupId:GroupId, UserId: UserId, Money: float64(Money), Gold: float64(Money), State: mconst.SaveMoneyState_5_OK,
		VoucherUrl: "", PayState: 0, AuditorId: AuditorId, AuditorName: AuditorName}
	e := aTtSaveMoney.Add(o)
	if e != nil {
		return e
	}

	goldInfo := gBox.AddGoldInfo{GroupId:GroupId, UserId: UserId, Gold: float64(Money), T: userConst.Account_08_AddMoney,
		Des:  fmt.Sprintf("[%s]上分%d。", AuditorName, Money),
		Des2: fmt.Sprintf("[%s]Nạp điểm%d。", AuditorName, Money),
	}
	_, e = GInstance.GetUserRpcClient().AddGold(goldInfo)
	return e
}

/**
下分
*/
func DrawMoney(UserId, Money, AuditorId int, AuditorName string) error {
	o := orm.NewOrm()
	goldInfo := gBox.AddGoldInfo{UserId: UserId, Gold: float64(Money), T: userConst.Account_09_DecMoney,
		Des:  fmt.Sprintf("[%s]下分%d。", AuditorName, Money),
		Des2: fmt.Sprintf("[%s]Rút điểm%d。", AuditorName, Money),
	}
	_, e := GInstance.GetUserRpcClient().AddGold(goldInfo)
	if e == nil {
		aTtDrawMoney := models.TtDrawMoney{UserId: UserId, Gold: float64(Money), State: mconst.DrawMoneyState_4,
			AuditorId: AuditorId, AuditorName: AuditorName}
		e := aTtDrawMoney.Add(o)
		if e != nil {
			return e
		}
	}
	return e
}

/**
同意充值
*/
func AgreeSaveMoney(UserId, SaveMoneyId, AuditorId int, AuditorName string) error {
	o := orm.NewOrm()

	aTtSaveMoney := models.TtSaveMoney{Id: SaveMoneyId}
	e := o.Read(&aTtSaveMoney)
	if e != nil {
		return e
	}
	if aTtSaveMoney.State == mconst.SaveMoneyState_5_OK || aTtSaveMoney.UserId != UserId {
		return errors.New("数据变动，请刷新数据")
	}

	goldInfo := gBox.AddGoldInfo{UserId: UserId,
		Gold: aTtSaveMoney.Gold, T: userConst.Account_03_SaveMoney,
		Des:  fmt.Sprintf("充值%.2f", aTtSaveMoney.Gold),
		Des2: fmt.Sprintf("Lên điểm%.2f", aTtSaveMoney.Gold),
	}
	_, e = GInstance.GetUserRpcClient().AddGold(goldInfo)
	if e == nil {
		aTtSaveMoney.State = mconst.SaveMoneyState_5_OK
		aTtSaveMoney.AuditorId = AuditorId
		aTtSaveMoney.AuditorName = AuditorName

		e = aTtSaveMoney.Update(o, "State", "AuditorName", "AuditorId")
		if e != nil {
			return e
		}
	}
	return e
}

func AgreeDrawMoney(DrawMoneyId, AuditorId int, AuditorName string) error {
	o := orm.NewOrm()
	aDrawMoney := models.TtDrawMoney{Id: DrawMoneyId}
	e := o.Read(&aDrawMoney)
	if e != nil {
		return e
	}

	aDrawMoney.State = mconst.DrawMoneyState_4
	aDrawMoney.AuditorId = AuditorId
	aDrawMoney.AuditorName = AuditorName
	e = aDrawMoney.Update(o, "State", "AuditorName", "AuditorId", "UpdatedAt")

	return e
}

type AccountEx2 struct {
	Id          int
	UserId      int
	AccountType int
	StrType     string
	Des         string
	CurUserGold float64
	Gold        float64
	CreatedAt   time.Time
	UserName    string
	FullName    string
	UserType    int
}

func (d AccountEx2) MarshalJSON() ([]byte, error) {
	type Alias AccountEx2
	return json.Marshal(&struct {
		Alias
		CreatedAt string
	}{
		Alias:     (Alias)(d),
		CreatedAt: d.CreatedAt.Format("06-01-02 15:04:05"),
	})
}

type GroupAccountEx2 struct {
	C       int
	FirstId int
	Gold    float64
}

func GetAccountList(curAgentId int, beginDay, endDay, userName string, userId, accountType, pageIndex, pageSize, FirstId int) (int, []AccountEx2, GroupAccountEx2) {
	PageTotal := 0
	o := orm.NewOrm()
	aGroupAccountEx2 := GroupAccountEx2{}
	arrData := make([]AccountEx2, 0)
	sqlArgs := make([]interface{}, 0)

	maxId := FirstId
	if FirstId == 0 {
		sqlCount := fmt.Sprintf(`select max(id) as maxid from tt_account a `)
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&maxId)
		if e != nil {
			ttLog.LogError(e)
		}
	}

	sqlWhere := " and a.id<=?"
	sqlArgs = append(sqlArgs, maxId)
	if curAgentId != 0 {
		sqlWhere = sqlWhere + " and b.agent_user_id=?"
		sqlArgs = append(sqlArgs, curAgentId)
	}

	if userName != "" {
		sqlWhere = sqlWhere + ` and locate(?,b.user_name)>0`
		sqlArgs = append(sqlArgs, userName)
	}

	if beginDay != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.created_at,'%[1]s') >= ? `, "%Y-%m-%d")
		sqlArgs = append(sqlArgs, beginDay)
	}
	if endDay != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.created_at,'%[1]s') <= ? `, "%Y-%m-%d")
		sqlArgs = append(sqlArgs, endDay)
	}
	if accountType != 0 {
		sqlWhere += " and a.account_type=?"
		sqlArgs = append(sqlArgs, accountType)
	}
	if userId != 0 {
		sqlWhere += " and a.user_id=?"
		sqlArgs = append(sqlArgs, userId)
	}

	sqlCount := fmt.Sprintf(`select count(1) c from 
tt_account a, tt_game_user b where a.user_id=b.id %s`, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&PageTotal)
	if e != nil {
		return PageTotal, arrData, aGroupAccountEx2
	}

	if FirstId != 0 {
		if accountType != 0 {
			sqlWhere += " and a.id<=?"
			sqlArgs = append(sqlArgs, FirstId)
		}
	}

	sqlCount = fmt.Sprintf(`select count(1) c from 
tt_account a, tt_game_user b where a.user_id=b.id %s`, sqlWhere)
	e = o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupAccountEx2)
	if e != nil {
		return PageTotal, arrData, aGroupAccountEx2
	}

	offset, _ := sqlKit.GetOffset(PageTotal, pageSize, pageIndex)
	sqlWhere = sqlWhere + ` order by a.created_at desc, a.id desc LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, pageSize)

	sql := fmt.Sprintf(`select a.*, b.user_name, b.head_imgurl, b.full_name,b.user_type  from 
tt_account a, tt_game_user b where a.user_id=b.id %s `,
		sqlWhere)
	o.Raw(sql, sqlArgs).QueryRows(&arrData)

	if len(arrData) > 0 && FirstId == 0 {
		aGroupAccountEx2.FirstId = arrData[0].Id
	} else {
		aGroupAccountEx2.FirstId = FirstId
	}

	aGroupAccountEx2.FirstId = maxId
	return PageTotal, arrData, aGroupAccountEx2
}

//func GetAccountListFirstId(beginDay, endDay, userName string, accountType, pageIndex, pageSize, FirstId int) (int, []AccountEx2, GroupAccountEx2) {
//	PageTotal := 0
//	o := orm.NewOrm()
//	aGroupAccountEx2 := GroupAccountEx2{}
//	arrData := make([]AccountEx2, 0)
//	sqlArgs := make([]interface{}, 0)
//
//	sqlWhere := ""
//	if userName != "" {
//		sqlWhere = sqlWhere + ` and locate(?,b.user_name)>0`
//		sqlArgs = append(sqlArgs, userName)
//	}
//
//	if beginDay != "" {
//		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.created_at,'%[1]s') >= ? `, "%Y-%m-%d")
//		sqlArgs = append(sqlArgs, beginDay)
//	}
//	if endDay != "" {
//		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.created_at,'%[1]s') <= ? `, "%Y-%m-%d")
//		sqlArgs = append(sqlArgs, endDay)
//	}
//	if accountType != 0 {
//		sqlWhere += " and a.account_type=?"
//		sqlArgs = append(sqlArgs, accountType)
//	}
//
//	sqlCount := fmt.Sprintf(`select count(1) c from
//tt_account a, tt_game_user b where a.user_id=b.id %s`, sqlWhere)
//	e := o.Raw(sqlCount, sqlArgs).QueryRow(&PageTotal)
//	if e != nil {
//		return PageTotal, arrData, aGroupAccountEx2
//	}
//
//	if FirstId!=0{
//		if accountType != 0 {
//			sqlWhere += " and a.id<=?"
//			sqlArgs = append(sqlArgs, FirstId)
//		}
//	}
//
//	sqlCount = fmt.Sprintf(`select count(1) c from
//tt_account a, tt_game_user b where a.user_id=b.id %s`, sqlWhere)
//	e = o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupAccountEx2)
//	if e != nil {
//		return PageTotal, arrData, aGroupAccountEx2
//	}
//
//	offset, _ := sqlKit.GetOffset(PageTotal, pageSize, pageIndex)
//	sqlWhere = sqlWhere + ` order by a.created_at desc, a.id desc LIMIT ?,?`
//	sqlArgs = append(sqlArgs, offset, pageSize)
//
//	sql := fmt.Sprintf(`select a.*, b.user_name, b.head_imgurl, b.full_name,b.user_type  from
//tt_account a, tt_game_user b where a.user_id=b.id %s `,
//		sqlWhere)
//	o.Raw(sql, sqlArgs).QueryRows(&arrData)
//
//	if len(arrData) >0 && FirstId==0{
//		aGroupAccountEx2.FirstId = arrData[0].Id
//	} else {
//		aGroupAccountEx2.FirstId = FirstId
//	}
//
//	return PageTotal, arrData, aGroupAccountEx2
//}

type TtDrawMoneyEx struct {
	Id          int
	UserId      int
	Gold        float64
	State       int
	SN          string
	AuditorId   int
	AuditorName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserName    string
	FullName    string

	UserType int
}

func (d TtDrawMoneyEx) MarshalJSON() ([]byte, error) {
	type Alias TtDrawMoneyEx
	return json.Marshal(&struct {
		Alias
		CreatedAt string
		UpdatedAt string
	}{
		Alias:     (Alias)(d),
		CreatedAt: d.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: d.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}

type GroupTtDrawMoneyEx struct {
	C       int
	Gold    float64
	FirstId int
}

func DrawMoneyList(curAgentId int, beginDay, endDay, userName string, userId, pageIndex, pageSize, FirstId int) (int, []TtDrawMoneyEx, GroupTtDrawMoneyEx) {
	aGroupTtDrawMoneyEx := GroupTtDrawMoneyEx{}
	o := orm.NewOrm()
	arrData := make([]TtDrawMoneyEx, 0)
	sqlArgs := make([]interface{}, 0)
	maxId := FirstId
	if FirstId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, mconst.TableName_TtDrawMoney)
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&maxId)
		if e != nil {
			ttLog.LogError(e)
		}
	}

	sqlWhere := "and b.agent_user_id=? and a.id<=? and a.state<>? and b.user_type<>? "
	sqlArgs = append(sqlArgs, curAgentId, maxId, mconst.DrawMoneyState_1_Apply, userConst.UserType_3)
	if userName != "" {
		sqlWhere = sqlWhere + ` and locate(?,b.user_name)>0`
		sqlArgs = append(sqlArgs, userName)
	}

	if beginDay != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.created_at,'%[1]s') >= ? `, "%Y-%m-%d")
		sqlArgs = append(sqlArgs, beginDay)
	}
	if endDay != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.created_at,'%[1]s') <= ? `, "%Y-%m-%d")
		sqlArgs = append(sqlArgs, endDay)
	}
	if userId != 0 {
		sqlWhere += " and a.user_id=?"
		sqlArgs = append(sqlArgs, userId)
	}

	sqlCount := fmt.Sprintf(`select count(1) c, sum(a.gold) as gold from 
tt_draw_money a, tt_game_user b where a.user_id=b.id  %s`, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupTtDrawMoneyEx)
	if e != nil {
		return aGroupTtDrawMoneyEx.C, arrData, aGroupTtDrawMoneyEx
	}

	offset, _ := sqlKit.GetOffset(aGroupTtDrawMoneyEx.C, pageSize, pageIndex)
	sqlWhere = sqlWhere + ` order by a.created_at DESC LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, pageSize)

	sql := fmt.Sprintf(`select a.*, b.user_name, b.full_name, b.user_type from 
tt_draw_money a, tt_game_user b where a.user_id=b.id  %s `,
		sqlWhere)
	o.Raw(sql, sqlArgs).QueryRows(&arrData)

	aGroupTtDrawMoneyEx.FirstId = maxId
	return aGroupTtDrawMoneyEx.C, arrData, aGroupTtDrawMoneyEx
}

type TtSaveMoneyEx struct {
	Id          int
	UserId      int
	Money       float64
	Gold        float64
	State       int
	VoucherUrl  string
	PayState    int
	AuditorId   int
	AuditorName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserName    string
	FullName    string
	UserType    int
}

func (d TtSaveMoneyEx) MarshalJSON() ([]byte, error) {
	type Alias TtSaveMoneyEx
	return json.Marshal(&struct {
		Alias
		UpdatedAt string
		CreatedAt string
	}{
		Alias:     (Alias)(d),
		UpdatedAt: d.UpdatedAt.Format("2006-01-02 15:04:05"),
		CreatedAt: d.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

type GruopTtSaveMoneyEx struct {
	C       int
	Money   float64
	Gold    float64
	FirstId int
}

func SaveMoneyList(curAgentId int, beginDay, endDay, userName string, userId, pageIndex, pageSize, FirstId int) (int, []TtSaveMoneyEx, GruopTtSaveMoneyEx) {
	aGruopTtSaveMoneyEx := GruopTtSaveMoneyEx{}
	o := orm.NewOrm()
	arrData := make([]TtSaveMoneyEx, 0)
	sqlArgs := make([]interface{}, 0)

	maxId := FirstId
	if FirstId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, "tt_save_money")
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&maxId)
		if e != nil {
			ttLog.LogError(e)
		}
	}

	sqlWhere := "and b.agent_user_id=? and a.id<=? and a.state<>? and b.user_type<>? "
	sqlArgs = append(sqlArgs, curAgentId, maxId, mconst.SaveMoneyState_1_Apply, userConst.UserType_3)
	if userName != "" {
		sqlWhere = sqlWhere + ` and locate(?,b.user_name)>0`
		sqlArgs = append(sqlArgs, userName)
	}

	if beginDay != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.created_at,'%[1]s') >= ? `, "%Y-%m-%d")
		sqlArgs = append(sqlArgs, beginDay)
	}
	if endDay != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.created_at,'%[1]s') <= ? `, "%Y-%m-%d")
		sqlArgs = append(sqlArgs, endDay)
	}
	if userId != 0 {
		sqlWhere += " and a.user_id=?"
		sqlArgs = append(sqlArgs, userId)
	}

	sqlCount := fmt.Sprintf(`select count(1) c, sum(a.money) as money, sum(a.gold) as gold from 
tt_save_money a, tt_game_user b  where a.user_id=b.id  %s`, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGruopTtSaveMoneyEx)
	if e != nil {
		return aGruopTtSaveMoneyEx.C, arrData, aGruopTtSaveMoneyEx
	}

	offset, _ := sqlKit.GetOffset(aGruopTtSaveMoneyEx.C, pageSize, pageIndex)
	sqlWhere = sqlWhere + ` order by a.created_at DESC LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, pageSize)

	sql := fmt.Sprintf(`select a.*, b.user_name, b.full_name, b.user_type from 
tt_save_money a, tt_game_user b where a.user_id=b.id %s `,
		sqlWhere)
	o.Raw(sql, sqlArgs).QueryRows(&arrData)

	aGruopTtSaveMoneyEx.FirstId = maxId
	return aGruopTtSaveMoneyEx.C, arrData, aGruopTtSaveMoneyEx
}

func DelSaveMoney(id, userId int) (int, error) {
	o := orm.NewOrm()
	aSaveMoney := models.TtSaveMoney{Id: id}
	e := o.Read(&aSaveMoney)
	if e != nil {
		return 0, errors.New("没有对应的数据")
	}
	if aSaveMoney.UserId != userId && userId != 0 {
		return 0, errors.New("你没权限删除此数据")
	}
	if aSaveMoney.State == mconst.SaveMoneyState_5_OK {
		return 0, errors.New("数据更新，请刷新后重试。")
	}
	_, e = o.Delete(&aSaveMoney)
	if e != nil {
		return 0, errors.New("没有对应的数据")
	}
	typeScore := int(aSaveMoney.Gold * 100)

	GetSaveDrawApplyInfoAndSave()
	return typeScore, nil
}

func DelDrawMoney(drawMoneyId, userId int, Excuse string) error {
	o := orm.NewOrm()
	//o.Begin()
	aTtDrawMoney := models.TtDrawMoney{Id: drawMoneyId}
	e := o.Read(&aTtDrawMoney)
	if e != nil {
		//o.Rollback()
		return errors.New("没有对应的数据")
	}

	if aTtDrawMoney.UserId != userId && userId != 0 {
		//o.Rollback()
		return errors.New("你没权限删除此数据")
	}
	if aTtDrawMoney.State != mconst.DrawMoneyState_1_Apply {
		return errors.New("数据更新，请刷新后重试。")
	}
	aGoldInfo := gBox.AddGoldInfo{UserId: aTtDrawMoney.UserId, Gold: aTtDrawMoney.Gold,
		T:    userConst.Account_07_DrawMoneyR,
		Des:  fmt.Sprintf("提现拒绝%.0f,原因：[%s]。", aTtDrawMoney.Gold, Excuse),
		Des2: fmt.Sprintf("rút điểm bị từ chối%.0f,Lý do：[%s]。", aTtDrawMoney.Gold, Excuse),
	}

	_, e = GInstance.GetUserRpcClient().AddGold(aGoldInfo)
	if e != nil {
		//o.Rollback()
		return e
	}
	_, e = o.Delete(&aTtDrawMoney)
	if e != nil {
		//o.Rollback()
		return e
	}
	//o.Commit()
	//todo GetSaveDrawApplyInfoAndSave()
	return e
}

type RebateInfo struct {
	Id            int
	UserId        int
	UserName      string
	RebateType    int
	StrType       string
	Des           string
	Rebate        float64
	CurUserRebate float64
	CreatedAt     time.Time
}

func (d RebateInfo) MarshalJSON() ([]byte, error) {
	type Alias RebateInfo
	return json.Marshal(&struct {
		Alias
		CreatedAt string
	}{
		Alias:     (Alias)(d),
		CreatedAt: d.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

type GroupRebateInfo struct {
	C       int
	Rebate  float64
	FirstId int
}

/**
获取 返利信息
*/
func GetRebateList(curAgentId, userId, RebateType int, name, beginDay, endDay string, pageIndex, pageSize, FirstId int) (int, []RebateInfo, GroupRebateInfo) {
	o := orm.NewOrm()
	aGroupUser := GroupRebateInfo{}
	arrData := make([]RebateInfo, 0)
	sqlArgs := make([]interface{}, 0)
	maxId := FirstId
	if FirstId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, "tt_rebate_info")
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&maxId)
		if e != nil {
			ttLog.LogError(e)
		}
	}

	sqlWhere := " and a.id<=? "
	sqlArgs = append(sqlArgs, maxId)
	if curAgentId != 0 {
		sqlWhere += ` and b.agent_user_id=? `
		sqlArgs = append(sqlArgs, curAgentId)
	}
	if RebateType != 0 {
		sqlWhere += ` and a.rebate_type=? `
		sqlArgs = append(sqlArgs, RebateType)
	}
	if userId != 0 {
		sqlWhere += ` and a.user_id=? `
		sqlArgs = append(sqlArgs, userId)
	}

	if name != "" {
		sqlWhere += ` and locate(?,b.user_name)>0`
		sqlArgs = append(sqlArgs, name)
	}

	if beginDay != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.created_at,'%[1]s') >= ? `, "%Y-%m-%d")
		sqlArgs = append(sqlArgs, beginDay)
	}
	if endDay != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.created_at,'%[1]s') <= ? `, "%Y-%m-%d")
		sqlArgs = append(sqlArgs, endDay)
	}

	//------------------------------------
	sqlCount := fmt.Sprintf(`select count(1) c, sum(a.rebate) rebate
from tt_rebate_info a, tt_game_user b where a.user_id = b.id %s`, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupUser)
	if e != nil {
		return 0, arrData, aGroupUser
	}

	offset, _ := sqlKit.GetOffset(aGroupUser.C, pageSize, pageIndex)
	sqlWhere = sqlWhere + ` order by a.created_at DESC, a.id DESC LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, pageSize)

	sql := fmt.Sprintf(` select b.user_name, a.id, a.user_id,a.rebate_type, a.str_type, a.des,a.rebate
, a.cur_user_rebate, a.created_at from tt_rebate_info a, tt_game_user b where a.user_id = b.id %s`, sqlWhere)
	_, e = o.Raw(sql, sqlArgs).QueryRows(&arrData)
	if e != nil {
		return 0, arrData, aGroupUser
	}
	aGroupUser.FirstId = maxId
	return aGroupUser.C, arrData, aGroupUser
}

/**
获取 返利信息
*/
//func GetRebateListFirstId(RebateType int, name, beginDay, endDay string, pageIndex, pageSize,FirstId int) (int, []RebateInfo, GroupRebateInfo) {
//	o := orm.NewOrm()
//	aGroupUser := GroupRebateInfo{}
//	arrData := make([]RebateInfo, 0)
//	sqlArgs := make([]interface{}, 0)
//
//	sqlWhere := ""
//	if RebateType != 0 {
//		sqlWhere += ` and a.rebate_type=? `
//		sqlArgs = append(sqlArgs, RebateType)
//	}
//
//	if name != "" {
//		sqlWhere += ` and locate(?,b.user_name)>0`
//		sqlArgs = append(sqlArgs, name)
//	}
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
//	//------------------------------------
//	sqlCount := fmt.Sprintf(`select count(1) c, sum(a.rebate) rebate
//from tt_rebate_info a, tt_game_user b where a.user_id = b.id a  %s`, sqlWhere)
//	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupUser)
//	if e != nil {
//		return 0, arrData, aGroupUser
//	}
//
//	offset, _ := sqlKit.GetOffset(aGroupUser.C, pageSize, pageIndex)
//	sqlWhere = sqlWhere + ` order by a.created_at DESC LIMIT ?,?`
//	sqlArgs = append(sqlArgs, offset, pageSize)
//
//	sql := fmt.Sprintf(` select b.user_name, a.id, a.user_id,a.rebate_type, a.str_type, a.des,a.rebate
//, a.cur_user_rebate, a.created_at from tt_rebate_info a, tt_game_user b where a.user_id = b.id %s`, sqlWhere)
//	_, e = o.Raw(sql, sqlArgs).QueryRows(&arrData)
//	if e != nil {
//		return 0, arrData, aGroupUser
//	}
//	return aGroupUser.C, arrData, aGroupUser
//}

func GetFinanceAccount() (models.TtAgentPayInfo, error) {
	aAgentPayInfo, e := models.GetAgentPayInfo(1)
	aAgentPayInfo.WXReceiptUrl = httpKit.GetImgUrl(GConfig.GetGConfig().SApiBaseP, aAgentPayInfo.WXReceiptUrl)
	aAgentPayInfo.AlipayUrl = httpKit.GetImgUrl(GConfig.GetGConfig().SApiBaseP, aAgentPayInfo.AlipayUrl)
	return aAgentPayInfo, e
}

type DrawMoneyAccountInfo struct {
	DrawMoneyId  int
	WXSKCodeUrl  string //微信收款码Url
	YHName       string //银行名
	CardNum      string //卡号
	YHUserName   string //银行用户名
	YHUserTel    string //银行预留电话
	Addr         string //银行预留地址
	Cate         string //身份证
	Remark       string //银行预留信息
	ZFBSKCodeUrl string //支付宝二维码Url
	ZFBSKName    string //支付宝名
}

func GetDrawMoneyAccountInfo(id int) (DrawMoneyAccountInfo, error) {
	aDrawMoneyAccountInfo := DrawMoneyAccountInfo{}
	o := orm.NewOrm()
	sql := fmt.Sprintf(`select a.id as draw_money_id, b.w_x_s_k_code_url, b.y_h_name,b.card_num 
,b.y_h_user_name,b.y_h_user_tel,b.addr,b.cate,b.remark,b.z_f_b_s_k_code_url,b.z_f_b_s_k_name
from tt_draw_money a, tt_game_user b where a.user_id=b.id and a.id=?`)
	e := o.Raw(sql, id).QueryRow(&aDrawMoneyAccountInfo)
	if e != nil {
		return aDrawMoneyAccountInfo, errors.New("没有对应的数据")
	}
	return aDrawMoneyAccountInfo, nil
}
