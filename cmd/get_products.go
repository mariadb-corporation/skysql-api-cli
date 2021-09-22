package cmd

import (
	"fmt"
	"net/http"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getProductCmd = &cobra.Command{
		Use:     PRODUCTS + " [PRODUCT]",
		Aliases: []string{PRODUCT},
		Short:   fmt.Sprintf("Retrieve MariaDB SkySQL %s information", PRODUCT),
		Long:    fmt.Sprintf("Queries information for %s offerings from MariaDB SkySQL", PRODUCT),
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			limit := viper.GetInt("limit")

			var res *http.Response
			var err error
			res, err = client.ReadProducts(cmd.Context(), &skysql.ReadProductsParams{
				Limit: &limit,
			})

			checkAndPrint(res, err, PRODUCTS)
		},
	}
)

func init() {
	getCmd.AddCommand(getProductCmd)

	getProductCmd.Flags().IntP(LIMIT, LIMIT_SHORTHAND, DEFAULT_GET_LIMIT, HINT_LIMIT)
	viper.BindPFlag(LIMIT, getProductCmd.Flags().Lookup(LIMIT))
}
