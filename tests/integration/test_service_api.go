package integration

import (
	"net/http"
	"sort"

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
	s.Len(responseBody.Data, 1)
	s.Equal("test_1", responseBody.Data[0].Name)
}

func (s *IntegrationTestSuite) Test_List_Services_with_Empty_Filter() {
	req := s.createRequest(http.MethodGet, "/v1/service/", nil, map[string]string{})
	s.Router.ServeHTTP(s.Recorder, req)
	response := s.Recorder.Result()
	defer response.Body.Close()
	s.Equal(http.StatusOK, response.StatusCode)

	var responseBody api.ListServicesResponse
	s.loadResponse(&responseBody)
	s.True(responseBody.Total == 5, "expected at least one service")
}

func (s *IntegrationTestSuite) Test_List_Services_with_Partial_Filter() {
	req := s.createRequest(http.MethodGet, "/v1/service/?name=ser", nil, map[string]string{})
	s.Router.ServeHTTP(s.Recorder, req)
	response := s.Recorder.Result()
	defer response.Body.Close()
	s.Equal(http.StatusOK, response.StatusCode)

	var responseBody api.ListServicesResponse
	s.loadResponse(&responseBody)
	s.True(responseBody.Total >= 1, "expected partial match to return some services")
	for _, svc := range responseBody.Data {
		s.Contains(svc.Name, "ser")
	}
}

func (s *IntegrationTestSuite) Test_List_Services_with_No_Matches() {
	req := s.createRequest(http.MethodGet, "/v1/service/?name=nonexistent", nil, map[string]string{})
	s.Router.ServeHTTP(s.Recorder, req)
	response := s.Recorder.Result()
	defer response.Body.Close()
	s.Equal(http.StatusOK, response.StatusCode)

	var responseBody api.ListServicesResponse
	s.loadResponse(&responseBody)
	s.Equal(0, responseBody.Total)
	s.Empty(responseBody.Data)
}

func (s *IntegrationTestSuite) Test_List_Services_with_Pagination() {
	// Assume DB has at least 3 services
	req := s.createRequest(http.MethodGet, "/v1/service/?size=2&page=3", nil, map[string]string{})
	s.Router.ServeHTTP(s.Recorder, req)
	response := s.Recorder.Result()
	defer response.Body.Close()
	s.Equal(http.StatusOK, response.StatusCode)

	var responseBody api.ListServicesResponse
	s.loadResponse(&responseBody)
	s.Equal(1, len(responseBody.Data))
}

func (s *IntegrationTestSuite) Test_List_Services_Sorted_By_Name() {
	req := s.createRequest(http.MethodGet, "/v1/service/?size=10", nil, map[string]string{})
	s.Router.ServeHTTP(s.Recorder, req)
	response := s.Recorder.Result()
	defer response.Body.Close()
	s.Equal(http.StatusOK, response.StatusCode)

	var responseBody api.ListServicesResponse
	s.loadResponse(&responseBody)
	s.True(sort.SliceIsSorted(responseBody.Data, func(i, j int) bool {
		return responseBody.Data[i].Name < responseBody.Data[j].Name
	}), "services should be sorted by name ascending")
}
