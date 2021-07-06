package weather_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"weather"

	"github.com/google/go-cmp/cmp"
)

func TestEndToEnd(t *testing.T) {
	type testCase struct {
		name            string
		location        string
		expectedURL     string
		expectedWeather weather.CurrentWeather
	}
	//goland:noinspection SpellCheckingInspection
	testCases := []testCase{
		{
			name:        "London",
			location:    "London, UK",
			expectedURL: "/weather?q=London%2C+UK&appid=fakeAPIKey&units=imperial",
			expectedWeather: weather.CurrentWeather{
				Temp:    56.21,
				Summary: "Clouds",
			},
		},
		{
			name:        "Millbrae",
			location:    "Millbrae, CA, USA",
			expectedURL: "/weather?q=Millbrae%2C+CA%2C+USA&appid=fakeAPIKey&units=imperial",
			expectedWeather: weather.CurrentWeather{
				Temp:    70.8,
				Summary: "Sunny",
			},
		},
	}
	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			server := httptest.NewTLSServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				if request.RequestURI != test.expectedURL {
					t.Fatalf("got \n%s\n but expected \n%s", request.RequestURI, test.expectedURL)
				}
				var weatherFile = "testData/weather.json"
				if strings.Contains(request.RequestURI, "Millbrae") {
					weatherFile = "testData/weather_millbrae.json"
				}

				response, err := os.Open(weatherFile)
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
			response, err := service.GetWeather(test.location)
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(test.expectedWeather, response) {
				t.Fatal(cmp.Diff(test.expectedWeather, response))
			}
		})
	}
}

func TestRequestGeneration(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name            string
		input           string
		expectedRequest string
	}
	tests := []testCase{{
		name:            "London",
		input:           "London, UK",
		expectedRequest: "https://api.openweathermap.org/data/2.5/weather?q=London%2C+UK&appid=fakeAPIKey&units=imperial",
	},
		{
			name:            "Millbrae",
			input:           "Millbrae, CA",
			expectedRequest: "https://api.openweathermap.org/data/2.5/weather?q=Millbrae%2C+CA&appid=fakeAPIKey&units=imperial",
		},
	}
	for _, test := range tests {
		test := test
		service := weather.New("fakeAPIKey",
			weather.WithBaseURL("https://api.openweathermap.org/data/2.5"),
		)
		url := service.MakeURL(test.input)
		if !cmp.Equal(url, test.expectedRequest) {
			t.Fatalf("got \n%s\n but expected \n%s", url, test.expectedRequest)
		}
	}
}

func TestCreateWeatherService(t *testing.T) {
	t.Parallel()
	weatherService := weather.New("fakeApiKey")
	if weatherService.APIKey != "fakeApiKey" {
		t.Fatal("Failed to create weather service with provided APIKey")
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
