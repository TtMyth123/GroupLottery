package routers

import (
	"github.com/astaxie/beego"
	"ttmyth123/GroupLottery/Staff/controllers"
)

func Init() {
	beego.AutoRouter(&controllers.ChatController{})
	beego.AutoRouter(&controllers.OtherController{})
	beego.AutoRouter(&controllers.LoginRegController{})
	beego.AutoRouter(&controllers.TestController{})
	beego.AutoRouter(&controllers.ChatExController{})

	beego.Router("/", &controllers.MainController{})
}
