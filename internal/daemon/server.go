package daemon

import (
	"google.golang.org/grpc"
	"net"
	"poker/internal/service"
)

const MAX_SIZE = 128 * 1024 * 1024

type Daemon struct {
	service.UnimplementedDaemonServer
}

func Server(port string) error {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	s := grpc.NewServer(grpc.MaxRecvMsgSize(MAX_SIZE))
	service.RegisterDaemonServer(s, &Daemon{})

	if err := s.Serve(l); err != nil {
		return err
	}
	return nil
}
