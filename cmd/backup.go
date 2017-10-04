package cmd

import (
       
       "github.com/spf13/cobra"
   //    "github.com/spf13/viper"
       
)

var backup = &cobra.Command{
    	   Use:     "backup",
	   Aliases: []string{"bkg","back","bak"},
    	   Short:   "Asimov start backup. There are different possibilities to make the backup",
    	   Long:    `The backup command starts the backup of all the Robots contains in the configuration file.`,
	   }
