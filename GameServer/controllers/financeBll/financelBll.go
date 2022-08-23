package financeBll

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TtMyth123/GameServer/CacheData"
	"github.com/TtMyth123/GameServer/GConfig"
	"github.com/TtMyth123/GameServer/GInstance"
	"github.com/TtMyth123/GameServer/GInstance/GTtHint"
	"github.com/TtMyth123/GameServer/controllers/base/TtError"
	"github.com/TtMyth123/GameServer/models"
	"github.com/TtMyth123/GameServer/models/mconst"
	"github.com/TtMyth123/UserInfoRpc/GData/gBox"
	userConst "github.com/TtMyth123/UserInfoRpc/models/mconst"
	"github.com/TtMyth123/kit/httpKit"
	"github.com/TtMyth123/kit/pwdKit"
	"github.com/TtMyth123/kit/sqlKit"
	"github.com/TtMyth123/kit/timeKit"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"time"
)

func SaveMoneyApply(userId int, Gold float64) (int, error) {
	id := 0
	o := orm.NewOrm()
	c, e := o.QueryTable(mconst.TableName_TtSaveMoney).Filter("UserId", userId).
		Filter("State", mconst.SaveMoneyState_1_Apply).Count()
	if e != nil {
		return id, e
	}

	if c < 5 {
		aTtDrawSaveSet := models.GetTtDrawSaveSet()
		if Gold < float64(aTtDrawSaveSet.MinSave) {

			//return id, errors.New(fmt.Sprintf("最少充值%d", aTtDrawSaveSet.MinSave))
			return id, TtError.New("最少充值%d", aTtDrawSaveSet.MinSave)
		}

		curT := time.Now()
		aYhSaveMoney := models.TtSaveMoney{UserId: userId, Money: Gold, Gold: Gold, State: mconst.SaveMoneyState_1_Apply,
			VoucherUrl: "", PayState: 0, AuditorId: 0, AuditorName: "", CreatedAt: curT, UpdatedAt: curT}

		e := aYhSaveMoney.Add(o)
		if e != nil {
			return id, e
		} else {
			//newGold := int(Gold * 100)
			//CacheData.SetScoreRequisition(newGold)
		}
		id = aYhSaveMoney.Id
	} else {
		//return id, errors.New("未批准的申请过多。")
		return id, TtError.New("未批准的申请过多。")
	}

	GetSaveDrawApplyInfoAndSave()
	return id, nil
}

type SaveDrawApplyInfo struct {
	SaveCount int
	DrawCount int
	SaveId    int
	DrawId    int
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
%s a, tt_game_user b where a.user_id=b.id `, mconst.TableName_TtSaveMoney)

	e := o.Raw(sqlCount).QueryRow(&aTmp)
	if e != nil {
		return aSaveDrawApplyInfo, e
	}
	aSaveDrawApplyInfo.SaveId = aTmp.MaxId
	aSaveDrawApplyInfo.SaveCount = aTmp.C

	sqlCount = fmt.Sprintf(`select count(1) as c,max(a.id) as max_id from 
%s a, tt_game_user b where a.user_id=b.id`, mconst.TableName_TtDrawMoney)
	e = o.Raw(sqlCount).QueryRow(&aTmp)
	if e != nil {
		return aSaveDrawApplyInfo, e
	}
	aSaveDrawApplyInfo.DrawCount = aTmp.C
	aSaveDrawApplyInfo.DrawId = aTmp.MaxId

	key := GetSaveDrawApplyInfoKey()
	CacheData.GetBeegoCache().SetCache(key, &aSaveDrawApplyInfo, 60*5)

	return aSaveDrawApplyInfo, nil
}

func GetSaveDrawApplyInfoKey() string {
	return fmt.Sprintf("SaveDrawApplyInfoKey")
}

/**
提现申请
*/
//func DrawMoneyApply(userId int, Gold float64, Pwd string) error {
func DrawMoneyApply(GroupId, userId int, Gold float64, Pwd string) error {
	if Gold < 0 {
		return TtError.New("金额不能为负数")
	}
	o := orm.NewOrm()
	aUser, e := GInstance.GetUserRpcClient().GetUser(userId)
	if e != nil {
		//return errors.New(fmt.Sprintf("[%d]用户信息不正确。", userId))
		return errors.New(GTtHint.GetTtHint().GetHint("用户信息不正确"))
	}

	aNewPwd := ""
	if Pwd != "" {
		aNewPwd = pwdKit.Sha1ToStr(Pwd)
	}
	if aUser.DrawMoneyPwd != aNewPwd {
		return errors.New(GTtHint.GetTtHint().GetHint("提现密码不正确"))
	}

	aTtDrawSaveSet := models.GetTtDrawSaveSet()
	if Gold < float64(aTtDrawSaveSet.MinDraw) {
		//return errors.New(fmt.Sprintf("最少提现%d。", aTtDrawSaveSet.MinDraw))
		return errors.New(GTtHint.GetTtHint().GetHint("低于最少提现金额"))
	}
	t := time.Now()
	curHr := t.Hour()
	if aTtDrawSaveSet.DrawWorkHr <= 0 {
		return errors.New(GTtHint.GetTtHint().GetHint("客服休息时间，不能提现"))
	} else {
		h := aTtDrawSaveSet.DrawBeginHr + aTtDrawSaveSet.DrawWorkHr - 1
		if h <= 23 {
			if curHr < aTtDrawSaveSet.DrawBeginHr || curHr > h {
				return errors.New(GTtHint.GetTtHint().GetHint("客服休息时间，不能提现"))
			}
		} else {
			h1 := h - 23
			if curHr > h1 && curHr < aTtDrawSaveSet.DrawBeginHr {
				return errors.New(GTtHint.GetTtHint().GetHint("客服休息时间，不能提现"))
			}
		}
	}

	aGoldInfo := gBox.AddGoldInfo{GroupId: GroupId, UserId: userId, Gold: Gold, T: userConst.Account_04_DrawMoney,
		Des:  fmt.Sprintf("提现申请扣除：%g。", Gold),
		Des2: GTtHint.GetTtHint().GetHint("提现申请扣除：%g。"), DesMp: GTtHint.GetTtHint().GetMpString(Gold)}
	_, e = GInstance.GetUserRpcClient().AddGold(aGoldInfo)

	if e != nil {
		//o.Rollback()
		return e
	}
	curT := time.Now()
	aYhDrawMoney := models.TtDrawMoney{GroupId: GroupId, UserId: userId, Gold: Gold, State: mconst.DrawMoneyState_1_Apply,
		AuditorId: 0, AuditorName: "", CreatedAt: curT, UpdatedAt: curT}
	e = aYhDrawMoney.Add(o)
	if e != nil {
		//o.Rollback()
		return e
	}

	GetSaveDrawApplyInfoAndSave()
	return e
}

func DelSaveMoney(id, userId int) (int, error) {
	o := orm.NewOrm()
	aSaveMoney := models.TtSaveMoney{Id: id}
	e := o.Read(&aSaveMoney)
	if e != nil {
		return 0, errors.New(GTtHint.GetTtHint().GetHint("没有对应的数据"))
	}
	if aSaveMoney.UserId != userId && userId != 0 {
		return 0, errors.New(GTtHint.GetTtHint().GetHint("你没权限删除此数据"))
	}
	if aSaveMoney.State == mconst.SaveMoneyState_5_OK {
		return 0, errors.New(GTtHint.GetTtHint().GetHint("数据更新，请刷新后重试。"))
	}
	_, e = o.Delete(&aSaveMoney)
	if e != nil {
		return 0, errors.New(GTtHint.GetTtHint().GetHint("没有对应的数据"))
	}
	typeScore := int(aSaveMoney.Gold * 100)

	GetSaveDrawApplyInfoAndSave()
	return typeScore, nil
}

func DelDrawMoney(drawMoneyId, userId int) error {
	o := orm.NewOrm()
	//o.Begin()
	aTtDrawMoney := models.TtDrawMoney{Id: drawMoneyId}
	e := o.Read(&aTtDrawMoney)
	if e != nil {
		//o.Rollback()
		return errors.New(GTtHint.GetTtHint().GetHint("没有对应的数据"))
	}

	if aTtDrawMoney.UserId != userId && userId != 0 {
		//o.Rollback()
		return errors.New(GTtHint.GetTtHint().GetHint("你没权限删除此数据"))
	}
	if aTtDrawMoney.State != mconst.DrawMoneyState_1_Apply {
		return errors.New(GTtHint.GetTtHint().GetHint("数据更新，请刷新后重试。"))
	}
	aGoldInfo := gBox.AddGoldInfo{GroupId: aTtDrawMoney.GroupId, UserId: aTtDrawMoney.UserId, Gold: aTtDrawMoney.Gold,
		T:    userConst.Account_07_DrawMoneyR,
		Des:  fmt.Sprintf("提现拒绝%0.2f", aTtDrawMoney.Gold),
		Des2: fmt.Sprintf("Rút điểm bị từ chối%0.2f", aTtDrawMoney.Gold),
	}
	_, e = GInstance.GetUserRpcClient().AddGold(aGoldInfo)
	if e != nil {
		//o.Rollback()
		return errors.New(GTtHint.GetTtHint().GetHint("没有对应的数据"))
	}
	_, e = o.Delete(&aTtDrawMoney)
	if e != nil {
		//o.Rollback()
		return e
	}

	GetSaveDrawApplyInfoAndSave()
	return e
}

type DrawSaveMoney struct {
	T         int `json:"2"`
	Id        int
	Gold      float64   `json:"3"`
	Money     float64   `json:"4"`
	State     int       `json:"5"`
	CreatedAt time.Time `json:"1"`
}

func (d DrawSaveMoney) MarshalJSON() ([]byte, error) {
	type Alias DrawSaveMoney
	StrCreatedAt := ""
	TimeLag := GConfig.GetGConfig().TimeLag
	if TimeLag == 0 {
		StrCreatedAt = d.CreatedAt.Format(timeKit.DateTimeLayout)
	} else {
		StrCreatedAt = d.CreatedAt.Add(TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
	}

	return json.Marshal(&struct {
		Alias
		CreatedAt string `json:"1"`
	}{
		Alias:     (Alias)(d),
		CreatedAt: StrCreatedAt,
	})
}

func GetDrawSaveMoneyList(userId, pageIndex, pageSize int) (int, []DrawSaveMoney) {
	PageTotal := 0
	o := orm.NewOrm()
	arrData := make([]DrawSaveMoney, 0)

	sqlCount := fmt.Sprintf(`
select count(1) c from (
SELECT 2 as t, a.id, a.gold, a.state, a.created_at, if(a.state=0,0, 1) s from tt_draw_money a where a.user_id=?
UNION
SELECT 1 as t,a.id, a.gold, a.state, a.created_at, if(a.state=0,0, 1) s from tt_save_money a where a.user_id =?
) a
`)
	e := o.Raw(sqlCount, userId, userId).QueryRow(&PageTotal)
	if e != nil {
		return PageTotal, arrData
	}

	offset, _ := sqlKit.GetOffset(PageTotal, pageSize, pageIndex)
	sql := `
select  a.t, a.id, a.gold, a.money, a.state, a.created_at from (
SELECT 2 as t, a.id, a.gold, a.money, a.state, a.created_at, if(a.state=0,0, 1) s from tt_draw_money a where a.user_id=?
UNION
SELECT 1 as t,a.id, a.gold, a.money, a.state, a.created_at, if(a.state=0,0, 1) s from tt_save_money a where a.user_id=?
) a  order by a.s, a.created_at DESC LIMIT ?,?
`
	o.Raw(sql, userId, userId, offset, pageSize).QueryRows(&arrData)

	return PageTotal, arrData
}

func UploadSaveMoneyVoucher(userId, SaveMoneyId, State int, VoucherUrl string) error {
	o := orm.NewOrm()
	aSaveMoney := models.TtSaveMoney{Id: SaveMoneyId}
	e := o.Read(&aSaveMoney)
	if e != nil {
		return errors.New(GTtHint.GetTtHint().GetHint("没有对应的数据"))
	}

	if userId != 0 {
		if aSaveMoney.UserId != userId {
			return errors.New(GTtHint.GetTtHint().GetHint("数据异常，用户信息不匹配"))
		}
	}
	if aSaveMoney.State >= State {
		return errors.New(GTtHint.GetTtHint().GetHint("数据异常,刷新后重试"))
	}

	aSaveMoney.State = State
	if aSaveMoney.State == mconst.SaveMoneyState_4 {
		aSaveMoney.VoucherUrl = VoucherUrl
		e = aSaveMoney.Update(o, "State", "UpdatedAt", "VoucherUrl")
	} else {
		e = aSaveMoney.Update(o, "State", "UpdatedAt")
	}

	return e
}

func SetDrawMoneyPwd(UserId int, OldPwd, newPwd string) error {
	aUser, e := GInstance.GetUserRpcClient().GetUser(UserId)
	if e != nil {
		return e
	}
	aOldPwd := pwdKit.Sha1ToStr(OldPwd)

	if aUser.DrawMoneyPwd != "" && aUser.DrawMoneyPwd != aOldPwd {
		return errors.New(GTtHint.GetTtHint().GetHint("原密码不正确"))
	}
	if newPwd == "" {
		return errors.New(GTtHint.GetTtHint().GetHint("密码不能为空"))
	}
	aNewPwd := pwdKit.Sha1ToStr(newPwd)
	infos := make([]gBox.UpdateDataInfo, 0)
	infos = append(infos, gBox.UpdateDataInfo{FieldName: "DrawMoneyPwd", Value: aNewPwd, Type: 0})

	_, e = GInstance.GetUserRpcClient().UpdateUserInfo(UserId, infos)
	return e
}

type Report struct {
	AddScore float64 `json:"1"`
	DecScore float64 `json:"2"`
	BetM     float64 `json:"3"`
	WinM     float64 `json:"4"`
	//UpdatedAt int `json:"updated_at"`
}

func GetReport(gameType, userId int, BeginDay, EndDay string) Report {
	aUscReport := Report{}
	o := orm.NewOrm()

	sql := fmt.Sprintf(
		`SELECT sum(gold) from %s a where a.user_id = ? 
and date_format(a.updated_at,'%s') >= ? and date_format(a.updated_at,'%s') <= ? and a.state = ?`,
		mconst.TableName_TtSaveMoney,
		"%Y-%m-%d", "%Y-%m-%d")
	addS := 0.0
	e := o.Raw(sql, userId, BeginDay, EndDay, mconst.SaveMoneyState_5_OK).QueryRow(&addS)
	if e != nil {
		ttLog.LogError(e)
	}

	sql = fmt.Sprintf(
		`SELECT sum(gold) from %s a where a.user_id = ? 
and date_format(a.updated_at,'%s') >= ? and date_format(a.updated_at,'%s') <= ?`,
		mconst.TableName_TtDrawMoney,
		"%Y-%m-%d", "%Y-%m-%d")
	decS := 0.0
	e = o.Raw(sql, userId, BeginDay, EndDay).QueryRow(&decS)
	if e != nil {
		ttLog.LogError(e)
	}

	sqlArgs := make([]interface{}, 0)

	sqlWhere := fmt.Sprintf(` where a.user_id = ? and a.status=?
and date_format(a.created_at,'%s') >= ? and date_format(a.created_at,'%s') <= ?`,
		"%Y-%m-%d", "%Y-%m-%d")
	sqlArgs = append(sqlArgs, userId, mconst.Bet_Status_2, BeginDay, EndDay)

	if gameType != 0 {
		sqlWhere = " and game_type =? "
		sqlArgs = append(sqlArgs, gameType)
	}

	sql = fmt.Sprintf(`SELECT sum(bet_m) as bet_m,sum(win) as win from %s a  %s`,
		mconst.TableName_LoBetInfo, sqlWhere)

	type Tmp struct {
		BetM float64
		Win  float64
	}
	aTmpUsc := Tmp{}
	e = o.Raw(sql).QueryRow(&aTmpUsc)
	if e != nil {
		ttLog.LogError(e)
	}

	aUscReport.AddScore = addS
	aUscReport.DecScore = decS
	aUscReport.BetM = aTmpUsc.BetM
	aUscReport.WinM = aTmpUsc.Win

	return aUscReport
}

func GetAgentPayInfo(Area string) (models.TtAgentPayInfo, error) {
	data, e := models.GetAgentPayInfo(1)

	data.AlipayUrl = httpKit.GetImgUrl(GConfig.GetGConfig().SApiBaseP, data.AlipayUrl)
	data.WXReceiptUrl = httpKit.GetImgUrl(GConfig.GetGConfig().SApiBaseP, data.WXReceiptUrl)

	return data, e
}
func GetAgentPayInfoEx(Area string) (map[string]interface{}, error) {
	data := make(map[string]interface{}, 0)
	aAgentPayInfo, e := models.GetAgentPayInfo(1)
	if e != nil {
		return nil, e
	}
	data["PayWay"] = aAgentPayInfo.PayWay
	switch aAgentPayInfo.PayWay {
	case 0:
		data["WXReceiptUrl"] = httpKit.GetImgUrl(GConfig.GetGConfig().SApiBaseP, aAgentPayInfo.WXReceiptUrl)
	case 1:
		data["BankCard"] = aAgentPayInfo.BankCard
		data["BankName"] = aAgentPayInfo.BankName
		data["BankUser"] = aAgentPayInfo.BankUser
		data["UserMobile"] = aAgentPayInfo.UserMobile
		data["BankAddr"] = aAgentPayInfo.BankAddr
	case 2:
		data["AlipayUrl"] = httpKit.GetImgUrl(GConfig.GetGConfig().SApiBaseP, aAgentPayInfo.AlipayUrl)
		data["AlipayName"] = aAgentPayInfo.AlipayName
		data["OnlinePay"] = aAgentPayInfo.OnlinePay
	}

	return data, e
}

/**
上分
*/
func SaveMoney(UserId int, Money float64, AuditorId int, AuditorName string) error {
	o := orm.NewOrm()
	aTtSaveMoney := models.TtSaveMoney{UserId: UserId, Money: Money, Gold: float64(Money), State: mconst.SaveMoneyState_5_OK,
		VoucherUrl: "", PayState: 0, AuditorId: AuditorId, AuditorName: AuditorName}
	e := aTtSaveMoney.Add(o)
	if e != nil {
		return e
	}

	goldInfo := gBox.AddGoldInfo{UserId: UserId, Gold: float64(Money), T: userConst.Account_08_AddMoney,
		Des:  fmt.Sprintf("[%s]上分%g。", AuditorName, Money),
		Des2: fmt.Sprintf("[%s]Nạp điểm%g。", AuditorName, Money),
	}
	_, e = GInstance.GetUserRpcClient().AddGold(goldInfo)
	return e
}

/**
下分
*/
func DrawMoney(GroupId, UserId int, Money float64, AuditorId int, AuditorName string) error {
	o := orm.NewOrm()
	aTtDrawMoney := models.TtDrawMoney{GroupId: GroupId, UserId: UserId, Gold: Money, State: mconst.DrawMoneyState_4,
		AuditorId: AuditorId, AuditorName: AuditorName}
	e := aTtDrawMoney.Add(o)
	if e != nil {
		return e
	}

	goldInfo := gBox.AddGoldInfo{GroupId: GroupId, UserId: UserId, Gold: Money, T: userConst.Account_09_DecMoney,
		Des:  fmt.Sprintf("[%s]下分%g。", AuditorName, Money),
		Des2: fmt.Sprintf("[%s]Rút điểm%g。", AuditorName, Money),
	}
	_, e = GInstance.GetUserRpcClient().AddGold(goldInfo)
	return e
}

/**
下分
*/
func DrawMoneying(GroupId, UserId int, Score, Money float64, State int, OrderId string) error {
	switch State {
	case 1:
		o := orm.NewOrm()
		if Money == 0 {
			Money = Score
		}
		c, e := o.QueryTable(mconst.TableName_TtDrawMoney).Filter("OrderId", OrderId).Count()
		if e != nil {
			return e
		}
		if c > 0 {
			return errors.New("订单号已存在。")
		}
		if Score <= 0 {
			return errors.New("金额不能小于或等于0")
		}

		aTtDrawMoney := models.TtDrawMoney{GroupId: GroupId, UserId: UserId, Gold: Score, Money: Money, State: mconst.DrawMoneyState_2,
			AuditorId: 0, AuditorName: "", OrderId: OrderId}
		e = aTtDrawMoney.Add(o)
		if e != nil {
			return e
		}

		goldInfo := gBox.AddGoldInfo{GroupId: GroupId, UserId: UserId, Gold: Score, T: userConst.Account_09_DecMoney,
			Des:  fmt.Sprintf("[%s]下分%g。", "", Score),
			Des2: fmt.Sprintf("[%s]Rút điểm%g。", "", Score),
		}
		_, e = GInstance.GetUserRpcClient().AddGold(goldInfo)
		return e
	case 2:
		o := orm.NewOrm()
		arr := make([]models.TtDrawMoney, 0)
		o.QueryTable(mconst.TableName_TtDrawMoney).Filter("OrderId", OrderId).Filter("UserId", UserId).Filter("GroupId", GroupId).All(&arr)
		if len(arr) != 1 {
			return errors.New("数据不唯一。")
		}
		arr[0].State = mconst.DrawMoneyState_4
		e := arr[0].Update(o, "State", "UpdatedAt")

		return e
	case 3:
		o := orm.NewOrm()
		arr := make([]models.TtDrawMoney, 0)
		o.QueryTable(mconst.TableName_TtDrawMoney).Filter("OrderId", OrderId).Filter("UserId", UserId).Filter("GroupId", GroupId).All(&arr)
		if len(arr) != 1 {
			return errors.New("数据不唯一。")
		}

		aGoldInfo := gBox.AddGoldInfo{GroupId: GroupId, UserId: arr[0].UserId, Gold: arr[0].Gold,
			T:    userConst.Account_07_DrawMoneyR,
			Des:  fmt.Sprintf("提现拒绝%0.2f", arr[0].Gold),
			Des2: fmt.Sprintf("Rút điểm bị từ chối%0.2f", arr[0].Gold),
		}
		_, e := GInstance.GetUserRpcClient().AddGold(aGoldInfo)
		if e != nil {
			return errors.New(GTtHint.GetTtHint().GetHint("没有对应的数据"))
		}
		_, e = o.Delete(&arr[0])

		return e
	}

	return errors.New("有问题的数据类型")
}
