package controllers

import (
	"github.com/TtMyth123/Admin/controllers/UserBll"
	"github.com/TtMyth123/Admin/controllers/base/enums"
	"github.com/TtMyth123/UserInfoRpc/GData/gBox"
)

/**
获取用户列表
*/
func (c *ApiController) GetGameUserList() {
	state, _ := c.GetInt("state", 0)
	userName := c.GetString("userName")
	pageIndex, _ := c.GetInt("pageIndex", 0)
	pageSize, _ := c.GetInt("pageSize", 0)
	userType, _ := c.GetInt("userType", 0)
	FirstId, _ := c.GetInt("FirstId", 0)
	userId, _ := c.GetInt("userId", 0)

	curAgentId := c.CurSysUserEx().GameId
	if c.CurSysUserEx().IsSuper {
		curAgentId = 0
	}

	PageTotal, data, groupData := UserBll.GetGameUserList(curAgentId, userId, userType, state, userName, pageIndex, pageSize, FirstId)

	c.JsonPageResultGroup(enums.JRCodeSucc, "", PageTotal, data, groupData)
}

/**
获取用户列表
*/
//func (c *ApiController) GetGameUserListFirstId() {
//	state, _ := c.GetInt("state", 0)
//	userName := c.GetString("userName")
//	pageIndex, _ := c.GetInt("pageIndex", 0)
//	pageSize, _ := c.GetInt("pageSize", 0)
//	userType, _ := c.GetInt("userType", 0)
//	FirstId, _ := c.GetInt("FirstId", 0)
//
//	PageTotal, data, groupData := UserBll.GetGameUserListFirstId(userType, state, userName, pageIndex, pageSize,FirstId)
//
//	c.JsonPageResultGroup(enums.JRCodeSucc, "", PageTotal, data, groupData)
//}
/**
获取下级用户列表
*/
func (c *ApiController) GetJuniorGameUserList() {
	state, _ := c.GetInt("state", 0)
	userName := c.GetString("userName")
	pageIndex, _ := c.GetInt("pageIndex", 0)
	pageSize, _ := c.GetInt("pageSize", 0)
	pid, _ := c.GetInt("pid", 0)
	userType, _ := c.GetInt("userType", 0)
	FirstId, _ := c.GetInt("FirstId", 0)

	PageTotal, data, groupData := UserBll.GetJuniorGameUserList(pid, userType, state, userName, pageIndex, pageSize, FirstId)
	c.JsonPageResultGroup(enums.JRCodeSucc, "", PageTotal, data, groupData)
}

/**
获取下级用户列表
*/
func (c *ApiController) GetJuniorGameUserListFirstId() {
	pid, _ := c.GetInt("pid", 0)
	state, _ := c.GetInt("state", 0)
	userName := c.GetString("userName")
	pageIndex, _ := c.GetInt("pageIndex", 0)
	pageSize, _ := c.GetInt("pageSize", 0)
	userType, _ := c.GetInt("userType", 0)
	FirstId, _ := c.GetInt("FirstId", 0)

	PageTotal, data, groupData := UserBll.GetJuniorGameUserListFirstId(pid, userType, state, userName, pageIndex, pageSize, FirstId)
	c.JsonPageResultGroup(enums.JRCodeSucc, "", PageTotal, data, groupData)
}

/**
修改密码
*/
func (c *ApiController) UpdateUserPwd() {
	UserId, _ := c.GetInt("UserId", 0)
	Pwd := c.GetString("Pwd")
	e := UserBll.UpdateUserPwd(UserId, Pwd)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *ApiController) SetUserState() {
	UserId, _ := c.GetInt("UserId")
	State, _ := c.GetInt("State")

	e := UserBll.SetUserState(UserId, State)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *ApiController) GetUserDetailInfo() {
	UserId, _ := c.GetInt("UserId")

	data, e := UserBll.GetUserDetailInfo(UserId)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", data)
}

func (c *ApiController) SaveUserDetailInfo() {
	UserId, _ := c.GetInt("Id")

	Nickname := c.GetString("Nickname")
	//OnlinePay := c.GetString("OnlinePay")
	//AlipayName := c.GetString("AlipayName")
	//BankCard := c.GetString("BankCard")
	//BankName := c.GetString("BankName")
	//BankUser := c.GetString("BankUser")
	//BankAddr := c.GetString("BankAddr")
	//UserMobile := c.GetString("UserMobile")
	//WXReceiptUrl := c.GetString("WXReceiptUrl")
	IdentityCard := c.GetString("IdentityCard")
	FullName := c.GetString("FullName")
	Tel := c.GetString("Tel")

	WXSKCodeUrl := c.GetString("WXSKCodeUrl")
	YHName := c.GetString("YHName")
	YHUserName := c.GetString("YHUserName")
	YHUserTel := c.GetString("YHUserTel")
	Addr := c.GetString("Addr")
	Cate := c.GetString("Cate")
	Remark := c.GetString("Remark")
	ZFBSKCodeUrl := c.GetString("ZFBSKCodeUrl")
	CardNum := c.GetString("CardNum")
	ZFBSKName := c.GetString("ZFBSKName")
	RealNameState, _ := c.GetInt("RealNameState")
	UserType, _ := c.GetInt("UserType", 0)
	IsReferrer, _ := c.GetInt("IsReferrer", 0)

	infos := make([]gBox.UpdateDataInfo, 0)
	infos = append(infos, gBox.UpdateDataInfo{FieldName: "Nickname", Type: 0, Value: Nickname})
	//infos = append(infos, gBox.UpdateDataInfo{FieldName: "OnlinePay", Type: 0, Value: OnlinePay})
	//infos = append(infos, gBox.UpdateDataInfo{FieldName: "AlipayName", Type: 0, Value: AlipayName})
	//infos = append(infos, gBox.UpdateDataInfo{FieldName: "BankCard", Type: 0, Value: BankCard})
	//infos = append(infos, gBox.UpdateDataInfo{FieldName: "BankName", Type: 0, Value: BankName})
	//infos = append(infos, gBox.UpdateDataInfo{FieldName: "BankUser", Type: 0, Value: BankUser})
	//infos = append(infos, gBox.UpdateDataInfo{FieldName: "BankAddr", Type: 0, Value: BankAddr})
	//infos = append(infos, gBox.UpdateDataInfo{FieldName: "UserMobile", Type: 0, Value: UserMobile})
	//infos = append(infos, gBox.UpdateDataInfo{FieldName: "WXReceiptUrl", Type: 0, Value: WXReceiptUrl})
	infos = append(infos, gBox.UpdateDataInfo{FieldName: "IdentityCard", Type: 0, Value: IdentityCard})
	infos = append(infos, gBox.UpdateDataInfo{FieldName: "Tel", Type: 0, Value: Tel})
	infos = append(infos, gBox.UpdateDataInfo{FieldName: "RealNameState", Type: 0, Value: RealNameState})
	infos = append(infos, gBox.UpdateDataInfo{FieldName: "UserType", Type: 0, Value: UserType})
	infos = append(infos, gBox.UpdateDataInfo{FieldName: "FullName", Type: 0, Value: FullName})
	infos = append(infos, gBox.UpdateDataInfo{FieldName: "IsReferrer", Type: 0, Value: IsReferrer})

	infos = append(infos, gBox.UpdateDataInfo{FieldName: "WXSKCodeUrl", Type: 0, Value: WXSKCodeUrl})
	infos = append(infos, gBox.UpdateDataInfo{FieldName: "YHName", Type: 0, Value: YHName})
	infos = append(infos, gBox.UpdateDataInfo{FieldName: "CardNum", Type: 0, Value: CardNum})
	infos = append(infos, gBox.UpdateDataInfo{FieldName: "YHUserName", Type: 0, Value: YHUserName})
	infos = append(infos, gBox.UpdateDataInfo{FieldName: "YHUserTel", Type: 0, Value: YHUserTel})

	infos = append(infos, gBox.UpdateDataInfo{FieldName: "Addr", Type: 0, Value: Addr})
	infos = append(infos, gBox.UpdateDataInfo{FieldName: "Cate", Type: 0, Value: Cate})
	infos = append(infos, gBox.UpdateDataInfo{FieldName: "Remark", Type: 0, Value: Remark})
	infos = append(infos, gBox.UpdateDataInfo{FieldName: "ZFBSKCodeUrl", Type: 0, Value: ZFBSKCodeUrl})
	infos = append(infos, gBox.UpdateDataInfo{FieldName: "ZFBSKName", Type: 0, Value: ZFBSKName})

	e := UserBll.SaveUserDetailInfo(UserId, infos)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}
