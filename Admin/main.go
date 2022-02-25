package main

import (
	"github.com/astaxie/beego/plugins/cors"
	"ttmyth123/GroupLottery/Admin/routers"
	"ttmyth123/GroupLottery/Admin/sysInit"

	"github.com/astaxie/beego"
)

func main() {
	sysInit.Init()
	routers.Init()
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"token", "Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		AllowCredentials: true,
	}))
	beego.Run()
}
