package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
)

var (
	getDatabaseCmd = &cobra.Command{
		Use:     "databases [DATABASE]",
		Aliases: []string{"database"},
		Short:   "Retrieve database information",
		Long:    `Queries for information about deployed database resources in SkySQL. Specify a single database using the database id (e.g. db00000000)`,
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var res *http.Response
			var err error
			if len(args) == 1 {
				dbid := args[0]
				res, err = client.ReadDatabase(cmd.Context(), dbid)
			} else {
				res, err = client.ListDatabases(cmd.Context(), &skysql.ListDatabasesParams{
					Limit: &limit,
				})
			}

			checkErr(err, "unable to retrieve databases from SkySQL")
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			checkErr(err, "unable to read response from SkySQL")

			if res.StatusCode != http.StatusOK {
				crash(fmt.Sprintf("unable to retrieve databases from SkySQL: Status: %v, Body: %v", res.StatusCode, string(body)))
			}

			fmt.Println(string(body))
		},
	}
)

func init() {
	getCmd.AddCommand(getDatabaseCmd)
}
