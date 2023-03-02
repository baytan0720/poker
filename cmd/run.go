/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"poker/alert"
	"poker/container"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use: "run IMAGE [COMMAND] [ARG...]",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			alert.Fatal("need image")
		}
		return nil
	},
	Short: "Run a command with a new container",
	Run:   run,
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolP("interactive", "i", false, "Keep STDIN open even if not attached")
	runCmd.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY")
	runCmd.Flags().BoolP("detach", "d", false, "Run container in background and print container ID")
	runCmd.Flags().String("name", "", "Assign a name to the container")
}

func run(cmd *cobra.Command, args []string) {
	if args[0] != "base" {
		alert.Error("no such image, use base")
		return
	}
	isInteractive, _ := cmd.Flags().GetBool("interactive")
	isTty, _ := cmd.Flags().GetBool("tty")
	isDetach, _ := cmd.Flags().GetBool("detach")
	container.CreateInitialcontainer(args[0], isInteractive, isTty, isDetach, args[1:]...)
}
