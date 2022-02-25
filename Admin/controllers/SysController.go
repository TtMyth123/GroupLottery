package controllers

import (
	"ttmyth123/GroupLottery/Admin/controllers/SysBll"
	"ttmyth123/GroupLottery/Admin/controllers/base"
	"ttmyth123/GroupLottery/Admin/controllers/base/enums"
	"ttmyth123/GroupLottery/Admin/models"
	"ttmyth123/GroupLottery/Admin/models/mconst"
)

type SysController struct {
	base.AuthorBaseController
}

func (c *SysController) GetMenuList() {

	sysUser := c.CurSysUserEx()
	data := SysBll.GetSysMenuList(sysUser.Id, sysUser.IsSuper)
	c.JsonResult(enums.JRCodeSucc, "", data)
}

func (c *SysController) GetSysUserList() {
	userId, _ := c.GetInt("userId", 0)
	userName := c.GetString("userName")
	pageIndex, _ := c.GetInt("pageIndex", 0)
	pageSize, _ := c.GetInt("pageSize", 0)
	curSysUserId := c.CurSysUserEx().Id

	PageTotal, data := SysBll.GetSysUserList(curSysUserId, userName, userId, pageIndex, pageSize)
	c.JsonPageResult(enums.JRCodeSucc, "", PageTotal, data)
}

func (c *SysController) UpdateSysUser() {
	Id, _ := c.GetInt("Id")
	UserName := c.GetString("UserName")

	aSysUser, e := models.GetSysUser(Id)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	aSysUser.UserName = UserName
	e = SysBll.UpdateSysUser(aSysUser, "UserName")
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *SysController) AddSysUser() {
	UserName := c.GetString("UserName")
	Pwd := c.GetString("Pwd")
	if UserName == "" {
		c.JsonResult(enums.JRCodeFailed, "用户名不能为空", "")
	}
	if Pwd == "" {
		c.JsonResult(enums.JRCodeFailed, "密码不能为空", "")
	}

	aSysUser := models.SysUser{UserName: UserName, Password: Pwd, Pid: c.CurSysUserEx().Id}
	e := SysBll.AddSysUser(aSysUser)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *SysController) AddAgentSysUser() {
	UserName := c.GetString("UserName")
	Pwd := c.GetString("Pwd")
	GameId, _ := c.GetInt("GameId", 0)

	if UserName == "" {
		c.JsonResult(enums.JRCodeFailed, "用户名不能为空", "")
	}
	if Pwd == "" {
		c.JsonResult(enums.JRCodeFailed, "密码不能为空", "")
	}

	curAgentId := c.CurSysUserEx().GameId

	aSysUser := models.SysUser{UserName: UserName, Password: Pwd, GameId: GameId, UserType: mconst.SysUserUserType_1_Agent, Pid: c.CurSysUserEx().Id}
	e := SysBll.AddAgentSysUser(curAgentId, aSysUser)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}
func (c *SysController) UpdateAgentSysUser() {
	Id, _ := c.GetInt("Id")
	UserName := c.GetString("UserName")
	//GameId,_ := c.GetInt("GameId",0)
	aSysUser, e := models.GetSysUser(Id)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	aSysUser.UserName = UserName
	//aSysUser.GameId = GameId
	e = SysBll.UpdateAgentSysUser(aSysUser, "UserName")
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *SysController) AddManageSysUser() {
	UserName := c.GetString("UserName")
	Pwd := c.GetString("Pwd")
	if UserName == "" {
		c.JsonResult(enums.JRCodeFailed, "用户名不能为空", "")
	}
	if Pwd == "" {
		c.JsonResult(enums.JRCodeFailed, "密码不能为空", "")
	}

	aSysUser := models.SysUser{UserName: UserName, Password: Pwd, UserType: mconst.SysUserUserType_0_Manage, Pid: c.CurSysUserEx().Id}
	e := SysBll.AddManageSysUser(aSysUser)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *SysController) DelSysUser() {
	id, _ := c.GetInt("id", 0)
	e := SysBll.DelSysUser(id)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *SysController) GetPermissionTree() {
	UserId, _ := c.GetInt("UserId")
	permissionTree := SysBll.GetFunMenuPermission(UserId)
	c.JsonResult(enums.JRCodeSucc, "", permissionTree)
}

func (c *SysController) SavePermissionTree() {
	type TmpA struct {
		UserId int
		Data   []SysBll.FunMenuPermission
	}
	aTmp := TmpA{}
	e := c.GetJsonData(&aTmp)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}

	permissionTree := SysBll.SaveFunMenuPermission(c.CurSysUserEx().Id, aTmp.UserId, aTmp.Data)
	c.JsonResult(enums.JRCodeSucc, "", permissionTree)
}

func (c *SysController) ModifyPwd() {
	OldPwd := c.GetString("OldPwd")
	NewPwd := c.GetString("NewPwd")
	sysUser := c.CurSysUserEx()
	e := SysBll.ModifyPwd(sysUser.Id, OldPwd, NewPwd)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}
func (c *SysController) ModifySysUserInfo() {
	FieldName := c.GetString("FieldName")
	FieldValue := c.GetString("FieldValue")
	sysUser := c.CurSysUserEx()
	e := SysBll.ModifySysUserInfo(sysUser.Id, FieldName, FieldValue)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}
