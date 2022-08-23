package models

//

import (
	"fmt"
	"github.com/TtMyth123/Admin/Langs"
	"github.com/TtMyth123/Admin/models/mconst"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"time"
)

type SysMenu struct {
	Id       int
	Pid      int
	Icon     string `orm:"size(256)"`
	Title    string `orm:"size(256)"`
	UrlFor   string `orm:"size(256)";json:"index"`
	MenuType int

	IconM   string `orm:"size(256)"`
	TitleM  string `orm:"size(256)"`
	UrlForM string `orm:"size(256)"`

	Seq       int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
	Level     int       `orm:"-"` //第几级，从0开始
}

func (a *SysMenu) TableName() string {
	return mconst.TableName_SysMenu
}

func iniSysMenu(o orm.Ormer) error {
	c, _ := o.QueryTable(mconst.TableName_SysMenu).Count()
	if c == 0 {
		arrSysMenu := make([]SysMenu, 0)
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 1000, Id: 1000, Pid: 0, Icon: "el-icon-lx-home", IconM: "", Title: "权限管理", UrlForM: "", UrlFor: "", MenuType: mconst.MenuType_Group})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 1001, Id: 1001, Pid: 1000, Icon: "el-icon-lx-file", IconM: "home", Title: "系统用户列表", UrlForM: "SysUserList", UrlFor: "SysUserList", MenuType: mconst.MenuType_Fun})

		//--------------------------------------------------------------------------------------------------------------
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 2000, Id: 2000, Pid: 0, Icon: "el-icon-lx-file", IconM: "", Title: "财务管理", UrlFor: "", MenuType: mconst.MenuType_Group})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 2001, Id: 2001, Pid: 2000, Icon: "el-icon-lx-file", IconM: "home", Title: "充值申请管理", UrlForM: "SaveMoneyApplyList", UrlFor: "SaveMoneyApplyList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 2002, Id: 2002, Pid: 2000, Icon: "el-icon-lx-file", IconM: "home", Title: "提现申请管理", UrlForM: "DrawMoneyApplyList", UrlFor: "DrawMoneyApplyList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 2003, Id: 2003, Pid: 2000, Icon: "el-icon-lx-file", IconM: "home", Title: "充值列表", UrlForM: "SaveMoneyList", UrlFor: "SaveMoneyList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 2004, Id: 2004, Pid: 2000, Icon: "el-icon-lx-file", IconM: "home", Title: "提现列表", UrlForM: "DrawMoneyList", UrlFor: "DrawMoneyList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 2005, Id: 2005, Pid: 2000, Icon: "el-icon-lx-file", IconM: "home", Title: "流水查询", UrlForM: "AccountList", UrlFor: "AccountList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 2006, Id: 2006, Pid: 2000, Icon: "el-icon-lx-file", IconM: "home", Title: "财务收款信息", UrlFor: "FinanceAccount", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 2007, Id: 2007, Pid: 2000, Icon: "el-icon-lx-file", IconM: "home", Title: "返佣设置", UrlForM: "RebateSet", UrlFor: "RebateSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 2008, Id: 2008, Pid: 2000, Icon: "el-icon-lx-file", IconM: "home", Title: "返利流水", UrlForM: "RebateList", UrlFor: "RebateList", MenuType: mconst.MenuType_Fun})
		//--------------------------------------------------------------------------------------------------------
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 3000, Id: 3000, Pid: 0, Icon: "el-icon-lx-file", IconM: "", Title: "用户管理", UrlFor: "", MenuType: mconst.MenuType_Group})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 3001, Id: 3001, Pid: 3000, Icon: "el-icon-lx-file", IconM: "home", Title: "用户列表", UrlForM: "GameUserList", UrlFor: "UserList", MenuType: mconst.MenuType_Fun})
		//arrSysMenu = append(arrSysMenu, SysMenu{Seq:3002, Id: 3002, Pid: 3000, Icon: "el-icon-lx-file",IconM: "home", Title: "用户报表", UrlFor: "UserReportList", MenuType: mconst.MenuType_Fun})
		//--------------------------------------------------------------------------------------------------------

		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11000, Id: 11000, Pid: 0, Icon: "el-icon-lx-file", IconM: "", Title: "赔率设置", UrlFor: "", MenuType: mconst.MenuType_Group})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11001, Id: 11001, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "南部彩赔率", UrlForM: "Wsx1NbcOddsInfoSet", UrlFor: "Wsx1NbcOddsInfoSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11002, Id: 11002, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "北部彩赔率", UrlForM: "Wsx2BbcOddsInfoSet", UrlFor: "Wsx2BbcOddsInfoSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11003, Id: 11003, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "中部彩赔率", UrlForM: "Wsx3ZbcOddsInfoSet", UrlFor: "Wsx3ZbcOddsInfoSet", MenuType: mconst.MenuType_Fun})

		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11003, Id: 11004, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "北京28赔率", UrlForM: "Zg28BjOddsInfoSet", UrlFor: "Zg28BjSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11003, Id: 11005, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "加拿大28赔率", UrlForM: "Zg28JndOddsInfoSet", UrlFor: "Zg28JndSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11003, Id: 11006, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "新加坡28赔率", UrlForM: "Zg28XjpOddsInfoSet", UrlFor: "Zg28XjpSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11003, Id: 11022, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "香港28赔率", UrlForM: "Zg28XgOddsInfoSet", UrlFor: "Zg28XgSet", MenuType: mconst.MenuType_Fun})

		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11007, Id: 11007, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "重庆时时彩赔率", UrlForM: "Usc5Cqssc103OdssInfoSet", UrlFor: "Usc5Cqssc103OdssInfoSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11008, Id: 11008, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "极速时时彩赔率", UrlForM: "Usc5Jsssc111OdssInfoSet", UrlFor: "Usc5Jsssc111OdssInfoSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11009, Id: 11009, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "英国幸运彩赔率", UrlForM: "Usc5Ygcyc114OdssInfoSet", UrlFor: "Usc5Ygcyc114OdssInfoSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11010, Id: 11010, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "英国时时彩赔率", UrlForM: "Usc5Ygssc120OdssInfoSet", UrlFor: "Usc5Ygssc120OdssInfoSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11011, Id: 11011, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "澳洲幸运5赔率", UrlForM: "Usc5Gzxy5116OdssInfoSet", UrlFor: "Usc5Gzxy5116OdssInfoSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11012, Id: 11012, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "腾讯分分彩赔率", UrlForM: "Usc5Yxssc118OdssInfoSet", UrlFor: "Usc5Yxssc118OdssInfoSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11013, Id: 11013, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "重庆幸运农场赔率", UrlForM: "Usc8Cqxync107OdssInfoSet", UrlFor: "Usc8Cqxync107OdssInfoSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11014, Id: 11014, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "广东快乐十分赔率", UrlForM: "Usc8Gdkl10f102OdssInfoSet", UrlFor: "Usc8Gdkl10f102OdssInfoSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11015, Id: 11015, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "北京赛车赔率", UrlForM: "Usc10Bjsc104OdssInfoSet", UrlFor: "Usc10Bjsc104OdssInfoSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11016, Id: 11016, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "幸运飞艇赔率", UrlForM: "Usc10Xyft108OdssInfoSet", UrlFor: "Usc10Xyft108OdssInfoSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11017, Id: 11017, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "极速快车赔率", UrlForM: "Usc10Jskc109OdssInfoSet", UrlFor: "Usc10Jskc109OdssInfoSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11018, Id: 11018, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "极速赛车赔率", UrlForM: "Usc10Jssc112OdssInfoSet", UrlFor: "Usc10Jssc112OdssInfoSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11019, Id: 11019, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "ESP赛马赔率", UrlForM: "Usc10Espsm113OdssInfoSet", UrlFor: "Usc10Espsm113OdssInfoSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11020, Id: 11020, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "英国赛车赔率", UrlForM: "Usc10Ygxyft115OdssInfoSet", UrlFor: "Usc10Ygxyft115OdssInfoSet", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 11021, Id: 11021, Pid: 11000, Icon: "el-icon-lx-file", IconM: "home", Title: "澳洲幸运10赔率", UrlForM: "Usc10Gzxy10117OdssInfoSet", UrlFor: "Usc10Gzxy10117OdssInfoSet", MenuType: mconst.MenuType_Fun})

		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 12000, Id: 12000, Pid: 0, Icon: "el-icon-lx-file", IconM: "", Title: "报表", UrlFor: "", MenuType: mconst.MenuType_Group})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 12001, Id: 12001, Pid: 12000, Icon: "el-icon-lx-file", IconM: "home", Title: "主投注记录", UrlForM: "GroupBetList", UrlFor: "GroupBetList", MenuType: mconst.MenuType_Fun})

		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13000, Id: 13000, Pid: 0, Icon: "el-icon-lx-file", IconM: "", Title: "开奖结果", UrlFor: "", MenuType: mconst.MenuType_Group})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13001, Id: 13001, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "南部彩开奖结果", UrlForM: "Wsx1AwardList", UrlFor: "Wsx1AwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13002, Id: 13002, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "北部彩开奖结果", UrlForM: "Wsx2AwardList", UrlFor: "Wsx2AwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13003, Id: 13003, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "中部彩开奖结果", UrlForM: "Wsx3AwardList", UrlFor: "Wsx3AwardList", MenuType: mconst.MenuType_Fun})

		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13003, Id: 13004, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "加拿大28开奖结果", UrlForM: "Zg28JndAwardList", UrlFor: "Zg28JndAwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13004, Id: 13005, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "北京28开奖结果", UrlForM: "Zg28BjAwardList", UrlFor: "Zg28BjAwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13005, Id: 13006, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "新加坡28开奖结果", UrlForM: "Zg28XjpAwardList", UrlFor: "Zg28XjpAwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13005, Id: 13023, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "香港28开奖结果", UrlForM: "Zg28XgAwardList", UrlFor: "Zg28XgAwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13005, Id: 13024, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "设置香港28开奖结果", UrlForM: "SetZg28XgAwardList", UrlFor: "SetZg28XgAwardList", MenuType: mconst.MenuType_Fun})

		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13007, Id: 13007, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "重庆时时彩开奖结果", UrlForM: "UscCqssc103AwardList", UrlFor: "UscCqssc103AwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13008, Id: 13008, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "极速时时彩开奖结果", UrlForM: "UscJsssc111AwardList", UrlFor: "UscJsssc111AwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13009, Id: 13009, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "英国幸运彩开奖结果", UrlForM: "UscYgcyc114AwardList", UrlFor: "UscYgcyc114AwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13010, Id: 13010, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "英国时时彩开奖结果", UrlForM: "UscYgssc120AwardList", UrlFor: "UscYgssc120AwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13011, Id: 13011, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "澳洲幸运5开奖结果", UrlForM: "UscGzxy5116AwardList", UrlFor: "UscGzxy5116AwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13012, Id: 13012, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "腾讯分分彩开奖结果", UrlForM: "UscYxssc118AwardList", UrlFor: "UscYxssc118AwardList", MenuType: mconst.MenuType_Fun})

		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13013, Id: 13013, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "北京赛车开奖结果", UrlForM: "UscBjsc104AwardList", UrlFor: "UscBjsc104AwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13014, Id: 13014, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "幸运飞艇开奖结果", UrlForM: "UscXyft108AwardList", UrlFor: "UscXyft108AwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13015, Id: 13015, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "极速快车开奖结果", UrlForM: "UscJskc109AwardList", UrlFor: "UscJskc109AwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13016, Id: 13016, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "极速赛车开奖结果", UrlForM: "UscJssc112AwardList", UrlFor: "UscJssc112AwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13017, Id: 13017, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "ESP赛马开奖结果", UrlForM: "UscEspsm113AwardList", UrlFor: "UscEspsm113AwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13018, Id: 13018, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "英国幸运飞艇开奖结果", UrlForM: "UscYgxyft115AwardList", UrlFor: "UscYgxyft115AwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13019, Id: 13019, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "英国赛车开奖结果", UrlForM: "UscYgsc119AwardList", UrlFor: "UscYgsc119AwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13020, Id: 13020, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "澳洲幸运10开奖结果", UrlForM: "UscCqssc103AwardList", UrlFor: "UscCqssc103AwardList", MenuType: mconst.MenuType_Fun})

		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13021, Id: 13021, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "重庆幸运农场开奖结果", UrlForM: "UscCqxync107AwardList", UrlFor: "UscCqxync107AwardList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 13022, Id: 13022, Pid: 13000, Icon: "el-icon-lx-file", IconM: "home", Title: "广东快乐十分开奖结果", UrlForM: "UscGdkl10f102AwardList", UrlFor: "UscGdkl10f102AwardList", MenuType: mconst.MenuType_Fun})

		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 14000, Id: 14000, Pid: 0, Icon: "el-icon-lx-file", IconM: "", Title: "文章公告", UrlFor: "", MenuType: mconst.MenuType_Group})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 14001, Id: 14001, Pid: 14000, Icon: "el-icon-lx-file", IconM: "home", Title: "公告", UrlForM: "NoticeList", UrlFor: "NoticeList", MenuType: mconst.MenuType_Fun})
		arrSysMenu = append(arrSysMenu, SysMenu{Seq: 14002, Id: 14002, Pid: 14000, Icon: "el-icon-lx-file", IconM: "home", Title: "幻灯片", UrlForM: "SlideshowList", UrlFor: "SlideshowList", MenuType: mconst.MenuType_Fun})

		_, e := o.InsertMulti(len(arrSysMenu), &arrSysMenu)
		ttLog.LogDebug("添加菜单 完成 e:", e)

		iniSysUserMenuRel(o)

		return e
	}
	return nil
}

type GroupMenu struct {
	Id          int
	Title       string `json:"title"`
	Icon        string `json:"icon"`
	Seq         int
	UrlFor      string `json:"path";orm:"size(256)"`
	IconM       string
	UrlForM     string
	SubMenuList []*GroupMenu `json:"children"`
	MenuType    int
}

func GetSysMenuList(userId int, IsSuper bool) []*GroupMenu {
	arrGroupMenu := GetGroupMenu(0)
	if IsSuper {
		return arrGroupMenu
	}

	mpMenu := GetSysUserMenuRel(userId)
	newGroupMenu := GetUserMenuRelList(arrGroupMenu, mpMenu)
	return newGroupMenu
}

func GetUserMenuRelList(arrGroupMenu []*GroupMenu, mpMenu map[int]int) []*GroupMenu {
	arrGroupMenu1 := make([]*GroupMenu, 0)
	for i := 0; i < len(arrGroupMenu); i++ {
		if v, ok := mpMenu[arrGroupMenu[i].Id]; ok {
			if v > 0 {
				arrGroupMenu1 = append(arrGroupMenu1, arrGroupMenu[i])
			}
		}
	}
	for i := 0; i < len(arrGroupMenu1); i++ {
		arrGroupMenu1[i].SubMenuList = GetUserMenuRelList(arrGroupMenu1[i].SubMenuList, mpMenu)
		//arrGroupMenu1[i].MenuType = mconst.MenuType_Group
		if len(arrGroupMenu1[i].SubMenuList) == 0 {
			arrGroupMenu1[i].SubMenuList = nil
			//arrGroupMenu1[i].MenuType = mconst.MenuType_Fun
		}
	}

	arrGroupMenu2 := make([]*GroupMenu, 0)
	for i := 0; i < len(arrGroupMenu1); i++ {
		if arrGroupMenu1[i].MenuType == mconst.MenuType_Group {
			if len(arrGroupMenu1[i].SubMenuList) > 0 {
				arrGroupMenu2 = append(arrGroupMenu2, arrGroupMenu1[i])
			}
		} else {
			arrGroupMenu2 = append(arrGroupMenu2, arrGroupMenu1[i])
		}
	}

	return arrGroupMenu2
}

func GetGroupMenu(Pid int) []*GroupMenu {
	arrGroupMenu := make([]*GroupMenu, 0)
	sql := `SELECT * from sys_menu where pid = ? order by seq `
	o := orm.NewOrm()
	_, e := o.Raw(sql, Pid).QueryRows(&arrGroupMenu)
	if Langs.GetLangs().Lang != "" {
		aLang := Langs.GetLangs()
		for i := 0; i < len(arrGroupMenu); i++ {
			arrGroupMenu[i].Title = aLang.Langs[aLang.Lang][arrGroupMenu[i].Id]
		}
	}

	if e != nil {
		ttLog.LogDebug(e)
	}
	for _, gMenu := range arrGroupMenu {
		gMenu.SubMenuList = GetGroupMenu(gMenu.Id)
		//gMenu.MenuType = mconst.MenuType_Group
		if len(gMenu.SubMenuList) == 0 {
			//gMenu.MenuType = mconst.MenuType_Fun
			gMenu.SubMenuList = nil
		}
		if gMenu.UrlFor == "" {
			gMenu.UrlFor = fmt.Sprintf("A%d", gMenu.Id)
		}
	}

	return arrGroupMenu
}
