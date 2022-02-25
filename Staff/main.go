package main

import (
	"github.com/astaxie/beego"
	_ "ttmyth123/GroupLottery/Staff/routers"
	"ttmyth123/GroupLottery/Staff/sysInit"
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
