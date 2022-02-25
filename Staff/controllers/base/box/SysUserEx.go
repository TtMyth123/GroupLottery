package box

type SysUserEx struct {
	BaseSysUser
	CurToken string
}

func GetSysUserEx(aTtStaff BaseSysUser, CurToken string) SysUserEx {
	aSysUserEx := SysUserEx{BaseSysUser: aTtStaff, CurToken: CurToken}
	return aSysUserEx
}
