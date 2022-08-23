package controllers

import (
	"github.com/TtMyth123/GameServer/CacheData"
	"github.com/TtMyth123/GameServer/GInstance"
	"github.com/TtMyth123/GameServer/GInstance/GTtHint"
	"github.com/TtMyth123/GameServer/controllers/base"
	"github.com/TtMyth123/GameServer/controllers/base/enums"
	"github.com/TtMyth123/GameServer/controllers/regBll"
	"github.com/TtMyth123/UserInfoRpc/models/mconst"
	"time"
)

type RegLoginController struct {
	base.BaseController
}

/**
注册
Code：推荐码。Pwd：密码。UserName：用户名。
*/
func (this *RegLoginController) Reg() {
	type TmpArgs struct {
		base.ArgsBox
		Code     int
		Pwd      string
		UserName string
		UserType int

		Tel          string `orm:"size(256);column(tel)" description:"电话"`
		YHName       string `orm:"size(512);column(y_h_name)" description:"银行名"`             //银行名
		CardNum      string `orm:"size(512);column(card_num)" description:"卡号"`              //卡号
		YHUserName   string `orm:"size(512);column(y_h_user_name)" description:"银行用户名"`      //银行用户名
		YHUserTel    string `orm:"size(512);column(y_h_user_tel)" description:"银行预留电话"`      //银行预留电话
		Addr         string `orm:"size(512);column(addr)" description:"银行预留地址"`              //银行预留地址
		Cate         string `orm:"size(512);column(cate)" description:"身份证"`                 //身份证
		Remark       string `orm:"size(512);column(remark)" description:"银行预留信息 (备注)"`       //银行预留信息
		DrawMoneyPwd string `orm:"size(512);column(draw_money_pwd)" description:"资金密码（提现密码）` //银行预留信息
	}

	args := TmpArgs{}
	e := this.GetJsonData(&args)
	if e != nil {
		this.ErrJsonResultEx(e)
		return
	}

	mapFiledName := make(map[string]string)
	mapFiledName["Tel"] = args.Tel
	mapFiledName["YHName"] = args.YHName
	mapFiledName["CardNum"] = args.CardNum
	mapFiledName["YHUserName"] = args.YHUserName
	mapFiledName["YHUserTel"] = args.YHUserTel
	mapFiledName["Addr"] = args.Addr
	mapFiledName["Cate"] = args.Cate
	mapFiledName["Remark"] = args.Remark

	aYhGameUser, e := regBll.Reg(args.Area, args.UserName, args.Pwd, args.DrawMoneyPwd, args.Code, args.UserType, mapFiledName)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	this.JsonResult(enums.JRCodeSucc, "", aYhGameUser, nil)
}

/**
登录
UserName：用户名。Pwd：密码。VisitorId:游客ID，0返回一下新的游客。IsVisitor：1表示游客登录。（这时用户名与密码无效）
返回:
TtGameUser:用户信息
*/
func (this *RegLoginController) Login() {
	type TmpArgs struct {
		Pwd       string
		UserName  string
		VisitorId int
		IsVisitor int
	}

	args := TmpArgs{}
	e := this.GetJsonData(&args)
	if e != nil {
		this.ErrJsonResultEx(e)
		return
	}

	if args.IsVisitor == 1 {
		u, e := CacheData.GetVisitor(args.VisitorId)

		if e != nil {
			this.ErrJsonResultEx(e)
		} else {
			this.JsonResult(enums.JRCodeSucc, "", u, nil)
		}
	} else {
		aUser, e := GInstance.GetUserRpcClient().GetUserByNamePwd(args.UserName, args.Pwd)
		if e != nil {
			this.JsonResult(enums.JRCodeFailed, GTtHint.GetTtHint().GetHint(e.Error()), "", nil)
			//this.ErrJsonResultEx(e)
		} else {
			if aUser.State == mconst.User_State_2 {
				this.JsonResult(enums.JRCodeFailed, GTtHint.GetTtHint().GetHint("被禁用无法登录"), "", nil)
			}
			aUser.LoginTime = time.Now()
			aUser.Update(nil, "LoginTime")
			this.JsonResult(enums.JRCodeSucc, "", aUser, nil)
		}
	}

}

/**
登出
*/
func (this *RegLoginController) Logout() {
	type TmpArgs struct {
		UserId    int
		IsVisitor int
	}

	args := TmpArgs{}
	e := this.GetJsonData(&args)
	if e != nil {
		this.ErrJsonResultEx(e)
		return
	}
	CacheData.ClearVisitor(args.UserId)
	CacheData.DelUserSid(args.UserId)

	this.JsonResult(enums.JRCodeSucc, "", "", nil)
}

func (this *RegLoginController) LostConnection() {
	type TmpArgs struct {
		base.ArgsBox
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", "", nil)
}

//func (this *RegLoginController) RegUser2() {
//	type TmpArgs struct {
//		base.ArgsBox
//		UserName     string `orm:"size(512);column(user_name)" description:"用户名"`
//		ReferrerCode int    `orm:"column(referrer_code)" description:"推荐人ID"` //推荐码
//		Pwd          string `orm:"size(256);column(pwd)" description:"密码"`
//		Tel          string `orm:"size(256);column(tel)" description:"电话"`
//		YHName       string `orm:"size(512);column(y_h_name)" description:"银行名"`        //银行名
//		CardNum      string `orm:"size(512);column(card_num)" description:"卡号"`         //卡号
//		YHUserName   string `orm:"size(512);column(y_h_user_name)" description:"银行用户名"` //银行用户名
//		YHUserTel    string `orm:"size(512);column(y_h_user_tel)" description:"银行预留电话"` //银行预留电话
//		Addr         string `orm:"size(512);column(addr)" description:"银行预留地址"`         //银行预留地址
//		Cate         string `orm:"size(512);column(cate)" description:"身份证"`            //身份证
//		Remark       string `orm:"size(512);column(remark)" description:"银行预留信息 (备注)"`  //银行预留信息
//		MoneyPwd 	string
//	}
//
//	args := TmpArgs{}
//	e := this.GetJsonData(&args)
//	if e != nil {
//		this.ErrJsonResultEx(e)
//		return
//	}
//
//	aUser, e := regBll.RegUser2(args.Area,args.UserName, args.Pwd,args.MoneyPwd,
//		args.Tel,args.YHName,args.CardNum,args.YHUserName,args.YHUserTel,args.Addr,args.Cate,args.Remark,
//		args.ReferrerCode)
//	if e != nil {
//		this.ErrJsonResultEx(e)
//	}
//
//	this.JsonResult(enums.JRCodeSucc, "", aUser, nil)
//}
