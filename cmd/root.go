package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stockstoker",
	Short: "Hey developers! Check your stock in command line",
}

func init() {
	rootCmd.AddCommand(marketCmd)
	rootCmd.AddCommand(initCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
