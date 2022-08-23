package main

import (
	"github.com/TtMyth123/LotteryResultSite/routers"
	"github.com/TtMyth123/LotteryResultSite/sysInit"
	"github.com/astaxie/beego"
)

func main() {
	sysInit.Init()
	routers.Init()
	beego.Run()
}
