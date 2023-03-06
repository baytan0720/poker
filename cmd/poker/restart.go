/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"errors"
	"poker/alert"
	"poker/internal/service"

	"github.com/spf13/cobra"
)

// restartCmd represents the restart command
var restartCmd = &cobra.Command{
	Use:   "restart CONTAINER [CONTAINER...]",
	Short: "Restart one or more containers",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			alert.Error(errors.New("enter at least one container"))
		}
		return nil
	},
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
