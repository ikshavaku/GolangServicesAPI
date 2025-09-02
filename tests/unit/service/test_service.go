package crud

import (
	"errors"
	"time"

	"github.com/ikshavaku/catalogue/service"
	"github.com/ikshavaku/catalogue/store"
	"github.com/stretchr/testify/assert"
)

var now = time.Now().Format(time.RFC3339)

func (s *ServiceUnitTestSuite) TestListServices_Success() {
	daoResult := store.PaginatedResponse[store.ServiceDAO]{
		Total: 2,
		Page:  1,
		Size:  10,
		Data: []store.ServiceDAO{
			{ID: "1", Name: "svc1", Description: "first", CreatedAt: now, UpdatedAt: now},
			{ID: "2", Name: "svc2", Description: "second", CreatedAt: now, UpdatedAt: now},
		},
	}
	params := service.ListServicesParams{Page: 1, Size: 10}

	s.ServiceRepo.On("ListServices", s.Ctx, params.ToDBParams()).Return(daoResult, nil)

	result, err := s.service.ListServices(s.Ctx, params)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 2, result.Total)
	assert.Len(s.T(), result.Data, 2)
	assert.Equal(s.T(), "svc1", result.Data[0].Name)
	s.ServiceRepo.AssertExpectations(s.T())
}

func (s *ServiceUnitTestSuite) TestListServices_EmptyResult() {
	daoResult := store.PaginatedResponse[store.ServiceDAO]{
		Total: 0,
		Page:  1,
		Size:  10,
		Data:  []store.ServiceDAO{},
	}
	params := service.ListServicesParams{Page: 1, Size: 10}

	s.ServiceRepo.On("ListServices", s.Ctx, params.ToDBParams()).Return(daoResult, nil)

	result, err := s.service.ListServices(s.Ctx, params)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 0, result.Total)
	assert.Empty(s.T(), result.Data)
	s.ServiceRepo.AssertExpectations(s.T())
}

func (s *ServiceUnitTestSuite) TestListServices_Error() {
	params := service.ListServicesParams{Page: 1, Size: 10}
	expectedErr := errors.New("db error")

	s.ServiceRepo.On("ListServices", s.Ctx, params.ToDBParams()).Return(store.PaginatedResponse[store.ServiceDAO]{}, expectedErr)

	result, err := s.service.ListServices(s.Ctx, params)

	assert.Error(s.T(), err)
	assert.Equal(s.T(), expectedErr, err)
	assert.Empty(s.T(), result.Data)
	s.ServiceRepo.AssertExpectations(s.T())
}

// ---- GetServiceByID ----

func (s *ServiceUnitTestSuite) TestGetServiceByID_Success() {

	dao := store.ServiceDAO{ID: "1", Name: "svc1", Description: "first", CreatedAt: now, UpdatedAt: now}
	params := service.GetServiceByIDParams{ID: "1"}

	s.ServiceRepo.On("GetServiceByID", s.Ctx, params.ToDBParams()).Return(dao, nil)

	result, err := s.service.GetServiceByID(s.Ctx, params)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "1", result.Service.ID)
	assert.Equal(s.T(), "svc1", result.Service.Name)
	s.ServiceRepo.AssertExpectations(s.T())
}

func (s *ServiceUnitTestSuite) TestGetServiceByID_NotFound() {
	params := service.GetServiceByIDParams{ID: "doesnotexist"}
	expectedErr := errors.New("not found")

	s.ServiceRepo.On("GetServiceByID", s.Ctx, params.ToDBParams()).Return(store.ServiceDAO{}, expectedErr)

	result, err := s.service.GetServiceByID(s.Ctx, params)

	assert.Error(s.T(), err)
	assert.Equal(s.T(), expectedErr, err)
	assert.Empty(s.T(), result.Service)
	s.ServiceRepo.AssertExpectations(s.T())
}

// ---- ListServiceVersions ----

func (s *ServiceUnitTestSuite) TestListServiceVersions_Success() {

	dao := []store.ServiceVersionDAO{
		{ServiceID: "1", VersionNumber: "v1", CreatedAt: now, UpdatedAt: now, IsDeleted: false},
		{ServiceID: "1", VersionNumber: "v2", CreatedAt: now, UpdatedAt: now, IsDeleted: true},
	}
	params := service.ListServiceVersionsByServiceIDParams{ID: "1"}

	s.ServiceRepo.On("ListServiceVersions", s.Ctx, params.ToDBParams()).Return(dao, nil)

	result, err := s.service.ListServiceVersions(s.Ctx, params)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 2, result.Total)
	assert.Len(s.T(), result.Data, 2)
	assert.Equal(s.T(), "v1", result.Data[0].VersionID)
	assert.True(s.T(), result.Data[1].Active) // because IsDeleted = true
	s.ServiceRepo.AssertExpectations(s.T())
}

func (s *ServiceUnitTestSuite) TestListServiceVersions_Empty() {
	params := service.ListServiceVersionsByServiceIDParams{ID: "1"}

	s.ServiceRepo.On("ListServiceVersions", s.Ctx, params.ToDBParams()).Return([]store.ServiceVersionDAO{}, nil)

	result, err := s.service.ListServiceVersions(s.Ctx, params)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 0, result.Total)
	assert.Empty(s.T(), result.Data)
	s.ServiceRepo.AssertExpectations(s.T())
}

func (s *ServiceUnitTestSuite) TestListServiceVersions_Error() {
	params := service.ListServiceVersionsByServiceIDParams{ID: "1"}
	expectedErr := errors.New("db error")

	s.ServiceRepo.On("ListServiceVersions", s.Ctx, params.ToDBParams()).Return(nil, expectedErr)

	result, err := s.service.ListServiceVersions(s.Ctx, params)

	assert.Error(s.T(), err)
	assert.Equal(s.T(), expectedErr, err)
	assert.Empty(s.T(), result.Data)
	s.ServiceRepo.AssertExpectations(s.T())
}
