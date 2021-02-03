// +build integration

package weather_test

import (
	"os"
	"testing"
	"weather"
)

func TestGetData(t *testing.T) {
	APIKey := os.Getenv("OPENWEATHERAPI")
	location := "London"
	data, err := weather.GetData(APIKey, location)
	if err != nil {
		t.Fatal(err)
	}
	_, err = weather.Decode(data)
	if err != nil {
		t.Fatal(err)
	}
}
