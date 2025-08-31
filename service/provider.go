package service

import "github.com/ikshavaku/catalogue/store"

func ProvideServicesService(repo store.IServiceRepository) *ServicesService {
	return &ServicesService{
		serviceReporitory: repo,
	}
}
