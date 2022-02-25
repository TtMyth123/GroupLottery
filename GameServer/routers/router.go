package routers

import (
	"github.com/astaxie/beego"
	"ttmyth123/GroupLottery/GameServer/controllers"
)

func Init() {
	beego.AutoRouter(&controllers.ApiController{})

	beego.AutoRouter(&controllers.RegLoginController{})
}
