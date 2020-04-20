package deps

import (
	"fmt"
	"strings"

	"github.com/danilvpetrov/weathersource/collection"
	"github.com/danilvpetrov/weathersource/internal/envvar"
	"github.com/dogmatiq/dodeca/config"
)

// ProvideClient provides a client to retrieve meteorological data from
// third-party API.
func ProvideClient() (*collection.Client, error) {
	token, err := config.Environment().Get("WEATHER_DARKSKY_TOKEN").AsString()
	if err != nil {
		return nil,
			fmt.Errorf("environment variable 'WEATHER_DARKSKY_TOKEN': %w", err)
	}

	// trim the token as it might come from various sources (strings, files, etc.)
	token = strings.TrimSpace(token)

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
