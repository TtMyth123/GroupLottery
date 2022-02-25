package main

import (
	"github.com/astaxie/beego"
	"ttmyth123/GroupLottery/UserInfoRpc/RpcServer"
	_ "ttmyth123/GroupLottery/UserInfoRpc/routers"
	"ttmyth123/GroupLottery/UserInfoRpc/sysInit"
)

func main() {
	sysInit.Init()
	strAddr := beego.AppConfig.String("UserRpcServerAddr")
	RpcServer.NewRpcServer(strAddr)
	beego.Run()
}
