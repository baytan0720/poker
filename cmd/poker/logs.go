/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"github.com/spf13/cobra"
	"poker/alert"
	"poker/internal/service"
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:    "logs CONTAINER",
	Short:  "Fetch the logs of a container",
	Args:   cobra.ExactArgs(1),
	Run:    logs,
	PreRun: Connect,
}

func logs(_ *cobra.Command, args []string) {
	r, err := client.LogsContainer(context.Background(), &service.LogsContainerReq{ContainerIdOrName: args[0]})
	checkErr(int32(r.Status), r.Msg, err)
	alert.Println(string(r.Logs))
}
