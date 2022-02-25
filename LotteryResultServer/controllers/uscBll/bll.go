package uscBll

import (
	"ttmyth123/GroupLottery/LotteryResultServer/GInstance"
	"ttmyth123/GroupLottery/LotteryResultServer/GInstance/lottery"
)

/**
6
*/
func SendGetResultByIndex(gameIndex int) lottery.NewAwardInfo {
	aNewAwardInfo := GInstance.GetServerByIndex(gameIndex).GetNewAwardInfo()
	return aNewAwardInfo
}

func SendGetResultByType(gameType int) lottery.NewAwardInfo {
	aNewAwardInfo := GInstance.GetServer(gameType).GetNewAwardInfo()
	return aNewAwardInfo
}

func UpdateResultUrl(gameType int, url string) error {
	e := GInstance.GetServer(gameType).UpdateResultUrl(url)
	return e
}
