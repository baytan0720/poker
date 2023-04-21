/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package command

import (
	"context"
	"github.com/spf13/cobra"
	"poker/pkg/alert"
	"poker/pkg/service"
	util2 "poker/tools/poker/util"
)

// LogsCmd represents the logs command
var LogsCmd = &cobra.Command{
	Use:    "logs CONTAINER",
	Short:  "Fetch the logs of a container",
	Args:   cobra.ExactArgs(1),
	Run:    logs,
	PreRun: util2.Connect,
}

func logs(_ *cobra.Command, args []string) {
	r, err := util2.Client.LogsContainer(context.Background(), &service.LogsContainerReq{ContainerIdOrName: args[0]})
	util2.CheckErr(r.Answer, err)
	alert.Println(string(r.Logs))
}
