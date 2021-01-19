package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/sungjunyoung/StockStoker/pkg/market"
	"strings"
)

var dataDir = ""
var all = false
var downloadCmd = &cobra.Command{
	Use: fmt.Sprintf("download [%s]", strings.Join(market.SupportedMarkets.Strings(), "|")),
	Short: fmt.Sprintf("Download current stocks database to local. %s supported for the argument",
		strings.Join(market.SupportedMarkets.Strings(), " and ")),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dataDir: " + dataDir)
		fmt.Println("args: " + strings.Join(args, ","))

		if len(args) == 0 && !all {
			fmt.Println("no market specified")
		}
	},
}

func init() {
	homeDir, err := homedir.Dir()
	if err != nil {
		dataDir = "/etc/.stockstoker/data/"
	} else {
		dataDir = fmt.Sprintf("%s/.stockstoker/data/", homeDir)
	}
	all = false

	downloadCmd.Flags().StringVarP(&dataDir, "data-dir", "d", dataDir, "Directory that database will be downloaded")
	downloadCmd.Flags().BoolVarP(&all, "all", "a", all, "Download all supported market's stock. this flag will overwrite argument")
}
