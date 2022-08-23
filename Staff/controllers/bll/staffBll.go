package bll

import (
	"errors"
	"github.com/TtMyth123/Staff/GInstance"
	"github.com/TtMyth123/Staff/OtherServer/httpGameServer"
	"github.com/TtMyth123/UserInfoRpc/GData/gBox"
	"github.com/TtMyth123/UserInfoRpc/models"
	"github.com/TtMyth123/UserInfoRpc/models/mconst"
	"github.com/TtMyth123/kit/pwdKit"
	"github.com/astaxie/beego/orm"
)

func Login(userName, pwd string) (models.TtGameUser, error) {
	o := orm.NewOrm()
	aTtGameUser := models.TtGameUser{}

	newPwd := pwdKit.Sha1ToStr(pwd)
	e := o.QueryTable(mconst.TableName_TtGameUser).Filter("UserName", userName).Filter("UserType", mconst.UserType_4).One(&aTtGameUser)
	if e != nil {
		return aTtGameUser, errors.New("用户名不存在。")
	}
	if aTtGameUser.Pwd != newPwd {
		return aTtGameUser, errors.New("密码不正确")
	}

	return aTtGameUser, nil
}

func ModifyPwd(id int, newPwd string) error {
	_, e := GInstance.GetUserRpcClient().GetUser(id)
	if e != nil {
		return errors.New("用户不存在。")
	}
	e = httpGameServer.ChangePaw(id, newPwd)
	if e != nil {
		return errors.New("修改密码失败，原因：远程服务出错")
	}

	updateData := make([]gBox.UpdateDataInfo, 0)
	newPwd1 := pwdKit.Sha1ToStr(newPwd)
	updateData = append(updateData, gBox.UpdateDataInfo{FieldName: "Pwd", Type: 0, Value: newPwd1}, gBox.UpdateDataInfo{FieldName: "Pwd2", Type: 0, Value: newPwd})

	_, e = GInstance.GetUserRpcClient().UpdateUserInfo(id, updateData)
	return e
}
