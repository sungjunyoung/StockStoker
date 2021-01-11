package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: `Initialize your local configurations`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Init")
	},
}
