package sysInit

import (
	"ttmyth123/GroupLottery/LotteryResultSite/models"
	"ttmyth123/kit/ttLog"
)

func Init() {
	ttLog.InitLogs()
	models.Init()
}
