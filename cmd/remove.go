package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Find stock in favorite list and remove",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Remove favorate stock")
	},
}
