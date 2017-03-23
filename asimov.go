package main

import (
	"bufio"
	"fmt"
	"os"
	"robot"
	"strings"
	"time"
	"sync"
	"log"
	
	"path/filepath"
	"encoding/json"
	"io/ioutil"
	
	"github.com/afeldman/cli"
)

const FILE_PATH = "backup.conf"

const ASIMOV_VERSION = "0.0.1"
const ASIMOV_NAME = "ASIMOV"


type BackupConf struct{
	Destination string
	Version		string
	Robots		robot.RobotList
}

func Init() *BackupConf{
	var bak BackupConf
	
	data, err := ioutil.ReadFile(FILE_PATH)
	
	if os.IsNotExist(err) {
		f, err := os.Create(FILE_PATH)
		check(err)
		fmt.Println("%s config file created.",FILE_PATH)
		f.Close()
		
		reader := bufio.NewReader(os.Stdin)		
		fmt.Println("Where should backups be stored?")
		dest, err := reader.ReadString('\n')
		check(err)
		dest = strings.TrimSpace(dest)
		
		bak.Destination = dest
		bak.Version = ASIMOV_VERSION
		bak.Save()		
	} else {
		json.Unmarshal([]byte(data), &bak)
	}
	
	return &bak

}

func (r *BackupConf) Save(){

 	b, err := json.Marshal(r)
 	check(err)
 	err = ioutil.WriteFile(FILE_PATH, b, 0644)
 	check(err)
 	fmt.Println("Project saved.")
}

func (r *BackupConf) Load(){
	data, err := ioutil.ReadFile(FILE_PATH)

	if os.IsNotExist(err) {
		f, err := os.Create(FILE_PATH)
		check(err)
 		fmt.Println("backupfile created.")
 		f.Close()
	} else {
		json.Unmarshal([]byte(data), &r)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (bak *BackupConf)addRobot(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please provide a name for the robot (e.g. R1):")
	name, err := reader.ReadString('\n')
	check(err)
	name = strings.TrimSpace(name)

	fmt.Printf("What is %s's IP address?\n", name)
	ip, err := reader.ReadString('\n')
	check(err)
	ip = strings.TrimSpace(ip)

	robot := robot.Robot{name,ip}
	bak.Robots.AddRobot(robot)

	bak.Save()
}

func (bak *BackupConf)removeRobot(){
	if len(bak.Robots.Robots) <= 0 {
		fmt.Println("There is not Robot to Remove.")
		return
	}

	for id, robot := range bak.Robots.Robots {
		fmt.Printf("%d. %s %s\n", id+1, robot.Name, robot.Host)
	}

	fmt.Println("\nWhich robot do you want to remove?")

	var id int
	_, err := fmt.Scanf("%d", &id)

	check(err)
	if bak.Robots.RemoveRobot(id) != nil {
		bak.Save()
	} else {
		fmt.Println("\nWhich robot do you want to remove?")
		return
	}
}

func (bak *BackupConf)Backup(filter func(string) bool, name string){
	if len(bak.Robots.Robots) <= 0 {
		fmt.Println("There is no robot to connect to")
		return
	}

	t := time.Now()

	fmt.Println("Backing up project...")

	dest := bak.Destination + "/" + fmt.Sprintf("%d-%02d-%02dT%02d-%02d-%02d_%s", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), name)

	var wg sync.WaitGroup
	for _, rob := range bak.Robots.Robots {
		log.Println("Backed up robot %s", rob.Name)
		wg.Add(1)
		go rob.Backup(filter, dest, &wg)
	}
	wg.Wait()
	log.Println("Backed up all robots in %v", time.Since(t))
}

func main() {

	bak := Init()
	
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
	app.Run(os.Args)
}
