package cmd

import (
	"github.com/spf13/cobra"
)

var backup_root = &cobra.Command{
	Use:     "backup",
	Aliases: []string{"bkg", "back", "bak"},
	Short:   "Asimov start backup. There are different possibilities to make the backup",
	Long:    `The backup command starts the backup of all the Robots contains in the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	backup_root.AddCommand(backup_all)
	backup_root.AddCommand(backup_bin)
	backup_root.AddCommand(backup_app)
	backup_root.AddCommand(backup_vision)
}
