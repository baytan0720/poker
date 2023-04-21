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

// RmCmd represents the rm command
var RmCmd = &cobra.Command{
	Use:    "rm CONTAINER [CONTAINER...]",
	Short:  "Remove one or more containers",
	Args:   cobra.MinimumNArgs(1),
	Run:    rm,
	PreRun: util.Connect,
}

func rm(_ *cobra.Command, args []string) {
	r, err := util.Client.RemoveContainers(context.Background(), &service.RemoveContainersReq{ContainerIdsOrNames: args})
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
