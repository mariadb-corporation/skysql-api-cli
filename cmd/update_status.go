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
		Use:   fmt.Sprintf("%s [%s] [Start|Stop]", STATUS, strings.ToUpper(DATABASE)),
		Short: fmt.Sprintf("Update %s for %s", STATUS, DATABASE),
		Long:  fmt.Sprintf("Updates %s for %s belonging user in MariaDB SkySQL.", STATUS, DATABASE),
		Args:  cobra.ExactArgs(2),
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlag(NAME, cmd.Flags().Lookup(NAME))
			viper.BindPFlag(CONFIG_JSON, cmd.Flags().Lookup(CONFIG_JSON))
		},
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
}
