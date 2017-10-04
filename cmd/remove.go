package cmd

import "github.com/spf13/cobra"

var remove = &cobra.Command{
    	   Use:   "Remove [name]",
    	   Short: "Remove a Robot from the Robot List",
    	   Long:  `remove a robot from the list of robots. This Robot can not be backuped anymore
    
		Example: asimov remove foo_bar `,
    	    Run:   func(cmd *cobra.Command, args []string){
    	   	    if len(args) < 1 {
		       er ("remove needs two arguments")
		    }

		    name := args[0]

		    bfg.DelRobot(name)
		    bfg.save(bfg.Config_Path)
    	   	    },
     	    }
