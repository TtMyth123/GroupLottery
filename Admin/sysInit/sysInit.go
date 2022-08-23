package sysInit

import (
	"github.com/TtMyth123/Admin/CacheData"
	"github.com/TtMyth123/Admin/GInstance"
	"github.com/TtMyth123/Admin/models"
	"github.com/TtMyth123/kit/ttLog"
)

func Init() {
	ttLog.InitLogs()
	CacheData.Init()
	models.Init()
	GInstance.Init()
}
