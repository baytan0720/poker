package main

import (
	"context"
	"flag"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"poker/pkg/proto"
)

var socketPath string

func init() {
	flag.StringVar(&socketPath, "socket", "/var/run/poker.sock", "socket path")
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial("unix://"+socketPath, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc dial fail, error: %s\n", err)
		return
	}
	defer conn.Close()

	c := proto.NewPokerServiceClient(conn)
	resp, err := c.Version(context.Background(), &proto.VersionReq{})
	if err != nil {
		fmt.Printf("Cannot connect to the Poker daemon at unix://%s. Is the poker daemon running?\n", socketPath)
		return
	}

	fmt.Println("Poker:")
	fmt.Println("  Version:", resp.PokerVersion)
	fmt.Println("  GoVersion:", resp.GoVersion)
	fmt.Println("  GitRevision:", resp.GitRevision)
	fmt.Println("  BuildTime:", resp.BuildTime)
}
