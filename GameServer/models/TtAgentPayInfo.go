package models

import (
	"github.com/astaxie/beego/orm"
	"ttmyth123/GroupLottery/GameServer/models/mconst"
	"ttmyth123/kit/ttLog"
)

type TtAgentPayInfo struct {
	Id          int
	AddIsTs     int     //添加用户是推手
	PayBetRatio float64 //充值投注比例

	PayWay int //0:微信，1:银行，2：支付宝

	OnlinePay  string `orm:"size(255)"` //在线支付号(支付宝号)
	AlipayName string `orm:"size(255)"` //在线支付号(用户姓名)
	AlipayUrl  string `orm:"size(255)"` //在线支付号(用户姓名)

	BankCard   string `orm:"size(255)"` //银行卡号
	BankName   string `orm:"size(255)"` //开户银行
	BankUser   string `orm:"size(255)"` //账号名称(姓名)
	UserMobile string `orm:"size(40)"`  //手机号
	BankAddr   string `orm:"size(255)"` //开户地址

	WXReceiptUrl string `orm:"size(255)"` //微信URL支付
}

func (this *TtAgentPayInfo) TableName() string {
	return mconst.TableName_TtAgentPayInfo
}

func InitAgentPayInfo() error {
	o := orm.NewOrm()
	c, _ := o.QueryTable(mconst.TableName_TtAgentPayInfo).Count()
	if c == 0 {
		aAgentPayInfo := TtAgentPayInfo{Id: 1}

		_, e := o.Insert(&aAgentPayInfo)
		return e
	}
	return nil
}

func GetAgentPayInfo(id int) (TtAgentPayInfo, error) {
	aAgentPayInfo := TtAgentPayInfo{}
	aAgentPayInfo.Id = id

	o := orm.NewOrm()
	e := o.Read(&aAgentPayInfo)
	if e != nil {
		ttLog.LogError(e)
	}
	return aAgentPayInfo, e
}

func (this *TtAgentPayInfo) Update(col ...string) error {
	o := orm.NewOrm()
	_, e := o.Update(this, col...)
	if e != nil {
		ttLog.LogError(e)
	}
	return e
}
