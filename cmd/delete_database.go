package cmd

import (
	"github.com/spf13/cobra"
)

var (
	deleteDatabaseCmd = &cobra.Command{
		Use:   DATABASE + " [DATABASE]",
		Short: "Delete an existing database",
		Long:  "Submits request to MariaDB SkySQL to delete an existing database. " + HINT_DB_ID,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dbid := args[0]

			res, err := client.DeleteDatabase(cmd.Context(), dbid)

			checkAndPrint(res, err, DATABASES)
		},
	}
)

func init() {
	deleteCmd.AddCommand(deleteDatabaseCmd)
}
