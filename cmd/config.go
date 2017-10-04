package cmd

import (
       "log"
       "fmt"

       "robot"
)

type BackupConfig struct {
     Version string
     Robots  []robot.Robot
}

var bfg BackupConfig

func (r *BackupConfig) AddRobot(robot robot.Robot){
	r.Robots = append(r.Robots,robot)
}

func (r *BackupConfig) AddRobotByName(name string, ip string){
	robot := robot.InitRobot(name, ip)
	r.Robots = append(r.Robots,*robot)
}

func (r *BackupConfig) RemoveRobot(idx int) *robot.Robot{
	if len(r.Robots) <= 0 {
 		log.Println("No robot in the list use 'asimov add' to add the robot.")
 		return nil
 	}

	if (idx < 0) || (idx > len(r.Robots) - 1) {
		return nil
	}

	for id, robot := range r.Robots {
 		log.Printf(fmt.Sprintf("%d. %s %s\n", id, robot.Name, robot.Host))
 	}

	tmpRobot := r.Robots[idx]

	r.Robots = append(r.Robots[:idx], r.Robots[idx+1:]...)

	return &tmpRobot
}
