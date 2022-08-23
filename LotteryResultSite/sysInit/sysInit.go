package sysInit

import (
	"github.com/TtMyth123/LotteryResultSite/models"
	"github.com/TtMyth123/kit/ttLog"
)

func Init() {
	ttLog.InitLogs()
	models.Init()
}
