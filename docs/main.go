package main

import (
	"log"

	"github.com/mariadb-corporation/skysqlcli/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	rootCmd := cmd.SkySQLCLICommand()
	rootCmd.DisableAutoGenTag = true
	err := doc.GenMarkdownTree(rootCmd, "./docs/md/")
	if err != nil {
		log.Fatal(err)
	}
	err = doc.GenYamlTree(rootCmd, "./docs/yaml/")
	if err != nil {
		log.Fatal(err)
	}
	err = doc.GenReSTTree(rootCmd, "./docs/rst/")
	if err != nil {
		log.Fatal(err)
	}
	err = doc.GenManTree(rootCmd, &doc.GenManHeader{
		Title:   "SkySQL CLI Client",
		Source:  "MariaDB Corporation",
		Section: "3",
	}, "./docs/man/")
	if err != nil {
		log.Fatal(err)
	}
}
