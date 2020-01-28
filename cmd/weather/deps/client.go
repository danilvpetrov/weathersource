package deps

import (
	"fmt"
	"strconv"

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

	fs, err := envvar.StringRequired("WEATHER_LOC_LATITUDE")
	if err != nil {
		return nil, err
	}

	lat, err := strconv.ParseFloat(fs, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot convert WEATHER_LOC_LATITUDE to float: %w", err)
	}

	fs, err = envvar.StringRequired("WEATHER_LOC_LONGITUDE")
	if err != nil {
		return nil, err
	}

	lon, err := strconv.ParseFloat(fs, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot convert WEATHER_LOC_LONGITUDE to float: %w", err)
	}

	cli := collection.NewClient(
		token,
		collection.WithLatAndLong(lat, lon),
	)

	return cli, nil
}
