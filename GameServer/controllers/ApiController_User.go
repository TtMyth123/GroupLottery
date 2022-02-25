package controllers

import (
	"ttmyth123/GroupLottery/GameServer/GInstance"
	"ttmyth123/GroupLottery/GameServer/controllers/base"
	"ttmyth123/GroupLottery/GameServer/controllers/base/enums"
	"ttmyth123/GroupLottery/GameServer/controllers/userBll"
	"ttmyth123/GroupLottery/UserInfoRpc/GData/gBox"
	"ttmyth123/kit"
)

/**
getuserinfo 用户基本信息
*/
func (this *ApiController) GetUserInfo() {
	type TmpArgs struct {
		base.ArgsBox
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	data, e := userBll.GetUserInfo(aTmpArgs.UserId)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}

func (this *ApiController) MyReferrerCode() {
	type TmpArgs struct {
		base.ArgsBox
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	data, e := userBll.GetMyReferrerCode(aTmpArgs.UserId)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}

/**
getjunioruserinfo 我的下线
*/
func (this *ApiController) GetJuniorUserInfo() {
	type TmpArgs struct {
		base.ArgsBox
		RootUserId int
		PageIndex  int
		PageSize   int
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	data, e := userBll.GetJuniorUserInfo(aTmpArgs.RootUserId, aTmpArgs.UserId, aTmpArgs.PageIndex, aTmpArgs.PageSize, aTmpArgs.LastId)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	this.JsonResult(enums.JRCodeSucc, "", data, nil)

}

/**
getjunioruserinfo 我的下线
*/
func (this *ApiController) GetJuniorUserInfoWsx() {
	type TmpArgs struct {
		base.ArgsBox
		RootUserId int
		PageIndex  int
		PageSize   int
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	data, e := userBll.GetJuniorUserInfoWsx(aTmpArgs.RootUserId, aTmpArgs.UserId, aTmpArgs.PageIndex, aTmpArgs.PageSize, aTmpArgs.LastId)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	this.JsonResult(enums.JRCodeSucc, "", data, nil)
}

func (this *ApiController) UpdateUserInfo() {
	type TmpArgs struct {
		base.ArgsBox
		ChangeInfo map[string]interface{}
	}
	aTmpArgs := TmpArgs{}
	e := this.GetJsonData(&aTmpArgs)
	if e != nil {
		this.ErrJsonResultEx(e)
	}
	updateData := make([]gBox.UpdateDataInfo, 0)
	for FieldName, v := range aTmpArgs.ChangeInfo {
		switch FieldName {
		case "VoucherFile", "ZFBSKCodeUrl", "ZFBSKName", "YHName", "CardNum", "YHUserName", "YHUserTel", "Addr", "Cate", "Remark":
			updateData = append(updateData, gBox.UpdateDataInfo{FieldName: FieldName, Type: 0, Value: kit.GetInterface2Str(v, "")})
		case "Nickname":
			updateData = append(updateData, gBox.UpdateDataInfo{FieldName: FieldName, Type: 0, Value: kit.GetInterface2Str(v, "")})
		case "IdentityCard", "FullName", "Tel":
			updateData = append(updateData, gBox.UpdateDataInfo{FieldName: FieldName, Type: 0, Value: kit.GetInterface2Str(v, "")})
		}
	}
	_, e = GInstance.GetUserRpcClient().UpdateUserInfo(aTmpArgs.UserId, updateData)
	if e != nil {
		this.ErrJsonResultEx(e)
	}

	this.JsonResult(enums.JRCodeSucc, "", "", nil)
}
