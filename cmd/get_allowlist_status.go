package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	getAllowlistStatusCmd = &cobra.Command{
		Use:   fmt.Sprintf("%s [%s]", ALLOWLIST_STATUS, strings.ToUpper(SERVICE)),
		Short: fmt.Sprintf("Retrieve status of allowlist for a %s", SERVICE),
		Long:  fmt.Sprintf("Queries for the provisioning status of an allowlist for a specific %s", SERVICE),
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				crash(fmt.Sprintf("Missing required %s ID", SERVICE))
			}

			svcid := args[0]
			res, err := client.ReadAllowlistStatus(cmd.Context(), svcid)
			checkAndPrint(res, err, ALLOWLIST_STATUS)
		},
	}
)

func init() {
	getCmd.AddCommand(getAllowlistStatusCmd)
}
