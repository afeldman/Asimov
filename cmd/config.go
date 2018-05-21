package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/afeldman/go-robot"
)

type BackupConfig struct {
	Config_Path string `yaml:"asimov_config_path"`
	Version     string `yaml:"asimov_version"`
	Robots      []robot.Robot
	Destination string `yaml:"asimov_destination"`
}

var bfg BackupConfig

func (r *BackupConfig) save(path string) {
	d, err := yaml.Marshal(r)
	if err != nil {
		log.Fatal("cannot yamalize config")
	}

	if err := ioutil.WriteFile(path, d, 0640); err != nil {
		log.Fatal("cannot write configuration into file")
	}
}

func (r *BackupConfig) contains(name string) int {
	if len(r.Robots) <= 0 {
		return -1
	}

	for idx, robo := range r.Robots {
		if robo.Name == strings.TrimSpace(name) {
			return idx
		}
	}

	return -1
}

func (r *BackupConfig) AddRobot(robot robot.Robot) {
	r.Robots = append(r.Robots, robot)
}

func (r *BackupConfig) AddRobotByName(name string, ip string) {
	if idx := r.contains(name); idx < 0 {
		robot := robot.InitRobot(name, ip)
		r.Robots = append(r.Robots, *robot)
	} else {
		r.Robots[idx].Host = ip
	}

}

func (r *BackupConfig) GetRobot(name string) *robot.Robot {
	if len(r.Robots) <= 0 {
		return nil
	}

	idx := r.contains(name)

	if idx == -1 {
		return nil
	}

	return &(r.Robots[idx])

}

func (r *BackupConfig) DelRobot(name string) *robot.Robot {
	if len(r.Robots) <= 0 {
		return nil
	}

	idx := r.contains(name)

	if idx == -1 {
		return nil
	}

	return r.RemoveRobot(idx)

}

func (r *BackupConfig) RemoveRobot(idx int) *robot.Robot {
	if len(r.Robots) <= 0 {
		log.Println("No robot in the list use 'asimov add' to add the robot.")
		return nil
	}

	if (idx < 0) || (idx > len(r.Robots)-1) {
		return nil
	}

	for id, robot := range r.Robots {
		log.Printf(fmt.Sprintf("%d. %s %s\n", id, robot.Name, robot.Host))
	}

	tmpRobot := r.Robots[idx]

	r.Robots = append(r.Robots[:idx], r.Robots[idx+1:]...)

	return &tmpRobot
}
