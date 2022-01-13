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
		Short:   fmt.Sprintf("Retrieve list of MariaDB SkySQL %s %s", SERVICE, TOPOLOGIES),
		Long:    fmt.Sprintf("Retrieves list of %s %s available for use with MariaDB SkySQL", SERVICE, TOPOLOGIES),
		Args:    cobra.NoArgs,
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlag(SERVICE_TYPE, cmd.PersistentFlags().Lookup(SERVICE_TYPE))
		},
		Run: func(cmd *cobra.Command, args []string) {
			limit := viper.GetInt(LIMIT)
			offset := viper.GetInt(OFFSET)
			serviceType := skysql.ReadTopologiesParamsServiceType(viper.GetString(SERVICE_TYPE))

			var res *http.Response
			var err error
			res, err = client.ReadTopologies(cmd.Context(), &skysql.ReadTopologiesParams{
				Limit:       &limit,
				Offset:      &offset,
				ServiceType: serviceType,
			})

			checkAndPrint(res, err, TOPOLOGIES)
		},
	}
)

func init() {
	getCmd.AddCommand(getTopologyCmd)

	getTopologyCmd.PersistentFlags().StringP(SERVICE_TYPE, "t", "", fmt.Sprintf("MariaDB SkySQL %s used to filter list of %s", SERVICE_TYPE, TOPOLOGIES))
}
