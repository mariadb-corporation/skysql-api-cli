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
		Short:   fmt.Sprintf("Retrieve list of MariaDB SkySQL %s %s", SERVICE, VERSIONS),
		Long:    fmt.Sprintf("Retrieves list of %s %s available for use with MariaDB SkySQL", SERVICE, VERSIONS),
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			limit := viper.GetInt(LIMIT)
			offset := viper.GetInt(OFFSET)

			var res *http.Response
			var err error
			res, err = client.ReadVersions(cmd.Context(), &skysql.ReadVersionsParams{
				Limit:  &limit,
				Offset: &offset,
			})

			checkAndPrint(res, err, VERSIONS)
		},
	}
)

func init() {
	getCmd.AddCommand(getVersionCmd)
}
