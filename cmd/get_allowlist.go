package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
)

var (
	getAllowlistCmd = &cobra.Command{
		Use:   fmt.Sprintf("%s [%s]", ALLOWLIST, strings.ToUpper(DATABASE)),
		Short: fmt.Sprintf("Retrieve list of allowed IPs for a %s", DATABASE),
		Long:  fmt.Sprintf("Queries for list of allowed IP addresses for a specific %s", DATABASE),
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				crash(fmt.Sprintf("Missing required %s ID", DATABASE))
			}

			dbid := args[0]
			res, err := client.ListAllowedAddresses(cmd.Context(), dbid, &skysql.ListAllowedAddressesParams{})
			checkErr(err, fmt.Sprintf("unable to retrieve %s from SkySQL", ALLOWLIST))
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			checkErr(err, "unable to read response from SkySQL")

			if res.StatusCode != http.StatusOK {
				crash(fmt.Sprintf("unable to retrieve %s from SkySQL: Status: %v, Body: %v", ALLOWLIST, res.StatusCode, string(body)))
			}

			fmt.Println(string(body))
		},
	}
)

func init() {
	getCmd.AddCommand(getAllowlistCmd)
}
