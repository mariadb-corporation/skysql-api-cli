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
		Run: func(cmd *cobra.Command, args []string) {
			limit := viper.GetInt(LIMIT)

			var res *http.Response
			var err error
			res, err = client.ReadRegions(cmd.Context(), &skysql.ReadRegionsParams{
				Limit: &limit,
			})

			checkAndPrint(res, err, REGIONS)
		},
	}
)

func init() {
	getCmd.AddCommand(getRegionCmd)

	getRegionCmd.Flags().IntP(LIMIT, LIMIT_SHORTHAND, DEFAULT_GET_LIMIT, HINT_LIMIT)
	viper.BindPFlag(LIMIT, getRegionCmd.Flags().Lookup(LIMIT))
}
