package controllers

import (
	"ttmyth123/GroupLottery/Staff/controllers/base"
	"ttmyth123/GroupLottery/Staff/controllers/base/enums"
	"ttmyth123/GroupLottery/Staff/controllers/bll"
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
