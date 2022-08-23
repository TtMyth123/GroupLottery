package controllers

import (
	"fmt"
	"github.com/TtMyth123/Staff/CacheData"
	"github.com/TtMyth123/Staff/controllers/base"
	"github.com/TtMyth123/Staff/controllers/base/box"
	"github.com/TtMyth123/Staff/controllers/base/enums"
	"github.com/TtMyth123/Staff/controllers/bll"
	"strings"
	"time"
)

type LoginRegController struct {
	base.ABaseController
}

func (this *LoginRegController) DoLogin() {
	username := this.GetString("username")
	password := this.GetString("pwd")

	username = strings.TrimSpace(username)
	//userpwd := strings.ToLower(strings.TrimSpace(password))

	if len(username) == 0 {
		this.JsonResult(enums.JRCodeFailed, "用户名和密码不正确", "")
	}
	aTtStaff, e := bll.Login(username, password)
	if e != nil {
		this.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}

	aToken, err := this.GetToken(box.BaseSysUser{Id: int64(aTtStaff.Id), Name: aTtStaff.UserName})
	if err != nil {
		this.JsonResult(enums.JRCodeFailed, err.Error(), "")
	}
	aSysUserEx := box.GetSysUserEx(box.BaseSysUser{Id: int64(aTtStaff.Id), Name: aTtStaff.UserName}, aToken)

	CacheData.ReLoadSysUserInfo(aSysUserEx)
	this.JsonResult(enums.JRCodeSucc, "", aSysUserEx)
}

func (this *LoginRegController) DoLogout() {
	fmt.Println("DoLogout T2:", time.Now())
	this.JsonResult(enums.JRCodeSucc, "", "")
}
