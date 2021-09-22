package cmd

import (
	"github.com/spf13/cobra"
)

var (
	createValidArgs = []string{DATABASE}
	createCmd       = &cobra.Command{
		Use:       CREATE,
		Short:     "Create a resource in MariaDB SkySQL",
		Long:      `Commands which deploy resources into MariaDB SkySQL`,
		ValidArgs: createValidArgs,
	}
)

func init() {
	rootCmd.AddCommand(createCmd)
}
