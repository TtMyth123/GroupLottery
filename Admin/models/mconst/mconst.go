package mconst

const (
	TableName_SysUser        = "sys_user"
	TableName_SysLog         = "sys_log"
	TableName_SysUserMenuRel = "sys_user_menu_rel"
	TableName_SysMenu        = "sys_menu"
)
const (
	SysUserUserType_0_Manage = 0 //管理员
	SysUserUserType_1_Agent  = 1 //代理
)

const (
	SysUserRootId  = 0
	SysUserAdminId = 1
)
const (
	MenuType_Group = 1 //组
	MenuType_Fun   = 2 //功能
)

const (
	Power_N = 0 //"不能看"
	Power_R = 1 //"读"
	Power_W = 2 //"写"
)
