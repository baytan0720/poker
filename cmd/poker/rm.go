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

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:    "rm CONTAINER [CONTAINER...]",
	Short:  "Remove one or more containers",
	Args:   cobra.MinimumNArgs(1),
	Run:    rm,
	PreRun: Connect,
}

func rm(_ *cobra.Command, args []string) {
	r, err := client.RemoveContainers(context.Background(), &service.RemoveContainersReq{ContainerIdsOrNames: args})
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
