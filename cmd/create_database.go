package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	releaseVersion    string
	topology          string
	size              string
	storage           string
	maxscaleConfig    string
	createDbName      string
	region            string
	replRegion        string
	provider          string
	replicas          string
	monitor           string
	maxscaleProxy     string
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

			checkErr(err, "unable to create database")
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			checkErr(err, "unable to read response from MariaDB SkySQL")

			if res.StatusCode != http.StatusOK {
				crash(fmt.Sprintf("unable to create database: Status: %v, Body: %v", res.StatusCode, string(body)))
			}

			fmt.Println(string(body))
		},
	}
)

func init() {
	createCmd.AddCommand(createDatabaseCmd)

	createDatabaseCmd.Flags().StringVarP(&releaseVersion, "release-version", "v", "MariaDB Enterprise Server 10.5.9-6", "Database version to deploy")
	viper.BindPFlag("release_version", createDatabaseCmd.Flags().Lookup("release-version"))

	createDatabaseCmd.Flags().StringVarP(&topology, "topology", "t", "Standalone", "Database topology to select")
	viper.BindPFlag("topology", createDatabaseCmd.Flags().Lookup("topology"))

	createDatabaseCmd.Flags().StringVarP(&size, "size", "s", "Sky-2x4", "Size of the nodes running the database")
	viper.BindPFlag("size", createDatabaseCmd.Flags().Lookup("size"))

	createDatabaseCmd.Flags().StringVarP(&storage, "storage", "g", "100", "Size of the persistent storage disk")
	viper.BindPFlag("storage", createDatabaseCmd.Flags().Lookup("storage"))

	createDatabaseCmd.Flags().StringVarP(&maxscaleConfig, "maxscale-config", "m", "", "Configurations for maxscale")
	viper.BindPFlag("maxscale_config", createDatabaseCmd.Flags().Lookup("maxscale-config"))

	createDatabaseCmd.Flags().StringVarP(&createDbName, "name", "n", "", "Name used to identify the database")
	viper.BindPFlag("name", createDatabaseCmd.Flags().Lookup("name"))

	createDatabaseCmd.Flags().StringVarP(&region, "region", "", "r", "Geographic region to deploy the database")
	viper.BindPFlag("region", createDatabaseCmd.Flags().Lookup("region"))

	createDatabaseCmd.Flags().StringVar(&replRegion, "repl-region", "", "Replica region for the database")
	viper.BindPFlag("repl_region", createDatabaseCmd.Flags().Lookup("repl-region"))

	createDatabaseCmd.Flags().StringVarP(&provider, "provider", "p", "", "Cloud provider to host the database")
	viper.BindPFlag("provider", createDatabaseCmd.Flags().Lookup("provider"))

	createDatabaseCmd.Flags().StringVar(&replicas, "replicas", "0", "Number of replicas to deploy")
	viper.BindPFlag("replicas", createDatabaseCmd.Flags().Lookup("replicas"))

	createDatabaseCmd.Flags().StringVar(&monitor, "monitor", "false", "Whether to deploy a monitoring cluster alongside the database")
	viper.BindPFlag("monitor", createDatabaseCmd.Flags().Lookup("monitor"))

	createDatabaseCmd.Flags().StringVar(&maxscaleProxy, "maxscale-proxy", "false", "Whether to set up a proxy for maxscale")
	viper.BindPFlag("maxscale_proxy", createDatabaseCmd.Flags().Lookup("maxscale-proxy"))
}
