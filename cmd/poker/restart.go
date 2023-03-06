/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"poker/alert"
	"poker/internal/service"

	"github.com/spf13/cobra"
)

// restartCmd represents the restart command
var restartCmd = &cobra.Command{
	Use:    "restart CONTAINER [CONTAINER...]",
	Short:  "Restart one or more containers",
	Args:   cobra.MinimumNArgs(1),
	Run:    restart,
	PreRun: Connect,
}

func restart(_ *cobra.Command, args []string) {
	r, err := client.RestartContainer(context.Background(), &service.RestartContainersReq{ContainerIdsOrNames: args})
	if err != nil {
		alert.Error(err)
	}
	for _, info := range r.StartNStopContainerInfo {
		if info.Status == 0 {
			alert.Print(info.ContainerIdOrName + " ")
		}
	}
	for _, info := range r.StartNStopContainerInfo {
		if info.Status != 0 {
			alert.Warn(info.ContainerIdOrName + ": " + info.Msg)
		}
	}
}
