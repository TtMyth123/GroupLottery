package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"ttmyth123/GroupLottery/GameServer/models/mconst"
)

//LoUserInfo2
//电话，用户名，密码，推荐人，身份证（可选），银行卡信息（可选）

type LoUserInfo2 struct {
	Id           int       `orm:"column(id)" description:"ID"`
	UserName     string    `orm:"size(512);column(user_name)" description:"用户名"`
	ReferrerCode int       `orm:"column(referrer_code)" description:"推荐人ID"` //推荐码
	Pwd          string    `orm:"size(256);column(pwd)" description:"密码"`
	Tel          string    `orm:"size(256);column(tel)" description:"电话"`
	YHName       string    `orm:"size(512);column(y_h_name)" description:"银行名"`        //银行名
	CardNum      string    `orm:"size(512);column(card_num)" description:"卡号"`         //卡号
	YHUserName   string    `orm:"size(512);column(y_h_user_name)" description:"银行用户名"` //银行用户名
	YHUserTel    string    `orm:"size(512);column(y_h_user_tel)" description:"银行预留电话"` //银行预留电话
	Addr         string    `orm:"size(512);column(addr)" description:"银行预留地址"`         //银行预留地址
	Cate         string    `orm:"size(512);column(cate)" description:"身份证"`            //身份证
	Remark       string    `orm:"size(512);column(remark)" description:"银行预留信息 (备注)"`  //银行预留信息
	MoneyPwd     string    `orm:"size(256);column(money_pwd)" description:"资金密码"`
	CreatedAt    time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt    time.Time `orm:"auto_now;type(datetime)"`
}

func (this *LoUserInfo2) TableName() string {
	return mconst.TableName_LoUserInfo2
}
func (this *LoUserInfo2) Add(o orm.Ormer) error {
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

func (this *LoUserInfo2) Update(o orm.Ormer, cols ...string) error {
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
