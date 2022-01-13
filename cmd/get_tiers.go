package cmd

import (
	"fmt"
	"net/http"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getTierCmd = &cobra.Command{
		Use:     TIERS,
		Aliases: []string{TIER},
		Short:   fmt.Sprintf("Retrieve list of MariaDB SkySQL %s %s", ACCOUNT, TIERS),
		Long:    fmt.Sprintf("Retrieves list of %s %s available for use with MariaDB SkySQL", ACCOUNT, TIERS),
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			limit := viper.GetInt(LIMIT)
			offset := viper.GetInt(OFFSET)

			var res *http.Response
			var err error
			res, err = client.ReadTiers(cmd.Context(), &skysql.ReadTiersParams{
				Limit:  &limit,
				Offset: &offset,
			})

			checkAndPrint(res, err, TIERS)
		},
	}
)

func init() {
	getCmd.AddCommand(getTierCmd)
}
