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

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:    "stop CONTAINER [CONTAINER...]",
	Short:  "Stop one or more running containers",
	Args:   cobra.MinimumNArgs(1),
	Run:    stop,
	PreRun: Connect,
}

func stop(_ *cobra.Command, args []string) {
	r, err := client.StopContainer(context.Background(), &service.StopContainersReq{ContainerIdsOrNames: args})
	if err != nil {
		alert.Error(err)
	}
	for _, answer := range r.Answers {
		if answer.Status == 0 {
			alert.Print(answer.ContainerIdOrName + " ")
		}
	}
	for _, answer := range r.Answers {
		if answer.Status != 0 {
			alert.Warn(answer.ContainerIdOrName + ": " + answer.Msg)
		}
	}
}
