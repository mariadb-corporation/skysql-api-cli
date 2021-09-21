package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	limit         int
	getValidArgs  = []string{DATABASES, QUOTAS}
	getArgAliases = []string{DATABASE, QUOTA}
	getCmd        = &cobra.Command{
		Use:        GET,
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
