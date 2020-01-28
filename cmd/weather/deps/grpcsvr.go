package deps

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/danilvpetrov/weathersource/api/apigrpc"
	"github.com/danilvpetrov/weathersource/api/apigrpc/services"
	"github.com/danilvpetrov/weathersource/internal/envvar"
	"github.com/danilvpetrov/weathersource/storage"
	"google.golang.org/grpc"
)

// ProvideGRPCServer provides a gRPC Server as a service to run.
func ProvideGRPCServer(
	dataAccessor storage.DataAccessor,
	logger *log.Logger,
) (*apigrpc.Server, error) {
	port, err := strconv.ParseInt(
		envvar.StringDefault(
			"WEATHER_GRPC_PORT",
			"8081",
		), 10, 64,
	)
	if err != nil {
		return nil,
			fmt.Errorf(
				"cannot parse WEATHER_GRPC_PORT as an integer: %w",
				err,
			)
	}

	fs := services.NewForecastServer(dataAccessor)

	return &apigrpc.Server{
		Listen: func() (net.Listener, error) {
			return net.Listen(
				"tcp",
				fmt.Sprintf(":%d", port),
			)
		},
		Register: func(s *grpc.Server) {
			services.RegisterForecastServer(s, fs)
		},
		Logger: logger,
	}, nil
}
