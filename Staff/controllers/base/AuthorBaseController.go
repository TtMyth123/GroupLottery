package base

import (
	"github.com/TtMyth123/Staff/CacheData"
	"github.com/TtMyth123/Staff/GConfig"
	"github.com/TtMyth123/Staff/controllers/base/box"
	"github.com/TtMyth123/Staff/controllers/base/enums"
)

var (
	Secret     = "dong_tech" // 加盐
	ExpireTime = 3600        // token有效期
)

const (
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin    = "请重新登陆"
)

type AuthorBaseController struct {
	ABaseController
	mSysUser box.SysUserEx
}

func (this *AuthorBaseController) Prepare() {
	if GConfig.GetGConfig().IsDev {
		return
	}
	//strToken := this.GetString("token", "")
	strToken := this.Ctx.Input.Header("token")

	aUser := CacheData.GetToken(strToken)
	if aUser == nil {
		this.JsonResult(enums.JRCode401, "授权过期", nil)
	}

	this.mSysUser = *aUser
}

func (this *AuthorBaseController) CurSysUserEx() box.SysUserEx {
	return this.mSysUser
}
