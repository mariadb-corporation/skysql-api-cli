package cmd

import (
	"fmt"
	"net/http"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getVersionCmd = &cobra.Command{
		Use:     VERSIONS,
		Aliases: []string{VERSION},
		Short:   fmt.Sprintf("Retrieve list of MariaDB SkySQL %s %s", DATABASE, VERSIONS),
		Long:    fmt.Sprintf("Retrieves list of %s %s available for use with MariaDB SkySQL", DATABASE, VERSIONS),
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			limit := viper.GetInt(LIMIT)

			var res *http.Response
			var err error
			res, err = client.ReadVersions(cmd.Context(), &skysql.ReadVersionsParams{
				Limit: &limit,
			})

			checkAndPrint(res, err, VERSIONS)
		},
	}
)

func init() {
	getCmd.AddCommand(getVersionCmd)
}
