package main

import (
       "log"

	"asimov/cmd"
)

func main() {

     if err := cmd.Asimov.Execute(); err != nil {
     	log.Fatal(err)
     }

	/*bak := Init()
	
	app := cli.NewApp()
 	app.Name = ASIMOV_NAME
 	app.Usage = "Easy FANUC Backup"
 	app.Version = ASIMOV_VERSION
 	app.Author = "Anton Feldmann"
 	app.Email = "anton.feldmann@fanuc.eu"
	app.Commands = []cli.Command{
		{
 			Name:      "add",
 			ShortName: "a",
 			Usage:     "add a robot",
 			Action: func(c *cli.Context) {
				bak.addRobot()
			},
		},
		{
 			Name:      "backup",
 			ShortName: "b",
 			Usage:     "Backup the robotfiles",
 			Subcommands: []cli.Command{
				{
					Name:  "all",
					Usage: "*.*",
					Action: func(c *cli.Context) {
						bak.Backup(func(filename string) bool { return true }, "all")
					},
				},
				{
					Name:  "bin",
					Usage: "*.zip, *.sv, *.tp, *.vr",
					Action: func(c *cli.Context) {
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
