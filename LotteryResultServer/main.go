package main

import (
	"github.com/astaxie/beego"
	"ttmyth123/GroupLottery/LotteryResultServer/GInstance"
	_ "ttmyth123/GroupLottery/LotteryResultServer/routers"
)

func main() {
	GInstance.Init()
	beego.Run()
}
