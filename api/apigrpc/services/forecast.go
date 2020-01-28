package services

import (
	"context"

	"github.com/danilvpetrov/weathersource/storage"
	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// ForecastServer is the server API for Forecast service.
type forecastServer struct {
	DataAccessor storage.DataAccessor
}

// NewForecastServer returns an instance that implements ForecastServer.
func NewForecastServer(
	dataAccessor storage.DataAccessor,
) ForecastServer {
	return &forecastServer{
		DataAccessor: dataAccessor,
	}
}

// Get retrieves the latest weather forecast from the forecast service API.
func (s *forecastServer) Get(
	ctx context.Context,
	req *ForecastRequest,
) (*ForecastResponse, error) {
	var resp ForecastResponse

	if err := resp.Forecast.CopyFrom(ctx, s.DataAccessor); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &resp, nil
}
