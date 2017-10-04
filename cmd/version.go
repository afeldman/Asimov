package cmd

import (
       "fmt"
       
       "github.com/spf13/cobra"
)


const ASIMOV_VERSION = "0.1.0"

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "print version number of Asimov",
    Long:  `asimove has different software versions. to specify
    	   the current use this version command`,
    Run:   func(cmd *cobra.Command, args []string){
    	   fmt.Println("Asimov Version is: ", ASIMOV_VERSION)
    },
}

func init() {
     Asimov.AddCommand(versionCmd)
}