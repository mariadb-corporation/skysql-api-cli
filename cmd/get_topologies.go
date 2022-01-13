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
			viper.BindPFlag(LIMIT, cmd.PersistentFlags().Lookup(LIMIT))
			viper.BindPFlag(OFFSET, cmd.PersistentFlags().Lookup(OFFSET))
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
	getTopologyCmd.PersistentFlags().IntP(LIMIT, LIMIT_SHORTHAND, DEFAULT_GET_LIMIT, "Number of records to return. Can be used for paginating results in conjuntion with offset.")
	getTopologyCmd.PersistentFlags().IntP(OFFSET, OFFSET_SHORTHAND, DEFAULT_GET_OFFSET, "Number of records to skip when retrieved. Can be used for paginating results in conjunction with limit.")
}
