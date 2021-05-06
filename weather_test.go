package weather_test

import (
	"os"
	"testing"
	"weather"

	"github.com/google/go-cmp/cmp"
)

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