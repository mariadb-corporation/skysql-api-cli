package cmd

import (
	"fmt"
	"strings"

	"github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
)

var (
	deleteAllowedAddressCmd = &cobra.Command{
		Use:   fmt.Sprintf("%s [%s] [%s]", ALLOWED_ADDRESS, strings.ToUpper(DATABASE), strings.ToUpper(ALLOWED_ADDRESS)),
		Short: fmt.Sprintf("Delete an %s from a %s", strings.Replace(ALLOWED_ADDRESS, "-", " ", -1), DATABASE),
		Long:  fmt.Sprintf("Deletes an %s for user on %s in MariaDB SkySQL.", strings.Replace(ALLOWED_ADDRESS, "-", " ", -1), DATABASE),
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			databaseId := string(args[0])
			ipAddress := string(args[1])
			reqQuery := skysql.RemoveAllowedAddressParams{
				Address: &ipAddress,
			}

			res, err := client.RemoveAllowedAddress(cmd.Context(), databaseId, &reqQuery)
			checkAndPrint(res, err, ALLOWED_ADDRESS)
		},
	}
)

func init() {
	deleteCmd.AddCommand(deleteAllowedAddressCmd)
}
