package sysInit

import (
	"ttmyth123/GroupLottery/GameServer/CacheData"
	"ttmyth123/GroupLottery/GameServer/GConfig"
	"ttmyth123/GroupLottery/GameServer/GInstance"
	"ttmyth123/GroupLottery/GameServer/GInstance/GTtHint"
	"ttmyth123/GroupLottery/GameServer/OtherServer/httpGameServer"
	"ttmyth123/GroupLottery/GameServer/models"
	"ttmyth123/kit/ttLog"
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
