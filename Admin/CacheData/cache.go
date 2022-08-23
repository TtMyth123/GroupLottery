package CacheData

import (
	"fmt"
	"github.com/TtMyth123/Admin/models"
	"github.com/TtMyth123/kit/beegoCacheKit"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego"
)

var mBeegoCache *beegoCacheKit.BeegoCache

func Init() {
	host := beego.AppConfig.String("redis::IP")
	pwd := beego.AppConfig.String("redis::Pwd")
	dbIndex, _ := beego.AppConfig.Int("redis::Index")
	DBName := beego.AppConfig.String("redis::DBName")
	mBeegoCache = beegoCacheKit.NewBeegoCache(DBName)
	mBeegoCache.InitCache(host, pwd, dbIndex)
}

func GetBeegoCache() *beegoCacheKit.BeegoCache {

	return mBeegoCache
}
func getUserInfoKey(userId int) string {
	return fmt.Sprintf("userInfo:%d", userId)
}

func GetSysUserInfo(userId int) models.SysUserEx {
	aUser := models.SysUserEx{}
	key := getUserInfoKey(userId)
	e := GetBeegoCache().GetCache(key, &aUser)
	if e != nil {
		ttLog.LogWarning(e)
	}

	return aUser
}

func ReLoadSysUserInfo(aUser models.SysUserEx) models.SysUserEx {
	key := getUserInfoKey(aUser.Id)
	e := GetBeegoCache().SetCache(key, aUser, beegoCacheKit.TimeoutOneDay)
	if e != nil {
		ttLog.LogError(e)
	}
	return aUser
}
func GetScoreRequisitionKey(typeScore int) string {
	key := fmt.Sprintf("bGetScoreRequisition:%d", typeScore)
	return key
}

func DelScoreRequisition(typeScore int) {
	key := GetScoreRequisitionKey(typeScore)
	GetBeegoCache().DelCache(key)
}
