package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

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

			checkErr(err, "unable to delete database")
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			checkErr(err, "unable to read response from MariaDB SkySQL")

			if res.StatusCode != http.StatusOK {
				crash(fmt.Sprintf("unable to delete database: Status: %v, Body: %v", res.StatusCode, string(body)))
			}

			fmt.Println(string(body))
		},
	}
)

func init() {
	deleteCmd.AddCommand(deleteDatabaseCmd)
}
