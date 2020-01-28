package deps

import (
	"fmt"
	"log"
	"time"

	backoff "github.com/cenkalti/backoff/v4"
	"github.com/danilvpetrov/weathersource/collection"
	"github.com/danilvpetrov/weathersource/internal/envvar"
	"github.com/danilvpetrov/weathersource/storage"
)

// ProvideCollector provides *collection.Collector.
func ProvideCollector(
	client *collection.Client,
	persister storage.Persister,
	logger *log.Logger,
) (*collection.Collector, error) {
	d, err := time.ParseDuration(
		envvar.StringDefault(
			"WEATHER_UPDATE_INTERVAL",
			"5m",
		),
	)
	if err != nil {
		return nil,
			fmt.Errorf(
				"cannot parse WEATHER_UPDATE_INTERVAL as time.Duration: %w",
				err,
			)
	}

	bexp := backoff.NewExponentialBackOff()
	bexp.MaxElapsedTime = d // retry only for the duration of interval.
	bexp.MaxInterval = 1 * time.Minute
	bexp.RandomizationFactor = 0.2

	return &collection.Collector{
		Client:    client,
		Retry:     bexp,
		Persister: persister,
		Interval:  d,
		Logger:    logger,
	}, nil
}
