package cmd

import (
	"github.com/spf13/cobra"
)

var (
	deleteServiceCmd = &cobra.Command{
		Use:   SERVICE + " [SERVICE]",
		Short: "Delete an existing service",
		Long:  "Submits request to MariaDB SkySQL to delete an existing service. " + HINT_SVC_ID,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			svcid := args[0]

			res, err := client.DeleteService(cmd.Context(), svcid)

			checkAndPrint(res, err, SERVICES)
		},
	}
)

func init() {
	deleteCmd.AddCommand(deleteServiceCmd)
}
