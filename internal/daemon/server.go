package daemon

import (
	"log"
	"net"
	"poker/internal/service"

	"google.golang.org/grpc"
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

	log.Println("daemon is running now, listen on :10720")
	if err := s.Serve(l); err != nil {
		return err
	}
	return nil
}
