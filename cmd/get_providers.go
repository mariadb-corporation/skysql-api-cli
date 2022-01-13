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
		PreRun: func(cmd *cobra.Command, arg []string) {
			viper.BindPFlag(LIMIT, cmd.PersistentFlags().Lookup(LIMIT))
			viper.BindPFlag(OFFSET, cmd.PersistentFlags().Lookup(OFFSET))
		},
		Run: func(cmd *cobra.Command, args []string) {
			limit := viper.GetInt(LIMIT)
			offset := viper.GetInt(OFFSET)

			var res *http.Response
			var err error
			res, err = client.ReadProviders(cmd.Context(), &skysql.ReadProvidersParams{
				Limit:  &limit,
				Offset: &offset,
			})

			checkAndPrint(res, err, PROVIDERS)
		},
	}
)

func init() {
	getCmd.AddCommand(getProviderCmd)

	getProviderCmd.PersistentFlags().IntP(LIMIT, LIMIT_SHORTHAND, DEFAULT_GET_LIMIT, "Number of records to return. Can be used for paginating results in conjuntion with offset.")
	getProviderCmd.PersistentFlags().IntP(OFFSET, OFFSET_SHORTHAND, DEFAULT_GET_OFFSET, "Number of records to skip when retrieved. Can be used for paginating results in conjunction with limit.")
}
