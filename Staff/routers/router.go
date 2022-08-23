package routers

import (
	"github.com/TtMyth123/Staff/controllers"
	"github.com/astaxie/beego"
)

func Init() {
	beego.AutoRouter(&controllers.ChatController{})
	beego.AutoRouter(&controllers.OtherController{})
	beego.AutoRouter(&controllers.LoginRegController{})
	beego.AutoRouter(&controllers.TestController{})
	beego.AutoRouter(&controllers.ChatExController{})

	beego.Router("/", &controllers.MainController{})
}
