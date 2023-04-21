/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"errors"
	"net"
	"poker/internal/pty"
	"poker/pkg/alert"
	"poker/pkg/service"
	"strings"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:    "run IMAGE [COMMAND] [ARG...]",
	Short:  "Run a command in a new container",
	Args:   cobra.MinimumNArgs(1),
	Run:    run,
	PreRun: Connect,
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
	if len(name) > 16 {
		alert.Error(errors.New("name is too long, the max length is 16"))
	}
	if tty && detach {
		alert.Error(errors.New("tty and detach can only choose one"))
	}
	if args[0] != "base" {
		alert.Error(errors.New("image not found, try base"))
	}

	command := strings.Join(args[1:], " ")
	r, err := client.CreateContainer(context.Background(), &service.CreateContainerReq{
		Image:   args[0],
		Name:    name,
		Command: command,
	})
	checkErr(r.Answer, err)

	containerId := r.Answer.ContainerIdOrName

	if detach {
		r, err := client.StartContainer(context.Background(), &service.StartContainersReq{
			ContainerIdsOrNames: []string{containerId},
		})
		checkErr(r.Answers[0], err)
		alert.Println(containerId)
		return
	}

	r2, err := client.RunContainer(context.Background(), &service.RunContainerReq{ContainerId: containerId})
	checkErr(r.Answer, err)

	conn, err := net.Dial("tcp", ":"+r2.PtyPort)
	if err != nil {
		alert.Error(err)
	}
	pty.NewPty(conn, interactive)
}
