// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package deps

import (
	commondeps "github.com/danilvpetrov/weathersource/internal/deps"
	"github.com/danilvpetrov/weathersource/internal/run"
	"github.com/danilvpetrov/weathersource/storage"
	"github.com/danilvpetrov/weathersource/storage/mysqlstorage"
	"github.com/google/wire"
)

func Deps() ([]run.Service, func(), error) {
	panic(
		wire.Build(
			commondeps.CommonDepsSet,
			ProvideClient,
			ProvideStorage,
			wire.Bind(new(storage.Persister), new(*mysqlstorage.Storage)),
			wire.Bind(new(storage.DataAccessor), new(*mysqlstorage.Storage)),
			ProvideCollector,
			ProvideHTTPServer,
			ProvideGRPCServer,
			ProvideServices,
		),
	)
}
