package api

import (
	gin "github.com/baytan0720/gin-grpc"
	"poker/daemon/api/rpc"
	"poker/daemon/api/rpc_util"
)

func AddRpcRoutes(e *gin.Engine) {
	rpc_util.AddFunctions(e, &rpc.ServerRpc{})
	rpc_util.AddFunctions(e, &rpc.ContainerRpc{})
}
