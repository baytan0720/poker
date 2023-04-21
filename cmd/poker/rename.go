/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"errors"
	"poker/pkg/alert"
	"poker/pkg/service"

	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:    "rename CONTAINER NEW_NAME",
	Short:  "Rename a container",
	Args:   cobra.ExactArgs(2),
	Run:    rename,
	PreRun: Connect,
}

func rename(_ *cobra.Command, args []string) {
	if len(args[1]) > 16 || args[1] == "" {
		alert.Error(errors.New("name is too long, the max length is 16"))
	}

	r, err := client.RenameContainer(context.Background(), &service.RenameContainerReq{
		ContainerIdOrName: args[0],
		NewName:           args[1],
	})
	checkErr(r.Answer, err)
	alert.Println("Success.")
}
