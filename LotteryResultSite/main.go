package main

import (
	"github.com/astaxie/beego"
	"ttmyth123/GroupLottery/LotteryResultSite/routers"
	"ttmyth123/GroupLottery/LotteryResultSite/sysInit"
)

func main() {
	sysInit.Init()
	routers.Init()
	beego.Run()
}
