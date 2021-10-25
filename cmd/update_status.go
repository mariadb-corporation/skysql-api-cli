package cmd

import (
	"fmt"
	"strings"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	updateStatusCmd = &cobra.Command{
		Use:   fmt.Sprintf("%s [%s] [Start|Stop]", STATUS, strings.ToUpper(SERVICE)),
		Short: fmt.Sprintf("Update %s for %s", STATUS, SERVICE),
		Long:  fmt.Sprintf("Updates %s for %s belonging user in MariaDB SkySQL.", STATUS, SERVICE),
		Args:  cobra.ExactArgs(2),
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlag(NAME, cmd.Flags().Lookup(NAME))
			viper.BindPFlag(CONFIG_JSON, cmd.Flags().Lookup(CONFIG_JSON))
		},
		Run: func(cmd *cobra.Command, args []string) {
			serviceID := args[0]
			status := args[1]

			reqBody := skysql.UpdateStatusJSONRequestBody{
				Action: status,
			}

			res, err := client.UpdateStatus(cmd.Context(), serviceID, reqBody)

			checkAndPrint(res, err, STATUS)
		},
	}
)

func init() {
	updateCmd.AddCommand(updateStatusCmd)
}
