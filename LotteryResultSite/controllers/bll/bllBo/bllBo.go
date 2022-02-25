package bllBo

import (
	"encoding/json"
	"strings"
	"time"
)

type AwardInfo struct {
	Id         int
	GameType   int
	LotteryStr string //期号
	ResultNums string //开奖结果数据

	CurLotteryTime time.Time //当前开期时间

	GameName string
}

func (d AwardInfo) MarshalJSON() ([]byte, error) {
	type Alias AwardInfo
	switch d.GameType {
	default:
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
}

type GroupAwardInfo struct {
	C         int
	MaxId     int
	PageCount int
}
