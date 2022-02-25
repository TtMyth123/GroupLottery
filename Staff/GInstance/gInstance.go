package GInstance

import (
	"github.com/astaxie/beego"
	"ttmyth123/GroupLottery/UserInfoRpc/UserRpcClient"
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
