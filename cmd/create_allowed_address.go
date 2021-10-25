package cmd

import (
	"fmt"
	"strings"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	createAllowedAddressCmd = &cobra.Command{
		Use:   fmt.Sprintf("%s [%s] [%s]", ALLOWED_ADDRESS, strings.ToUpper(SERVICE), strings.ToUpper(IP_ADDRESS)),
		Short: fmt.Sprintf("Add a new %s to a %s", strings.Replace(ALLOWED_ADDRESS, "-", " ", -1), SERVICE),
		Long:  fmt.Sprintf("Adds a new %s for %s in MariaDB SkySQL.", strings.Replace(ALLOWED_ADDRESS, "-", " ", -1), SERVICE),
		Args:  cobra.ExactArgs(2),
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlag(COMMENT, cmd.Flags().Lookup(COMMENT))
		},
		Run: func(cmd *cobra.Command, args []string) {
			serviceID := string(args[0])
			ip_address := string(args[1])
			comment := viper.GetString(COMMENT)

			reqBody := skysql.AddAllowedAddressJSONRequestBody{
				Comment:   &comment,
				IpAddress: ip_address,
			}
			res, err := client.AddAllowedAddress(cmd.Context(), serviceID, reqBody)
			checkAndPrint(res, err, ALLOWED_ADDRESS)
		},
	}
)

func init() {
	createCmd.AddCommand(createAllowedAddressCmd)

	createAllowedAddressCmd.Flags().String(COMMENT, "", fmt.Sprintf("Additional %s to help identify address", COMMENT))
}
