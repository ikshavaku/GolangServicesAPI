package integration

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/ikshavaku/catalogue/api"
	service_api "github.com/ikshavaku/catalogue/api/service"
	"github.com/ikshavaku/catalogue/service"
	"github.com/ikshavaku/catalogue/store"
	"github.com/ikshavaku/catalogue/store/postgres"
	"github.com/ikshavaku/catalogue/store/postgres/repositories"
	"github.com/ikshavaku/catalogue/tests"
	"github.com/ikshavaku/catalogue/utils"
	"github.com/stretchr/testify/suite"
)

type IntegrationTestSuite struct {
	suite.Suite
	Recorder *httptest.ResponseRecorder
	Router   *gin.Engine
	repo     store.IServiceRepository
	Ctx      context.Context
}

func (s *IntegrationTestSuite) SetupTest() {
	log.Println("[TEST] Setting up integration suite...")
	cfg := tests.LoadTestConfig()
	LoadTestFixtures(cfg)
	url := postgres.ProvidePostgresURL(cfg.Postgres)
	conn := postgres.ProvidePostgresConnection(url)
	query := postgres.ProvidePostgresQueries(conn)
	s.repo = repositories.ProvideServicesRepository(query)
	servicesService := service.ProvideServicesService(s.repo)
	serviceAPIController := service_api.ProvideServiceAPIController(servicesService)
	s.Recorder = httptest.NewRecorder()
	s.Router = utils.Build()
	api.InitRoutes(s.Router, api.WithPathController("/v1/service", serviceAPIController))
	log.Println("[TEST] Integration suite setup completed")
	s.Ctx = context.Background()
}

func (s *IntegrationTestSuite) createRequest(method, path string, body any, headers map[string]string) *http.Request {
	// Set new recorder for each request
	s.Recorder = httptest.NewRecorder()

	request, err := tests.CreateRequest(method, path, body, headers)
	s.NoError(err)
	return request
}

func (s *IntegrationTestSuite) loadResponse(target any) {
	err := tests.LoadResponse(s.Recorder.Body, target)
	s.NoError(err)
}
