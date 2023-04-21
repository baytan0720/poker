/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package command

import (
	"context"
	"errors"
	"net"
	"poker/pkg/alert"
	"poker/pkg/pty"
	"poker/pkg/service"
	util2 "poker/tools/poker/util"
	"strings"

	"github.com/spf13/cobra"
)

// RunCmd represents the run command
var RunCmd = &cobra.Command{
	Use:    "run IMAGE [COMMAND] [ARG...]",
	Short:  "Run a command in a new container",
	Args:   cobra.MinimumNArgs(1),
	Run:    run,
	PreRun: util2.Connect,
}

func init() {
	RunCmd.Flags().StringP("name", "n", "", "Assign a name to the container")
	RunCmd.Flags().BoolP("interactive", "i", false, "Keep STDIN open even if not attached")
	RunCmd.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY")
	RunCmd.Flags().BoolP("detach", "d", false, "Run container in background and print container ID")
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
	r, err := util2.Client.CreateContainer(context.Background(), &service.CreateContainerReq{
		Image:   args[0],
		Name:    name,
		Command: command,
	})
	util2.CheckErr(r.Answer, err)

	containerId := r.Answer.ContainerIdOrName

	if detach {
		r, err := util2.Client.StartContainer(context.Background(), &service.StartContainersReq{
			ContainerIdsOrNames: []string{containerId},
		})
		util2.CheckErr(r.Answers[0], err)
		alert.Println(containerId)
		return
	}

	r2, err := util2.Client.RunContainer(context.Background(), &service.RunContainerReq{ContainerId: containerId})
	util2.CheckErr(r.Answer, err)

	conn, err := net.Dial("tcp", ":"+r2.PtyPort)
	if err != nil {
		alert.Error(err)
	}
	pty.NewPty(conn, interactive)
}
