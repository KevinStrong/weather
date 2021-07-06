package weather_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"weather"

	"github.com/google/go-cmp/cmp"
)

func TestEndToEnd(t *testing.T) {
	t.Parallel()
	server := httptest.NewTLSServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//goland:noinspection SpellCheckingInspection
		expected := "/weather?q=Millbrae%2C+CA%2C+USA&appid=fakeAPIKey&units=imperial"
		if request.RequestURI != expected {
			t.Fatalf("got \n%s\n but expected \n%s", request.RequestURI, expected)
		}
		response, err := os.Open("testData/weather.json")
		if err != nil {
			t.Fatal(err)
		}
		_, err = io.Copy(writer, response)
		if err != nil {
			t.Fatal(err)
		}
	}))
	service := weather.New(
		"fakeAPIKey",
		weather.WithBaseURL(server.URL),
		weather.WithHTTPClient(server.Client()),
	)
	response, err := service.GetWeather("Millbrae, CA, USA")
	if err != nil {
		t.Fatal(err)
	}
	expect := weather.CurrentWeather{Temp: 56.21, Summary: "Clouds"}
	if !cmp.Equal(expect, response) {
		t.Fatal(cmp.Diff(expect, response))
	}
}

func TestRequestGeneration(t *testing.T) {
	t.Parallel()
	expected := `https://api.openweathermap.org/data/2.5/weather?q=London%2C+UK&appid=fakeAPIKey&units=imperial`
	input := "London, UK"
	service := weather.New("fakeAPIKey",
		weather.WithBaseURL("https://api.openweathermap.org/data/2.5"),
	)
	url := service.MakeURL(input)
	if !cmp.Equal(url, expected) {
		t.Fatalf("got \n%s\n but expected \n%s", url, expected)
	}
}

func TestMillbrae(t *testing.T) {
	t.Parallel()
	expected := `https://api.openweathermap.org/data/2.5/weather?q=Millbrae%2C+CA%2C+USA&appid=fakeAPIKey&units=imperial`
	input := "Millbrae, CA, USA"
	service := weather.New("fakeAPIKey",
		weather.WithBaseURL("https://api.openweathermap.org/data/2.5"),
	)
	url := service.MakeURL(input)
	if !cmp.Equal(url, expected) {
		t.Fatalf("got \n%s\n but expected \n%s", url, expected)
	}
}

func TestCreateWeatherService(t *testing.T) {
	t.Parallel()
	weatherService := weather.New("fakeApiKey")
	if weatherService.APIKey != "fakeApiKey" {
		t.Fatal("Failed to create weather service with provided ApiKey")
	}
}

func TestThatWeCanDecodeAOpenApiResponse(t *testing.T) {
	t.Parallel()
	response, err := os.Open("testData/weather.json")
	if err != nil {
		t.Fatal(err)
	}
	ourStruct, err := weather.ConvertWeatherOpenAPIResponseToStruct(response)
	if err != nil {
		t.Fatal(err)
	}
	want := weather.CurrentWeather{Temp: 56.21, Summary: "Clouds"}
	if !cmp.Equal(want, ourStruct) {
		diff := cmp.Diff(want, ourStruct)
		t.Fatal(diff)
	}
}
