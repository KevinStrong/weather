// Package Weather provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.0 DO NOT EDIT.
package Weather

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

const (
	App_idScopes = "app_id.Scopes"
)

// Defines values for Lang.
const (
	Ar Lang = "ar"

	Bg Lang = "bg"

	Ca Lang = "ca"

	Cz Lang = "cz"

	De Lang = "de"

	El Lang = "el"

	En Lang = "en"

	Es Lang = "es"

	Fa Lang = "fa"

	Fi Lang = "fi"

	Fr Lang = "fr"

	Gl Lang = "gl"

	Hr Lang = "hr"

	Hu Lang = "hu"

	It Lang = "it"

	Ja Lang = "ja"

	Kr Lang = "kr"

	La Lang = "la"

	Lt Lang = "lt"

	Mk Lang = "mk"

	Nl Lang = "nl"

	Pl Lang = "pl"

	Pt Lang = "pt"

	Ro Lang = "ro"

	Ru Lang = "ru"

	Se Lang = "se"

	Sk Lang = "sk"

	Sl Lang = "sl"

	Tr Lang = "tr"

	Ua Lang = "ua"

	Vi Lang = "vi"

	ZhCn Lang = "zh_cn"

	ZhTw Lang = "zh_tw"
)

// Defines values for Mode.
const (
	Html Mode = "html"

	Json Mode = "json"

	Xml Mode = "xml"
)

// Defines values for Units.
const (
	Imperial Units = "imperial"

	Metric Units = "metric"

	Standard Units = "standard"
)

// N200 defines model for 200.
type N200 struct {

	// Internal parameter
	Base   *string `json:"base,omitempty"`
	Clouds *Clouds `json:"clouds,omitempty"`

	Coord *Coord `json:"coord,omitempty"`

	// Time of data calculation, unix, UTC
	Dt *int32 `json:"dt,omitempty"`

	// City ID
	Id   *int32  `json:"id,omitempty"`
	Main *Main   `json:"main,omitempty"`
	Name *string `json:"name,omitempty"`
	Rain *Rain   `json:"rain,omitempty"`
	Snow *Snow   `json:"snow,omitempty"`
	Sys  *Sys    `json:"sys,omitempty"`

	// Visibility, meter
	Visibility *int `json:"visibility,omitempty"`

	// (more info Weather condition codes)
	Weather *[]Weather `json:"weather,omitempty"`
	Wind    *Wind      `json:"wind,omitempty"`
}

// Clouds defines model for Clouds.
type Clouds struct {

	// Cloudiness, %
	All *int32 `json:"all,omitempty"`
}

// Coord defines model for Coord.
type Coord struct {

	// City geo location, latitude
	Lat *float32 `json:"lat,omitempty"`

	// City geo location, longitude
	Lon *float32 `json:"lon,omitempty"`
}

// Main defines model for Main.
type Main struct {

	// Atmospheric pressure on the ground level, hPa
	GrndLevel *float64 `json:"grnd_level,omitempty"`

	// Humidity, %
	Humidity *int32 `json:"humidity,omitempty"`

	// Atmospheric pressure (on the sea level, if there is no sea_level or grnd_level data), hPa
	Pressure *int32 `json:"pressure,omitempty"`

	// Atmospheric pressure on the sea level, hPa
	SeaLevel *float64 `json:"sea_level,omitempty"`

	// Temperature. Unit Default: Kelvin, Metric: Celsius, Imperial: Fahrenheit.
	Temp *float64 `json:"temp,omitempty"`

	// Maximum temperature at the moment. This is deviation from current temp that is possible for large cities and megalopolises geographically expanded (use these parameter optionally). Unit Default: Kelvin, Metric: Celsius, Imperial: Fahrenheit.
	TempMax *float64 `json:"temp_max,omitempty"`

	// Minimum temperature at the moment. This is deviation from current temp that is possible for large cities and megalopolises geographically expanded (use these parameter optionally). Unit Default: Kelvin, Metric: Celsius, Imperial: Fahrenheit.
	TempMin *float64 `json:"temp_min,omitempty"`
}

// Rain defines model for Rain.
type Rain struct {

	// Rain volume for the last 3 hours
	N3h *int32 `json:"3h,omitempty"`
}

// Snow defines model for Snow.
type Snow struct {

	// Snow volume for the last 3 hours
	N3h *float32 `json:"3h,omitempty"`
}

// Sys defines model for Sys.
type Sys struct {

	// Country code (GB, JP etc.)
	Country *string `json:"country,omitempty"`

	// Internal parameter
	Id *int32 `json:"id,omitempty"`

	// Internal parameter
	Message *float32 `json:"message,omitempty"`

	// Sunrise time, unix, UTC
	Sunrise *int32 `json:"sunrise,omitempty"`

	// Sunset time, unix, UTC
	Sunset *int32 `json:"sunset,omitempty"`

	// Internal parameter
	Type *int32 `json:"type,omitempty"`
}

// Weather defines model for Weather.
type Weather struct {

	// Weather condition within the group
	Description *string `json:"description,omitempty"`

	// Weather icon id
	Icon *string `json:"icon,omitempty"`

	// Weather condition id
	Id *int32 `json:"id,omitempty"`

	// Group of weather parameters (Rain, Snow, Extreme etc.)
	Main *string `json:"main,omitempty"`
}

// Wind defines model for Wind.
type Wind struct {

	// Wind direction, degrees (meteorological)
	Deg *int32 `json:"deg,omitempty"`

	// Wind speed. Unit Default: meter/sec, Metric: meter/sec, Imperial: miles/hour.
	Speed *float32 `json:"speed,omitempty"`
}

// Id defines model for id.
type Id string

// Lang defines model for lang.
type Lang string

// Lat defines model for lat.
type Lat string

// Lon defines model for lon.
type Lon string

// Mode defines model for mode.
type Mode string

// Q defines model for q.
type Q string

// Units defines model for units.
type Units string

// Zip defines model for zip.
type Zip string

// CurrentWeatherDataParams defines parameters for CurrentWeatherData.
type CurrentWeatherDataParams struct {

	// **City name**. *Example: London*. You can call by city name, or by city name and country code. The API responds with a list of results that match a searching word. For the query value, type the city name and optionally the country code divided by comma; use ISO 3166 country codes.
	Q *Q `json:"q,omitempty"`

	// **City ID**. *Example: `2172797`*. You can call by city ID. API responds with exact result. The List of city IDs can be downloaded [here](http://bulk.openweathermap.org/sample/). You can include multiple cities in parameter &mdash; just separate them by commas. The limit of locations is 20. *Note: A single ID counts as a one API call. So, if you have city IDs. it's treated as 3 API calls.*
	Id *Id `json:"id,omitempty"`

	// **Latitude**. *Example: 35*. The latitude cordinate of the location of your interest. Must use with `lon`.
	Lat *Lat `json:"lat,omitempty"`

	// **Longitude**. *Example: 139*. Longitude cordinate of the location of your interest. Must use with `lat`.
	Lon *Lon `json:"lon,omitempty"`

	// **Zip code**. Search by zip code. *Example: 95050,us*. Please note if country is not specified then the search works for USA as a default.
	Zip *Zip `json:"zip,omitempty"`

	// **Units**. *Example: imperial*. Possible values: `standard`, `metric`, and `imperial`. When you do not use units parameter, format is `standard` by default.
	Units *CurrentWeatherDataParamsUnits `json:"units,omitempty"`

	// **Language**. *Example: en*. You can use lang parameter to get the output in your language. We support the following languages that you can use with the corresponded lang values: Arabic - `ar`, Bulgarian - `bg`, Catalan - `ca`, Czech - `cz`, German - `de`, Greek - `el`, English - `en`, Persian (Farsi) - `fa`, Finnish - `fi`, French - `fr`, Galician - `gl`, Croatian - `hr`, Hungarian - `hu`, Italian - `it`, Japanese - `ja`, Korean - `kr`, Latvian - `la`, Lithuanian - `lt`, Macedonian - `mk`, Dutch - `nl`, Polish - `pl`, Portuguese - `pt`, Romanian - `ro`, Russian - `ru`, Swedish - `se`, Slovak - `sk`, Slovenian - `sl`, Spanish - `es`, Turkish - `tr`, Ukrainian - `ua`, Vietnamese - `vi`, Chinese Simplified - `zh_cn`, Chinese Traditional - `zh_tw`.
	Lang *CurrentWeatherDataParamsLang `json:"lang,omitempty"`

	// **Mode**. *Example: html*. Determines format of response. Possible values are `xml` and `html`. If mode parameter is empty the format is `json` by default.
	Mode *CurrentWeatherDataParamsMode `json:"mode,omitempty"`
}

// CurrentWeatherDataParamsUnits defines parameters for CurrentWeatherData.
type CurrentWeatherDataParamsUnits string

// CurrentWeatherDataParamsLang defines parameters for CurrentWeatherData.
type CurrentWeatherDataParamsLang string

// CurrentWeatherDataParamsMode defines parameters for CurrentWeatherData.
type CurrentWeatherDataParamsMode string

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// CurrentWeatherData request
	CurrentWeatherData(ctx context.Context, params *CurrentWeatherDataParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) CurrentWeatherData(ctx context.Context, params *CurrentWeatherDataParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCurrentWeatherDataRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewCurrentWeatherDataRequest generates requests for CurrentWeatherData
func NewCurrentWeatherDataRequest(server string, params *CurrentWeatherDataParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/weather")
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	queryValues := queryURL.Query()

	if params.Q != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "q", runtime.ParamLocationQuery, *params.Q); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Id != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "id", runtime.ParamLocationQuery, *params.Id); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Lat != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "lat", runtime.ParamLocationQuery, *params.Lat); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Lon != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "lon", runtime.ParamLocationQuery, *params.Lon); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Zip != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "zip", runtime.ParamLocationQuery, *params.Zip); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Units != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "units", runtime.ParamLocationQuery, *params.Units); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Lang != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "lang", runtime.ParamLocationQuery, *params.Lang); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Mode != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "mode", runtime.ParamLocationQuery, *params.Mode); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// CurrentWeatherData request
	CurrentWeatherDataWithResponse(ctx context.Context, params *CurrentWeatherDataParams, reqEditors ...RequestEditorFn) (*CurrentWeatherDataResponse, error)
}

type CurrentWeatherDataResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *N200
}

// Status returns HTTPResponse.Status
func (r CurrentWeatherDataResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CurrentWeatherDataResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// CurrentWeatherDataWithResponse request returning *CurrentWeatherDataResponse
func (c *ClientWithResponses) CurrentWeatherDataWithResponse(ctx context.Context, params *CurrentWeatherDataParams, reqEditors ...RequestEditorFn) (*CurrentWeatherDataResponse, error) {
	rsp, err := c.CurrentWeatherData(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCurrentWeatherDataResponse(rsp)
}

// ParseCurrentWeatherDataResponse parses an HTTP response from a CurrentWeatherDataWithResponse call
func ParseCurrentWeatherDataResponse(rsp *http.Response) (*CurrentWeatherDataResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CurrentWeatherDataResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest N200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Call current weather data for one location
	// (GET /weather)
	CurrentWeatherData(ctx echo.Context, params CurrentWeatherDataParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// CurrentWeatherData converts echo context to params.
func (w *ServerInterfaceWrapper) CurrentWeatherData(ctx echo.Context) error {
	var err error

	ctx.Set(App_idScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params CurrentWeatherDataParams
	// ------------- Optional query parameter "q" -------------

	err = runtime.BindQueryParameter("form", true, false, "q", ctx.QueryParams(), &params.Q)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter q: %s", err))
	}

	// ------------- Optional query parameter "id" -------------

	err = runtime.BindQueryParameter("form", true, false, "id", ctx.QueryParams(), &params.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Optional query parameter "lat" -------------

	err = runtime.BindQueryParameter("form", true, false, "lat", ctx.QueryParams(), &params.Lat)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter lat: %s", err))
	}

	// ------------- Optional query parameter "lon" -------------

	err = runtime.BindQueryParameter("form", true, false, "lon", ctx.QueryParams(), &params.Lon)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter lon: %s", err))
	}

	// ------------- Optional query parameter "zip" -------------

	err = runtime.BindQueryParameter("form", true, false, "zip", ctx.QueryParams(), &params.Zip)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter zip: %s", err))
	}

	// ------------- Optional query parameter "units" -------------

	err = runtime.BindQueryParameter("form", true, false, "units", ctx.QueryParams(), &params.Units)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter units: %s", err))
	}

	// ------------- Optional query parameter "lang" -------------

	err = runtime.BindQueryParameter("form", true, false, "lang", ctx.QueryParams(), &params.Lang)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter lang: %s", err))
	}

	// ------------- Optional query parameter "mode" -------------

	err = runtime.BindQueryParameter("form", true, false, "mode", ctx.QueryParams(), &params.Mode)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter mode: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CurrentWeatherData(ctx, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/weather", wrapper.CurrentWeatherData)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
