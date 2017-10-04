package cmd

import (
       "fmt"
       "net"

       "github.com/spf13/cobra"
      // "github.com/spf13/viper"      
)

var (
    add = &cobra.Command{
    Use:   "add [name] [ip-address]",
    Short: "Add a Robot to the Robot List",
    Long:  `Add adds a robot to the list of the backupable robots.
    	   	The list will contain the name and the robot ip adress.

	    Example: asimov add foo_bar 127.0.0.1`,
    Run:   func(cmd *cobra.Command, args []string){
    	   	    if len(args) < 2 {
		       er ("add needs two arguments")
		    }

		    name := args[0]
		    ip   := args[1]

		    if net.ParseIP(ip) == nil {
		       er (fmt.Sprintf("check the ip adress: %s", ip))
		    }

		    bfg.AddRobotByName(name,ip)
		    bfg.save(bfg.Config_Path)
    	   },
     }

)
