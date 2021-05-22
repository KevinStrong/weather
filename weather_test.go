package weather_test

import (
	"os"
	"testing"
	"weather"

	"github.com/google/go-cmp/cmp"
)

func TestThatWeCAnConvertOutInputsIntoWeatherRequest(t *testing.T) {
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

func TestCreateOurWeatherObject(t *testing.T) {
	weatherService := weather.New("fakeApiKey")
	if weatherService.ApiKey != "fakeApiKey" {
		t.Fatal("Failed to create weather service with provided ApiKey")
	}
}

func TestThatWeCanDecodeAWeatherResponse(t *testing.T){
	response, err := os.Open("testData/weather.json")
	if err != nil {
		t.Fatal(err)
	}
	ourStruct, err := weather.ConvertWeatherOpenApiResponseToStruct(response)
	if err != nil {
		t.Fatal(err)
	}
	want := weather.Weather{Cod: "200", List: []weather.Forcast{
		{weather.Main{Temp: 295.72}},
		{weather.Main{Temp: 285.78}},
	}}
	if !cmp.Equal(want, ourStruct) {
		diff := cmp.Diff(want, ourStruct)
		t.Fatal(diff)
	}
}