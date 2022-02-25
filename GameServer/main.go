package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"ttmyth123/GroupLottery/GameServer/routers"
	"ttmyth123/GroupLottery/GameServer/sysInit"
	"ttmyth123/kit"
)

func main() {
	fmt.Println("aaa")
	str, _ := os.Getwd()
	fmt.Println("当前目录：", str)
	sysInit.Init()
	routers.Init()
	beego.Run()
	aa := kit.GetGuid()
	fmt.Println("当前目录：", aa)
}
