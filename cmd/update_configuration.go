package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	updateConfigurationCmd = &cobra.Command{
		Use:   fmt.Sprintf("%s [%s]", CONFIGURATION, strings.ToUpper(CONFIGURATION)),
		Short: fmt.Sprintf("Update a %s", CONFIGURATION),
		Long:  fmt.Sprintf("Updates a %s for user in MariaDB SkySQL.", CONFIGURATION),
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			configNumber := string(args[0])
			configJson := map[string]interface{}{}
			configJsonString := viper.GetString(CONFIG_JSON)
			if err := json.Unmarshal([]byte(configJsonString), &configJson); err != nil {
				crash(fmt.Sprintf("Invalid JSON provided for %s: %s", CONFIG_JSON, err.Error()))
			}

			name := viper.GetString(NAME)
			reqBody := skysql.UpdateConfigurationJSONRequestBody{
				ConfigJson: &configJson,
				Name:       &name,
			}

			res, err := client.UpdateConfiguration(cmd.Context(), configNumber, reqBody)

			checkAndPrint(res, err, CONFIGURATION)
		},
	}
)

func init() {
	updateCmd.AddCommand(updateConfigurationCmd)

	updateConfigurationCmd.Flags().StringP(NAME, "n", "", fmt.Sprintf("Name used to identify the %s", CONFIGURATION))
	updateConfigurationCmd.Flags().StringP(CONFIG_JSON, "j", "", fmt.Sprintf("JSON object containing %s", CONFIGURATION))

	viper.BindPFlag(NAME, updateConfigurationCmd.Flags().Lookup(NAME))
	viper.BindPFlag(CONFIG_JSON, updateConfigurationCmd.Flags().Lookup(CONFIG_JSON))
}
