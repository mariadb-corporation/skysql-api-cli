package cmd

import (
	"fmt"
	"strings"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	createServiceCmd = &cobra.Command{
		Use:   SERVICE,
		Short: "Create a new service",
		Long:  `Submits request to MariaDB SkySQL to deploy a new service`,
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
			maxscaleConfig := viper.GetString(MAXSCALE_CONFIG)
			replRegion := viper.GetString(REPL_REGION)
			monitor := viper.GetString(MONITOR)
			maxscaleProxy := viper.GetString(MAXSCALE_PROXY)
			volumeIOPS := viper.GetString(VOLUME_IOPS)
			volumeType := skysql.ServiceInVolumeType(viper.GetString(VOLUME_TYPE))
			ssltls := skysql.ServiceInSslTls(viper.GetString(SSL_TLS))
			reqBody := skysql.CreateServiceJSONRequestBody{
				ReleaseVersion: viper.GetString(RELEASE_VERSION),
				Topology:       skysql.ServiceInTopology(viper.GetString(TOPOLOGY)),
				Size:           viper.GetString(SIZE),
				TxStorage:      viper.GetString(STORAGE),
				MaxscaleConfig: &maxscaleConfig,
				Name:           viper.GetString(NAME),
				Region:         viper.GetString(REGION),
				ReplRegion:     &replRegion,
				Provider:       skysql.SnowProviders(viper.GetString(PROVIDER)),
				Replicas:       viper.GetString(REPLICAS),
				Monitor:        &monitor,
				MaxscaleProxy:  &maxscaleProxy,
				VolumeIops:     &volumeIOPS,
				VolumeType:     &volumeType,
				Tier:           skysql.ServiceInTier(viper.GetString(TIER)),
				SslTls:         &ssltls,
			}

			res, err := client.CreateService(cmd.Context(), reqBody)

			checkAndPrint(res, err, SERVICES)
		},
	}
)

func init() {
	createCmd.AddCommand(createServiceCmd)

	createServiceCmd.Flags().StringP(RELEASE_VERSION, "v", DEFAULT_CREATE_SERVICE_RELEASE_VERSION, "Release version to deploy")
	createServiceCmd.Flags().StringP(TOPOLOGY, "t", DEFAULT_CREATE_SERVICE_TOPOLOGY, "Service topology to select")
	createServiceCmd.Flags().StringP(SIZE, "s", DEFAULT_CREATE_SERVICE_SIZE, "Size of the nodes running the service")
	createServiceCmd.Flags().StringP(STORAGE, "g", DEFAULT_CREATE_SERVICE_TXSTORAGE, "Size of the persistent storage disk")
	createServiceCmd.Flags().StringP(MAXSCALE_CONFIG, "m", DEFAULT_CREATE_SERVICE_MAXSCALE_CONFIG, "Configurations for maxscale")
	createServiceCmd.Flags().StringP(NAME, "n", DEFAULT_CREATE_SERVICE_NAME, "Name used to identify the service")
	createServiceCmd.Flags().StringP(REGION, "r", DEFAULT_CREATE_SERVICE_REGION, "Geographic region to deploy the service")
	createServiceCmd.Flags().String(REPL_REGION, DEFAULT_CREATE_SERVICE_REPL_REGION, "Replica region for the service")
	createServiceCmd.Flags().StringP(PROVIDER, "p", DEFAULT_CREATE_SERVICE_PROVIDER, "Cloud provider to host the service")
	createServiceCmd.Flags().String(REPLICAS, DEFAULT_CREATE_SERVICE_REPLICAS, "Number of replicas to deploy")
	createServiceCmd.Flags().String(MONITOR, DEFAULT_CREATE_SERVICE_MONITOR, "Whether to deploy a monitoring cluster alongside the service")
	createServiceCmd.Flags().String(MAXSCALE_PROXY, DEFAULT_CREATE_SERVICE_MAXSCALE_PROXY, "Whether to set up a proxy for maxscale")
	createServiceCmd.Flags().String(VOLUME_IOPS, DEFAULT_CREATE_SERVICE_VOLUME_IOPS, "Amount of IOPS for the volume (e.g. 100). Required for Amazon AWS")
	createServiceCmd.Flags().String(VOLUME_TYPE, DEFAULT_CREATE_SERVICE_VOLUME_TYPE, "Type of volume to use (e.g. io1, gp3). Required for Amazon AWS")
	createServiceCmd.Flags().String(TIER, DEFAULT_CREATE_SERVICE_TIER, fmt.Sprintf("%s in which to provision %s", strings.Title(TIER), SERVICE))
	createServiceCmd.Flags().String(SSL_TLS, DEFAULT_CREATE_SERVICE_SSL_TLS, "Specify whether to use SSL/TLS encryption")
}
