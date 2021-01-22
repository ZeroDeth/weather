package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type APIResponse struct{
	Weather []struct{
		Main string
	}
	Main struct{
		Temp float64
	}
}

func Decode(data []byte) (string, error) {
	var result APIResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %.1fC", result.Weather[0].Main, result.Main.Temp-273.15), nil
}

func GetData(APIKey, location string) ([]byte, error) {
	URL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", location, APIKey)
	resp, err := http.Get(URL)
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