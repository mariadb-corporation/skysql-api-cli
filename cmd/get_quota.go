package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

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
			checkErr(err, fmt.Sprintf("unable to retrieve %s from SkySQL", QUOTAS))
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			checkErr(err, "unable to read response from SkySQL")

			if res.StatusCode != http.StatusOK {
				crash(fmt.Sprintf("unable to retrieve %s from SkySQL: Status: %v, Body: %v", QUOTAS, res.StatusCode, string(body)))
			}

			fmt.Println(string(body))
		},
	}
)

func init() {
	getCmd.AddCommand(getQuotaCmd)
}