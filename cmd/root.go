package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "stockstoker",
	Short: "Developers! Check your stock price in your favorite terminal",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("NAVER 39,3000    \t+12%\nSAMSUNG 10,0000    \t+8%")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(removeCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
