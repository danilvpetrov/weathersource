package deps

import (
	"github.com/danilvpetrov/weathersource/api/apigrpc"
	"github.com/danilvpetrov/weathersource/api/apihttp"
	"github.com/danilvpetrov/weathersource/collection"
	"github.com/danilvpetrov/weathersource/internal/run"
)

// ProvideServices provides a slice of run.Service implementations.
func ProvideServices(
	collector *collection.Collector,
	grpcsvr *apigrpc.Server,
	httpsvr *apihttp.Server,
) []run.Service {
	return []run.Service{
		collector,
		grpcsvr,
		httpsvr,
	}
}
