package sysInit

import (
	"ttmyth123/GroupLottery/Admin/CacheData"
	"ttmyth123/GroupLottery/Admin/GInstance"
	"ttmyth123/GroupLottery/Admin/models"
	"ttmyth123/kit/ttLog"
)

func Init() {
	ttLog.InitLogs()
	CacheData.Init()
	models.Init()
	GInstance.Init()
}
