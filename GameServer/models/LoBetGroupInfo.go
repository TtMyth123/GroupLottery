package models

import (
	"fmt"
	"github.com/TtMyth123/GameServer/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

type LoBetGroupInfo struct {
	Id          int
	BetSn       string `orm:"size(56)"`
	BetStr      string `orm:"size(1024)"`
	Status      int
	Period      int64
	StrPeriod   string `orm:"size(20)"`
	ResultNums  string `orm:"size(1024)"`
	BetM        int
	Win         float64
	GameType    int
	UserId      int
	GroupUserId int
	GameName    string    `orm:"size(100)"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time `orm:"auto_now;type(datetime)"`
}

func (this *LoBetGroupInfo) TableName() string {
	return mconst.TableName_LoBetGroupInfo
}
func (this *LoBetGroupInfo) Add(o orm.Ormer) (int, error) {
	if o == nil {
		o = orm.NewOrm()
	}
	this.GameName = mconst.GetGameName(this.GameType)
	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	if e == nil {
		this.Id = int(id)
	}
	this.BetSn = fmt.Sprintf("%s%06d", this.CreatedAt.Format("060102"), this.Id)
	this.Update(o, "BetSn")

	return int(id), e
}

func (this *LoBetGroupInfo) Update(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	this.UpdatedAt = time.Now()
	id, e := o.Update(this, cols...)
	if e == nil {
		this.Id = int(id)
	}

	return e
}

func GeLoBetGroupInfo(o orm.Ormer, Id int) (LoBetGroupInfo, error) {
	aLoBetGroupInfo := LoBetGroupInfo{Id: Id}
	e := o.Read(&aLoBetGroupInfo)
	return aLoBetGroupInfo, e
}
