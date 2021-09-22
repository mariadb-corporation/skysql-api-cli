package cmd

import (
	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	createDatabaseCmd = &cobra.Command{
		Use:   DATABASE,
		Short: "Create a new database",
		Long:  `Submits request to MariaDB SkySQL to deploy a new database`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			reqBody := skysql.CreateDatabaseJSONRequestBody{
				ReleaseVersion: viper.GetString("release_version"),
				Topology:       viper.GetString("topology"),
				Size:           viper.GetString("size"),
				TxStorage:      viper.GetString("storage"),
				MaxscaleConfig: viper.GetString("maxscale_config"),
				Name:           viper.GetString("name"),
				Region:         viper.GetString("region"),
				ReplRegion:     viper.GetString("repl_region"),
				Provider:       viper.GetString("provider"),
				Replicas:       viper.GetString("replicas"),
				Monitor:        viper.GetString("monitor"),
				MaxscaleProxy:  viper.GetString("maxscale_proxy"),
			}

			res, err := client.CreateDatabase(cmd.Context(), reqBody)

			checkAndPrint(res, err, DATABASES)
		},
	}
)

func init() {
	createCmd.AddCommand(createDatabaseCmd)

	createDatabaseCmd.Flags().StringP("release-version", "v", "MariaDB Enterprise Server 10.5.9-6", "Database version to deploy")
	viper.BindPFlag("release_version", createDatabaseCmd.Flags().Lookup("release-version"))

	createDatabaseCmd.Flags().StringP("topology", "t", "Standalone", "Database topology to select")
	viper.BindPFlag("topology", createDatabaseCmd.Flags().Lookup("topology"))

	createDatabaseCmd.Flags().StringP("size", "s", "Sky-2x4", "Size of the nodes running the database")
	viper.BindPFlag("size", createDatabaseCmd.Flags().Lookup("size"))

	createDatabaseCmd.Flags().StringP("storage", "g", "100", "Size of the persistent storage disk")
	viper.BindPFlag("storage", createDatabaseCmd.Flags().Lookup("storage"))

	createDatabaseCmd.Flags().StringP("maxscale-config", "m", "", "Configurations for maxscale")
	viper.BindPFlag("maxscale_config", createDatabaseCmd.Flags().Lookup("maxscale-config"))

	createDatabaseCmd.Flags().StringP("name", "n", "", "Name used to identify the database")
	viper.BindPFlag("name", createDatabaseCmd.Flags().Lookup("name"))

	createDatabaseCmd.Flags().StringP("region", "", "r", "Geographic region to deploy the database")
	viper.BindPFlag("region", createDatabaseCmd.Flags().Lookup("region"))

	createDatabaseCmd.Flags().String("repl-region", "", "Replica region for the database")
	viper.BindPFlag("repl_region", createDatabaseCmd.Flags().Lookup("repl-region"))

	createDatabaseCmd.Flags().StringP("provider", "p", "", "Cloud provider to host the database")
	viper.BindPFlag("provider", createDatabaseCmd.Flags().Lookup("provider"))

	createDatabaseCmd.Flags().String("replicas", "0", "Number of replicas to deploy")
	viper.BindPFlag("replicas", createDatabaseCmd.Flags().Lookup("replicas"))

	createDatabaseCmd.Flags().String("monitor", "false", "Whether to deploy a monitoring cluster alongside the database")
	viper.BindPFlag("monitor", createDatabaseCmd.Flags().Lookup("monitor"))

	createDatabaseCmd.Flags().String("maxscale-proxy", "false", "Whether to set up a proxy for maxscale")
	viper.BindPFlag("maxscale_proxy", createDatabaseCmd.Flags().Lookup("maxscale-proxy"))
}
