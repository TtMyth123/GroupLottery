package main

import (
	"github.com/TtMyth123/LotteryResultServer/GInstance"
	_ "github.com/TtMyth123/LotteryResultServer/routers"
	"github.com/astaxie/beego"
)

func main() {
	GInstance.Init()
	beego.Run()
}
