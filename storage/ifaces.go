package storage

import (
	"context"

	"github.com/danilvpetrov/weathersource/data"
)

// Persister persists the meterological data to the storage.
type Persister interface {
	// Save saves the meterological data to the storage.
	Save(ctx context.Context, data []*data.Data) error
}

// DataAccessor retrieves the latest meteorological data from the storage.
type DataAccessor interface {
	// GetLatest retrieves the latest meteorological reading and and all
	// available forecasts from the storage.
	//
	// TO-DO: specify the type of the forecast and geolocation coordinates
	// to retrieve.
	GetLatest(ctx context.Context) ([]*data.Data, error)
}
