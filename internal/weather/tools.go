package weather

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"

	_http "github.com/kieranroneill/oarbit/internal/platform/http"
)

//==========================================
// get weather by location tool
//==========================================

type CurrentForecast struct {
	RelativeHumidity2m float64 `json:"relative_humidity_2m"`
	Temperature2m      float64 `json:"temperature_2m"`
	WeatherCode        int     `json:"weather_code"`
}

type CurrentWeatherResult struct {
	RelativeHumidity2m float64 `json:"relativeHumidity2m"`
	Temperature2m      float64 `json:"temperature2m"`
	WeatherCode        int     `json:"weatherCode"`
}

type ForecastResponseBody struct {
	Current CurrentForecast `json:"current"`
}

type GeoLocation struct {
	Name      string  `json:"name"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeoLocationResponseBody struct {
	Results []GeoLocation `json:"results"`
}

type GetWeatherByLocationParams struct {
	City        string `json:"city" jsonschema:"description=The city to look up,required"`
	CountryCode string `json:"countryCode,omitempty" jsonschema:"description=Optional ISO country code, for example GB"`
}

type GetWeatherByLocationResult struct {
	City      string               `json:"city"`
	Country   string               `json:"country,omitempty"`
	Current   CurrentWeatherResult `json:"current"`
	Error     string               `json:"error,omitempty"`
	Latitude  float64              `json:"latitude"`
	Longitude float64              `json:"longitude"`
}

func getWeatherByLocation(
	ctx context.Context,
	params GetWeatherByLocationParams,
) (*GetWeatherByLocationResult, error) {
	var geoLocationResponseBody GeoLocationResponseBody
	var weatherResponseBody ForecastResponseBody

	httpClient := _http.NewClient(nil)
	geoParams := url.Values{}
	geoParams.Set("name", params.City)
	geoParams.Set("count", "1")
	geoParams.Set("language", "en")
	geoParams.Set("format", "json")

	if params.CountryCode != "" {
		geoParams.Set("countryCode", params.CountryCode)
	}

	_, err := httpClient.DoGetRequest(ctx, "https://geocoding-api.open-meteo.com/v1/search?"+geoParams.Encode(), nil, &geoLocationResponseBody)

	if err != nil {
		slog.Error("geocoding request failed", "error", err)

		return nil, fmt.Errorf("geocoding request failed: %w", err)
	}

	if len(geoLocationResponseBody.Results) == 0 {
		return &GetWeatherByLocationResult{
			Error: fmt.Sprintf("could not find location: %s", params.City),
		}, nil
	}

	location := geoLocationResponseBody.Results[0]
	weatherParams := url.Values{}
	weatherParams.Set("latitude", fmt.Sprintf("%f", location.Latitude))
	weatherParams.Set("longitude", fmt.Sprintf("%f", location.Longitude))
	weatherParams.Set(
		"current",
		"temperature_2m,relative_humidity_2m,weather_code",
	)
	_, err = httpClient.DoGetRequest(ctx, "https://api.open-meteo.com/v1/forecast?"+weatherParams.Encode(), nil, &weatherResponseBody)

	if err != nil {
		return nil, fmt.Errorf("weather request failed: %w", err)
	}

	return &GetWeatherByLocationResult{
		City:      location.Name,
		Country:   location.Country,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Current: CurrentWeatherResult{
			RelativeHumidity2m: weatherResponseBody.Current.RelativeHumidity2m,
			Temperature2m:      weatherResponseBody.Current.Temperature2m,
			WeatherCode:        weatherResponseBody.Current.WeatherCode,
		},
	}, nil
}

func NewGetWeatherByLocationTool() (tool.InvokableTool, error) {
	return utils.InferTool(
		"get_weather_by_location",
		"Get current weather for a city, optionally constrained by an ISO country code.",
		getWeatherByLocation,
	)
}
