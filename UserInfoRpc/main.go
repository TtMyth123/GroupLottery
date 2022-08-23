package main

import (
	"github.com/TtMyth123/UserInfoRpc/RpcServer"
	_ "github.com/TtMyth123/UserInfoRpc/routers"
	"github.com/TtMyth123/UserInfoRpc/sysInit"
	"github.com/astaxie/beego"
)

func main() {
	sysInit.Init()
	strAddr := beego.AppConfig.String("UserRpcServerAddr")
	RpcServer.NewRpcServer(strAddr)
	beego.Run()
}
