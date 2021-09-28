package cmd

import (
	"encoding/json"
	"fmt"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	createConfigurationCmd = &cobra.Command{
		Use:   CONFIGURATION,
		Short: fmt.Sprintf("Create a new %s", CONFIGURATION),
		Long:  fmt.Sprintf("Creates a new %s for user in MariaDB SkySQL.", CONFIGURATION),
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			configJson := map[string]interface{}{}
			configJsonString := viper.GetString(CONFIG_JSON)
			if err := json.Unmarshal([]byte(configJsonString), &configJson); err != nil {
				crash(fmt.Sprintf("Invalid JSON provided for %s: %s", CONFIG_JSON, err.Error()))
			}

			reqBody := skysql.CreateConfigurationJSONRequestBody{
				ConfigJson: &configJson,
				Name:       viper.GetString(NAME),
				Topology:   viper.GetString(TOPOLOGY),
			}

			res, err := client.CreateConfiguration(cmd.Context(), reqBody)

			checkAndPrint(res, err, CONFIGURATION)
		},
	}
)

func init() {
	createCmd.AddCommand(createConfigurationCmd)

	createConfigurationCmd.Flags().StringP(TOPOLOGY, "t", "", "Configuration topology to select")
	createConfigurationCmd.Flags().StringP(NAME, "n", "", fmt.Sprintf("Name used to identify the %s", CONFIGURATION))
	createConfigurationCmd.Flags().StringP(CONFIG_JSON, "j", "", fmt.Sprintf("JSON object containing %s", CONFIGURATION))

	viper.BindPFlag(TOPOLOGY, createConfigurationCmd.Flags().Lookup(TOPOLOGY))
	viper.BindPFlag(NAME, createConfigurationCmd.Flags().Lookup(NAME))
	viper.BindPFlag(CONFIG_JSON, createConfigurationCmd.Flags().Lookup(CONFIG_JSON))
}
