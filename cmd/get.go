package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getValidArgs  = []string{SERVICES, QUOTAS}
	getArgAliases = []string{SERVICE, QUOTA}
	getCmd        = &cobra.Command{
		Use:        GET,
		Short:      "Query MariaDB SkySQL for resource information",
		Long:       `Commands which retrieve data about resources deployed into MariaDB SkySQL`,
		ValidArgs:  getValidArgs,
		ArgAliases: getArgAliases,
	}
)

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.PersistentFlags().IntP("limit", "l", DEFAULT_GET_LIMIT, "Number of records to return")
	viper.BindPFlag("limit", getCmd.PersistentFlags().Lookup("limit"))
}
