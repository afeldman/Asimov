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

						bak.Backup(func(filename string) bool {
							switch filepath.Ext(filename) {
							case ".zip", ".sv", ".tp", ".vr":
								return true
							}
							return false
						}, "bin")
					},
				},
				{
					Name:  "app",
					Usage: "*.tp, numreg.vr, posreg.vr",
					Action: func(c *cli.Context) {
						bak.Backup(func(filename string) bool {
							switch filepath.Ext(filename) {
							case ".tp":
								return true
							}
							switch filename {
							case "numreg.vr", "posreg.vr":
								return true
							}
							return false
						}, "app")
					},
				},
				{
					Name:  "vision",
					Usage: "*.vd, *.vda, *.zip",
					Action: func(c *cli.Context) {
						bak.Backup(func(filename string) bool {
							switch filepath.Ext(filename) {
							case ".vd", ".vda", ".zip":
								return true
							}
							return false
						}, "vision")
					},
				},
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
