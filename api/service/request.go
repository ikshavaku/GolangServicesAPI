package service

import (
	"github.com/ikshavaku/catalogue/service"
)

type GetServiceByIDRequest struct {
	ServiceID string `uri:"service_id"`
}

func (r *GetServiceByIDRequest) ToParams() service.GetServiceByIDParams {
	return service.GetServiceByIDParams{
		ID: r.ServiceID,
	}
}

type ListServiceVersionsByServiceIDRequest struct {
	ServiceID string `uri:"service_id"`
}

func (r *ListServiceVersionsByServiceIDRequest) ToParams() service.ListServiceVersionsByServiceIDParams {
	return service.ListServiceVersionsByServiceIDParams{
		ID: r.ServiceID,
	}
}

type ListServicesRequest struct {
	Page int     `form:"page" binding:"omitempty,min=1"`
	Size int     `form:"size" binding:"omitempty,min=1,max=100"`
	Name *string `form:"name"`
}

func (r *ListServicesRequest) ToParams() service.ListServicesParams {
	params := service.ListServicesParams{
		Page: r.Page,
		Size: r.Size,
		Name: r.Name,
	}

	if params.Page < 1 {
		params.Page = 1
	}

	if params.Size < 1 || params.Size > 100 {
		params.Size = 10
	}
	return params
}
