package cmd

import (
	"fmt"
	"net/http"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getProviderCmd = &cobra.Command{
		Use:     PROVIDERS,
		Aliases: []string{PROVIDER},
		Short:   fmt.Sprintf("Retrieve list of cloud %s", PROVIDERS),
		Long:    fmt.Sprintf("Queries a list of cloud %s supported by MariaDB SkySQL", PROVIDERS),
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			limit := viper.GetInt(LIMIT)

			var res *http.Response
			var err error
			res, err = client.ReadProviders(cmd.Context(), &skysql.ReadProvidersParams{
				Limit: &limit,
			})

			checkAndPrint(res, err, PROVIDERS)
		},
	}
)

func init() {
	getCmd.AddCommand(getProviderCmd)
}
