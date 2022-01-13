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

	getCmd.PersistentFlags().IntP("limit", "l", DEFAULT_GET_LIMIT, "Number of records to return. Can be used for paginating results in conjuntion with offset.")
	getCmd.PersistentFlags().IntP("offset", "o", DEFAULT_GET_OFFSET, "Number of records to skip when retrieved. Can be used for paginating results in conjunction with limit.")

	viper.BindPFlag("limit", getCmd.PersistentFlags().Lookup("limit"))
	viper.BindPFlag("offset", getCmd.PersistentFlags().Lookup("offset"))
}
