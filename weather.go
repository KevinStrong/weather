package weather

import (
	"encoding/json"
	"io"
)

type Main struct {
	Temp float64
}

type Forcast struct {
	Main Main
}

type Weather struct {
	Cod  string
	List []Forcast
}

func ConvertWeatherOpenApiResponseToStruct(r io.Reader) (Weather, error) {
	weatherResponse := &Weather{}
	err := json.NewDecoder(r).Decode(weatherResponse)
	if err != nil {
		return Weather{}, err
	}
	return *weatherResponse, nil
}