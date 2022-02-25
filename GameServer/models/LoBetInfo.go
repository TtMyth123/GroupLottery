package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"ttmyth123/GroupLottery/GameServer/models/mconst"
)

type LoBetInfo struct {
	Id          int
	GroupId     int     //主ID
	GroupBetSn  string  `orm:"size(56)"`   //主投注流水号
	BetSn       string  `orm:"size(56)"`   //投注流水号
	BetStr      string  `orm:"size(1024)"` //投注说明
	Status      int     //状态
	Period      int64   //期号
	StrPeriod   string  `orm:"size(20)"` //期号
	Odds        float64 //赔率
	StrOdds     string  `orm:"size(1024)"` //赔率,{"2":1,"3":5,"4":40}
	OddsName    string  `orm:"size(1000)"`
	OddsType    int     //赔率类型
	BigOddsType int     //大赔率类型
	Nums        string
	BetM        int       //投注金额
	Win         float64   //赢得金额
	RcRatio     float64   //返现比例
	RcMoney     float64   //返现金额
	ResultNums  string    `orm:"size(1024)"`
	GameType    int       //游戏类型
	GameName    string    `orm:"size(100)"`
	UserId      int       //用户ID
	GroupUserId int
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time `orm:"auto_now;type(datetime)"`
}

func (this *LoBetInfo) TableName() string {
	return mconst.TableName_LoBetInfo
}
func (this *LoBetInfo) Add(o orm.Ormer) error {
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

	return e
}

func (this *LoBetInfo) Update(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	this.UpdatedAt = time.Now()
	arrC := []string{"UpdatedAt"}
	arrC = append(arrC, cols...)

	id, e := o.Update(this, arrC...)
	if e == nil {
		this.Id = int(id)
	}

	return e
}

func InsertMultiBetInfo(o orm.Ormer, arrBetInfo []LoBetInfo) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.InsertMulti(len(arrBetInfo), arrBetInfo)
	return e
}
