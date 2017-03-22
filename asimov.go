package main

import (
    "fmt"
	"robot"
    "github.com/afeldman/goftp"
)

func main() {
    var err error
    var ftp *goftp.FTP

    // For debug messages: goftp.ConnectDbg("ftp.server.com:21")
    if ftp, err = goftp.Connect("127.0.0.1:21"); err != nil {
        panic(err)
    }

    defer ftp.Close()
    fmt.Println("Successfully connected to", ftp)

    var curpath string
    if curpath, err = ftp.Pwd(); err != nil {
        panic(err)
    }

    fmt.Printf("Current path: %s", curpath)

    // Get directory listing
    var files []string
    if files, err = ftp.List(""); err != nil {
        panic(err)
    }
    fmt.Println("\nDirectory listing:\n", files)

}