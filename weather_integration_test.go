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
	client := weather.NewClient(APIKey)
	data, err := client.GetData(location)
	if err != nil {
		t.Fatal(err)
	}
	_, err = weather.Decode(data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConditionsIntegration(t *testing.T) {
	t.Parallel()
	APIKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	conditions, err := weather.Conditions("London", APIKey)
	if err != nil {
		t.Fatal(err)
	}
	if conditions == "" {
		t.Error("conditions was empty")
	}
}