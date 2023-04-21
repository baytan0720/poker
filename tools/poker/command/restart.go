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

// RestartCmd represents the restart command
var RestartCmd = &cobra.Command{
	Use:    "restart CONTAINER [CONTAINER...]",
	Short:  "Restart one or more containers",
	Args:   cobra.MinimumNArgs(1),
	Run:    restart,
	PreRun: util.Connect,
}

func restart(_ *cobra.Command, args []string) {
	r, err := util.Client.RestartContainer(context.Background(), &service.RestartContainersReq{ContainerIdsOrNames: args})
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
