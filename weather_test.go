package weather_test

import (
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"weather"
)

var rawJSON = `{"coord":{"lon":-0.1257,"lat":51.5085},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"base":"stations","main":{"temp":280.57,"feels_like":276.26,"temp_min":279.26,"temp_max":281.48,"pressure":994,"humidity":65},"visibility":10000,"wind":{"speed":3.6,"deg":250},"clouds":{"all":0},"dt":1611321852,"sys":{"type":1,"id":1414,"country":"GB","sunrise":1611301933,"sunset":1611333099},"timezone":0,"id":2643743,"name":"London","cod":200}`

func CloseEnough(a, b float64) bool {
	return math.Abs(a-b) < 0.001
}

func TestDecode(t *testing.T) {
	t.Parallel()
	wantSummary := "Clear"
	wantTemp := 7.42
	gotSummary, gotTemp, err := weather.Decode([]byte(rawJSON))
	if err != nil {
		t.Fatal(err)
	}
	if wantSummary != gotSummary {
		t.Errorf("want %q, got %q", wantSummary, gotSummary)
	}
	if !CloseEnough(wantTemp, gotTemp) {
		t.Errorf("want %f, got %f", wantTemp, gotTemp)
	}
}

func TestGetData(t *testing.T) {
	t.Parallel()
	APIKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	ts := httptest.NewTLSServer(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		wantURL := "/data/2.5/weather?q=London&appid="
		gotURL := fmt.Sprintf("%s?%s", r.URL.Path, r.URL.RawQuery)
		if !strings.HasPrefix(gotURL, wantURL) {
			t.Errorf("want at least %q, got %q", wantURL, gotURL)
		}
		fmt.Fprintln(w, rawJSON)
	}))
	location := "London"
	client := weather.NewClient(APIKey)
	client.HTTPClient = ts.Client()
	client.URL = ts.URL
	data, err := client.GetData(location)
	if err != nil {
		t.Fatal(err)
	}
	wantSummary := "Clear"
	wantTemp := 7.42
	gotSummary, gotTemp, err := weather.Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	if wantSummary != gotSummary {
		t.Errorf("want %q, got %q", wantSummary, gotSummary)
	}
	if !CloseEnough(wantTemp, gotTemp) {
		t.Errorf("want %f, got %f", wantTemp, gotTemp)
	}
}

func TestEmoji (t *testing.T) {
  t.Parallel()
  input := "Sunny"
  want := "☀️"
  got := weather.Emoji(input)

  if want != got {
    t.Errorf("want %q, got %q", want, got)
  }
}

