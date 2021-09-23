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

			var res *http.Response
			var err error
			res, err = client.ReadSizes(cmd.Context(), &skysql.ReadSizesParams{
				Limit: &limit,
			})

			checkAndPrint(res, err, SIZES)
		},
	}
)

func init() {
	getCmd.AddCommand(getSizeCmd)
}
