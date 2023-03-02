/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"poker/alert"
	"poker/container"

	"github.com/spf13/cobra"
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs CONTAINER",
	Short: "Fetch the logs of a container",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			alert.Fatal("need Container ID")
		}
		return nil
	},
	Run: logs,
}

func init() {
	rootCmd.AddCommand(logsCmd)
}

func logs(cmd *cobra.Command, args []string) {
	container.PrintContainerLogs(args[0])
}
