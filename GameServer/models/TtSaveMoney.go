package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"ttmyth123/GroupLottery/GameServer/GConfig"
	"ttmyth123/GroupLottery/GameServer/models/mconst"
	"ttmyth123/kit/sqlKit"
	"ttmyth123/kit/timeKit"
)

/**
充值数据
*/
type TtSaveMoney struct {
	Id          int
	UserId      int
	GroupId int
	Money       float64
	Gold        float64
	State       int
	VoucherUrl  string    `orm:"size(512)"` //支付凭证
	PayState    int       //支付状态（只是用来接收客户端消息，不做是否支付判断）
	AuditorId   int       //审核人Id
	AuditorName string    `orm:"size(256)"` //审核人名称
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time `orm:"auto_now;type(datetime)"`
}

func (d TtSaveMoney) MarshalJSON() ([]byte, error) {
	type Alias TtSaveMoney

	TimeLag := GConfig.GetGConfig().TimeLag
	StrCreatedAt := ""
	StrUpdatedAt := ""
	if TimeLag == 0 {
		StrCreatedAt = d.CreatedAt.Format(timeKit.DateTimeLayout)
		StrUpdatedAt = d.UpdatedAt.Format(timeKit.DateTimeLayout)
	} else {
		StrCreatedAt = d.CreatedAt.Add(TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
		StrUpdatedAt = d.UpdatedAt.Add(TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
	}

	return json.Marshal(&struct {
		Alias
		StrCreatedAt string
		StrUpdatedAt string
	}{
		Alias:        (Alias)(d),
		StrCreatedAt: StrCreatedAt,
		StrUpdatedAt: StrUpdatedAt,
	})
}

func (this *TtSaveMoney) TableName() string {
	return mconst.TableName_TtSaveMoney
}

func (this *TtSaveMoney) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = int(id)
	return e
}

func (this *TtSaveMoney) Update(o orm.Ormer, cols ...string) error {
	_, e := o.Update(this, cols...)
	return e
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
}

func (d TtSaveMoneyEx) MarshalJSON() ([]byte, error) {
	type Alias TtSaveMoneyEx

	StrTime := ""
	TimeLag := GConfig.GetGConfig().TimeLag
	if TimeLag == 0 {
		StrTime = d.UpdatedAt.Format(timeKit.DateTimeLayout)
	} else {
		StrTime = d.UpdatedAt.Add(TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
	}

	return json.Marshal(&struct {
		Alias
		UpdatedAt string
	}{
		Alias:     (Alias)(d),
		UpdatedAt: StrTime,
	})
}

func SaveMoneyListByPage(beginDay, endDay, userName string, userId, pageIndex, pageSize int) (int, []TtSaveMoneyEx) {
	PageTotal := 0
	o := orm.NewOrm()
	arrData := make([]TtSaveMoneyEx, 0)
	sqlArgs := make([]interface{}, 0)

	sqlWhere := " and a.state<>? "
	sqlArgs = append(sqlArgs, mconst.SaveMoneyState_1_Apply)
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

	sqlCount := fmt.Sprintf(`select count(1) c from 
tt_save_money a, tt_game_user b where a.user_id=b.id %s`, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&PageTotal)
	if e != nil {
		return PageTotal, arrData
	}

	offset, _ := sqlKit.GetOffset(PageTotal, pageSize, pageIndex)
	sqlWhere = sqlWhere + ` LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, pageSize)

	sql := fmt.Sprintf(`select a.*, b.user_name from 
tt_save_money a, tt_game_user b where a.user_id=b.id %s `,
		sqlWhere)
	o.Raw(sql, sqlArgs).QueryRows(&arrData)

	return PageTotal, arrData
}
