package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var marketCmd = &cobra.Command{
	Use:   "market",
	Short: `Manage your markets. "nasdaq" and "kospi" supported`,
}

func init() {
	downloadCmd := &cobra.Command{
		Use:   "download",
		Short: "Download current market stocks",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Download")
		},
	}
	marketCmd.AddCommand(downloadCmd)
}
