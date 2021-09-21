package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	skysql "github.com/mariadb-corporation/skysql-api-go"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

// TokenResponse is the model for the /token endpoint from MariaDB ID.
// There are other fields in the response, but we don't use them.
type tokenResponse struct {
	Token string `json:"token"`
}

func checkErr(err interface{}, msg string) {
	if err != nil {
		crash(fmt.Sprintf(msg+": %v", err))
	}
}

func crash(msg string) {
	fmt.Fprintln(os.Stderr, "Error:", msg)
	os.Exit(1)
}

func checkAndPrint(res *http.Response, err error, entityName string) {
	checkErr(err, fmt.Sprintf("unable to retrieve %s from SkySQL", entityName))
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	checkErr(err, "unable to read response from SkySQL")

	if res.StatusCode != http.StatusOK {
		crash(fmt.Sprintf("unable to retrieve %s from SkySQL: Status: %v, Body: %v", entityName, res.StatusCode, string(body)))
	}

	fmt.Println(string(body))
}

var (
	cfgFile  string
	apiKey   string
	host     string
	mdbid    string
	client   *skysql.Client
	skipAuth = map[string]bool{ // no need to do auth for subcommands like completion generation
		"bash":       true,
		"fish":       true,
		"powershell": true,
		"zsh":        true,
		"version":    true,
	}
	validArgs = []string{GET}
	rootCmd   = &cobra.Command{
		Use:       CLI_NAME,
		Short:     "CLI client to interact with the SkySQL API",
		Long:      `A command line tool for managing resources deployed into MariaDB SkySQL`,
		ValidArgs: validArgs,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if _, ok := skipAuth[cmd.Use]; ok {
				return
			}

			apiKey := viper.GetString("api_key")
			if apiKey == "" {
				crash("required flag \"api-key\" not set " + cmd.Use)
			}

			mdbid = viper.GetString("mdbid")
			url, err := url.Parse(mdbid)
			checkErr(err, "unable to parse mdbid url")
			if url.String() == "" {
				checkErr(url, "unable to parse mdbid url")
			}
			url.Path = path.Join(url.Path, "/api/v1/token")

			req, _ := http.NewRequest("POST", url.String(), nil)
			req.Header.Add("Authorization", "Token "+apiKey)

			httpClient := &http.Client{}
			res, err := httpClient.Do(req)
			checkErr(err, "unable to send request for token to MariaDB ID")

			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				body, err := ioutil.ReadAll(res.Body)
				checkErr(err, "unable to read response from MariaDB ID")
				crash(fmt.Sprintf("unable to retrieve token from MariaDB ID: URL: %v, Status: %v, Body: %v", url, res.StatusCode, string(body)))
			}

			var resData tokenResponse
			err = json.NewDecoder(res.Body).Decode(&resData)
			checkErr(err, "unable to decode response from MariaDB ID")

			slt := resData.Token
			if slt == "" {
				crash("failed to retrieve short lived token from MariaDB ID")
			}

			bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken(slt)
			checkErr(err, "unable to initialize bearer auth security provider")
			client, err = skysql.NewClient(
				host,
				skysql.WithRequestEditorFn(bearerTokenProvider.Intercept),
			)
			checkErr(err, "unable to initialize skysql client")
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// PersistentPreRun: https://medium.com/@calavera/stop-saving-credential-tokens-in-text-files-65e840a237bb
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default $HOME/.skysqlcli.yaml)")
	rootCmd.PersistentFlags().StringVar(&apiKey, "api-key", "", "Long-lived JWT issued from MariaDB ID")
	rootCmd.PersistentFlags().StringVar(&host, "host", SKYSQL_API, "URL for the SkySQL API")
	rootCmd.PersistentFlags().StringVar(&mdbid, "mdbid", MARIADB_ID, "URL for MariaDB ID")

	viper.BindPFlag("api_key", rootCmd.PersistentFlags().Lookup("api-key"))
	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("mdbid", rootCmd.PersistentFlags().Lookup("mdbid"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".skysqlcli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(fmt.Sprintf(".%s", CLI_NAME))
	}

	viper.SetEnvPrefix(PROJECT_NAME)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
