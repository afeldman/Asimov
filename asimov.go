package main

import (
       "log"

	"asimov/cmd"
)

func main() {

     if err := cmd.Asimov.Execute(); err != nil {
     	log.Fatal(err)
     }

	/*

				{
					Name:  "ascii",
					Usage: "*.ls, *.va, *.dat, *.dg, *.xml",
					Action: func(c *cli.Context) {
						bak.Backup(func(filename string) bool {
							switch filepath.Ext(filename) {
							case ".ls", ".va", ".dat", ".dg", ".xml":
								return true
							}
							return false
						}, "ascii")
					},
				},
			},
		},
 		{
 			Name:      "remove",
 			ShortName: "r",
 			Usage:     "remove a robot",
 			Action: func(c *cli.Context) {
				bak.removeRobot()
 			},
 		},
	}
	app.Run(os.Args)*/
}
