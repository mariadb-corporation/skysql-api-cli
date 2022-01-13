package cmd

import (
	"github.com/spf13/cobra"
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
}
