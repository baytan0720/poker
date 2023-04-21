/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package command

import (
	"poker/internal/daemon"
	"poker/pkg/alert"

	"github.com/spf13/cobra"
)

const PORT = "10720"

// StartCmd represents the start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the daemon",
	Run:   start,
}

func start(*cobra.Command, []string) {
	if err := daemon.Server(PORT); err != nil {
		alert.Error(err)
	}
}
