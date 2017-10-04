package cmd

import (
       "fmt"
       "log"
       "time"
       "sync"
)

func (r *BackupConfig) Backup(filter func(string) bool, name string){
	if len(r.Robots) <= 0 {
		fmt.Println("There is no robot to connect to")
		return
	}

	t := time.Now()

	log.Println("Backing up project...")
	
	bak_dir := fmt.Sprintf("%d-%02d-%02dT%02d-%02d-%02d_%s", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), name)

	dest := r.Destination + "/" + bak_dir

	var wg sync.WaitGroup
	
	robots_in_project := len(r.Robots)
	log.Println("Roboters in Project %d",robots_in_project)
	for _, rob := range r.Robots {
		wg.Add(1)
		go rob.Backup(filter, dest, &wg)
	}
	wg.Wait()

	log.Println("Backed up all robots in %d", time.Since(t))

}