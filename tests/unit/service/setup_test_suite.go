package crud

import (
	"context"
	"log"

	"github.com/ikshavaku/catalogue/service"
	"github.com/ikshavaku/catalogue/tests"
	testmocks "github.com/ikshavaku/catalogue/tests/mocks"
	"github.com/ikshavaku/catalogue/tests/unit"
	"github.com/stretchr/testify/suite"
)

type ServiceUnitTestSuite struct {
	suite.Suite
	service     service.IServicesService
	ServiceRepo *testmocks.IServiceRepository
	Ctx         context.Context
}

func (s *ServiceUnitTestSuite) SetupTest() {
	log.Println("[TEST] Setting up integration suite...")
	cfg := tests.LoadTestConfig()
	unit.LoadTestFixtures(cfg)
	repo := new(testmocks.IServiceRepository)
	s.ServiceRepo = repo
	s.service = service.ProvideServicesService(repo)
	log.Println("[TEST] Integration suite setup completed")
	s.Ctx = context.Background()
}
