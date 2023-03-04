/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/spf13/cobra"
	"os"
)

const PORT = "10720"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "poker-daemon",
	Short:   "A daemon for create and manage containers",
	Long:    "Daemon is for Poker",
	Version: "v0.1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(StartCmd, StopCmd)
}
