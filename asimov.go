package main

import (
	"github.com/afeldman/asimov/cmd"
	"log"
)

func main() {
	if err := cmd.Asimov.Execute(); err != nil {
		log.Fatal(err)
	}
}
