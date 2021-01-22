package weather_test

import (
	"testing"
	"weather"

	"github.com/google/go-cmp/cmp"
)

var rawJSON =`{"coord":{"lon":-0.1257,"lat":51.5085},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"base":"stations","main":{"temp":280.57,"feels_like":276.26,"temp_min":279.26,"temp_max":281.48,"pressure":994,"humidity":65},"visibility":10000,"wind":{"speed":3.6,"deg":250},"clouds":{"all":0},"dt":1611321852,"sys":{"type":1,"id":1414,"country":"GB","sunrise":1611301933,"sunset":1611333099},"timezone":0,"id":2643743,"name":"London","cod":200}`

func TestDecode(t *testing.T) {
	want := "Clear 7.4C"
	got, err := weather.Decode([]byte(rawJSON))
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestGetData(t *testing.T) {
	APIKey := "XXX"
	location := "London"
	got, err := weather.GetData(APIKey, location)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(rawJSON, string(got)) {
		t.Error(cmp.Diff(rawJSON, string(got)))
	}
}