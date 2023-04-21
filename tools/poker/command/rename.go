/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package command

import (
	"context"
	"errors"
	"poker/pkg/alert"
	"poker/pkg/service"
	util2 "poker/tools/poker/util"

	"github.com/spf13/cobra"
)

// RenameCmd represents the rename command
var RenameCmd = &cobra.Command{
	Use:    "rename CONTAINER NEW_NAME",
	Short:  "Rename a container",
	Args:   cobra.ExactArgs(2),
	Run:    rename,
	PreRun: util2.Connect,
}

func rename(_ *cobra.Command, args []string) {
	if len(args[1]) > 16 || args[1] == "" {
		alert.Error(errors.New("name is too long, the max length is 16"))
	}

	r, err := util2.Client.RenameContainer(context.Background(), &service.RenameContainerReq{
		ContainerIdOrName: args[0],
		NewName:           args[1],
	})
	util2.CheckErr(r.Answer, err)
	alert.Println("Success.")
}
