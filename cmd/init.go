/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"poker/alert"
	"poker/container"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use: "init CONTAINERID [COMMAND] [ARG...]",
	Run: Init,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			alert.Fatal("need ContainerID")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func Init(_ *cobra.Command, args []string) {
	container.CreateRuncontainer(args[0], args[1:])
}
