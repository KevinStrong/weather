package weather_test

import (
	"os"
	"testing"

	"weather"

	"github.com/google/go-cmp/cmp"
)

func TestConvertingWeatherRequestIntoOpenApiRequest(t *testing.T) {
	input := weather.Request{
		ZipCode: "75080",
	}
	ourUrl, err := weather.ConvertOurRequestStructToOpenApiRequest(input, "fakeApiKey")
	if err != nil {
		t.Fatal(err)
	}
	wantUrl := "http://api.openweathermap.org/data/2.5/forecast?zip=75080&appid=fakeApiKey"
	if ourUrl != wantUrl {
		diff := cmp.Diff(wantUrl, ourUrl)
		t.Fatal(diff)
	}
}

func TestValidatingInputsForWeatherRequest(t *testing.T) {
	emptyInput := weather.Request{
		ZipCode: "",
	}
	validInput := weather.Request{
		ZipCode: "75080",
	}
	_, err := weather.ConvertOurRequestStructToOpenApiRequest(emptyInput, "fakeApiKey")
	if err == nil {
		t.Fatal("Should error when zipcode is empty string")
	}
	_, err = weather.ConvertOurRequestStructToOpenApiRequest(validInput, "")
	if err == nil {
		t.Fatal("Should error when apiKey is empty string")
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
