package cmd

import (
       "path/filepath"
       "github.com/spf13/cobra"
)

var backup_app = &cobra.Command{
    	   Use:     "app",
    	   Short:   "Asimov start backup. There are different possibilities to make the backup",
    	   Long:    `The backup command starts the backup of all the Robots contains in the configuration file.`,
	   Run:	    func(cmd *cobra.Command, args []string){
	   	    	     bfg.Backup(func(filename string) bool {
			     
			        if filepath.Ext(filename) == ".tp"{
				   	return true
				}

				switch filename {
				   case "numreg.vr", "posreg.vr":
				   	return true
				   default:
				        return false
				}

				return false
				
			     }, "app")
	   	    },
	   }