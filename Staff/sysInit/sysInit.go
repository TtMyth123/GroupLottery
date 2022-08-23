package sysInit

import (
	"github.com/TtMyth123/Staff/CacheData"
	"github.com/TtMyth123/Staff/GConfig"
	"github.com/TtMyth123/Staff/GData"
	"github.com/TtMyth123/Staff/GData/chat"
	"github.com/TtMyth123/Staff/GInstance"
	"github.com/TtMyth123/Staff/OtherServer/httpGameServer"
	"github.com/TtMyth123/Staff/models"
	"github.com/TtMyth123/Staff/routers"
	"github.com/TtMyth123/kit/ttLog"
)

func Init() {
	//创建共享数据
	ttLog.InitLogs()
	CacheData.Init()
	GConfig.Init()
	httpGameServer.Init()
	GInstance.Init()
	//GTtHint.Init()
	//game_models.InitRegisterModel()
	models.Init()
	GData.InitStaffDataManage()

	//AreaConfig.Init()
	chat.Init()
	routers.Init()
}
