package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"weather"
)

func main() {
	API := os.Getenv("OPENWEATHERAPI")
	URL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=London&appid=%s", API)

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Unexpected HTTP response status: %q\n", resp.Status)
		os.Exit(1)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	weather, err := weather.Decode(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("The weather is: %s\n", weather)
}
