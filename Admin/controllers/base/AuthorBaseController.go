package base

import (
	"fmt"
	"github.com/TtMyth123/Admin/CacheData"
	"github.com/TtMyth123/Admin/GConfig"
	"github.com/TtMyth123/Admin/controllers/base/enums"
	"github.com/TtMyth123/Admin/models"
	"time"
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
	JWTClaims *JWTClaims
	mSysUser  models.SysUserEx
}

func (this *AuthorBaseController) Prepare() {
	if GConfig.GetGConfig().IsDev {
		return
	}
	//strToken := this.GetString("token", "")
	strToken := this.Ctx.Input.Header("token")
	JWTClaims, err := this.VerifyAction(strToken)
	if err != nil {
		this.JsonResult(enums.JRCode401, err.Error(), nil)
	}
	this.mSysUser = CacheData.GetSysUserInfo(JWTClaims.UserID)
	//if this.mSysUser.UserId == 0 || this.mSysUser.CurToken != strToken {
	if this.mSysUser.Id == 0 {
		this.JsonResult(enums.JRCode401, "授权过期", nil)
	}
	this.JWTClaims = JWTClaims
}

func (this *AuthorBaseController) CheckP(funId int, allowModify bool) {
	if v, ok := this.mSysUser.MenuUrlFor[funId]; ok {
		if allowModify {
			if v != 1 {
				this.JsonResult(enums.JRCode402, "权限不足", nil)
			}
		}
	} else {
		this.JsonResult(enums.JRCode402, "权限不足", nil)
	}
}

func (this *AuthorBaseController) CurSysUserEx() models.SysUserEx {
	return this.mSysUser
}

func (this *AuthorBaseController) Refresh() {
	strToken := this.GetString("token", "")
	claims, err := this.VerifyAction(strToken)
	if err != nil {
		this.JsonResult(enums.JRCode401, err.Error(), nil)
	}

	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	signedToken, err := this.GetToken(BaseSysUser{UserId: claims.UserID, Name: claims.UserName})
	if err != nil {
		this.JsonResult(enums.JRCode401, err.Error(), nil)
	}
	fmt.Println(signedToken)
}
