package cmd

import (
	"github.com/spf13/cobra"
)

var backup_all = &cobra.Command{
	Use:     "all",
	Aliases: []string{"*"},
	Short:   "Asimov start backup. There are different possibilities to make the backup",
	Long:    `The backup command starts the backup of all the Robots contains in the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		bfg.Backup(func(filename string) bool { return true }, "all")
	},
}
