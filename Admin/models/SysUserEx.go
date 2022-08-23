package models

import (
	"errors"
	"github.com/TtMyth123/Admin/models/mconst"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
)

type SysUserEx struct {
	SysUser
	MenuUrlFor map[int]int
	CurToken   string
}

func GetSysUserEx(aSysUser SysUser, CurToken string) SysUserEx {
	aSysUserEx := SysUserEx{SysUser: aSysUser}
	aSysUserEx.MenuUrlFor = GetSysUserMenuRel(aSysUser.Id)
	aSysUserEx.CurToken = CurToken
	return aSysUserEx
}

type SysMenuR struct {
	Id     int
	Title  string
	Field  string
	Modify int
}

type SysUserR struct {
	SysUser
	MenuList []SysMenuR
}

func GetSysUserR(userId int) SysUserR {
	aSysUserR := SysUserR{}

	aSysUser := SysUser{Id: userId}
	o := orm.NewOrm()
	e := o.Read(&aSysUser)
	if e != nil {
		aSysUser.Id = 0
	}
	aSysUserR.SysUser = aSysUser
	sql := `select a.title, a.id,  IFNULL(b.allow_modify, -1)+1 as modify from sys_menu a
LEFT JOIN sys_user_menu_rel b on (a.id = b.sys_menu_id and b.sys_user_id =?)
where menu_type = ?
`
	arr := make([]SysMenuR, 0)
	_, e = o.Raw(sql, userId, 2).QueryRows(&arr)
	if e != nil {
		ttLog.LogError(e)
	}

	aSysUserR.MenuList = arr
	return aSysUserR
}

func DelSysUser(sysUserId int) error {
	o := orm.NewOrm()
	if sysUserId == mconst.SysUserAdminId {
		return errors.New("不能删除管理员")
	}

	sql := `delete from sys_user where id = ?`
	_, e := o.Raw(sql, sysUserId).Exec()

	if e != nil {
		ttLog.LogError(e)
		return e
	}
	sql = `delete from sys_user_menu_rel where sys_user_id = ?`
	_, e = o.Raw(sql, sysUserId).Exec()
	if e != nil {
		ttLog.LogError(e)
	}
	return e
}

type SimpleSysUser struct {
	Id       int
	UserName string
	Mobile   string `orm:"size(16)"`
}

func GetSysUserList() []SimpleSysUser {
	arr := make([]SimpleSysUser, 0)
	o := orm.NewOrm()
	sql := `select a.id, a.user_name, a.mobile from sys_user a`

	_, e := o.Raw(sql).QueryRows(&arr)
	if e != nil {
		ttLog.LogError(e)
	}
	return arr
}

func SaveSysUser(aSysUserR SysUserR) error {
	if aSysUserR.Id == 0 {
		aSysUser := aSysUserR.SysUser
		e := aSysUser.AddSysUser()
		if e != nil {
			return e
		}

		aSysUser.SetMenu(aSysUserR.MenuList)
	} else {
		aSysUser := aSysUserR.SysUser
		e := aSysUser.Update("NickName", "UserName", "Mobile")
		if e != nil {
			return e
		}
		aSysUser.SetMenu(aSysUserR.MenuList)
	}

	return nil
}
