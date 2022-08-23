package SysBll

import (
	"errors"
	"fmt"
	"github.com/TtMyth123/Admin/GInstance"
	"github.com/TtMyth123/UserInfoRpc/GData/gBox"
	"github.com/astaxie/beego/orm"
	"time"

	"github.com/TtMyth123/kit/pwdKit"
	"github.com/TtMyth123/kit/ttLog"
	//"ttmyth123/AdminLite/models/mConst"
	"github.com/TtMyth123/Admin/CacheData"
	"github.com/TtMyth123/Admin/models"
	"github.com/TtMyth123/Admin/models/mconst"

	u_models "github.com/TtMyth123/UserInfoRpc/models"
)

func GetSysMenuList(userId int, IsSuper bool) []*models.GroupMenu {
	data := make([]*models.GroupMenu, 0)
	key := fmt.Sprintf("1GetMenuList:%d", userId)
	e := CacheData.GetBeegoCache().GetCache(key, &data)
	if e != nil {
		data = models.GetSysMenuList(userId, IsSuper)
		e := CacheData.GetBeegoCache().SetCache(key, data, 10)

		if e != nil {
			ttLog.LogDebug(e)
		}
	}

	return data
}

func GetSysUserList(sysUserId int, userName string, userId, pageIndex, pageSize int) (int, []models.SysUser) {
	if pageIndex <= 0 {
		pageIndex = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return models.GetSysUserListByPage(sysUserId, userName, userId, pageIndex, pageSize)
}

func UpdateSysUser(aSysUser models.SysUser, cols ...string) error {
	if aSysUser.Id == 1 && aSysUser.UserName != "admin" {
		return errors.New(fmt.Sprintf("不允许修改管理员用户名"))
	}

	e := aSysUser.Update(cols...)
	return e
}
func UpdateAgentSysUser(aSysUser models.SysUser, cols ...string) error {
	if aSysUser.Id == 1 && aSysUser.UserName != "admin" {
		return errors.New(fmt.Sprintf("不允许修改管理员用户名"))
	}
	o := orm.NewOrm()
	aSysUser.UserType = mconst.SysUserUserType_1_Agent
	aTtGameUser := u_models.TtGameUser{Id: aSysUser.GameId}
	e := o.Read(&aTtGameUser)
	if e != nil {
		return fmt.Errorf("[%d]游戏用户不存在。", aSysUser.GameId)
	}
	e = aSysUser.Update(cols...)
	return e
}

func AddAgentSysUser(curAgentId int, aSysUser models.SysUser) error {
	o := orm.NewOrm()
	aSysUser.UserType = mconst.SysUserUserType_1_Agent
	aTtGameUser := u_models.TtGameUser{Id: aSysUser.GameId}
	e := o.Read(&aTtGameUser)
	if e != nil {
		return fmt.Errorf("[%d]游戏用户不存在。", aSysUser.GameId)
	}

	aTtGameUserAgent := u_models.TtGameUser{Id: curAgentId}
	e = o.Read(&aTtGameUserAgent)
	if e != nil {
		return fmt.Errorf("你不是代理商。")
	}
	if aTtGameUser.Pid != curAgentId {
		return fmt.Errorf("[%d]游戏用户【%s】下级。", aSysUser.GameId, aTtGameUserAgent.UserName)
	}
	if e == nil {
		infos := make([]gBox.UpdateDataInfo, 0)
		infos = append(infos, gBox.UpdateDataInfo{FieldName: "IsAgent", Type: 0, Value: 1})
		_, e = GInstance.GetUserRpcClient().UpdateUserInfo(aSysUser.GameId, infos)
		if e != nil {
			return e
		}

		GInstance.GetUserRpcClient().UpdateUserAgentId(aSysUser.GameId)
	}

	e = aSysUser.AddSysUser()

	return e
}

func AddManageSysUser(aSysUser models.SysUser) error {
	aSysUser.GameId = 0
	aSysUser.UserType = mconst.SysUserUserType_0_Manage
	e := aSysUser.AddSysUser()
	return e
}
func AddSysUser(aSysUser models.SysUser) error {
	e := aSysUser.AddSysUser()
	return e
}

func DelSysUser(id int) error {
	o := orm.NewOrm()
	if id == mconst.SysUserAdminId {
		return errors.New("不能删除管理员用户")
	}
	aSysUser := models.SysUser{Id: id}
	_, e := o.Delete(&aSysUser)
	return e
}

type FunMenuPermission struct {
	Id       int
	TypeName string
	Title    string
	Power    int
}

func GetFunMenuPermission(userId int) []FunMenuPermission {
	arr := make([]FunMenuPermission, 0)
	o := orm.NewOrm()
	sql := `select a.id, b.title as type_name, a.title, IFNULL(e.power,0) as power from sys_menu a
LEFT JOIN sys_menu b on (a.pid = b.id)
LEFT JOIN 
(select c.sys_menu_id, c.power from sys_user_menu_rel c, sys_user d where d.id = c.sys_user_id and d.id=?) e on (e.sys_menu_id=a.id)
where a.menu_type = ? order by b.seq
`
	_, e := o.Raw(sql, userId, mconst.MenuType_Fun).QueryRows(&arr)
	if e != nil {
		ttLog.LogError(e)
	}

	return arr
}

func SaveFunMenuPermission(PUserId, userId int, arr []FunMenuPermission) error {
	o := orm.NewOrm()
	if len(arr) == 0 {
		return errors.New("数据为空无法保存。")
	}
	sql := fmt.Sprintf(`delete from %s where sys_user_id=?`, mconst.TableName_SysUserMenuRel)
	_, e := o.Raw(sql, userId).Exec()
	if e != nil {
		return e
	}

	sqlP := fmt.Sprintf(`select * from %s where sys_user_id=? `, mconst.TableName_SysUserMenuRel)
	arrPData := make([]models.SysUserMenuRel, 0)
	o.Raw(sqlP, PUserId).QueryRows(&arrPData)
	mpMenu := make(map[int]models.SysUserMenuRel)
	for _, r := range arrPData {
		mpMenu[r.SysMenuId] = r
	}

	curTime := time.Now()

	arrData := make([]models.SysUserMenuRel, 0)
	for _, menu := range arr {
		Power := menu.Power
		if mpMenu[menu.Id].Power < menu.Power {
			Power = mpMenu[menu.Id].Power
		}
		arrData = append(arrData, models.SysUserMenuRel{SysUserId: userId, SysMenuId: menu.Id, Power: Power, CreatedAt: curTime})
	}
	_, e = o.InsertMulti(len(arrData), &arrData)

	return e
}
func ModifyPwd(userId int, OldPwd, NewPwd string) error {
	sysU, e := models.GetSysUser(userId)
	if e != nil {
		return e
	}
	aNewPwd := pwdKit.Sha1ToStr(OldPwd)
	if sysU.Password != aNewPwd {
		return errors.New("原密码不正确.")
	}

	sysU.Password = pwdKit.Sha1ToStr(NewPwd)
	e = sysU.Update("Password")
	return e
}

func ModifySysUserInfo(userId int, FieldName, FieldValue string) error {
	sysU, e := models.GetSysUser(userId)
	if e != nil {
		return e
	}

	switch FieldName {
	case "NickName":
		sysU.NickName = FieldValue
		e = sysU.Update("NickName")
		return e
	case "Mobile":
		sysU.Mobile = FieldValue
		e = sysU.Update("Mobile")
		return e
	case "Email":
		sysU.Email = FieldValue
		e = sysU.Update("Email")
		return e
	}

	return errors.New("有问题的参数")
}
