/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"errors"
	"poker/alert"
	"poker/internal/service"
	"strings"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run IMAGE [COMMAND] [ARG...]",
	Short: "Run a command in a new container",
	Run:   run,
}

func init() {
	runCmd.Flags().StringP("name", "n", "", "Assign a name to the container")
}

func run(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	command := strings.Join(args[1:], " ")
	r, err := client.CreateContainer(context.Background(), &service.CreateContainerReq{
		Image:   args[0],
		Name:    name,
		Command: command,
	})
	checkErr(int32(r.Status), r.Msg, err)

	containerId := r.ContainerId
	alert.Println(containerId)

	r2, err := client.StartContainer(context.Background(), &service.StartContainersReq{
		ContainerIds: []string{containerId},
	})
	checkErr(int32(r2.StartNStopContainerInfo[0].Status), r2.StartNStopContainerInfo[0].Msg, err)
}

func checkErr(status int32, msg string, err error) {
	if err != nil {
		alert.Error(err)
	}
	if status != 0 {
		alert.Error(errors.New(msg))
	}
}
