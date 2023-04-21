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

// StopCmd represents the stop command
var StopCmd = &cobra.Command{
	Use:    "stop CONTAINER [CONTAINER...]",
	Short:  "Stop one or more running containers",
	Args:   cobra.MinimumNArgs(1),
	Run:    stop,
	PreRun: util.Connect,
}

func stop(_ *cobra.Command, args []string) {
	r, err := util.Client.StopContainer(context.Background(), &service.StopContainersReq{ContainerIdsOrNames: args})
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
