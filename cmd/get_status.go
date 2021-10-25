package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var (
	getStatusCmd = &cobra.Command{
		Use:   fmt.Sprintf("%s [%s]", STATUS, strings.ToUpper(SERVICE)),
		Short: fmt.Sprintf("Get current %s for a %s", STATUS, SERVICE),
		Long:  fmt.Sprintf("Get the current %s of a %s in MariaDB SkySQL", STATUS, SERVICE),
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			serviceID := args[0]

			var res *http.Response
			var err error
			res, err = client.ReadStatus(cmd.Context(), serviceID)

			checkAndPrint(res, err, STATUS)
		},
	}
)

func init() {
	getCmd.AddCommand(getStatusCmd)
}
