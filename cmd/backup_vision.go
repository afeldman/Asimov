package cmd

import (
       "path/filepath"
       "github.com/spf13/cobra"
)

var backup_vision = &cobra.Command{
    	   Use:     "vision",
    	   Short:   "Asimov start backup. There are different possibilities to make the backup",
    	   Long:    `The backup command starts the backup of all the Robots contains in the configuration file.`,
	   Run:	    func(cmd *cobra.Command, args []string){
	   	    	     bfg.Backup(func(filename string) bool {

				switch filepath.Ext(filename) {
				   case ".vd",".vda",".zip":
				   	return true
				   default:
				        return false
				}

				
			     }, "vision")
	   	    },
	   }