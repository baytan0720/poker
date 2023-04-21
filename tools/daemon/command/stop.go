/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// StopCmd represents the stop command
var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop if daemon is running",
	Run:   stop,
}

func stop(cmd *cobra.Command, args []string) {
	fmt.Println("stop called")
}
