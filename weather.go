package weather

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	weatherAPI "weather/gen"
)

// CurrentWeather is part of weather api response
type CurrentWeather struct {
	Temp    float64
	Summary string
}

type Service struct {
	APIKey   string
	baseURL  string
	client   *http.Client
	readIn   io.Reader
	writeOut io.Writer
}

func (s *Service) GetWeather(location string) (CurrentWeather, error) {
	targetURL := s.MakeURL(location)

	response, err := s.client.Get(targetURL)

	if err != nil {
		return CurrentWeather{}, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer response.Body.Close()
	if response.StatusCode == http.StatusNotFound {
		return CurrentWeather{}, fmt.Errorf("no weather data found for %s", location)
	}
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

func New(opts ...Option) (*Service, error) {
	service := &Service{
		baseURL: "https://api.openweathermap.org/data/2.5",
		APIKey:  "",
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
		readIn:   os.Stdin,
		writeOut: os.Stdout,
	}
	for _, o := range opts {
		o(service)
	}
	if service.APIKey == "" {
		return nil, errors.New("API key must not be empty")
	}
	return service, nil
}

func WithAPIKey(APIKey string) Option {
	return func(s *Service) {
		s.APIKey = APIKey
	}
}

func WithConfigFromEnv() Option {
	return func(s *Service) {
		s.APIKey = os.Getenv("WEATHER_API")
	}
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

func WithReader(reader io.Reader) Option {
	return func(s *Service) {
		s.readIn = reader
	}
}

func WithWriter(writer io.Writer) Option {
	return func(s *Service) {
		s.writeOut = writer
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

func (s *Service) getLocationFromTerminal() string {
	_, _ = fmt.Fprint(s.writeOut, "Enter in a location to get it's weather\n")
	reader := bufio.NewReader(s.readIn)
	text, _ := reader.ReadString('\n')
	return text
}

func (s *Service) GetLocation(args []string) string {
	var location string
	if len(args) > 0 {
		// Trailing \n is required for USA locations if only the city and state are specified.
		location = strings.Join(args, " ") + "\n"
	} else {
		location = s.getLocationFromTerminal()
	}
	return location
}
