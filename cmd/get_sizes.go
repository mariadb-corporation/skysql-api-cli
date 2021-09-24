package cmd

import (
	"fmt"
	"net/http"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getSizeCmd = &cobra.Command{
		Use:     SIZES,
		Aliases: []string{SIZE},
		Short:   fmt.Sprintf("Retrieve list of %s %s", STORAGE, SIZES),
		Long:    fmt.Sprintf("Retrieves list of %s %s available for use with MariaDB SkySQL", STORAGE, SIZES),
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			limit := viper.GetInt(LIMIT)
			product := viper.GetString(PRODUCT)
			provider := viper.GetString(PROVIDER)

			var res *http.Response
			var err error
			res, err = client.ReadSizes(cmd.Context(), &skysql.ReadSizesParams{
				Limit:    &limit,
				Product:  &product,
				Provider: &provider,
			})

			checkAndPrint(res, err, SIZES)
		},
	}
)

func init() {
	getCmd.AddCommand(getSizeCmd)

	getSizeCmd.PersistentFlags().String(PRODUCT, "", fmt.Sprintf("MariaDB SkySQL %s to query for %s %s", PRODUCT, STORAGE, SIZES))
	getSizeCmd.PersistentFlags().String(PROVIDER, "", fmt.Sprintf("MariaDB SkySQL %s to query for %s %s", PROVIDER, STORAGE, SIZES))

	viper.BindPFlag(PRODUCT, getSizeCmd.PersistentFlags().Lookup(PRODUCT))
	viper.BindPFlag(PROVIDER, getSizeCmd.PersistentFlags().Lookup(PROVIDER))

}
