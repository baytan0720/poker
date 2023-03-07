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

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:    "start CONTAINER [CONTAINER...]",
	Short:  "Start one or more exited containers",
	Args:   cobra.MinimumNArgs(1),
	Run:    start,
	PreRun: Connect,
}

func start(_ *cobra.Command, args []string) {
	r, err := client.StartContainer(context.Background(), &service.StartContainersReq{ContainerIdsOrNames: args})
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
