package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
)

var (
	getConfigurationCmd = &cobra.Command{
		Use:     "configurations [CONFIGURATION NUMBER]",
		Aliases: []string{"configuration"},
		Short:   "Retrieve stored configurations",
		Long:    `Retrieves one or more custom configurations owned by the user`,
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var res *http.Response
			var err error
			if len(args) == 1 {
				configNumber := args[0]
				res, err = client.ReadConfiguration(cmd.Context(), configNumber)
			} else {
				res, err = client.ListConfigurations(cmd.Context(), &skysql.ListConfigurationsParams{
					Limit: &limit,
				})
			}

			checkErr(err, "unable to retrieve configurations from SkySQL")
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			checkErr(err, "unable to read response from SkySQL")

			if res.StatusCode != http.StatusOK {
				crash(fmt.Sprintf("unable to retrieve configurations from SkySQL: Status: %v, Body: %v", res.StatusCode, string(body)))
			}

			fmt.Println(string(body))
		},
	}
)

func init() {
	getCmd.AddCommand(getConfigurationCmd)
}
