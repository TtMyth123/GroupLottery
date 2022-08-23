package main

import (
	"github.com/TtMyth123/Admin/routers"
	"github.com/TtMyth123/Admin/sysInit"
	"github.com/astaxie/beego/plugins/cors"

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
