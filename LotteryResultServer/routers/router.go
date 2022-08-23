package routers

import (
	"github.com/TtMyth123/LotteryResultServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.ApiController{})
	beego.Router("/", &controllers.MainController{})
}
