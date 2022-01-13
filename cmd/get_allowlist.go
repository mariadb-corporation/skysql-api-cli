package cmd

import (
	"fmt"
	"strings"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getAllowlistCmd = &cobra.Command{
		Use:   fmt.Sprintf("%s [%s]", ALLOWLIST, strings.ToUpper(SERVICE)),
		Short: fmt.Sprintf("Retrieve list of allowed IPs for a %s", SERVICE),
		Long:  fmt.Sprintf("Queries for list of allowed IP addresses for a specific %s", SERVICE),
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				crash(fmt.Sprintf("Missing required %s ID", SERVICE))
			}

			svcid := args[0]
			limit := viper.GetInt(LIMIT)
			offset := viper.GetInt(OFFSET)
			res, err := client.ListAllowedAddresses(cmd.Context(), svcid, &skysql.ListAllowedAddressesParams{Limit: &limit, Offset: &offset})
			checkAndPrint(res, err, ALLOWLIST)
		},
	}
)

func init() {
	getCmd.AddCommand(getAllowlistCmd)
}
