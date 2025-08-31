package service

import (
	"context"

	"github.com/ikshavaku/catalogue/store"
)

type ServicesService struct {
	serviceReporitory store.IServiceRepository
}

func ServiceToResult(s store.ServiceDAO) Service {
	return Service{
		ID:          s.ID,
		Name:        s.Name,
		Description: s.Description,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}

func ServiceVersionToResult(s store.ServiceVersionDAO) ServiceVersion {
	return ServiceVersion{
		ServiceID: s.ServiceID,
		VersionID: s.VersionNumber,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
		Active:    s.IsDeleted,
	}
}

func (s *ServicesService) ListServices(ctx context.Context, params ListServicesParams) (ListServicesResult, error) {
	dao, err := s.serviceReporitory.ListServices(ctx, params.ToDBParams())
	if err != nil {
		return ListServicesResult{}, err
	}
	data := make([]Service, 0, len(dao.Data))
	for _, s := range dao.Data {
		data = append(data, ServiceToResult(s))
	}
	return ListServicesResult{
		Total: int(dao.Total),
		Page:  dao.Page,
		Size:  dao.Size,
		Data:  data,
	}, nil
}

func (s *ServicesService) GetServiceByID(ctx context.Context, params GetServiceByIDParams) (GetServiceByIDResult, error) {
	dao, err := s.serviceReporitory.GetServiceByID(ctx, params.ToDBParams())
	if err != nil {
		return GetServiceByIDResult{}, err
	}
	return GetServiceByIDResult{
		Service: ServiceToResult(dao),
	}, nil
}

func (s *ServicesService) ListServiceVersions(ctx context.Context, params ListServiceVersionsByServiceIDParams) (ListServiceVersionsByServiceIDResult, error) {
	dao, err := s.serviceReporitory.ListServiceVersions(ctx, params.ToDBParams())
	if err != nil {
		return ListServiceVersionsByServiceIDResult{}, err
	}
	data := make([]ServiceVersion, 0, len(dao))
	for _, s := range dao {
		data = append(data, ServiceVersionToResult(s))
	}
	return ListServiceVersionsByServiceIDResult{
		Total: len(dao),
		Data:  data,
	}, nil
}
