package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"weather"
)

func main() {
	emoji := flag.Bool("emoji", false, "emoji mod")
	flag.Parse()

	APIKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if APIKey == "" {
		log.Fatal("OPENWEATHERMAP_API_KEY environment variable must be set")
	}
	location := flag.Arg(0)
	if location == "" {
		fmt.Printf("Usage: %s LOCATION\n\nExample: %s London\n", os.Args[0], os.Args[0])
		os.Exit(1)
	}
	fmt.Printf("requesting weather for %q\n", location)
	summary, temp, err := weather.Conditions(location, APIKey)
	if err != nil {
		log.Fatal(err)
	}

	if *emoji {
		emoji := weather.Emoji(summary)
		fmt.Printf("%s %.1fC\n", emoji, temp)
	} else {
		fmt.Printf("%s %.1fC\n", summary, temp)
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage: weather LOCATION (for example, 'weather London')")
	}
}
