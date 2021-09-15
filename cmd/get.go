package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	limit         int
	getValidArgs  = []string{"databases"}
	getArgAliases = []string{"database"}
	getCmd        = &cobra.Command{
		Use:        "get",
		Short:      "Query SkySQL for resource information",
		Long:       `Commands which retrieve data about resources deployed into MariaDB SkySQL`,
		ValidArgs:  getValidArgs,
		ArgAliases: getArgAliases,
	}
)

func init() {
	rootCmd.AddCommand(getCmd)

	getDatabaseCmd.PersistentFlags().IntVarP(&limit, "limit", "l", 10, "Number of records to return")
	viper.BindPFlag("limit", getDatabaseCmd.PersistentFlags().Lookup("limit"))
}
