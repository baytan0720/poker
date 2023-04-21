package util

import (
	"context"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"poker/pkg/alert"
	service2 "poker/pkg/service"
)

const (
	MAX_SIZE = 128 * 1024 * 1024
	HOST     = "127.0.0.1:10720"
)

var Client service2.DaemonClient

func Connect(*cobra.Command, []string) {
	opt := grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(MAX_SIZE))
	conn, err := grpc.Dial(HOST, grpc.WithTransportCredentials(insecure.NewCredentials()), opt)
	if err != nil {
		alert.Error(err)
	}
	Client = service2.NewDaemonClient(conn)
	if res, err := Client.Ping(context.Background(), &service2.PingReq{}); err != nil || res.Status != 0 {
		alert.Error(err)
	}
}
