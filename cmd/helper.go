package cmd

import "log"

func er(msg interface{}) {
     log.Fatal("Error: ", msg)
}