package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var (
	getStatusCmd = &cobra.Command{
		Use:   fmt.Sprintf("%s [%s]", STATUS, strings.ToUpper(DATABASE)),
		Short: fmt.Sprintf("Get current %s for a %s", STATUS, DATABASE),
		Long:  fmt.Sprintf("Get the current %s of a %s in MariaDB SkySQL", STATUS, DATABASE),
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			databaseId := args[0]

			var res *http.Response
			var err error
			res, err = client.ReadStatus(cmd.Context(), databaseId)

			checkAndPrint(res, err, STATUS)
		},
	}
)

func init() {
	getCmd.AddCommand(getStatusCmd)
}
