package cmd

import (
	"fmt"
	"net/http"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getRegionCmd = &cobra.Command{
		Use:     REGIONS,
		Aliases: []string{REGION},
		Short:   fmt.Sprintf("Retrieve list of %s %s", PROVIDER, REGIONS),
		Long:    fmt.Sprintf("Retrieves list of %s %s available for use with MariaDB SkySQL", PROVIDER, REGIONS),
		Args:    cobra.NoArgs,
		PreRun: func(cmd *cobra.Command, arg []string) {
			viper.BindPFlag(PROVIDER, cmd.Flags().Lookup(PROVIDER))
			viper.BindPFlag(LIMIT, cmd.PersistentFlags().Lookup(LIMIT))
			viper.BindPFlag(OFFSET, cmd.PersistentFlags().Lookup(OFFSET))
		},
		Run: func(cmd *cobra.Command, args []string) {
			limit := viper.GetInt(LIMIT)
			offset := viper.GetInt(OFFSET)
			provider := skysql.SnowProviders(viper.GetString(PROVIDER))

			var res *http.Response
			var err error
			res, err = client.ReadRegions(cmd.Context(), &skysql.ReadRegionsParams{
				Limit:    &limit,
				Offset:   &offset,
				Provider: provider,
			})

			checkAndPrint(res, err, REGIONS)
		},
	}
)

func init() {
	getCmd.AddCommand(getRegionCmd)

	getRegionCmd.Flags().String(PROVIDER, "", fmt.Sprintf("MariaDB SkySQL %s to query for %s", PROVIDER, REGIONS))
	getRegionCmd.PersistentFlags().IntP(LIMIT, LIMIT_SHORTHAND, DEFAULT_GET_LIMIT, "Number of records to return. Can be used for paginating results in conjuntion with offset.")
	getRegionCmd.PersistentFlags().IntP(OFFSET, OFFSET_SHORTHAND, DEFAULT_GET_OFFSET, "Number of records to skip when retrieved. Can be used for paginating results in conjunction with limit.")
}
