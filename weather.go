package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	weatherAPI "weather/gen"
)

// CurrentWeather is part of weather api response
type CurrentWeather struct {
	Temp    float64
	Summary string
}

type Service struct {
	APIKey  string
	baseURL string
	client  *http.Client
}

func (s *Service) GetWeather(location string) (CurrentWeather, error) {
	targetURL := s.MakeURL(location)

	response, err := s.client.Get(targetURL)

	if err != nil {
		return CurrentWeather{}, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return CurrentWeather{}, fmt.Errorf("got %d", response.StatusCode)
	}

	weather, err := ConvertWeatherOpenAPIResponseToStruct(response.Body)
	if err != nil {
		return CurrentWeather{}, err
	}
	return weather, nil
}

func (s *Service) MakeURL(city string) string {
	return fmt.Sprintf("%s/weather?q=%s&appid=%s&units=imperial", s.baseURL, url.QueryEscape(city), s.APIKey)
}

type Option func(*Service)

func New(apiKey string, opts ...Option) *Service {
	service := &Service{
		baseURL: "https://api.openweathermap.org/data/2.5",
		APIKey:  apiKey,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
	for _, o := range opts {
		o(service)
	}
	return service
}

func WithBaseURL(baseURL string) Option {
	return func(s *Service) {
		s.baseURL = baseURL
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(s *Service) {
		s.client = client
	}
}

func ConvertWeatherOpenAPIResponseToStruct(r io.Reader) (CurrentWeather, error) {
	weatherResponse := &weatherAPI.N200{}
	err := json.NewDecoder(r).Decode(weatherResponse)
	if err != nil {
		return CurrentWeather{}, err
	}
	response := CurrentWeather{}
	response.Temp = *weatherResponse.Main.Temp
	if len(*weatherResponse.Weather) == 0 {
		return CurrentWeather{}, fmt.Errorf("invalid weather response: %v", weatherResponse)
	}
	response.Summary = *(*weatherResponse.Weather)[0].Main
	return response, nil
}
