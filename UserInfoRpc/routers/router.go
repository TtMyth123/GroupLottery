package routers

import (
	"github.com/TtMyth123/UserInfoRpc/controllers"
	"github.com/astaxie/beego"
)

func Init() {
	beego.Router("/", &controllers.MainController{})

	beego.AutoRouter(&controllers.MainController{})
}
