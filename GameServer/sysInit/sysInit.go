package sysInit

import (
	"github.com/TtMyth123/GameServer/CacheData"
	"github.com/TtMyth123/GameServer/GConfig"
	"github.com/TtMyth123/GameServer/GInstance"
	"github.com/TtMyth123/GameServer/GInstance/GTtHint"
	"github.com/TtMyth123/GameServer/OtherServer/httpGameServer"
	"github.com/TtMyth123/GameServer/models"
	"github.com/TtMyth123/kit/ttLog"
)

func Init() {
	//创建共享数据
	ttLog.InitLogs()
	GConfig.Init()
	GTtHint.Init()
	models.Init()
	httpGameServer.Init()
	CacheData.Init()
	GInstance.Init()
}
