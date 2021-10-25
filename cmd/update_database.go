package cmd

import (
	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	updateServiceCmd = &cobra.Command{
		Use:   SERVICE + " [SERVICE]",
		Short: "Update an existing service",
		Long:  "Submits request to MariaDB SkySQL to update an existing service. " + HINT_SVC_ID,
		Args:  cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlag(NAME, cmd.Flags().Lookup(NAME))
		},
		Run: func(cmd *cobra.Command, args []string) {
			svcid := args[0]
			reqBody := skysql.UpdateServiceJSONRequestBody{
				Name: viper.GetString(NAME),
			}

			res, err := client.UpdateService(cmd.Context(), svcid, reqBody)

			checkAndPrint(res, err, SERVICES)
		},
	}
)

func init() {
	updateCmd.AddCommand(updateServiceCmd)

	updateServiceCmd.Flags().StringP(NAME, "n", DEFAULT_UPDATE_SERVICE_NAME, "Name used to identify the service")
}
