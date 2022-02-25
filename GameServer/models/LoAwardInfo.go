package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"ttmyth123/GroupLottery/GameServer/models/mconst"
)

type LoAwardInfo struct {
	Id             int
	GameType       int
	LotteryNum     int64  //期号
	LotteryStr     string //期号
	ResultNums     string `orm:"size(1024)"` //开奖结果数据
	OriginalResult string `orm:"size(1024)"` //原始接口数据

	NextLotteryStr  string
	NextLotteryTime time.Time //下一期开奖时间
	CurLotteryTime  time.Time //当前开期时间
	NextTime        string

	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func (this *LoAwardInfo) TableName() string {
	return mconst.TableName_LoAwardInfo
}
func (this *LoAwardInfo) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	if e == nil {
		this.Id = int(id)
	}

	return e
}

func (this *LoAwardInfo) Update(o orm.Ormer, cols ...string) error {
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
