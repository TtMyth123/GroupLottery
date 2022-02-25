package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"ttmyth123/GroupLottery/UserInfoRpc/models/mconst"
	"ttmyth123/kit/pwdKit"
	"ttmyth123/kit/timeKit"
	"ttmyth123/kit/ttLog"
)

type TtGameUser struct {
	Id            int
	Pid           int    `orm:"size(512)"` //父Id. 推荐人Id
	SPids         string //父Id. 推荐人Id
	HeadImgurl    string `orm:"size(512)"`
	UserName      string `orm:"size(512)"`
	Pwd           string `json:"-"`
	Pwd2          string `json:"-"`
	Nickname      string `orm:"size(256)"`
	LowerCount    int    //直接下级会员数
	AllLowerCount int    //全部的下级有会员数
	ReferrerCode  int    //推荐码
	IsReferrer    int    //是否推手
	State         int
	Gold          float64 //当前余额(金币)
	Silver        float64 //当前余额(银币)
	Copper        float64 //当前余额(铜币)

	SumSaveMoney float64 //累计充值金币
	MaxSaveMoney float64 //最大充值金币
	SumDrawMoney float64 //累计提现金币
	MaxDrawMoney float64 //最大提现金币
	SumBet       float64 //全部投注金额
	SumXmBet     float64 //全部洗码金额
	SumWin       float64 //全部赢得金额
	SumAddMoney  float64 //累计上分金额
	MaxAddMoney  float64 //最大上分金币
	SumDecMoney  float64 //累计下分金额
	MaxDecMoney  float64 //最大下分金额
	Sum2Rebate   float64 //累计给上级的返佣
	SumRebate    float64 //累计获取返佣
	Rebate       float64 //可提佣金
	MemberLevel  int

	SumTime time.Time `orm:"auto_now_add;type(datetime)"`

	CreatedAt    time.Time `orm:"auto_now_add;type(datetime)"`
	LoginTime    time.Time `orm:"auto_now_add;type(datetime)"`
	LoginOutTime time.Time `orm:"auto_now_add;type(datetime)"`

	VoucherFile  string
	WXSKCodeUrl  string //微信收款码Url
	YHName       string //银行名
	CardNum      string //卡号
	YHUserName   string //银行用户名
	YHUserTel    string //银行预留电话
	Addr         string //银行预留地址
	Cate         string //身份证
	Remark       string //银行预留信息
	ZFBSKCodeUrl string //支付宝二维码Url
	ZFBSKName    string //支付宝名

	Sid string `orm:"-`

	IdentityCard  string `orm:"size(255)"`
	FullName      string `orm:"size(255)"`
	Tel           string `orm:"size(255)"`
	RealNameState int

	UserType     int
	DrawMoneyPwd string `orm:"size(255)"`
	Area         string
	AreaId       int
	IsAgent      int
	AgentUserId  int
}

func (this *TtGameUser) TableName() string {
	return mconst.TableName_TtGameUser
}

func GetUserByCode(code int) (TtGameUser, error) {
	o := orm.NewOrm()
	aYhGameUser := TtGameUser{}
	e := o.QueryTable(mconst.TableName_TtGameUser).Filter("ReferrerCode", code).One(&aYhGameUser)
	return aYhGameUser, e
}

func InitUser() error {
	o := orm.NewOrm()
	c, _ := o.QueryTable(mconst.TableName_TtGameUser).Count()
	if c == 0 {
		RootReferrerCode, _ := beego.AppConfig.Int("RootReferrerCode")
		aYhGameUser := TtGameUser{Id: mconst.RootGameUserId, Pid: 0, SPids: "0", State: mconst.User_State_1,
			ReferrerCode: RootReferrerCode, UserType: mconst.UserType_3, IsAgent: 1}
		_, e := o.Insert(&aYhGameUser)
		if e != nil {
			ttLog.LogError("aaaaaaaaaaaa:", e)
			return e
		}

		aYhGameUser = TtGameUser{Id: 20000, Pid: 0, SPids: "0", ReferrerCode: 0}
		_, e = o.Insert(&aYhGameUser)
		if e != nil {
			ttLog.LogError("aaaaaaaaaaaa:", e)
			return e
		}

		_, e = o.Delete(&aYhGameUser)
		if e != nil {
			ttLog.LogError("aaaaaaaaaaaa:", e)
			return e
		}

		arr := make([]TtGameUser, 0)
		for i := 100; i < 900; i++ {
			aGameUser := TtGameUser{Id: i, State: mconst.User_State_1, UserType: mconst.UserType_3,
				Nickname: fmt.Sprintf("游客%d", i), UserName: fmt.Sprintf("游客%d", i)}
			//o.Insert(&aGameUser)
			arr = append(arr, aGameUser)
		}

		_, e = o.InsertMulti(len(arr), &arr)
		if e != nil {
			ttLog.LogError("aaaaaaaaaaaa:", e)
			return e
		}
	}
	return nil
}

func GetReferralCode(o orm.Ormer) int {
	if o == nil {
		o = orm.NewOrm()
	}
	r := timeKit.GlobaRand.Int63n(99999)
	for i := 0; i < 99999; i++ {
		c, _ := o.QueryTable(mconst.TableName_TtGameUser).Filter("ReferrerCode", r).Count()
		if c == 0 {
			return int(r)
		}
		c++
		if c > 99999 {
			c = 0
		}
	}

	return 0
}
func (this *TtGameUser) Add(o orm.Ormer, MarketID int) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.ReferrerCode = MarketID
	id, e := o.Insert(this)
	this.Id = int(id)
	return e
}
func (this *TtGameUser) Update(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.Update(this, cols...)
	return e
}
func (this *TtGameUser) Del() error {
	o := orm.NewOrm()
	_, err := o.Delete(this)
	return err
}

func AddGameUser(aArea TtArea, UserId int, UserName, pwd, MoneyPwd string, Code, MarketID, UserType int, InfoEx map[string]string) (TtGameUser, error) {
	ttLog.LogDebug("添加用户B:", aArea, UserId, UserName, pwd, MoneyPwd, Code, MarketID, UserType, InfoEx)
	o := orm.NewOrm()
	aYhGameUser := TtGameUser{Id: UserId}
	c, e := o.QueryTable(mconst.TableName_TtGameUser).Filter("UserName", UserName).Count()
	if e != nil {
		return aYhGameUser, e
	}
	if c != 0 {
		ttLog.LogDebug("用户名已存在:", UserName)
		return aYhGameUser, errors.New("用户名已存在")
	}

	pYhGameUser := TtGameUser{}
	e = o.QueryTable(mconst.TableName_TtGameUser).Filter("ReferrerCode", Code).One(&pYhGameUser)
	if e != nil {
		ttLog.LogError("推荐人不存在:", Code)
		return aYhGameUser, errors.New("推荐人不存在")
	}

	if UserType == 0 {
		aYhGameUser.UserType = mconst.UserType_1
	} else {
		aYhGameUser.UserType = UserType
	}
	aYhGameUser.UserName = UserName
	aYhGameUser.Pwd = pwdKit.Sha1ToStr(pwd)
	aYhGameUser.Nickname = UserName
	aYhGameUser.Pid = pYhGameUser.Id
	aYhGameUser.State = mconst.User_State_1
	aYhGameUser.DrawMoneyPwd = MoneyPwd

	aYhGameUser.Area = aArea.Area
	aYhGameUser.AreaId = aArea.Id

	for k, v := range InfoEx {
		switch k {
		case "Tel":
			aYhGameUser.Tel = v
		case "YHName":
			aYhGameUser.YHName = v
		case "CardNum":
			aYhGameUser.CardNum = v
		case "YHUserName":
			aYhGameUser.YHUserName = v
		case "YHUserTel":
			aYhGameUser.YHUserTel = v
		case "Addr":
			aYhGameUser.Addr = v
		case "Cate":
			aYhGameUser.Cate = v
		case "Remark":
			aYhGameUser.Remark = v
		}
	}
	aYhGameUser.AgentUserId = getAgentId(o, pYhGameUser)
	e = aYhGameUser.Add(o, MarketID)
	AddLowerCount(o, aYhGameUser.Pid)
	AddAllLowerCount(o, aYhGameUser.Pid)
	ttLog.LogDebug("添加用户:", aYhGameUser)
	return aYhGameUser, e
}
func getAgentId(o orm.Ormer, pYhGameUser TtGameUser) int {
	if pYhGameUser.IsAgent == 1 {
		return pYhGameUser.Id
	}
	aYhGameUser := TtGameUser{Id: pYhGameUser.Pid}
	e := o.Read(&aYhGameUser)
	if e != nil {
		return 0
	}
	return getAgentId(o, aYhGameUser)
}

func AddLowerCount(o orm.Ormer, pid int) {
	if o == nil {
		o = orm.NewOrm()
	}
	sql := `update tt_game_user set lower_count = lower_count+1 where id = ?`
	o.Raw(sql, pid).Exec()
}

func AddAllLowerCount(o orm.Ormer, pid int) {
	if pid == 0 {
		return
	}
	aYhGameUser := TtGameUser{Id: pid}
	e := o.Read(&aYhGameUser)
	if e == nil {
		sql := `update tt_game_user set all_lower_count = all_lower_count+1 where id = ?`
		o.Raw(sql, pid).Exec()
		AddAllLowerCount(o, aYhGameUser.Pid)
	}
}

func GetUserInfo(userId int) (TtGameUser, error) {
	o := orm.NewOrm()
	aGameUser := TtGameUser{Id: userId}
	e := o.Read(&aGameUser)
	return aGameUser, e
}

func GetGameUserByNamePwd(userName, pwd string) (TtGameUser, error) {
	o := orm.NewOrm()
	newPwd := pwdKit.Sha1ToStr(pwd)
	aYhGameUser := TtGameUser{}
	e := o.QueryTable(mconst.TableName_TtGameUser).Filter("UserName", userName).Filter("Pwd", newPwd).One(&aYhGameUser)
	if e != nil {
		return aYhGameUser, errors.New("用户名或密码不正确。")
	}
	return aYhGameUser, nil
}
