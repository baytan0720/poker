/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run IMAGE [COMMAND] [ARG...]",
	Short: "Run a command in a new container",
	Run:   run,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println("run called")
}
