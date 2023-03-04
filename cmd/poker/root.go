/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"os"
	"poker/alert"
	"poker/internal/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/spf13/cobra"
)

const (
	MAX_SIZE = 128 * 1024 * 1024
	PORT     = "10720"
)

var client service.DaemonClient

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "poker",
	Short:   "A container creation and running tool",
	Long:    "Poker is a container technology like docker",
	Version: "v0.1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(runCmd, startCmd, stopCmd)
	opt := grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(MAX_SIZE))
	conn, err := grpc.Dial("127.0.0.1:"+PORT, grpc.WithTransportCredentials(insecure.NewCredentials()), opt)
	if err != nil {
		alert.Error(err)
	}
	client = service.NewDaemonClient(conn)
	if res, err := client.Ping(context.Background(), &service.PingReq{}); err != nil || res.Status != 0 {
		alert.Error(err)
	}
}
