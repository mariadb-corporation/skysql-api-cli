package cmd

import (
	"fmt"
	"net/http"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getServiceTypesCmd = &cobra.Command{
		Use:     SERVICE_TYPES,
		Aliases: []string{SERVICE_TYPE},
		Short:   fmt.Sprintf("Retrieve MariaDB SkySQL %s information", SERVICE_TYPE),
		Long:    fmt.Sprintf("Queries information for %s offerings from MariaDB SkySQL", SERVICE_TYPE),
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			limit := viper.GetInt(LIMIT)

			var res *http.Response
			var err error
			res, err = client.ReadServiceTypes(cmd.Context(), &skysql.ReadServiceTypesParams{
				Limit: &limit,
			})

			checkAndPrint(res, err, SERVICE_TYPES)
		},
	}
)

func init() {
	getCmd.AddCommand(getServiceTypesCmd)
}
