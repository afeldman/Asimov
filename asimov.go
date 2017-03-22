package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/afeldman/cli"
	"robot"
	"strings"
	"ping"
	"time"
	"sync"
)

const ASIMOV_VERSION = "0.0.1"
const ASIMOV_NAME = "ASIMOV"

var destination string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func addRobot(r *robot.RobotList){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please provide a name for the robot (e.g. R1):")
	name, err := reader.ReadString('\n')
	check(err)
	name = strings.TrimSpace(name)

	fmt.Printf("What is %s's IP address?\n", name)
	ip, err := reader.ReadString('\n')
	check(err)
	ip = strings.TrimSpace(ip)

	if ping.Ping(ip,5) {
		robot := robot.Robot{name,ip}
		r.AddRobot(robot)

		r.Save()
	}else{
		return
	}
}

func removeRobot(r *robot.RobotList){
	if len(r.Robots) <= 0 {
		fmt.Println("There is not Robot to Remove.")
		return
	}

	for id, robot := range r.Robots {
		fmt.Printf("%d. %s %s\n", id+1, robot.Name, robot.Host)
	}

	fmt.Println("\nWhich robot do you want to remove?")

	var id int
	_, err := fmt.Scanf("%d", &id)

	check(err)
	if r.RemoveRobot(id) != nil {
		r.Save()
	} else {
		return
	}
}

func Backup(filter func(string) bool, name string, r *robot.RobotList){
	if len(r.Robots) <= 0 {
		fmt.Println("Your project does not have any robots. Please run `BackupTool add` to add one.")
		return
	}

	t := time.Now()

	fmt.Println("Backing up project...")

	dest := destination + "/" + fmt.Sprintf("%d-%02d-%02dT%02d-%02d-%02d_%s", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), name)

	var wg sync.WaitGroup
	for _, rob := range r.Robots {
		wg.Add(1)
		go rob.Backup(filter, dest, &wg)
	}
	wg.Wait()
	fmt.Printf("Backed up all robots in %v", time.Since(t))
}

func main() {

	r := robot.RobotList{}

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
				addRobot(&r)
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
						Backup(func(filename string) bool { return true }, "all", &r)
					},
				},
			},
		},
 		{
 			Name:      "remove",
 			ShortName: "r",
 			Usage:     "remove a robot",
 			Action: func(c *cli.Context) {
				removeRobot(&r)
 			},
 		},
	}
	app.Run(os.Args)
}
