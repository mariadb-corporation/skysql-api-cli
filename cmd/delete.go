package cmd

import (
	"github.com/spf13/cobra"
)

var (
	deleteValidArgs = []string{SERVICE}
	deleteCmd       = &cobra.Command{
		Use:       DELETE,
		Short:     "Delete a resource in MariaDB SkySQL",
		Long:      `Commands which delete existing resources in MariaDB SkySQL`,
		ValidArgs: deleteValidArgs,
	}
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}
