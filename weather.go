package weather

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Client struct {
	HTTPClient *http.Client
	URL        string
	APIKey     string
}

func NewClient(APIKey string) (Client, error) {
	if APIKey == "" {
		return Client{}, errors.New("API key must not be empty")
	}
	return Client{
		HTTPClient: http.DefaultClient,
		URL:        "https://api.openweathermap.org",
		APIKey:     APIKey,
	}, nil
}

type APIResponse struct {
	Weather []struct {
		Main string
	}
	Main struct {
		Temp float64
	}
}

func Decode(data []byte) (summary string, temp float64, err error) {
	var result APIResponse
	err = json.Unmarshal(data, &result)
	if err != nil {
		return "", 0, err
	}
	return result.Weather[0].Main, result.Main.Temp - 273.15, nil
}

func (c Client) GetData(location string) ([]byte, error) {
	URL := fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s", c.URL, location, c.APIKey)
	resp, err := c.HTTPClient.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected HTTP response status: %q\n", resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Conditions(location, APIKey string) (string, float64, error) {
	client, err := NewClient(APIKey)
	if err != nil {
		return "", 0, err
	}
	data, err := client.GetData(location)
	if err != nil {
		return "", 0, err
	}
	summary, temp, err := Decode(data)
	if err != nil {
		return "", 0, err
	}
	return summary, temp, nil
}

var emoji = map[string]string{
	"Sunny":  "☀️",
	"Clear":  "☀️",
	"Clouds": "☁️",
}

func Emoji(input string) string {
	return emoji[input]
}

func ParseArgs(args []string) (emojiMode, fahrenheitMode bool, location string) {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.BoolVar(&emojiMode, "emoji", false, "emoji mod")
	fs.Parse(args)
	location = fs.Arg(0)
	return emojiMode, fahrenheitMode, location
}

func RunCLI(args []string) error {
	APIKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if APIKey == "" {
		log.Fatal("OPENWEATHERMAP_API_KEY environment variable must be set")
	}
	emoji, fahrenheit, location := ParseArgs(args)
	if location == "" {
		return fmt.Errorf("Usage: weather LOCATION\n\nExample: weather London\n")
	}
	fmt.Printf("requesting weather for %q\n", location)
	summary, temp, err := Conditions(location, APIKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("fahrenheit mode", fahrenheit)
	if emoji {
		emoji := Emoji(summary)
		fmt.Printf("%s %.1fC\n", emoji, temp)
	} else {
		fmt.Printf("%s %.1fC\n", summary, temp)
	}
	return nil
}
