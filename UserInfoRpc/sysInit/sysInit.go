package sysInit

import (
	game_models "github.com/TtMyth123/GameServer/models"
	"github.com/TtMyth123/UserInfoRpc/GConfig"
	"github.com/TtMyth123/UserInfoRpc/GInstance/AreaConfig"
	"github.com/TtMyth123/UserInfoRpc/GInstance/GTtHint"
	"github.com/TtMyth123/UserInfoRpc/models"
	"github.com/TtMyth123/UserInfoRpc/routers"
	"github.com/TtMyth123/kit/ttLog"
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
