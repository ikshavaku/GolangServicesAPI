package service

import "context"

type IServicesService interface {
	ListServices(context.Context, ListServicesParams) (ListServicesResult, error)
	GetServiceByID(context.Context, GetServiceByIDParams) (GetServiceByIDResult, error)
	ListServiceVersions(context.Context, ListServiceVersionsByServiceIDParams) (ListServiceVersionsByServiceIDResult, error)
}
