package controllers

import (
	"strings"
	"ttmyth123/GroupLottery/Admin/CacheData"
	"ttmyth123/GroupLottery/Admin/GConfig"
	"ttmyth123/GroupLottery/Admin/controllers/base"
	"ttmyth123/GroupLottery/Admin/controllers/base/enums"
	"ttmyth123/GroupLottery/Admin/models"
	"ttmyth123/kit"
	"ttmyth123/kit/pwdKit"
)

type LoginRegController struct {
	base.ABaseController
}

func (this *LoginRegController) DoLogin() {
	username := this.GetString("username")
	password := this.GetString("password")

	username = strings.ToLower(strings.TrimSpace(username))
	userpwd := strings.ToLower(strings.TrimSpace(password))

	if len(username) == 0 {
		this.JsonResult(enums.JRCodeFailed, "用户名和密码不正确", "")
	}

	newUserpwd := pwdKit.Sha1ToStr(userpwd)
	user, err := models.SysUserOneByUserName(username, newUserpwd)
	if user != nil && err == nil {
		aToken, err := this.GetToken(base.BaseSysUser{UserId: user.Id, Name: user.UserName})
		if err != nil {
			this.JsonResult(enums.JRCodeFailed, err.Error(), "")
		} else {
			aSysUserEx := models.GetSysUserEx(*user, aToken)
			CacheData.ReLoadSysUserInfo(aSysUserEx)
			this.JsonResult(enums.JRCodeSucc, "", aSysUserEx)
		}
	} else {
		this.JsonResult(enums.JRCodeFailed, "用户名或者密码错误", "")
	}
}

func (this *LoginRegController) DoLogout() {
	if GConfig.GetGConfig().IsDev {
		return
	}
	//strToken := this.GetString("token", "")
	//strToken := this.Ctx.Input.Header("token")
	//JWTClaims, err := this.VerifyAction(strToken)
	//if err != nil {
	//	this.JsonResult(enums.JRCode401, err.Error(), nil)
	//}
	//this.mSysUser = cacheData.GetSysUserInfo(JWTClaims.UserID)
	////if this.mSysUser.UserId == 0 || this.mSysUser.CurToken != strToken {
	//if this.mSysUser.UserId == 0 {
	//	this.JsonResult(enums.JRCode401, "授权过期", nil)
	//}
	//this.JWTClaims = JWTClaims
	this.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *LoginRegController) GetGuid() {
	a := kit.GetGuid()
	c.JsonResult(enums.JRCodeSucc, "", a)
}
