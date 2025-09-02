package crud

import (
	"os"
	"testing"

	"github.com/ikshavaku/catalogue/tests"
	"github.com/stretchr/testify/suite"
)

func TestRunIntegrationSuite(t *testing.T) {
	suite.Run(t, new(ServiceUnitTestSuite))
}
func TestMain(m *testing.M) {
	tests.LoadTestConfig()
	os.Exit(m.Run())
}
