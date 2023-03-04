/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the Poker version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("poker version v0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
