package cmd

import (
	"fmt"
	"net/http"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getConfigurationCmd = &cobra.Command{
		Use:     fmt.Sprintf("%s [CONFIGURATION NUMBER]", CONFIGURATIONS),
		Aliases: []string{CONFIGURATION},
		Short:   fmt.Sprintf("Retrieve stored %s configurations", DATABASE),
		Long:    fmt.Sprintf("Retrieves one or more custom %s configurations owned by the user", DATABASE),
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			limit := viper.GetInt(LIMIT)

			var res *http.Response
			var err error
			if len(args) == 1 {
				configNumber := args[0]
				res, err = client.ReadConfiguration(cmd.Context(), configNumber)
			} else {
				res, err = client.ListConfigurations(cmd.Context(), &skysql.ListConfigurationsParams{
					Limit: &limit,
				})
			}

			checkAndPrint(res, err, CONFIGURATIONS)
		},
	}
)

func init() {
	getCmd.AddCommand(getConfigurationCmd)

	getConfigurationCmd.Flags().IntP("limit", "l", DEFAULT_GET_LIMIT, "Maximum number of configurations to retrieve")
	viper.BindPFlag(LIMIT, createDatabaseCmd.Flags().Lookup(LIMIT))
}
