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
		Use:   fmt.Sprintf("%s [%s] [%s]", STATUS, strings.ToUpper(DATABASE), strings.ToUpper(STATUS)),
		Short: fmt.Sprintf("Update %s for %s", STATUS, DATABASE),
		Long:  fmt.Sprintf("Updates %s for %s belonging user in MariaDB SkySQL.", STATUS, DATABASE),
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			databaseId := args[0]
			status := args[1]

			reqBody := skysql.UpdateStatusJSONRequestBody{
				Action: status,
			}

			res, err := client.UpdateStatus(cmd.Context(), databaseId, reqBody)

			checkAndPrint(res, err, STATUS)
		},
	}
)

func init() {
	updateCmd.AddCommand(updateStatusCmd)

	updateStatusCmd.Flags().StringP(NAME, "n", "", fmt.Sprintf("Name used to identify the %s", STATUS))
	updateStatusCmd.Flags().StringP(CONFIG_JSON, "j", "", fmt.Sprintf("JSON object containing %s", STATUS))

	viper.BindPFlag(NAME, updateStatusCmd.Flags().Lookup(NAME))
	viper.BindPFlag(CONFIG_JSON, updateStatusCmd.Flags().Lookup(CONFIG_JSON))
}
