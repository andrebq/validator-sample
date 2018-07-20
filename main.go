package main

import (
	"log"
	"os"
)

func main() {
	vl := (Project{
		Ops: Ops{},
	}).Validate()

	if vl != nil {
		log.Printf("%v", vl)
		os.Exit(vl.Level().ExitCode())
	}
}
