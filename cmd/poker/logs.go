/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"errors"
	"github.com/spf13/cobra"
	"poker/alert"
	"poker/internal/service"
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs CONTAINER",
	Short: "Fetch the logs of a container",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			alert.Error(errors.New("enter a container"))
		}
		return nil
	},
	Run:    logs,
	PreRun: Connect,
}

func logs(_ *cobra.Command, args []string) {
	r, err := client.LogsContainer(context.Background(), &service.LogsContainerReq{ContainerIdOrName: args[0]})
	checkErr(int32(r.Status), r.Msg, err)
	alert.Println(string(r.Logs))
}
