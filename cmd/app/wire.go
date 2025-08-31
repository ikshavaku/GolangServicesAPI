//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ikshavaku/catalogue/api"
	servicectlr "github.com/ikshavaku/catalogue/api/service"
	"github.com/ikshavaku/catalogue/service"
	"github.com/ikshavaku/catalogue/store"
	"github.com/ikshavaku/catalogue/store/postgres"
	"github.com/ikshavaku/catalogue/store/postgres/repositories"
)

//go:generate go run github.com/google/wire/cmd/wire

func InjectServicesController() api.IServiceAPIController {
	wire.Build(
		servicectlr.ProvideServiceAPIController,
		service.ProvideServicesService,
		wire.Bind(new(store.IServiceRepository), new(*repositories.SerivcesRepository)),
		repositories.ProvideServicesRepository,
		postgres.PostgresProviderSet,
		wire.Bind(new(service.IServicesService), new(*service.ServicesService)),
		wire.Bind(new(api.IServiceAPIController), new(*servicectlr.ServiceAPIController)),
	)
	return nil
}
