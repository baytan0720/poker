/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"errors"
	"net"
	"poker/alert"
	"poker/internal/pty"
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
	runCmd.Flags().BoolP("interactive", "i", false, "Keep STDIN open even if not attached")
	runCmd.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY")
	runCmd.Flags().BoolP("detach", "d", false, "Run container in background and print container ID")
}

func run(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	interactive, _ := cmd.Flags().GetBool("interactive")
	tty, _ := cmd.Flags().GetBool("tty")
	detach, _ := cmd.Flags().GetBool("detach")
	if tty && detach {
		alert.Error(errors.New("tty and detach can only choose one"))
	}

	command := strings.Join(args[1:], " ")
	r, err := client.CreateContainer(context.Background(), &service.CreateContainerReq{
		Image:   args[0],
		Name:    name,
		Command: command,
	})
	checkErr(int32(r.Status), r.Msg, err)

	containerId := r.ContainerId

	if detach || !tty {
		r, err := client.StartContainer(context.Background(), &service.StartContainersReq{
			ContainerIds: []string{containerId},
		})
		checkErr(int32(r.StartNStopContainerInfo[0].Status), r.StartNStopContainerInfo[0].Msg, err)
		alert.Println(containerId)
		return
	}

	r2, err := client.RunContainer(context.Background(), &service.RunContainerReq{ContainerId: containerId})
	checkErr(int32(r2.Status), r2.Msg, err)

	conn, err := net.Dial("tcp", ":"+r2.PtyPort)
	if err != nil {
		alert.Error(err)
	}
	pty.NewPty(conn, interactive)
}

func checkErr(status int32, msg string, err error) {
	if err != nil {
		alert.Error(err)
	}
	if status != 0 {
		alert.Error(errors.New(msg))
	}
}
