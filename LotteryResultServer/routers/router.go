package routers

import (
	"github.com/astaxie/beego"
	"ttmyth123/GroupLottery/LotteryResultServer/controllers"
)

func init() {
	beego.AutoRouter(&controllers.ApiController{})
	beego.Router("/", &controllers.MainController{})
}
