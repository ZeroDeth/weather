package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"weather"
)

func main() {
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=London&appid=3b814c61996538f2e8a2b921e23bbb0a")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Unexpected HTTP response status: %q\n", resp.Status)
		os.Exit(1)
	}
	defer resp.Body.Close()
	weather, err := weather.Decode(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("The weather is: %s\n", weather)
}