package RpcServer

import (
	"ttmyth123/GroupLottery/UserInfoRpc/RpcServer/RpcConst"
	"ttmyth123/GroupLottery/UserInfoRpc/RpcServer/UserServer"
	"ttmyth123/kit/ttLog"

	"github.com/smallnest/rpcx/server"
)

type RpcServer struct {
	*server.Server
}

func NewRpcServer(strAddr string) *RpcServer {
	aRpcServer := new(RpcServer)
	aRpcServer.Server = server.NewServer()

	aUserInfoServer := UserServer.NewUserServer()
	aRpcServer.RegisterName(RpcConst.ServerName, aUserInfoServer, "")
	go func() {
		ttLog.LogDebug("Run aUserInfoServer begin:", strAddr)
		ee := aRpcServer.Serve("tcp", strAddr)
		ttLog.LogDebug("Run aUserInfoServer ee:", ee)
		if ee == nil {
			ttLog.LogDebug("Run aUserInfoServer end OK")
		}
	}()
	return aRpcServer
}
