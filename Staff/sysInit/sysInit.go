package sysInit

import (
	"ttmyth123/GroupLottery/Staff/CacheData"
	"ttmyth123/GroupLottery/Staff/GConfig"
	"ttmyth123/GroupLottery/Staff/GData"
	"ttmyth123/GroupLottery/Staff/GData/chat"
	"ttmyth123/GroupLottery/Staff/GInstance"
	"ttmyth123/GroupLottery/Staff/OtherServer/httpGameServer"
	"ttmyth123/GroupLottery/Staff/models"
	"ttmyth123/GroupLottery/Staff/routers"
	"ttmyth123/kit/ttLog"
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
