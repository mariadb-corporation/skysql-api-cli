package cmd

import (
	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	dbUpdateName      string
	updateDatabaseCmd = &cobra.Command{
		Use:   DATABASE + " [DATABASE]",
		Short: "Update an existing database",
		Long:  "Submits request to MariaDB SkySQL to update an existing database. " + HINT_DB_ID,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dbid := args[0]
			reqBody := skysql.UpdateDatabaseJSONRequestBody{
				Name: viper.GetString("name"),
			}

			res, err := client.UpdateDatabase(cmd.Context(), dbid, reqBody)

			checkAndPrint(res, err, DATABASES)
		},
	}
)

func init() {
	updateCmd.AddCommand(updateDatabaseCmd)

	updateDatabaseCmd.Flags().StringVarP(&dbUpdateName, "name", "n", "", "Name used to identify the database")
	viper.BindPFlag("name", updateDatabaseCmd.Flags().Lookup("name"))
}
