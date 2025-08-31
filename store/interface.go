package store

import "context"

type PaginatedResponse[T any] struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Size  int   `json:"size"`
	Data  []T   `json:"data"`
}

type ListServicesParams struct {
	Page   int
	Size   int
	Name   *string
	SortBy string
	Order  string
}

type GetServiceByIDParams struct {
	ID string
}

type ListServiceVersionsByServiceIDParams struct {
	ID string
}

type IServiceRepository interface {
	ListServices(context.Context, ListServicesParams) (PaginatedResponse[ServiceDAO], error)
	GetServiceByID(context.Context, GetServiceByIDParams) (ServiceDAO, error)
	ListServiceVersions(context.Context, ListServiceVersionsByServiceIDParams) ([]ServiceVersionDAO, error)
}
