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

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop CONTAINER [CONTAINER...]",
	Short: "Stop one or more running containers",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			alert.Error(errors.New("enter at least one container"))
		}
		return nil
	},
	Run:    stop,
	PreRun: Connect,
}

func stop(_ *cobra.Command, args []string) {
	r, err := client.StopContainer(context.Background(), &service.StopContainersReq{ContainerIdsOrNames: args})
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
