package regBll

import (
	"errors"
	"fmt"
	"github.com/TtMyth123/GameServer/GInstance"
	"github.com/TtMyth123/GameServer/GInstance/GTtHint"
	"github.com/TtMyth123/GameServer/models"
	"github.com/TtMyth123/GameServer/models/mconst"
	"github.com/TtMyth123/UserInfoRpc/GData/gBox"
	userModels "github.com/TtMyth123/UserInfoRpc/models"
	userMconst "github.com/TtMyth123/UserInfoRpc/models/mconst"
	"github.com/astaxie/beego/orm"
)

func Reg(Area, UserName, pwd, MoneyPwd string, Code, UserType int, mpFiled map[string]string) (userModels.TtGameUser, error) {
	u, e := GInstance.GetUserRpcClient().AddUser(Area, UserName, pwd, MoneyPwd, Code, UserType, mpFiled)
	if e == nil {
		aTtDrawSaveSet := models.GetTtDrawSaveSet()

		//aTtDrawSaveSet.RegGiveCount
		des := fmt.Sprintf("注册奖励金额%d。", aTtDrawSaveSet.RegGiveCount)
		goldInfo := gBox.AddGoldInfo{GroupId: 0, UserId: u.Id, Gold: float64(aTtDrawSaveSet.RegGiveCount),
			T: userMconst.Account_05_Give, Des: des,
			Des2: GTtHint.GetTtHint().GetHint("注册奖励金额%d。"), DesMp: GTtHint.GetTtHint().GetMpString(aTtDrawSaveSet.RegGiveCount)}
		u, e = GInstance.GetUserRpcClient().AddGold(goldInfo)
	}

	return u, e
}

func RegUser2(Area, UserName, Pwd, MoneyPwd, Tel, YHName, CardNum, YHUserName, YHUserTel, Addr, Cate, Remark string, ReferrerCode int) (models.LoUserInfo2, error) {
	aLoUserInfo2 := models.LoUserInfo2{
		UserName:     UserName,
		Pwd:          Pwd,
		Tel:          Tel,
		YHName:       YHName,
		CardNum:      CardNum,
		YHUserName:   YHUserName,
		YHUserTel:    YHUserTel,
		Addr:         Addr,
		Cate:         Cate,
		Remark:       Remark,
		ReferrerCode: ReferrerCode,
		MoneyPwd:     MoneyPwd,
	}
	if UserName == "" || Tel == "" || Pwd == "" {
		return aLoUserInfo2, errors.New("用户名,电话,密码不能为空")
	}
	o := orm.NewOrm()
	c, e := o.QueryTable(mconst.TableName_LoUserInfo2).Filter("UserName", UserName).Count()
	if c != 0 {
		return aLoUserInfo2, errors.New("用户名已存在")
	}

	aPLoUserInfo2 := models.LoUserInfo2{}
	e = o.QueryTable(mconst.TableName_LoUserInfo2).Filter("ReferrerCode", ReferrerCode).One(&aPLoUserInfo2)
	if e != nil {
		return aLoUserInfo2, errors.New("推荐人不存在")
	}
	e = aLoUserInfo2.Add(nil)

	return aLoUserInfo2, e
}
