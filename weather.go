package weather

import (
	"encoding/json"
	"errors"
	"io"
)

// Main is part of Weather api response
type Main struct {
	Temp float64
}

// Forcast is part of Weather api response
type Forcast struct {
	Main Main
}

// Weather is part of Weather api response
type Weather struct {
	Cod  string
	List []Forcast
}

type Request struct {
	ZipCode string
}

func ConvertOurRequestStructToOpenApiRequest(request Request, apiKey string) (string, error) {
	if request.ZipCode == "" || apiKey == "" {
		return "", errors.New("please specify zipcode and apiKey")
	}
	return "http://api.openweathermap.org/data/2.5/forecast?zip="+ request.ZipCode + "&appid=" + apiKey, nil
}

func ConvertWeatherOpenApiResponseToStruct(r io.Reader) (Weather, error) {
	weatherResponse := &Weather{}
	err := json.NewDecoder(r).Decode(weatherResponse)
	if err != nil {
		return Weather{}, err
	}
	return *weatherResponse, nil
}