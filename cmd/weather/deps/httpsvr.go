package deps

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/danilvpetrov/weathersource/api/apihttp"
	"github.com/danilvpetrov/weathersource/api/apihttp/handlers"
	"github.com/danilvpetrov/weathersource/internal/envvar"
	"github.com/danilvpetrov/weathersource/storage"
)

// ProvideHTTPServer provides a HTTP Server as a service to run.
func ProvideHTTPServer(
	dataAccessor storage.DataAccessor,
	logger *log.Logger,
) (*apihttp.Server, error) {
	port, err := strconv.ParseInt(
		envvar.StringDefault(
			"WEATHER_HTTP_PORT",
			"8080",
		), 10, 64,
	)
	if err != nil {
		return nil,
			fmt.Errorf(
				"cannot parse WEATHER_HTTP_PORT as an integer: %w",
				err,
			)
	}

	return &apihttp.Server{
		Listen: func() (net.Listener, error) {
			return net.Listen(
				"tcp",
				fmt.Sprintf(":%d", port),
			)
		},
		Handler:      handlers.NewHandler(dataAccessor),
		Logger:       logger,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}, nil
}
