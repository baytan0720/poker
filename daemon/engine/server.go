package engine

import (
	"net"
	"poker/daemon/api"
	"poker/daemon/middleware"

	gin "github.com/baytan0720/gin-grpc"
	"poker/pkg/proto"
)

type Server struct {
	proto.UnimplementedPokerServiceServer
}

func serve(l net.Listener) error {
	e := gin.Default(&Server{})
	proto.RegisterPokerServiceServer(e, e.Srv.(*Server))

	e.Use(middleware.ErrorHandle)
	api.AddRpcRoutes(e)

	return e.Serve(l)
}
