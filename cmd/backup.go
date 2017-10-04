package cmd

import (
       
       "github.com/spf13/cobra"
       "github.com/spf13/viper"
       
)

var (

    Backup = &cobra.Command{
    	   Use:   "backup",
    	   Short: "Asimov start backup. There are different possibilities to make the backup",
    	   Long:  `A fast and flexible backup tool for robot FTP Server.
    	      	   Each Robot using FTP can be backuped.`,
}

     confFile string
     )

)