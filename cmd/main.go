package main

import (
	"log"
	"os"
	"weather"
)

func main() {
	log.Fatal(weather.RunCLI(os.Args[1:]))
}

