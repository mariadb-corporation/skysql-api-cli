package cmd

import (
	"net/http"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getDatabaseCmd = &cobra.Command{
		Use:     DATABASES + " [DATABASE]",
		Aliases: []string{DATABASE},
		Short:   "Retrieve database information",
		Long:    "Queries for information about deployed database resources in SkySQL. " + HINT_DB_ID,
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var res *http.Response
			var err error
			if len(args) == 1 {
				dbid := args[0]
				res, err = client.ReadDatabase(cmd.Context(), dbid)
			} else {
				limit := viper.GetInt("limit")
				res, err = client.ListDatabases(cmd.Context(), &skysql.ListDatabasesParams{
					Limit: &limit,
				})
			}
			checkAndPrint(res, err, DATABASES)
		},
	}
)

func init() {
	getCmd.AddCommand(getDatabaseCmd)
}
