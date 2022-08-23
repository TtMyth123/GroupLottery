package AwardBll

import (
	"encoding/json"
	"fmt"
	"github.com/TtMyth123/GameServer/LotteryResult/WsxBbcResultKit"
	"github.com/TtMyth123/GameServer/LotteryResult/WsxNbcResultKit"
	"github.com/TtMyth123/GameServer/LotteryResult/WsxZbcResultKit"
	"github.com/TtMyth123/GameServer/models"
	"github.com/TtMyth123/GameServer/models/mconst"
	"github.com/TtMyth123/kit/sqlKit"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"

	gmconst "github.com/TtMyth123/GameServer/models/mconst"
)

type GroupAwardInfo struct {
	C       int
	FirstId int
}
type AwardInfo struct {
	Id             int
	GameType       int
	LotteryNum     int64  //期号
	LotteryStr     string //期号
	ResultNums     string //开奖结果数据
	OriginalResult string //原始接口数据

	NextLotteryStr  string
	NextLotteryTime time.Time `json:"-"` //下一期开奖时间
	CurLotteryTime  time.Time `json:"-"` //当前开期时间

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	GameName string
}

func (d AwardInfo) MarshalJSON() ([]byte, error) {
	type Alias AwardInfo
	switch d.GameType {
	case gmconst.GameType_Wsx_201:
		aResultNums := WsxNbcResultKit.GetStr2ResultNums(d.ResultNums)
		return json.Marshal(&struct {
			Alias
			CurLotteryTime string
			NumList        map[string]interface{}
		}{
			Alias:          (Alias)(d),
			CurLotteryTime: d.CurLotteryTime.Format("2006-01-02 15:04:05"),
			NumList:        aResultNums,
		})

	case gmconst.GameType_Wsx_202:
		aResultNums := WsxBbcResultKit.GetStr2ResultNums(d.ResultNums)
		return json.Marshal(&struct {
			Alias
			CurLotteryTime string
			NumList        map[string]interface{}
		}{
			Alias:          (Alias)(d),
			CurLotteryTime: d.CurLotteryTime.Format("2006-01-02 15:04:05"),
			NumList:        aResultNums,
		})

	case gmconst.GameType_Wsx_203:
		aResultNums := WsxZbcResultKit.GetStr2ResultNums(d.ResultNums)
		return json.Marshal(&struct {
			Alias
			CurLotteryTime string
			NumList        map[string]interface{}
		}{
			Alias:          (Alias)(d),
			CurLotteryTime: d.CurLotteryTime.Format("2006-01-02 15:04:05"),
			NumList:        aResultNums,
		})

	case gmconst.GameType_G28_041, gmconst.GameType_G28_042, gmconst.GameType_G28_043:
		arrNums := strings.Split(d.ResultNums, ",")
		return json.Marshal(&struct {
			Alias
			CurLotteryTime string
			NumList        []string
		}{
			Alias:          (Alias)(d),
			CurLotteryTime: d.CurLotteryTime.Format("2006-01-02 15:04:05"),
			NumList:        arrNums,
		})
	}

	return json.Marshal(&struct {
		Alias
		CurLotteryTime string
	}{
		Alias:          (Alias)(d),
		CurLotteryTime: d.CurLotteryTime.Format("2006-01-02 15:04:05"),
	})
}

func GetAwardList(LotteryStr, beginDay, endDay string, gameType int, pageIndex, pageSize, FirstId int) (int, []AwardInfo, GroupAwardInfo) {
	o := orm.NewOrm()
	arrData := make([]AwardInfo, 0)
	aGroupAwardInfo := GroupAwardInfo{}
	sqlArgs := make([]interface{}, 0)

	maxId := FirstId
	if FirstId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, mconst.TableName_LoAwardInfo)
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&maxId)
		if e != nil {
			ttLog.LogError(e)
		}
	}

	sqlWhere := "and a.id<=? and a.game_type=? "
	sqlArgs = append(sqlArgs, maxId, gameType)
	if LotteryStr != "" {
		sqlWhere += ` and a.lottery_str=?`
		sqlArgs = append(sqlArgs, LotteryStr)
	}
	if beginDay != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.cur_lottery_time,'%[1]s') >= ? `, "%Y-%m-%d")
		sqlArgs = append(sqlArgs, beginDay)
	}
	if endDay != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(` and date_format(a.cur_lottery_time,'%[1]s') <= ? `, "%Y-%m-%d")
		sqlArgs = append(sqlArgs, endDay)
	}

	sqlCount := fmt.Sprintf(`select count(1) c from %s a,%s b where a.game_type=b.game_type %s`,
		mconst.TableName_LoAwardInfo, mconst.TableName_TtGameInfo, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupAwardInfo)
	if e != nil {
		return 0, arrData, aGroupAwardInfo
	}

	offset, _ := sqlKit.GetOffset(aGroupAwardInfo.C, pageSize, pageIndex)
	sqlWhere = sqlWhere + ` order by a.cur_lottery_time desc LIMIT ?,? `
	sqlArgs = append(sqlArgs, offset, pageSize)

	sql := fmt.Sprintf(`select a.*, b.game_name from %s a,%s b where a.game_type=b.game_type %s`,
		mconst.TableName_LoAwardInfo, mconst.TableName_TtGameInfo, sqlWhere)
	c, e := o.Raw(sql, sqlArgs).QueryRows(&arrData)
	if e != nil {
		ttLog.LogError(c, e)
		return 0, arrData, aGroupAwardInfo
	}

	aGroupAwardInfo.FirstId = maxId

	return aGroupAwardInfo.C, arrData, aGroupAwardInfo
}

//func GetAwardListFirstId(gameType int, Period string,pageIndex, pageSize, FirstId int) (int,[]AwardInfo, GroupAwardInfo) {
//	PageTotal:=0
//	o := orm.NewOrm()
//	arrData := make([]AwardInfo,0)
//	aGroupAwardInfo := GroupAwardInfo{}
//
//	sqlArgs := make([]interface{}, 0)
//
//	sqlWhere := " and a.game_type=? "
//	sqlArgs = append(sqlArgs, gameType)
//
//	if Period != "" {
//		sqlWhere += ` and locate(?,a.lottery_str)>0`
//		sqlArgs = append(sqlArgs, Period)
//	}
//
//	sqlCount := fmt.Sprintf(`select count(1) c from %s a,%s b where a.game_type=b.game_type %s`,
//		mconst.TableName_LoAwardInfo,mconst.TableName_TtGameInfo, sqlWhere)
//	e := o.Raw(sqlCount, sqlArgs).QueryRow(&PageTotal)
//	if e != nil {
//		return 0,arrData,aGroupAwardInfo
//	}
//
//	if FirstId!=0 {
//		sqlWhere += " and a.id<=? "
//		sqlArgs = append(sqlArgs, FirstId)
//	}
//
//	sqlCount = fmt.Sprintf(`select count(1) c from %s a,%s b where a.game_type=b.game_type %s`,
//		mconst.TableName_LoAwardInfo,mconst.TableName_TtGameInfo, sqlWhere)
//	e = o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupAwardInfo)
//	if e != nil {
//		return 0,arrData,aGroupAwardInfo
//	}
//
//	offset, _ := sqlKit.GetOffset(aGroupAwardInfo.C, pageSize, pageIndex)
//	sqlWhere = sqlWhere + ` order by a.cur_lottery_time desc,a.id desc LIMIT ?,? `
//	sqlArgs = append(sqlArgs, offset, pageSize)
//
//	sql := fmt.Sprintf(`select a.*, b.game_name from %s a,%s b where a.game_type=b.game_type %s`,
//		mconst.TableName_LoAwardInfo,mconst.TableName_TtGameInfo, sqlWhere)
//	c,e  := o.Raw(sql, sqlArgs).QueryRows(&arrData)
//	if e != nil {
//		ttLog.LogError(c, e)
//		return 0, arrData, aGroupAwardInfo
//	}
//
//	if len(arrData) >0 {
//		aGroupAwardInfo.FirstId = arrData[0].Id
//	}
//
//	return PageTotal,arrData,aGroupAwardInfo
//}

type AwardInfoDetail struct {
	models.LoAwardInfo

	Result map[string]interface{}
}

func GetAwardDetail(id int) (AwardInfoDetail, error) {
	aAwardInfoDetail := AwardInfoDetail{}
	aAwardInfoDetail.Id = id
	o := orm.NewOrm()
	aLoAwardInfo := models.LoAwardInfo{Id: id}
	e := o.Read(&aLoAwardInfo)
	if e != nil {
		return AwardInfoDetail{}, e
	}
	aAwardInfoDetail.LoAwardInfo = aLoAwardInfo
	Result := make(map[string]interface{})
	e = json.Unmarshal([]byte(aAwardInfoDetail.ResultNums), &Result)

	aAwardInfoDetail.Result = Result
	return aAwardInfoDetail, e
}

type SetAwardInfo struct {
	Id         int
	GameType   int
	LotteryNum int64  //期号
	LotteryStr string //期号
	ResultNums string

	CreatedAt time.Time
	UpdatedAt time.Time

	GameName string
}

func (d SetAwardInfo) MarshalJSON() ([]byte, error) {

	type Alias SetAwardInfo
	return json.Marshal(&struct {
		Alias
		CreatedAt string
		UpdatedAt string
	}{
		Alias:     (Alias)(d),
		CreatedAt: d.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: d.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}

type GroupSetAwardInfoInfo struct {
	C       int
	FirstId int
}

func GetSetAwardInfo(gameType int, pageIndex, pageSize, FirstId int, LotteryStr string) (int, []SetAwardInfo, GroupSetAwardInfoInfo) {
	o := orm.NewOrm()
	arrData := make([]SetAwardInfo, 0)
	aGroupAwardInfo := GroupSetAwardInfoInfo{}
	sqlArgs := make([]interface{}, 0)

	maxId := FirstId
	if FirstId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, mconst.TableName_LoSetAwardInfo)
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&maxId)
		if e != nil {
			ttLog.LogError(e)
		}
	}

	sqlWhere := "and a.id<=? and a.game_type=? "
	sqlArgs = append(sqlArgs, maxId, gameType)
	if LotteryStr != "" {
		sqlWhere += ` and a.lottery_str=?`
		sqlArgs = append(sqlArgs, LotteryStr)
	}

	sqlCount := fmt.Sprintf(`select count(1) c from %s a,%s b where a.game_type=b.game_type %s`,
		mconst.TableName_LoSetAwardInfo, mconst.TableName_TtGameInfo, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupAwardInfo)
	if e != nil {
		return 0, arrData, aGroupAwardInfo
	}

	offset, _ := sqlKit.GetOffset(aGroupAwardInfo.C, pageSize, pageIndex)
	sqlWhere = sqlWhere + ` order by a.lottery_str desc LIMIT ?,? `
	sqlArgs = append(sqlArgs, offset, pageSize)

	sql := fmt.Sprintf(`select a.*, b.game_name from %s a,%s b where a.game_type=b.game_type %s`,
		mconst.TableName_LoSetAwardInfo, mconst.TableName_TtGameInfo, sqlWhere)
	c, e := o.Raw(sql, sqlArgs).QueryRows(&arrData)
	if e != nil {
		ttLog.LogError(c, e)
		return 0, arrData, aGroupAwardInfo
	}

	aGroupAwardInfo.FirstId = maxId

	return aGroupAwardInfo.C, arrData, aGroupAwardInfo
}
