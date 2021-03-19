package main

import (
	"fmt"
	"log"
	"os"
	"weather"
)

func main() {
	APIKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if APIKey == "" {
		log.Fatal("OPENWEATHERMAP_API_KEY environment variable must be set")
	}
	if len(os.Args) < 2 {
		log.Fatal("Usage: weather LOCATION (for example, 'weather London')")
	}
	location := os.Args[1]
	summary, temp, err := weather.Conditions(location, APIKey)
	if err != nil {
		log.Fatal(err)
	}
	emoji := weather.Emoji(summary)
	fmt.Printf("%s %.1fC\n", emoji, temp)
}
