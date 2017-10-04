import (
       "ftp"
       "robot"
       "time"
)

func (r *BackupConfig) Backup(filter func(string) bool, name string){
	if len(r.Robots) <= 0 {
		fmt.Println("There is no robot to connect to")
		return
	}

	t := time.Now()

	fmt.Println("Backing up project...")
	
	bak_dir := fmt.Sprintf("%d-%02d-%02dT%02d-%02d-%02d_%s", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), name)
	dest := bak.Destination + "/" + bak_dir

	var wg sync.WaitGroup
	robots_in_project := len(bak.Robots.Robots)
	log.Println("Roboters in Project %d",robots_in_project)
	for _, rob := range bak.Robots.Robots {
		wg.Add(1)
		go rob.Backup(filter, dest, &wg)
	}
	wg.Wait()
	log.Println("Backed up all robots in %d", time.Since(t))
	
	//compress
/*	var hash2 sha1.FileHash
	var tmp_str string
	tarpath := compress.TarIt(dest,bak_dir)
	hash2.Hash(tarpath)
	for _, node := range hash2.Nodes {
		tmp_str = fmt.Sprintf("%s\t%x\n", node.Path, node.Hash)
	}
	compress.GzipIt(tarpath,tarpath,tmp_str)
	
	err := os.RemoveAll(dest)
	check(err)*/
}