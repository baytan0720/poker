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

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start CONTAINER [CONTAINER...]",
	Short: "Start one or more exited containers",
	Run:   start,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			alert.Error(errors.New("enter at least one container"))
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func start(cmd *cobra.Command, args []string) {
	r, err := client.StartContainer(context.Background(), &service.StartContainersReq{ContainerIds: args})
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
