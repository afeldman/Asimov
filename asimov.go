package main

import (
	"github.com/afeldman/Asimov/cmd"
	"log"
)

func main() {
	if err := cmd.Asimov.Execute(); err != nil {
		log.Fatal(err)
	}
}
