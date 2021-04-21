// +build integration

package weather_test

import (
	"os"
	"testing"
	"weather"
)

func TestGetDataIntegration(t *testing.T) {
	t.Parallel()
	APIKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	location := "London"
	client, err := weather.NewClient(APIKey)
	if err != nil {
		t.Fatal(err)
	}
	data, err := client.GetData(location)
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = weather.Decode(data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConditionsIntegration(t *testing.T) {
	t.Parallel()
	APIKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	summary, _, err := weather.Conditions("London", APIKey)
	if err != nil {
		t.Fatal(err)
	}
	if summary == "" {
		t.Error("summary was empty")
	}
}
