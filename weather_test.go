package weather_test

import (
	"os"
	"testing"

	"weather"

	"github.com/google/go-cmp/cmp"
)

func TestRequestGeneration(t *testing.T) {
	expected := `https://api.openweathermap.org/data/2.5/weather?q=London%2C+UK&appid=bar&units=imperial`
	input := "London, UK"
	service := weather.New("bar")
	url := service.MakeURL(input)
	if !cmp.Equal(url, expected) {
		t.Fatalf("got \n%s\n but expected \n%s", url, expected)
	}
}

func TestMillbrae(t *testing.T) {
	expected := `https://api.openweathermap.org/data/2.5/weather?q=Millbrae%2C+CA%2C+USA&appid=bar&units=imperial`
	input := "Millbrae, CA, USA"
	service := weather.New("bar")
	url := service.MakeURL(input)
	if !cmp.Equal(url, expected) {
		t.Fatalf("got \n%s\n but expected \n%s", url, expected)
	}
}

func TestCreateWeatherService(t *testing.T) {
	weatherService := weather.New("fakeApiKey")
	if weatherService.ApiKey != "fakeApiKey" {
		t.Fatal("Failed to create weather service with provided ApiKey")
	}
}

func TestThatWeCanDecodeAOpenApiResponse(t *testing.T) {
	response, err := os.Open("testData/weather.json")
	if err != nil {
		t.Fatal(err)
	}
	ourStruct, err := weather.ConvertWeatherOpenApiResponseToStruct(response)
	if err != nil {
		t.Fatal(err)
	}
	want := weather.CurrentWeather{Temp: 56.21}
	if !cmp.Equal(want, ourStruct) {
		diff := cmp.Diff(want, ourStruct)
		t.Fatal(diff)
	}
}
