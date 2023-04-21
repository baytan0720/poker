package poker

import (
	"os"

	"poker/tools/poker/command"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "poker",
	Short:   "A container creation and running tool",
	Long:    "Poker is a container technology as like docker",
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
	rootCmd.AddCommand(command.RunCmd, command.StartCmd, command.StopCmd, command.PsCmd, command.LogsCmd, command.RestartCmd, command.RenameCmd, command.RmCmd)
}
