package cmd

import (
	"fmt"
	"strings"

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
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlag(RELEASE_VERSION, cmd.Flags().Lookup(RELEASE_VERSION))
			viper.BindPFlag(TOPOLOGY, cmd.Flags().Lookup(TOPOLOGY))
			viper.BindPFlag(SIZE, cmd.Flags().Lookup(SIZE))
			viper.BindPFlag(STORAGE, cmd.Flags().Lookup(STORAGE))
			viper.BindPFlag(MAXSCALE_CONFIG, cmd.Flags().Lookup(MAXSCALE_CONFIG))
			viper.BindPFlag(NAME, cmd.Flags().Lookup(NAME))
			viper.BindPFlag(REGION, cmd.Flags().Lookup(REGION))
			viper.BindPFlag(REPL_REGION, cmd.Flags().Lookup(REPL_REGION))
			viper.BindPFlag(PROVIDER, cmd.Flags().Lookup(PROVIDER))
			viper.BindPFlag(REPLICAS, cmd.Flags().Lookup(REPLICAS))
			viper.BindPFlag(MONITOR, cmd.Flags().Lookup(MONITOR))
			viper.BindPFlag(MAXSCALE_PROXY, cmd.Flags().Lookup(MAXSCALE_PROXY))
			viper.BindPFlag(VOLUME_IOPS, cmd.Flags().Lookup(VOLUME_IOPS))
			viper.BindPFlag(VOLUME_TYPE, cmd.Flags().Lookup(VOLUME_TYPE))
			viper.BindPFlag(TIER, cmd.Flags().Lookup(TIER))
		},
		Run: func(cmd *cobra.Command, args []string) {
			reqBody := skysql.CreateDatabaseJSONRequestBody{
				ReleaseVersion: viper.GetString(RELEASE_VERSION),
				Topology:       viper.GetString(TOPOLOGY),
				Size:           viper.GetString(SIZE),
				TxStorage:      viper.GetString(STORAGE),
				MaxscaleConfig: viper.GetString(MAXSCALE_CONFIG),
				Name:           viper.GetString(NAME),
				Region:         viper.GetString(REGION),
				ReplRegion:     viper.GetString(REPL_REGION),
				Provider:       viper.GetString(PROVIDER),
				Replicas:       viper.GetString(REPLICAS),
				Monitor:        viper.GetString(MONITOR),
				MaxscaleProxy:  viper.GetString(MAXSCALE_PROXY),
				VolumeIops:     viper.GetString(VOLUME_IOPS),
				VolumeType:     viper.GetString(VOLUME_TYPE),
				Tier:           viper.GetString(TIER),
			}

			res, err := client.CreateDatabase(cmd.Context(), reqBody)

			checkAndPrint(res, err, DATABASES)
		},
	}
)

func init() {
	createCmd.AddCommand(createDatabaseCmd)

	createDatabaseCmd.Flags().StringP(RELEASE_VERSION, "v", DEFAULT_CREATE_DATABASE_RELEASE_VERSION, "Database version to deploy")
	createDatabaseCmd.Flags().StringP(TOPOLOGY, "t", DEFAULT_CREATE_DATABASE_TOPOLOGY, "Database topology to select")
	createDatabaseCmd.Flags().StringP(SIZE, "s", DEFAULT_CREATE_DATABASE_SIZE, "Size of the nodes running the database")
	createDatabaseCmd.Flags().StringP(STORAGE, "g", DEFAULT_CREATE_DATABASE_TXSTORAGE, "Size of the persistent storage disk")
	createDatabaseCmd.Flags().StringP(MAXSCALE_CONFIG, "m", DEFAULT_CREATE_DATABASE_MAXSCALE_CONFIG, "Configurations for maxscale")
	createDatabaseCmd.Flags().StringP(NAME, "n", DEFAULT_CREATE_DATABASE_NAME, "Name used to identify the database")
	createDatabaseCmd.Flags().StringP(REGION, "r", DEFAULT_CREATE_DATABASE_REGION, "Geographic region to deploy the database")
	createDatabaseCmd.Flags().String(REPL_REGION, DEFAULT_CREATE_DATABASE_REPL_REGION, "Replica region for the database")
	createDatabaseCmd.Flags().StringP(PROVIDER, "p", DEFAULT_CREATE_DATABASE_PROVIDER, "Cloud provider to host the database")
	createDatabaseCmd.Flags().String(REPLICAS, DEFAULT_CREATE_DATABASE_REPLICAS, "Number of replicas to deploy")
	createDatabaseCmd.Flags().String(MONITOR, DEFAULT_CREATE_DATABASE_MONITOR, "Whether to deploy a monitoring cluster alongside the database")
	createDatabaseCmd.Flags().String(MAXSCALE_PROXY, DEFAULT_CREATE_DATABASE_MAXSCALE_PROXY, "Whether to set up a proxy for maxscale")
	createDatabaseCmd.Flags().String(VOLUME_IOPS, DEFAULT_CREATE_DATABASE_VOLUME_IOPS, "Amount of IOPS for the volume (e.g. 100). Required for Amazon AWS")
	createDatabaseCmd.Flags().String(VOLUME_TYPE, DEFAULT_CREATE_DATABASE_VOLUME_TYPE, "Type of volume to use (e.g. io1, gp3). Required for Amazon AWS")
	createDatabaseCmd.Flags().String(TIER, DEFAULT_CREATE_DATABASE_VOLUME_TYPE, fmt.Sprintf("%s in which to provision %s", strings.ToTitle(TIER), DATABASE))
}
