package main

import (
	"log"
	"npd/prototype"
)

func main() {
	npd := &prototype.Npd{}

	log.Println("Starting server")

	log.Fatal(npd.Start())
}
