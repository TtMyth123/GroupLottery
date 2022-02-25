package routers

import (
	"github.com/astaxie/beego"
	"ttmyth123/GroupLottery/UserInfoRpc/controllers"
)

func Init() {
	beego.Router("/", &controllers.MainController{})
}
