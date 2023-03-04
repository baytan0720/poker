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

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop CONTAINER [CONTAINER...]",
	Short: "Stop one or more running containers",
	Run:   stop,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			alert.Error(errors.New("enter at least one container"))
		}
		return nil
	},
}

func stop(_ *cobra.Command, args []string) {
	r, err := client.StopContainer(context.Background(), &service.StopContainersReq{ContainerIds: args})
	if err != nil {
		alert.Error(err)
	}
	for _, info := range r.StartNStopContainerInfo {
		if info.Status == 0 {
			alert.Print(info.ContainerId + " ")
		}
	}
	for _, info := range r.StartNStopContainerInfo {
		if info.Status != 0 {
			alert.Warn(info.ContainerId + ": " + info.Msg)
		}
	}
}
