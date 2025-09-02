package crud

import (
	"context"

	"github.com/ikshavaku/catalogue/store/postgres"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/mock"
)

type MockQueries struct {
	mock.Mock
}

func (m *MockQueries) CountServices(ctx context.Context, name pgtype.Text) (int64, error) {
	args := m.Called(ctx, name)
	return int64(args.Int(0)), args.Error(1)
}

func (m *MockQueries) ListServices(ctx context.Context, params postgres.ListServicesParams) ([]postgres.Service, error) {
	args := m.Called(ctx, params)
	return args.Get(0).([]postgres.Service), args.Error(1)
}

func (m *MockQueries) GetServiceByID(ctx context.Context, id int32) (postgres.Service, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(postgres.Service), args.Error(1)
}

func (m *MockQueries) ListServiceVersionsByServiceID(ctx context.Context, id int32) ([]postgres.ServiceVersion, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]postgres.ServiceVersion), args.Error(1)
}
