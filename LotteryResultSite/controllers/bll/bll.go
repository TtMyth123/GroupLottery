package bll

import (
	"fmt"
	"github.com/TtMyth123/GameServer/models/mconst"
	"github.com/TtMyth123/LotteryResultSite/controllers/bll/bllBo"
	"github.com/TtMyth123/kit/sqlKit"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
)

func GetAwardList(LotteryStr, strDay string, gameType int, pageIndex, pageSize, MaxId int) ([]bllBo.AwardInfo, bllBo.GroupAwardInfo) {
	aGroup := bllBo.GroupAwardInfo{}
	arrData := make([]bllBo.AwardInfo, 0)
	if pageIndex < 1 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 10
	} else if pageSize > 100 {
		pageSize = 100
	}

	o := orm.NewOrm()
	sqlArgs := make([]interface{}, 0)
	if MaxId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, mconst.TableName_LoAwardInfo)
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&MaxId)
		if e != nil {
			ttLog.LogError(e)
		}
	}

	sqlWhere := " a.id<=? and a.game_type=? "
	sqlArgs = append(sqlArgs, MaxId, gameType)
	if LotteryStr != "" {
		sqlWhere += ` and a.lottery_str=?`
		sqlArgs = append(sqlArgs, LotteryStr)
	}

	if strDay != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.cur_lottery_time,'%[1]s') = ? `, "%Y-%m-%d")
		sqlArgs = append(sqlArgs, strDay)
	}

	sqlCount := fmt.Sprintf(`select count(1) c from %s a where %s`,
		mconst.TableName_LoAwardInfo, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroup)
	if e != nil {
		return arrData, aGroup
	}

	offset, pageC := sqlKit.GetOffset(aGroup.C, pageSize, pageIndex)
	sqlWhere = sqlWhere + ` order by a.cur_lottery_time desc LIMIT ?,? `
	sqlArgs = append(sqlArgs, offset, pageSize)

	sql := fmt.Sprintf(`select a.*  from %s a where %s`,
		mconst.TableName_LoAwardInfo, sqlWhere)
	c, e := o.Raw(sql, sqlArgs).QueryRows(&arrData)
	if e != nil {
		ttLog.LogError(c, e)
		return arrData, aGroup
	}

	aGroup.MaxId = MaxId
	aGroup.PageCount = pageC

	return arrData, aGroup
}
