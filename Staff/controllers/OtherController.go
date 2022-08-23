package controllers

import (
	"github.com/TtMyth123/Staff/controllers/base"
	"github.com/TtMyth123/Staff/controllers/base/enums"
	"github.com/TtMyth123/Staff/controllers/bll"
)

type OtherController struct {
	base.AuthorBaseController
}

func (this *OtherController) ModifyPwd() {
	NewPwd := this.GetString("NewPwd")
	curUser := this.CurSysUserEx()
	e := bll.ModifyPwd(int(curUser.Id), NewPwd)
	if e != nil {
		this.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}

	this.JsonResult(enums.JRCodeSucc, "", "")
}
