package routers

import (
	"github.com/TtMyth123/Admin/controllers"
	"github.com/astaxie/beego"
)

func Init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.MainController{})

	beego.AutoRouter(&controllers.ApiController{})
	beego.AutoRouter(&controllers.LoginRegController{})
	beego.AutoRouter(&controllers.SysController{})
	beego.AutoRouter(&controllers.ChatController{})
	beego.AutoRouter(&controllers.KitController{})
}
