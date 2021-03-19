package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
	URL string
	APIKey string
}

func NewClient(APIKey string) Client {
	return Client{
		HTTPClient: http.DefaultClient,
		URL: "https://api.openweathermap.org",
		APIKey: APIKey,
	}
}

type APIResponse struct{
	Weather []struct{
		Main string
	}
	Main struct{
		Temp float64
	}
}

func Decode(data []byte) (summary string, temp float64, err error) {
	var result APIResponse
	err = json.Unmarshal(data, &result)
	if err != nil {
		return "", 0, err
	}
	return result.Weather[0].Main, result.Main.Temp-273.15, nil
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
	client := NewClient(APIKey)
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
	"Sunny": "☀️",
	"Clouds": "☁️",
}

func Emoji(input string) string {
	return emoji[input]
}
