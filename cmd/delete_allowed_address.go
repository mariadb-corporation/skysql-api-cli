package cmd

import (
	"fmt"
	"strings"

	"github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
)

var (
	deleteAllowedAddressCmd = &cobra.Command{
		Use:   fmt.Sprintf("%s [%s] [%s]", ALLOWED_ADDRESS, strings.ToUpper(SERVICE), strings.ToUpper(ALLOWED_ADDRESS)),
		Short: fmt.Sprintf("Delete an %s from a %s", strings.Replace(ALLOWED_ADDRESS, "-", " ", -1), SERVICE),
		Long:  fmt.Sprintf("Deletes an %s for user on %s in MariaDB SkySQL.", strings.Replace(ALLOWED_ADDRESS, "-", " ", -1), SERVICE),
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			serviceID := string(args[0])
			ipAddress := string(args[1])
			reqQuery := skysql.RemoveAllowedAddressParams{
				Address: &ipAddress,
			}

			res, err := client.RemoveAllowedAddress(cmd.Context(), serviceID, &reqQuery)
			checkAndPrint(res, err, ALLOWED_ADDRESS)
		},
	}
)

func init() {
	deleteCmd.AddCommand(deleteAllowedAddressCmd)
}
