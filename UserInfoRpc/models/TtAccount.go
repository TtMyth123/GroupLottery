package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"ttmyth123/GroupLottery/UserInfoRpc/models/mconst"
)

/**
流水账目
*/
type TtAccount struct {
	Id          int
	UserId      int
	GroupId 	int
	AccountType int
	StrType     string `orm:"size(200)"`
	Des         string `orm:"size(512)"`
	CurUserGold float64
	Gold        float64
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	RefId       int

	Des2  string `orm:"size(512)"`
	DesMp string `orm:"size(512)"`
}

func (this *TtAccount) TableName() string {
	return mconst.TableName_TtAccount
}

func (this *TtAccount) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	id, e := o.Insert(this)
	this.Id = int(id)
	return e
}
