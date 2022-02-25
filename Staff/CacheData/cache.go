package CacheData

import (
	"fmt"
	"github.com/astaxie/beego"
	"ttmyth123/GroupLottery/Staff/controllers/base/box"
	"ttmyth123/kit/beegoCacheKit"
	"ttmyth123/kit/ttLog"
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
func getUserInfoKey(userId int64) string {
	return fmt.Sprintf("userInfo:%d", userId)
}

func GetSysUserInfo(userId int64) box.SysUserEx {
	aUser := box.SysUserEx{}
	key := getUserInfoKey(userId)
	e := GetBeegoCache().GetCache(key, &aUser)
	if e != nil {
		ttLog.LogWarning(e)
	}

	return aUser
}

func ReLoadSysUserInfo(aUser box.SysUserEx) {
	key := getUserInfoKey(aUser.Id)
	e := GetBeegoCache().SetCache(key, aUser, beegoCacheKit.TimeoutOneDay)
	if e != nil {
		ttLog.LogError(e)
	}
}

func getSetTokenKey(TokenKey string) string {
	return fmt.Sprintf("TokenKey:%s", TokenKey)
}

func DelToken(TokenKey string) error {
	key := getSetTokenKey(TokenKey)
	e := GetBeegoCache().DelCache(key)
	return e
}

func SetToken(TokenKey string, aUser box.SysUserEx) error {
	key := getSetTokenKey(TokenKey)
	e := GetBeegoCache().SetCache(key, aUser, beegoCacheKit.TimeoutOneWeek)
	if e != nil {
		ttLog.LogError(e)
	}

	return e
}

func GetToken(TokenKey string) *box.SysUserEx {
	key := getSetTokenKey(TokenKey)
	aUser := box.SysUserEx{}
	e := GetBeegoCache().GetCache(key, &aUser)
	if e != nil {
		return nil
	}
	return &aUser
}
