package main

import (
	"log"

	"github.com/MydroX/short-circuit/pkg/env"
	"github.com/MydroX/short-circuit/src/data"
)

func main() {
	log.Println("Starting program...")

	env.Load()

	log.Println("Loading data...")

	data.Load()

	log.Println("Adding properties...")

	//TODO

	log.Println("Creating points...")

	//TODO
}
