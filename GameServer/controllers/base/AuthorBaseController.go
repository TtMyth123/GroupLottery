package base

import "ttmyth123/GroupLottery/GameServer/controllers/base/enums"

var (
	Secret     = "dong_tech" // 加盐
	ExpireTime = 3600        // token有效期
)

const (
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin    = "请重新登陆"
)

type AuthorBaseController struct {
	BaseController
	JWTClaims *JWTClaims
}

func (this *AuthorBaseController) Prepare() {
	return
	strToken := this.Ctx.Input.Header("token")
	JWTClaims, err := this.verifyAction(strToken)
	if err != nil {
		this.JsonResult(enums.JRCode401, err.Error(), nil, nil)
	}
	//this.mSysUser = cacheData.GetSysUserInfo(JWTClaims.UserID)
	////if this.mSysUser.Id == 0 || this.mSysUser.CurToken != strToken {
	//if this.mSysUser.Id == 0 {
	//	this.JsonResult(enums.JRCode401, "授权过期", nil)
	//}
	this.JWTClaims = JWTClaims
}
