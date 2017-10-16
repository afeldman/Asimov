package main

import (
	"asimov/cmd"
	"log"
)

func main() {
	if err := cmd.Asimov.Execute(); err != nil {
		log.Fatal(err)
	}
}
