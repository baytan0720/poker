/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"poker/alert"
	"poker/internal/daemon"

	"github.com/spf13/cobra"
)

// StartCmd represents the start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the daemon",
	Run:   start,
}

func init() {
	rootCmd.AddCommand(StartCmd)
}

func start(*cobra.Command, []string) {
	if err := daemon.Server(PORT); err != nil {
		alert.Error(err)
	}
}
