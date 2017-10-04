package cmd

import (
       "log"
       "github.com/spf13/cobra"
   //    "github.com/spf13/viper"
       
)

var backup_root = &cobra.Command{
    	   Use:     "backup",
	   Aliases: []string{"bkg","back","bak"},
    	   Short:   "Asimov start backup. There are different possibilities to make the backup",
    	   Long:    `The backup command starts the backup of all the Robots contains in the configuration file.`,
	   Run:	    func(cmd *cobra.Command, args []string){
	   	    	     log.Println("To run the backup run a subcommand")
	   	    },
	   }

func init() {
     backup_root.AddCommand(backup_all)
}