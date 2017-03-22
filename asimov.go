package main

import (
    //"fmt"
	"ftp"
	"log"
)

const ASIMOV_VERSION = "0.0.1"

func main() {
	
	c := ftp.NewConnection("127.0.0.1", "21") 
 	c.Connect() 
	defer c.Quit()
	files, err := c.NameList() 
 	if err != nil { 
 		log.Println("error getting list of files") 
 		return 
 	} 
 
 	c.Type("I") 

	var errorList []error 
 	for _, file := range files { 
 			log.Printf("Downloading %s", file) 
 			err := c.Download(file, "all/") 
			if err != nil { 
 				errorList = append(errorList, err) 
 			} 

 	} 

	if len(errorList) > 0 { 
 		log.Printf("There were %d errors.\n", len(errorList)) 
 		for _, err := range errorList { 
 			log.Println(err) 
 		} 
 	} 

	
	// robots := robot.GetRobotList()
	// robots.fp = "test"

    // app := cli.NewApp() 
 	// app.Name = "BackupTool" 
 	// app.Usage = "Easy FANUC Backup" 
 	// app.Version = ASIMOV_VERSION 
 	// app.Author = "Anton Feldmann" 
 	// app.Email = "anton.feldmann@fanuc.eu" 
 	// app.Commands = []cli.Command{ 
		// { 
 			// Name:      "add", 
 			// ShortName: "a", 
 			// Usage:     "add a robot", 
 			// Action: func(c *cli.Context) { 
 			// },
		// },
		// { 
 			// Name:      "backup", 
 			// ShortName: "b", 
 			// Usage:     "Backup the robotfiles", 
 			// Action: func(c *cli.Context) {
				// Backup("all")
 			// },
		// },
 		// { 
 			// Name:      "remove", 
 			// ShortName: "r", 
 			// Usage:     "remove a robot", 
 			// Action: func(c *cli.Context) { 

 			// }, 
 		// },
	// }
	// app.Run(os.Args) 
}