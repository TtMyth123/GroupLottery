package models

import (
	"github.com/TtMyth123/UserInfoRpc/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

/**
流水账目
*/
type TtRebateInfo struct {
	Id            int
	UserId        int
	UserSid       int
	RebateType    int
	StrType       string `orm:"size(200)"`
	Des           string `orm:"size(512)"`
	Gold          float64
	Rebate        float64
	CurUserRebate float64

	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	RefId     int
	Des2      string `orm:"size(512)"`
	DesMp     string `orm:"size(512)"`
}

func (this *TtRebateInfo) TableName() string {
	return mconst.TableName_TtRebateInfo
}

func (this *TtRebateInfo) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	id, e := o.Insert(this)
	this.Id = int(id)
	return e
}
