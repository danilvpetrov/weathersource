package deps

import (
	"github.com/danilvpetrov/weathersource/collection"
	"github.com/danilvpetrov/weathersource/internal/envvar"
)

// ProvideClient provides a client to retrieve meteorological data from
// third-party API.
func ProvideClient() (*collection.Client, error) {
	token, err := envvar.StringRequired("WEATHER_DARKSKY_TOKEN")
	if err != nil {
		return nil, err
	}

	lat, err := envvar.Float64Required("WEATHER_LOC_LATITUDE")
	if err != nil {
		return nil, err
	}

	lon, err := envvar.Float64Required("WEATHER_LOC_LONGITUDE")
	if err != nil {
		return nil, err
	}

	cli := collection.NewClient(
		token,
		collection.WithLatAndLong(lat, lon),
	)

	return cli, nil
}
