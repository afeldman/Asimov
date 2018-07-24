package robot

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	sha1 "github.com/afeldman/go-util/crypto/sha1"
	ftp "github.com/afeldman/go-util/net/ftp"
)

type Robot struct {
	Name string
	Host string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func InitRobot(name string, ip string) *Robot {
	r := Robot{Name: name, Host: ip}
	return &r
}

func (r *Robot) SetName(name string) {
	r.Name = name
}

func (r *Robot) SetHost(host string) {
	r.Host = host

}

func (r *Robot) ToJson() string {
	b, err := json.Marshal(r)
	check(err)
	return string(b)
}

func FromJson(str string) Robot {
	var r Robot
	bytes := []byte(str)

	err := json.Unmarshal(bytes, &r)
	check(err)
	return r
}

func (r *Robot) Backup(filter func(filename string) bool, destination string, wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.Now()

	fmt.Println("Backing up: ", r.Name, "\t at host: ", r.Host)
	dirname := destination + "/" + r.Name
	err := os.MkdirAll(dirname, os.ModePerm)
	check(err)

	c := ftp.NewConnection(r.Host, "21")
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
		err := c.Download(file, dirname)
		if err != nil {
			errorList = append(errorList, err)
		}
	}

	if len(errorList) > 0 {
		for _, err := range errorList {
			log.Println(err)
		}
	}

	log.Printf("Finished backing up %s in %v\n", r.Name, time.Since(t))

	//hashing
	var hash sha1.FileHash
	hash.Hash(dirname)
	hashfile, err := os.Create(dirname + "/name.hash")
	check(err)
	var tmp_str string
	defer hashfile.Close()
	for _, node := range hash.Nodes {
		tmp_str = fmt.Sprintf("%s\t%x\n", node.Path, node.Hash)
		_, err := hashfile.WriteString(tmp_str)
		check(err)
	}
	hashfile.Sync()
}

func byteToGoString(c []byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}
