package service

import "github.com/ikshavaku/catalogue/service"

func ProvideServiceAPIController(servicesService service.IServicesService) *ServiceAPIController {
	return &ServiceAPIController{
		service: servicesService,
	}
}
