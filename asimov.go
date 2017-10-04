package main

import (
       "log"

	"asimov/cmd"
)

func main() {

     if err := cmd.Asimov.Execute(); err != nil {
     	log.Fatal(err)
     }
}
