/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package command

import (
	"context"
	"poker/pkg/alert"
	"poker/pkg/service"
	"poker/tools/poker/util"

	"github.com/spf13/cobra"
)

// StartCmd represents the start command
var StartCmd = &cobra.Command{
	Use:    "start CONTAINER [CONTAINER...]",
	Short:  "Start one or more exited containers",
	Args:   cobra.MinimumNArgs(1),
	Run:    start,
	PreRun: util.Connect,
}

func start(_ *cobra.Command, args []string) {
	r, err := util.Client.StartContainer(context.Background(), &service.StartContainersReq{ContainerIdsOrNames: args})
	if err != nil {
		alert.Error(err)
	}
	for _, answer := range r.Answers {
		if answer.Status == 0 {
			alert.Println(answer.ContainerIdOrName + " ")
		}
	}
	for _, answer := range r.Answers {
		if answer.Status != 0 {
			alert.Warn(answer.ContainerIdOrName + ": " + answer.Msg)
		}
	}
}
