package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	deleteConfigurationCmd = &cobra.Command{
		Use:   fmt.Sprintf("%s [%s]", CONFIGURATION, strings.ToUpper(CONFIGURATION)),
		Short: fmt.Sprintf("Delete a %s %s", DATABASE, CONFIGURATION),
		Long:  fmt.Sprintf("Deletes a %s %s for user in MariaDB SkySQL.", DATABASE, CONFIGURATION),
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			configNumber := string(args[0])
			res, err := client.DeleteConfiguration(cmd.Context(), configNumber)
			checkAndPrint(res, err, CONFIGURATION)
		},
	}
)

func init() {
	deleteCmd.AddCommand(deleteConfigurationCmd)
}
