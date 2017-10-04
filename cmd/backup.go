package cmd

import (
       
       "github.com/spf13/cobra"
   //    "github.com/spf13/viper"
       
)

var (

    backup = &cobra.Command{
    	   Use:   "backup ",
    	   Short: "Asimov start backup. There are different possibilities to make the backup",
    	   Long:  `test`,
	   }

     )

func init(){
     
}