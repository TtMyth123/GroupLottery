package main

import (
	_ "github.com/TtMyth123/Staff/routers"
	"github.com/TtMyth123/Staff/sysInit"
	"github.com/astaxie/beego"
)

func main() {
	sysInit.Init()
	beego.SetStaticPath("/css", "css")
	beego.SetStaticPath("/fonts", "fonts")
	beego.SetStaticPath("/icons", "icons")
	beego.SetStaticPath("/img", "img")
	beego.SetStaticPath("/js", "js")
	beego.SetStaticPath("/media", "media")

	beego.Run()
}
