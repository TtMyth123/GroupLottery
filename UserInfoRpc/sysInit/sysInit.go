package sysInit

import (
	game_models "ttmyth123/GroupLottery/GameServer/models"
	"ttmyth123/GroupLottery/UserInfoRpc/GConfig"
	"ttmyth123/GroupLottery/UserInfoRpc/GInstance/AreaConfig"
	"ttmyth123/GroupLottery/UserInfoRpc/GInstance/GTtHint"
	"ttmyth123/GroupLottery/UserInfoRpc/models"
	"ttmyth123/GroupLottery/UserInfoRpc/routers"
	"ttmyth123/kit/ttLog"
)

func Init() {
	//创建共享数据
	ttLog.InitLogs()

	GConfig.Init()
	GTtHint.Init()
	game_models.InitRegisterModel()
	models.Init()

	AreaConfig.Init()

	routers.Init()
}
