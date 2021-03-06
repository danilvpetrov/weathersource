// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package deps

import (
	"github.com/danilvpetrov/weathersource/internal/deps"
	"github.com/danilvpetrov/weathersource/internal/run"
)

// Injectors from deps.go:

func Deps() ([]run.Service, func(), error) {
	client, err := ProvideClient()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup, err := deps.ProvideDatabase()
	if err != nil {
		return nil, nil, err
	}
	storage := ProvideStorage(db)
	logger := deps.ProvideLogger()
	collector, err := ProvideCollector(client, storage, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	server, err := ProvideGRPCServer(storage, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	apihttpServer, err := ProvideHTTPServer(storage, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	v := ProvideServices(collector, server, apihttpServer)
	return v, func() {
		cleanup()
	}, nil
}
