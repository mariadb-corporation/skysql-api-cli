package cmd

import (
	"github.com/spf13/cobra"
)

var (
	getQuotaCmd = &cobra.Command{
		Use:     QUOTAS,
		Aliases: []string{QUOTA},
		Short:   "Retrieve quota information",
		Long:    `Queries for quota limits, and progress towards those quotas.`,
		Args:    cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			res, err := client.ReadQuotas(cmd.Context())
			checkAndPrint(res, err, QUOTAS)
		},
	}
)

func init() {
	getCmd.AddCommand(getQuotaCmd)
}
