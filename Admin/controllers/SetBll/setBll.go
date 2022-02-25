package SetBll

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"ttmyth123/GroupLottery/Admin/GInstance"
	"ttmyth123/GroupLottery/Admin/OtherServer/GameClientHttp"
	"ttmyth123/GroupLottery/GameServer/models"
	"ttmyth123/GroupLottery/GameServer/models/mconst"
	"ttmyth123/kit/sqlKit"
	"ttmyth123/kit/ttLog"
)

func GetOddsList(gameType int, arrBigType []int) ([]models.LoOddsInfo, error) {
	arr := make([]models.LoOddsInfo, 0)
	o := orm.NewOrm()
	sqlArgs := make([]interface{}, 0)

	sqlWhere := ` where a.game_type=? `
	sqlArgs = append(sqlArgs, gameType)

	if len(arrBigType) != 0 {
		sqlW := ""
		for _, bigType := range arrBigType {
			if bigType != 0 {
				sqlW += ` or a.big_type=? `
				sqlArgs = append(sqlArgs, bigType)
			}
		}
		sqlWhere += fmt.Sprintf(" and ( %s )", sqlW[3:])
	}

	sql := fmt.Sprintf(`select * from %s a %s`, mconst.TableName_LoOddsInfo, sqlWhere)
	_, e := o.Raw(sql, sqlArgs).QueryRows(&arr)

	return arr, e
}

func GetWsxZbcOddsList() ([]models.LoOddsInfo, error) {
	gameType := mconst.GameType_Wsx_203
	arr := make([]models.LoOddsInfo, 0)
	o := orm.NewOrm()
	sqlArgs := make([]interface{}, 0)

	sqlWhere := ` where a.game_type=? order by odds_type`
	sqlArgs = append(sqlArgs, gameType)

	sql := fmt.Sprintf(`select * from %s a %s`, mconst.TableName_LoOddsInfo, sqlWhere)
	_, e := o.Raw(sql, sqlArgs).QueryRows(&arr)

	return arr, e
}
func SaveWsxZbcOddsInfo(aOddsInfo models.LoOddsInfo) error {
	o := orm.NewOrm()
	e := aOddsInfo.Update(o, "Odds")
	if e != nil {
		return e
	}

	sql := fmt.Sprintf(`update %s set one_user_max_bet=?, one_user_min_bet=?, all_user_max_bet=?, updated_at=?
where game_type=? and big_type=?`, mconst.TableName_LoOddsInfo)
	_, e = o.Raw(sql, aOddsInfo.OneUserMaxBet, aOddsInfo.OneUserMinBet, aOddsInfo.AllUserMaxBet, time.Now(), mconst.GameType_Wsx_203,
		aOddsInfo.BigType).Exec()
	if e != nil {
		return e
	}
	e = GameClientHttp.OddsChange(mconst.GameType_Wsx_203)

	return e
}

type GroupLoOddsInfo struct {
	C             int
	FirstId       int
	AllUserMaxBet int
	OneUserMaxBet int
}

func GetOddsListFirstId(OddsDes string, gameType int, OddsType, BigType int, pageIndex, pageSize, FirstId int) (int, []models.LoOddsInfo, GroupLoOddsInfo) {
	PageTotal := 0
	arrData := make([]models.LoOddsInfo, 0)
	aGroupLoOddsInfo := GroupLoOddsInfo{}

	o := orm.NewOrm()
	sqlArgs := make([]interface{}, 0)

	sqlWhere := ` where a.game_type=? `
	sqlArgs = append(sqlArgs, gameType)

	if BigType != 0 {
		sqlWhere += " and a.big_type=?"
		sqlArgs = append(sqlArgs, BigType)
	}
	if OddsType != 0 {
		sqlWhere += " and a.odds_type=?"
		sqlArgs = append(sqlArgs, OddsType)
	}
	if OddsDes != "" {
		sqlWhere = sqlWhere + ` and locate(?,a.odds_des)>0`
		sqlArgs = append(sqlArgs, OddsDes)
	}

	sqlCount := fmt.Sprintf(`select count(1) c 
from %s a %s`, mconst.TableName_LoOddsInfo, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&PageTotal)
	if e != nil {
		ttLog.LogError(e)
		return PageTotal, arrData, aGroupLoOddsInfo
	}

	sqlCount = fmt.Sprintf(`select count(1) c, sum(a.all_user_max_bet) as all_user_max_bet 
,sum(a.one_user_max_bet) as one_user_max_bet
from %s a %s`, mconst.TableName_LoOddsInfo, sqlWhere)
	e = o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupLoOddsInfo)
	if e != nil {
		ttLog.LogError(e)
		return PageTotal, arrData, aGroupLoOddsInfo
	}

	offset, _ := sqlKit.GetOffset(PageTotal, pageSize, pageIndex)
	sqlWhere = sqlWhere + ` LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, pageSize)

	sql := fmt.Sprintf(`select a.* from %s a %s`, mconst.TableName_LoOddsInfo, sqlWhere)
	_, e = o.Raw(sql, sqlArgs).QueryRows(&arrData)
	if e != nil {
		ttLog.LogError(e)
		return PageTotal, arrData, aGroupLoOddsInfo
	}

	if len(arrData) > 0 && FirstId == 0 {
		aGroupLoOddsInfo.FirstId = arrData[0].Id
	} else {
		aGroupLoOddsInfo.FirstId = FirstId
	}
	return PageTotal, arrData, aGroupLoOddsInfo
}

func SaveBatchOdds(gameType, minKey, maxKey int, odds float64) error {
	o := orm.NewOrm()
	sql := fmt.Sprintf(`update %s set odds=? where game_type=? and odds_type>=? and odds_type<=?`, mconst.TableName_LoOddsInfo)
	_, e := o.Raw(sql, odds, gameType, minKey, maxKey).Exec()
	return e
}

func GetOddsInfo(Id int) (models.LoOddsInfo, error) {
	o := orm.NewOrm()
	aLoOddsInfo := models.LoOddsInfo{Id: Id}
	e := o.Read(&aLoOddsInfo)
	return aLoOddsInfo, e
}

func GetTtRebateSet(AreaId int) models.TtRebateSet {
	return models.GetTtRebateSet()
}

func SaveTtRebateSet(aTtRebateSet models.TtRebateSet) error {
	o := orm.NewOrm()
	e := aTtRebateSet.Update(o, "Level", "BetRebateRatio", "BetRebateRatio1", "BetRebateRatio2", "BetRebateRatio3")
	if e != nil {
		return e
	}
	e = GInstance.GetUserRpcClient().ReRebateSetConfig(0)
	return e
}
