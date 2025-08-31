package service

import "github.com/ikshavaku/catalogue/store"

type Service struct {
	ID          string
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string
}

type ListServicesResult struct {
	Total int
	Page  int
	Size  int
	Data  []Service
}

type GetServiceByIDParams struct {
	ID string
}

func (p *GetServiceByIDParams) ToDBParams() store.GetServiceByIDParams {
	return store.GetServiceByIDParams{
		ID: p.ID,
	}
}

type ListServiceVersionsByServiceIDParams struct {
	ID string
}

func (p *ListServiceVersionsByServiceIDParams) ToDBParams() store.ListServiceVersionsByServiceIDParams {
	return store.ListServiceVersionsByServiceIDParams{
		ID: p.ID,
	}
}

type ServiceVersion struct {
	ServiceID string
	VersionID string
	CreatedAt string
	UpdatedAt string
	Active    bool
}

type ListServiceVersionsByServiceIDResult struct {
	Total int
	Data  []ServiceVersion
}

type GetServiceByIDResult struct {
	Service
}

type ListServicesParams struct {
	Page   int
	Size   int
	Name   *string
	SortBy string
	Order  string
}

func (p *ListServicesParams) ToDBParams() store.ListServicesParams {
	return store.ListServicesParams{
		Page:   p.Page,
		Size:   p.Size,
		Name:   p.Name,
		SortBy: p.SortBy,
		Order:  p.Order,
	}
}
