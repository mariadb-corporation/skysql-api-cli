package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	getQuotaCmd = &cobra.Command{
		Use:     "quota",
		Aliases: []string{"quotas"},
		Short:   "Retrieve quota information",
		Long:    `Queries for quota limits, and progress towards those quotas.`,
		Args:    cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			res, err := client.ReadQuotas(cmd.Context())
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
	getCmd.AddCommand(getQuotaCmd)
}
