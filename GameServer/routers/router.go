package routers

import (
	"github.com/TtMyth123/GameServer/controllers"
	"github.com/astaxie/beego"
)

func Init() {
	beego.AutoRouter(&controllers.ApiController{})

	beego.AutoRouter(&controllers.RegLoginController{})
}
