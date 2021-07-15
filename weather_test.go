package weather_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"weather"

	"github.com/google/go-cmp/cmp"
)

func TestReturnGoodErrorFor404(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		http.Error(writer, "wrong city, probably", http.StatusNotFound)
	}))
	service, _ := weather.New(
		weather.WithAPIKey("fakeAPIKey"),
		weather.WithBaseURL(server.URL),
		weather.WithHTTPClient(server.Client()), // httptest self signs it's certs.  This client will trust httptest servers.
	)
	_, err := service.GetWeather("fake city, AR")
	if err == nil {
		t.Fatal("weather service should fail if it receives a 404 from weather server")
	}
	s := "no weather data found for fake city, AR"
	if err.Error() != s {
		t.Fatalf("Bad error message, want: %s, got %s", s, err.Error())
	}
}

func TestNewWithEmptyAPIKeyReturnsError(t *testing.T) {
	_, err := weather.New(weather.WithAPIKey(""))
	if err == nil {
		t.Error("Empty api key should error")
	}
}

func TestEndToEnd(t *testing.T) {
	type testCase struct {
		name            string
		location        string
		weatherFile     string
		expectedURL     string
		expectedWeather weather.CurrentWeather
	}
	//goland:noinspection SpellCheckingInspection
	testCases := []testCase{
		{
			name:        "London",
			location:    "London, UK",
			weatherFile: "testData/weather.json",
			expectedURL: "/weather?q=London%2C+UK&appid=fakeAPIKey&units=imperial",
			expectedWeather: weather.CurrentWeather{
				Temp:    56.21,
				Summary: "Clouds",
			},
		},
		{
			name:        "Millbrae",
			location:    "Millbrae, CA, USA",
			weatherFile: "testData/weather_millbrae.json",
			expectedURL: "/weather?q=Millbrae%2C+CA%2C+USA&appid=fakeAPIKey&units=imperial",
			expectedWeather: weather.CurrentWeather{
				Temp:    70.8,
				Summary: "Sunny",
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			server := httptest.NewTLSServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				if request.RequestURI != tc.expectedURL {
					t.Fatalf("got \n%s\n but expected \n%s", request.RequestURI, tc.expectedURL)
				}

				response, err := os.Open(tc.weatherFile)
				if err != nil {
					t.Fatal(err)
				}
				_, err = io.Copy(writer, response)
				if err != nil {
					t.Fatal(err)
				}
			}))
			service, _ := weather.New(
				weather.WithAPIKey("fakeAPIKey"),
				weather.WithBaseURL(server.URL),
				weather.WithHTTPClient(server.Client()), // httptest self signs it's certs.  This client will trust httptest servers.
			)
			response, err := service.GetWeather(tc.location)
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(tc.expectedWeather, response) {
				t.Fatal(cmp.Diff(tc.expectedWeather, response))
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
		service, _ := weather.New(
			weather.WithAPIKey("fakeAPIKey"),
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
	weatherService, _ := weather.New(weather.WithAPIKey("fakeApiKey"))
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

func TestGetLocationFromProgramArgs(t *testing.T) {
	service, _ := weather.New(weather.WithAPIKey("fakeAPIKey"))
	want := "Dallas, TX, USA\n"
	got := service.GetLocation([]string{"Dallas,", "TX,", "USA"})
	if want != got {
		t.Fatalf("Want: %q, got: %q", want, got)
	}
}

func TestGetLocationFromStdIn(t *testing.T) {
	want := "Dallas, TX, USA\n"
	writer := &bytes.Buffer{}
	service, _ := weather.New(
		weather.WithAPIKey("fakeAPIKey"),
		weather.WithReader(strings.NewReader(want)),
		weather.WithWriter(writer),
	)
	got := service.GetLocation([]string{})
	if want != got {
		t.Fatalf("Want: %q, got: %q", want, got)
	}
	want = "Enter in a location to get it's weather\n"
	got = writer.String()
	if want != got {
		t.Fatalf("Want: %q, got: %q", want, got)
	}
}
