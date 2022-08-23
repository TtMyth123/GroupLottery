package routers

import (
	"github.com/TtMyth123/LotteryResultSite/controllers"
	"github.com/astaxie/beego"
)

func Init() {
	beego.AutoRouter(&controllers.ApiController{})
	beego.Router("/", &controllers.MainController{})
}
