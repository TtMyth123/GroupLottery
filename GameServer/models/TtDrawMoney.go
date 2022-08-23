package models

import (
	"encoding/json"
	"fmt"
	"github.com/TtMyth123/GameServer/GConfig"
	"github.com/TtMyth123/GameServer/models/mconst"
	"github.com/TtMyth123/kit/sqlKit"
	"github.com/TtMyth123/kit/timeKit"
	"github.com/astaxie/beego/orm"
	"time"
)

/**
提现数据
*/
type TtDrawMoney struct {
	Id          int
	UserId      int
	GroupId     int
	Gold        float64
	Money       float64
	State       int
	SN          string `orm:"size(40)"`
	AuditorId   int    //审核人Id
	AuditorName string `orm:"size(256)"` //审核人名称
	OrderId     string `orm:"size(32)"`

	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func (d TtDrawMoney) MarshalJSON() ([]byte, error) {
	type Alias TtDrawMoney

	TimeLag := GConfig.GetGConfig().TimeLag
	StrUpdatedAt := ""
	if TimeLag == 0 {
		StrUpdatedAt = d.UpdatedAt.Format(timeKit.DateTimeLayout)
	} else {
		StrUpdatedAt = d.UpdatedAt.Add(TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
	}

	return json.Marshal(&struct {
		Alias
		StrCreatedAt string
	}{
		Alias:        (Alias)(d),
		StrCreatedAt: StrUpdatedAt,
	})
}

func (this *TtDrawMoney) TableName() string {
	return mconst.TableName_TtDrawMoney
}

func (this *TtDrawMoney) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	id, e := o.Insert(this)
	this.Id = int(id)
	return e
}

func (this *TtDrawMoney) Update(o orm.Ormer, cols ...string) error {
	this.UpdatedAt = time.Now()
	_, e := o.Update(this, cols...)

	return e
}

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
}

func (d TtDrawMoneyEx) MarshalJSON() ([]byte, error) {
	type Alias TtDrawMoneyEx

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

func DrawMoneyListByPage(beginDay, endDay, userName string, userId, pageIndex, pageSize int) (int, []TtDrawMoneyEx) {
	PageTotal := 0
	o := orm.NewOrm()
	arrData := make([]TtDrawMoneyEx, 0)
	sqlArgs := make([]interface{}, 0)

	sqlWhere := " and a.state<>?"
	sqlArgs = append(sqlArgs, mconst.DrawMoneyState_1_Apply)
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
tt_draw_money a, tt_game_user b where a.user_id=b.id %s`, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&PageTotal)
	if e != nil {
		return PageTotal, arrData
	}

	offset, _ := sqlKit.GetOffset(PageTotal, pageSize, pageIndex)
	sqlWhere = sqlWhere + ` order by a.created_at DESC LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, pageSize)

	sql := fmt.Sprintf(`select a.*, b.user_name from 
tt_draw_money a, tt_game_user b where a.user_id=b.id %s `,
		sqlWhere)
	o.Raw(sql, sqlArgs).QueryRows(&arrData)

	return PageTotal, arrData
}
