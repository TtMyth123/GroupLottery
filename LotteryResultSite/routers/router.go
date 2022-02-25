package routers

import (
	"github.com/astaxie/beego"
	"ttmyth123/GroupLottery/LotteryResultSite/controllers"
)

func Init() {
	beego.AutoRouter(&controllers.ApiController{})
	beego.Router("/", &controllers.MainController{})
}
