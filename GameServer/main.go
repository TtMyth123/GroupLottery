package main

import (
	"fmt"
	"github.com/TtMyth123/GameServer/routers"
	"github.com/TtMyth123/GameServer/sysInit"
	"github.com/TtMyth123/kit"
	"github.com/astaxie/beego"
	"os"
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
