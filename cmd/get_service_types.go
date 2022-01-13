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
		PreRun: func(cmd *cobra.Command, arg []string) {
			viper.BindPFlag(LIMIT, cmd.PersistentFlags().Lookup(LIMIT))
			viper.BindPFlag(OFFSET, cmd.PersistentFlags().Lookup(OFFSET))
		},
		Run: func(cmd *cobra.Command, args []string) {
			limit := viper.GetInt(LIMIT)
			offset := viper.GetInt(OFFSET)

			var res *http.Response
			var err error
			res, err = client.ReadServiceTypes(cmd.Context(), &skysql.ReadServiceTypesParams{
				Limit:  &limit,
				Offset: &offset,
			})

			checkAndPrint(res, err, SERVICE_TYPES)
		},
	}
)

func init() {
	getCmd.AddCommand(getServiceTypesCmd)

	getServiceTypesCmd.PersistentFlags().IntP(LIMIT, LIMIT_SHORTHAND, DEFAULT_GET_LIMIT, "Number of records to return. Can be used for paginating results in conjuntion with offset.")
	getServiceTypesCmd.PersistentFlags().IntP(OFFSET, OFFSET_SHORTHAND, DEFAULT_GET_OFFSET, "Number of records to skip when retrieved. Can be used for paginating results in conjunction with limit.")
}
