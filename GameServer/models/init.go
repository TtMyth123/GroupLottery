package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	user_models "ttmyth123/GroupLottery/UserInfoRpc/models"

	_ "github.com/go-sql-driver/mysql"
)

type InitDBConfig struct {
	dbType  string
	dbAlias string
	dbName  string
	dbUser  string
	dbPwd   string
	dbHost  string
	dbPort  string
}

func Init() {
	InitRegisterModel()
	user_models.InitRegisterModel()

	initDatabase()

	InitData()
}

func InitRegisterModel() {
	orm.RegisterModel(new(LoOddsInfo))
	orm.RegisterModel(new(LoBetGroupInfo))
	orm.RegisterModel(new(LoBetInfo))
	orm.RegisterModel(new(LoAwardInfo))
	orm.RegisterModel(new(LoSetAwardInfo))
	orm.RegisterModel(new(LoUserInfo2))

	orm.RegisterModel(new(TtDrawSaveSet))
	orm.RegisterModel(new(TtDrawMoney))
	orm.RegisterModel(new(TtSaveMoney))
	orm.RegisterModel(new(TtGameInfo))

	orm.RegisterModel(new(TtArticleInfo))
	orm.RegisterModel(new(TtServiceInfo))
	orm.RegisterModel(new(TtAgentPayInfo))
	orm.RegisterModel(new(TtRebateSet))

}
func InitData() {
	InitLoOddsInfo()
	InitTtDrawSaveSet()
	InitTtRebateSet()
	InitTtGameInfo()
	InitTtServiceInfo()
	InitArticleInfoList()
	InitAgentPayInfo()
}

//初始化数据连接
func initDatabase() {
	//读取配置文件，设置数据库参数
	//数据库类别
	dbType := "mysql"
	//连接名称
	dbAlias := beego.AppConfig.String(dbType + "::db_alias")
	//数据库名称
	dbName := beego.AppConfig.String(dbType + "::db_name")
	//数据库连接用户名
	dbUser := beego.AppConfig.String(dbType + "::db_user")
	//数据库连接用户名
	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")
	//数据库IP（域名）
	dbHost := beego.AppConfig.String(dbType + "::db_host")
	//数据库端口
	dbPort := beego.AppConfig.String(dbType + "::db_port")

	switch dbType {
	case "sqlite3":
		orm.RegisterDataBase(dbAlias, dbType, dbName)
	case "mysql":
		dbCharset := beego.AppConfig.String(dbType + "::db_charset")
		//dataS := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharset + "&loc=Asia%2FShanghai"
		dataS := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharset + "&loc=Local"
		//orm.RegisterDataBase(dbAlias, dbType, dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset="+dbCharset+"&loc=Asia%2FShanghai", 30)
		e := orm.RegisterDataBase(dbAlias, dbType, dataS, 30)
		if e != nil {
			panic(e)
		}
	}

	//如果是开发模式，则显示命令信息
	isDev := (beego.AppConfig.String("runmode") == "dev")
	//自动建表
	orm.RunSyncdb("default", false, true)
	if isDev {
		orm.Debug = isDev
	}
}
