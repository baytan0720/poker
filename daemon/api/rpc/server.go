package rpc

import (
	gin "github.com/baytan0720/gin-grpc"
	"poker/daemon/api/rpc_util"
	"poker/pkg/config"
	"poker/pkg/proto"
)

type ServerRpc struct{}

func (r *ServerRpc) Functions() rpc_util.Functions {
	return rpc_util.Functions{
		{
			Name:    "Ping",
			Handler: r.Ping,
		},
		{
			Name:    "Version",
			Handler: r.Version,
		},
	}
}

func (r *ServerRpc) Ping(c *gin.Context) {}

func (r *ServerRpc) Version(c *gin.Context) {
	c.Response(&proto.VersionRes{
		PokerVersion: config.GetPokerVersion(),
		GitRevision:  config.GetGitRevision(),
		GoVersion:    config.GetGoVersion(),
		BuildTime:    config.GetBuildTime(),
	})
}
