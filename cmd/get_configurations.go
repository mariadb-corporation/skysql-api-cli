package cmd

import (
	"fmt"
	"net/http"
	"strings"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getConfigurationCmd = &cobra.Command{
		Use:     fmt.Sprintf("%s [%s]", CONFIGURATIONS, strings.ToUpper(CONFIGURATION)),
		Aliases: []string{CONFIGURATION},
		Short:   fmt.Sprintf("Retrieve stored %s %s", SERVICE, CONFIGURATIONS),
		Long:    fmt.Sprintf("Retrieves one or more custom %s %s owned by the user", SERVICE, CONFIGURATIONS),
		Args:    cobra.MaximumNArgs(1),
		PreRun: func(cmd *cobra.Command, arg []string) {
			viper.BindPFlag(LIMIT, cmd.PersistentFlags().Lookup(LIMIT))
		},
		Run: func(cmd *cobra.Command, args []string) {
			var res *http.Response
			var err error
			if len(args) == 1 {
				configNumber := args[0]
				res, err = client.ReadConfiguration(cmd.Context(), configNumber)
			} else {
				limit := viper.GetInt(LIMIT)
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

	getConfigurationCmd.PersistentFlags().IntP(LIMIT, LIMIT_SHORTHAND, DEFAULT_GET_LIMIT, "Number of records to return. Can be used for paginating results in conjuntion with offset.")
}
