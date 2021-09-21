package cmd

import (
	"github.com/spf13/cobra"
)

var (
	updateValidArgs = []string{DATABASE}
	updateCmd       = &cobra.Command{
		Use:       UPDATE,
		Short:     "Update a resource in MariaDB SkySQL",
		Long:      `Commands which update existing resources into MariaDB SkySQL`,
		ValidArgs: updateValidArgs,
	}
)

func init() {
	rootCmd.AddCommand(updateCmd)
}
