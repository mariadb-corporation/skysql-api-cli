package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var (
	getCredentialsCmd = &cobra.Command{
		Use:   fmt.Sprintf("%s [%s]", CREDENTIALS, strings.ToUpper(SERVICE)),
		Short: "Retrieve default service credentials",
		Long:  "Queries for default credentials configured for a service. " + HINT_SVC_ID,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var res *http.Response
			var err error
			svcid := args[0]
			res, err = client.RetrieveDefaultCredentials(cmd.Context(), svcid)

			checkAndPrint(res, err, CREDENTIALS)
		},
	}
)

func init() {
	getCmd.AddCommand(getCredentialsCmd)
}
