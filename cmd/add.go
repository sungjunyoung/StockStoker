package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Find stock in database and add. registered stocks will be shown in root command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add favorate stock")
	},
}
