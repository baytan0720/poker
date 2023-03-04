/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"poker/alert"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start CONTAINER [CONTAINER...]",
	Short: "Start one or more stopped containers",
	Run:   start,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			alert.Error(errors.New("enter at least one container"))
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func start(cmd *cobra.Command, args []string) {
	fmt.Println("start called")
}
