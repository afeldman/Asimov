package cmd

import (
       "path/filepath"
       
       "github.com/spf13/cobra"
)

var backup_bin = &cobra.Command{
    	   Use:     "bin",
    	   Short:   "save all binary files from the ftp server",
    	   Long:    `The backup command starts the backup of all the Robots contains in the configuration file.`,
	   Run:	    func(cmd *cobra.Command, args []string){
	   	    	     bfg.Backup(func(filename string) bool {

	  	      	        switch filepath.Ext(filename){
				   case ".zip",".sv",".tp",".vr":
				   	return true
				   default:
					return false
			        }

			     }, "bin")
	   	    },
	   }