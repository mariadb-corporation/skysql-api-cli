package cmd

import (
	"fmt"
	"net/http"
	"strings"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getServiceCmd = &cobra.Command{
		Use:     fmt.Sprintf("%s [%s]", SERVICES, strings.ToUpper(SERVICE)),
		Aliases: []string{SERVICE},
		Short:   "Retrieve service information",
		Long:    "Queries for information about deployed service resources in SkySQL. " + HINT_SVC_ID,
		Args:    cobra.MaximumNArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlag(NAME, cmd.PersistentFlags().Lookup(NAME))
			viper.BindPFlag(LIMIT, cmd.PersistentFlags().Lookup(LIMIT))
			viper.BindPFlag(OFFSET, cmd.PersistentFlags().Lookup(OFFSET))
		},
		Run: func(cmd *cobra.Command, args []string) {
			var res *http.Response
			var err error
			if len(args) == 1 {
				svcid := args[0]
				res, err = client.ReadService(cmd.Context(), svcid)
			} else {
				limit := viper.GetInt(LIMIT)
				offset := viper.GetInt(OFFSET)
				name := viper.GetString("name")
				res, err = client.ListServices(cmd.Context(), &skysql.ListServicesParams{
					Limit:  &limit,
					Offset: &offset,
					Name:   &name,
				})
			}
			checkAndPrint(res, err, SERVICES)
		},
	}
)

func init() {
	getCmd.AddCommand(getServiceCmd)

	getServiceCmd.PersistentFlags().StringP(NAME, "n", "", "Search string to match any services containing the name")
	getServiceCmd.PersistentFlags().IntP(LIMIT, LIMIT_SHORTHAND, DEFAULT_GET_LIMIT, "Number of records to return. Can be used for paginating results in conjuntion with offset.")
	getServiceCmd.PersistentFlags().IntP(OFFSET, OFFSET_SHORTHAND, DEFAULT_GET_OFFSET, "Number of records to skip when retrieved. Can be used for paginating results in conjunction with limit.")
}
