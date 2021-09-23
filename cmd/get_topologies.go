package cmd

import (
	"fmt"
	"net/http"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getTopologyCmd = &cobra.Command{
		Use:     TOPOLOGIES,
		Aliases: []string{TOPOLOGY},
		Short:   fmt.Sprintf("Retrieve list of MariaDB SkySQL %s %s", DATABASE, TOPOLOGIES),
		Long:    fmt.Sprintf("Retrieves list of %s %s available for use with MariaDB SkySQL", DATABASE, TOPOLOGIES),
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			limit := viper.GetInt(LIMIT)

			var res *http.Response
			var err error
			res, err = client.ReadTopologies(cmd.Context(), &skysql.ReadTopologiesParams{
				Limit: &limit,
			})

			checkAndPrint(res, err, TOPOLOGIES)
		},
	}
)

func init() {
	getCmd.AddCommand(getTopologyCmd)
}
