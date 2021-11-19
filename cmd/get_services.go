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
		},
		Run: func(cmd *cobra.Command, args []string) {
			var res *http.Response
			var err error
			if len(args) == 1 {
				svcid := args[0]
				res, err = client.ReadService(cmd.Context(), svcid)
			} else {
				limit := viper.GetInt("limit")
				name := viper.GetString("name")
				res, err = client.ListServices(cmd.Context(), &skysql.ListServicesParams{
					Limit: &limit,
					Name:  &name,
				})
			}
			checkAndPrint(res, err, SERVICES)
		},
	}
)

func init() {
	getCmd.AddCommand(getServiceCmd)

	getServiceCmd.PersistentFlags().StringP(NAME, "n", "", "Search string to match any services containing the name")
}
