package integration

import (
	"net/http"

	api "github.com/ikshavaku/catalogue/api/service"
)

func (s *IntegrationTestSuite) Test_List_Services_with_Name_Filter() {
	req := s.createRequest(http.MethodGet, "/v1/service/?name=test", nil, map[string]string{})
	s.Router.ServeHTTP(s.Recorder, req)
	response := s.Recorder.Result()
	defer response.Body.Close()
	s.Equal(http.StatusOK, response.StatusCode)
	var responseBody api.ListServicesResponse
	s.loadResponse(&responseBody)
	s.Equal(1, responseBody.Total)
}
