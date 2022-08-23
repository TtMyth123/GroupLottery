package GInstance

import (
	"github.com/TtMyth123/UserInfoRpc/UserRpcClient"
	"github.com/astaxie/beego"
)

var (
	mUserRpcClient *UserRpcClient.RpcClient
)

func Init() {
	addr := beego.AppConfig.String("UserRpcClientAddr")
	mUserRpcClient = UserRpcClient.NewRpcClient(addr)
}

func GetUserRpcClient() *UserRpcClient.RpcClient {
	return mUserRpcClient
}
