package CacheData

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"
	"ttmyth123/GroupLottery/UserInfoRpc/models"
	"ttmyth123/kit/beegoCacheKit"
	"ttmyth123/kit/timeKit"
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

func GetVisitorKey(userId int) string {
	return fmt.Sprintf("Visitor:%d", userId)
}

func GetVisitor(Id int) (models.TtGameUser, error) {
	if Id == 0 {
		i := int(timeKit.GlobaRand.Int63n(1000 - 100))
		newUserId := 100 + i
		for j := 100; j < 1000; j++ {
			key := GetVisitorKey(newUserId)
			tmp := models.TtGameUser{}
			e := mBeegoCache.GetCache(key, &tmp)
			if e != nil {
				newUserId++
				break
			}
		}

		aYhGameUser := models.TtGameUser{Id: newUserId,
			Nickname: fmt.Sprintf("游客%d", newUserId),
			UserName: fmt.Sprintf("游客%d", newUserId),
		}

		key := GetVisitorKey(newUserId)
		mBeegoCache.SetCache(key, aYhGameUser, beegoCacheKit.TimeoutOneDay)
		return aYhGameUser, nil
	} else {
		key := GetVisitorKey(Id)
		tmp := models.TtGameUser{}
		mBeegoCache.GetCache(key, &tmp)
		return tmp, nil
	}
}

func ClearVisitor(Id int) error {
	key := GetVisitorKey(Id)
	err := mBeegoCache.DelCache(key)
	return err
}

func DelUserSid(gameId int) error {
	key := fmt.Sprintf("GetUserSid:%d", gameId)
	e := GetBeegoCache().DelCache(key)
	return e
}

//func SetUserSid(gameId int, sid string)error {
//	key := fmt.Sprintf("GetUserSid:%d", gameId)
//	e := GetBeegoCache().SetCache(key, sid, beegoCacheKit.TimeoutOneWeek)
//	return e
//}
//func GetUserSid(gameId int)string  {
//	key := fmt.Sprintf("GetUserSid:%d", gameId)
//	sid := ""
//	GetBeegoCache().GetCache(key, &sid)
//	return sid
//}

func GetOneDayLossWinKey(date time.Time, userId int) string {
	strDay := timeKit.GetStrDay(date)
	key := fmt.Sprintf("GetOneDayLossWin:%s:%d", strDay, userId)
	return key
}
func GetOneDayLossWin(date time.Time, userId int) float64 {
	key := GetOneDayLossWinKey(date, userId)
	LossWin := 0.0
	GetBeegoCache().GetCache(key, &LossWin)
	return LossWin
}

/**
获取 用户，游戏，期号，赔率 投注Key
*/
func GetUserGamePeriodOddsTypeKey(Period string, gameType, OddsType, userId int) string {
	key := fmt.Sprintf("a_CurPeriodBetM:%d:%s:%d:%d", gameType, Period, OddsType, userId)
	return key
}

/**
获取 游戏，期号，赔率 投注Key
*/
func GetGamePeriodOddsTypeKey(Period string, gameType, OddsType int) string {
	key := fmt.Sprintf("b_CurPeriodBetM:%d:%s:%d", gameType, Period, OddsType)
	return key
}

/**
获取 游戏、期号、用户 投注Key
*/
func GetGamePeriodUserKey(Period string, gameType, UserId int) string {
	key := fmt.Sprintf("c_CurPeriodBetM:%d:%s:%d", gameType, Period, UserId)
	return key
}

/**
获取 游戏、期号  投注Key
*/
func GetGamePeriodKey(Period string, gameType int) string {
	key := fmt.Sprintf("d_CurPeriodBetM:%d:%s", gameType, Period)
	return key
}

/**
获取
*/
func CurDayBetMKey(t time.Time) string {
	key := fmt.Sprintf("CurDayBetM:%s", t.Format("2006-01-02"))
	return key
}

/**
获取
*/
func CurDayWinMKey(t time.Time) string {
	key := fmt.Sprintf("CurDayWinM:%s", t.Format("2006-01-02"))
	return key
}

func AddCurPeriodUserWin(tt time.Time, gameType, userId, OddsType, bet int, Period string, win float64) {
	key := CurDayWinMKey(tt)
	curM := 0.0
	mBeegoCache.GetCache(key, &curM)
	curM += win
	mBeegoCache.SetCache(key, curM, beegoCacheKit.TimeoutOneDay*31)

	key = GetOneDayLossWinKey(tt, userId)
	allLossWin := 0.0
	GetBeegoCache().GetCache(key, &allLossWin)
	LossWin := win - float64(bet)
	allLossWin += LossWin
	GetBeegoCache().SetCache(key, allLossWin, beegoCacheKit.TimeoutOneDay*31)
}

func AddCurPeriodUserBet(gameType, userId, betM int, OddsType int, Period string) {
	tt := time.Now()
	key := CurDayBetMKey(tt)
	curBetM := 0
	mBeegoCache.GetCache(key, &curBetM)
	curBetM += betM
	mBeegoCache.SetCache(key, curBetM, beegoCacheKit.TimeoutOneDay*31)

	//用户、游戏、赔率、期号
	key = GetUserGamePeriodOddsTypeKey(Period, gameType, OddsType, userId)
	curBetM = 0
	mBeegoCache.GetCache(key, &curBetM)
	curBetM += betM
	mBeegoCache.SetCache(key, curBetM, beegoCacheKit.TimeoutOneDay*2)

	//游戏、赔率、期号
	key = GetGamePeriodOddsTypeKey(Period, gameType, OddsType)
	curBetM = 0
	mBeegoCache.GetCache(key, &curBetM)
	curBetM += betM
	mBeegoCache.SetCache(key, curBetM, beegoCacheKit.TimeoutOneDay*2)

	//游戏、期号、用户
	key = GetGamePeriodUserKey(Period, gameType, userId)
	curBetM = 0
	mBeegoCache.GetCache(key, &curBetM)
	curBetM += betM
	mBeegoCache.SetCache(key, curBetM, beegoCacheKit.TimeoutOneDay*2)

	//游戏、期号
	key = GetGamePeriodKey(Period, gameType)
	curBetM = 0
	mBeegoCache.GetCache(key, &curBetM)
	curBetM += betM
	mBeegoCache.SetCache(key, curBetM, beegoCacheKit.TimeoutOneDay*2)
}

func GetCacheSumBet(Period string, gameType, OddsType, userId int) (int, int) {
	key := GetUserGamePeriodOddsTypeKey(Period, gameType, OddsType, userId)
	betUserGamePeriodOddsType := 0
	mBeegoCache.GetCache(key, &betUserGamePeriodOddsType)

	betGamePeriodOddsType := 0
	key = GetGamePeriodOddsTypeKey(Period, gameType, OddsType)
	mBeegoCache.GetCache(key, &betGamePeriodOddsType)
	return betUserGamePeriodOddsType, betGamePeriodOddsType
}

/**
游戏、期号、用户 投注金额
*/
func GetCurUserPeriodBet(userId, gameType int, Period string) int {
	key := GetGamePeriodUserKey(Period, gameType, userId)
	curBetM := 0
	mBeegoCache.GetCache(key, &curBetM)
	return curBetM
}
