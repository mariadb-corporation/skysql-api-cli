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

	createDatabaseCmd.Flags().StringP("release-version", "v", DEFAULT_CREATE_DATABASE_RELEASE_VERSION, "Database version to deploy")
	viper.BindPFlag("release_version", createDatabaseCmd.Flags().Lookup("release-version"))

	createDatabaseCmd.Flags().StringP("topology", "t", DEFAULT_CREATE_DATABASE_TOPOLOGY, "Database topology to select")
	viper.BindPFlag("topology", createDatabaseCmd.Flags().Lookup("topology"))

	createDatabaseCmd.Flags().StringP("size", "s", DEFAULT_CREATE_DATABASE_SIZE, "Size of the nodes running the database")
	viper.BindPFlag("size", createDatabaseCmd.Flags().Lookup("size"))

	createDatabaseCmd.Flags().StringP("storage", "g", DEFAULT_CREATE_DATABASE_TXSTORAGE, "Size of the persistent storage disk")
	viper.BindPFlag("storage", createDatabaseCmd.Flags().Lookup("storage"))

	createDatabaseCmd.Flags().StringP("maxscale-config", "m", DEFAULT_CREATE_DATABASE_MAXSCALE_CONFIG, "Configurations for maxscale")
	viper.BindPFlag("maxscale_config", createDatabaseCmd.Flags().Lookup("maxscale-config"))

	createDatabaseCmd.Flags().StringP("name", "n", DEFAULT_CREATE_DATABASE_NAME, "Name used to identify the database")
	viper.BindPFlag("name", createDatabaseCmd.Flags().Lookup("name"))

	createDatabaseCmd.Flags().StringP("region", "r", DEFAULT_CREATE_DATABASE_REGION, "Geographic region to deploy the database")
	viper.BindPFlag("region", createDatabaseCmd.Flags().Lookup("region"))

	createDatabaseCmd.Flags().String("repl-region", DEFAULT_CREATE_DATABASE_REPL_REGION, "Replica region for the database")
	viper.BindPFlag("repl_region", createDatabaseCmd.Flags().Lookup("repl-region"))

	createDatabaseCmd.Flags().StringP("provider", "p", DEFAULT_CREATE_DATABASE_PROVIDER, "Cloud provider to host the database")
	viper.BindPFlag("provider", createDatabaseCmd.Flags().Lookup("provider"))

	createDatabaseCmd.Flags().String("replicas", DEFAULT_CREATE_DATABASE_REPLICAS, "Number of replicas to deploy")
	viper.BindPFlag("replicas", createDatabaseCmd.Flags().Lookup("replicas"))

	createDatabaseCmd.Flags().String("monitor", DEFAULT_CREATE_DATABASE_MONITOR, "Whether to deploy a monitoring cluster alongside the database")
	viper.BindPFlag("monitor", createDatabaseCmd.Flags().Lookup("monitor"))

	createDatabaseCmd.Flags().String("maxscale-proxy", DEFAULT_CREATE_DATABASE_MAXSCALE_PROXY, "Whether to set up a proxy for maxscale")
	viper.BindPFlag("maxscale_proxy", createDatabaseCmd.Flags().Lookup("maxscale-proxy"))
}
