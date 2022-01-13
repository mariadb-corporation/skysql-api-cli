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
		Short:   fmt.Sprintf("Retrieve list of %s %s", MACHINE, SIZES),
		Long:    fmt.Sprintf("Retrieves list of %s %s available for use with MariaDB SkySQL", MACHINE, SIZES),
		Args:    cobra.NoArgs,
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlag(SERVICE_TYPE, cmd.PersistentFlags().Lookup(SERVICE_TYPE))
			viper.BindPFlag(PROVIDER, cmd.PersistentFlags().Lookup(PROVIDER))
			viper.BindPFlag(TIER, cmd.PersistentFlags().Lookup(TIER))
			viper.BindPFlag(LIMIT, cmd.PersistentFlags().Lookup(LIMIT))
			viper.BindPFlag(OFFSET, cmd.PersistentFlags().Lookup(OFFSET))
		},
		Run: func(cmd *cobra.Command, args []string) {
			limit := viper.GetInt(LIMIT)
			offset := viper.GetInt(OFFSET)
			serviceType := skysql.ReadSizesParamsServiceType(viper.GetString(SERVICE_TYPE))
			provider := skysql.SnowProviders(viper.GetString(PROVIDER))
			tier := skysql.ReadSizesParamsTier(viper.GetString(TIER))

			var res *http.Response
			var err error
			res, err = client.ReadSizes(cmd.Context(), &skysql.ReadSizesParams{
				Limit:       &limit,
				Offset:      &offset,
				ServiceType: serviceType,
				Provider:    provider,
				Tier:        tier,
			})

			checkAndPrint(res, err, SIZES)
		},
	}
)

func init() {
	getCmd.AddCommand(getSizeCmd)

	getSizeCmd.PersistentFlags().String(SERVICE_TYPE, "t", fmt.Sprintf("MariaDB SkySQL %s to query for %s %s", SERVICE_TYPE, MACHINE, SIZES))
	getSizeCmd.PersistentFlags().String(PROVIDER, "", fmt.Sprintf("MariaDB SkySQL %s to query for %s %s", PROVIDER, MACHINE, SIZES))
	getSizeCmd.PersistentFlags().String(TIER, "", fmt.Sprintf("MariaDB SkySQL %s to query for %s %s", TIER, MACHINE, SIZES))
	getSizeCmd.PersistentFlags().IntP(LIMIT, LIMIT_SHORTHAND, DEFAULT_GET_LIMIT, "Number of records to return. Can be used for paginating results in conjuntion with offset.")
	getSizeCmd.PersistentFlags().IntP(OFFSET, OFFSET_SHORTHAND, DEFAULT_GET_OFFSET, "Number of records to skip when retrieved. Can be used for paginating results in conjunction with limit.")
}
