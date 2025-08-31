package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ikshavaku/catalogue/api"
	"github.com/ikshavaku/catalogue/service"
	"github.com/ikshavaku/catalogue/utils"
)

type ServiceAPIController struct {
	service service.IServicesService
}

func (c *ServiceAPIController) ListAllServices(g *gin.Context) {
	var request ListServicesRequest
	if err := utils.ParseRequest(g, &request); err != nil {
		err := api.APIError{Status: http.StatusBadRequest, Response: api.ErrorResponse{Details: err.Error(), Error: "bad request"}}
		g.AbortWithStatusJSON(err.Status, err.Response)
		return
	}
	result, err := c.service.ListServices(g, request.ToParams())
	if err != nil {
		err := api.APIError{Status: http.StatusBadRequest, Response: api.ErrorResponse{Details: err.Error(), Error: "bad request"}}
		g.AbortWithStatusJSON(err.Status, err.Response)
		return
	}
	var response ListServicesResponse
	services := make([]Service, 0, len(result.Data))
	for _, s := range result.Data {
		services = append(services, Service{
			ID:          s.ID,
			Name:        s.Name,
			Description: s.Description,
			CreatedAt:   s.CreatedAt,
			UpdatedAt:   s.UpdatedAt,
		})
	}
	response.Total = result.Total
	response.Size = result.Size
	response.Page = result.Page
	response.Data = services
	g.JSON(http.StatusOK, response)
}
func (c *ServiceAPIController) GetServiceByID(g *gin.Context) {
	var request GetServiceByIDRequest
	if err := utils.ParseRequest(g, &request); err != nil {
		err := api.APIError{Status: http.StatusBadRequest, Response: api.ErrorResponse{Details: err.Error(), Error: "bad request"}}
		g.AbortWithStatusJSON(err.Status, err.Response)
		return
	}
	result, err := c.service.GetServiceByID(g, request.ToParams())
	if err != nil {
		err := api.APIError{Status: http.StatusBadRequest, Response: api.ErrorResponse{Details: err.Error(), Error: "bad request"}}
		g.AbortWithStatusJSON(err.Status, err.Response)
		return
	}
	g.JSON(http.StatusOK, GetServiceResponse{
		Service: Service{
			ID:          result.ID,
			Name:        result.Name,
			Description: result.Description,
			CreatedAt:   result.CreatedAt,
			UpdatedAt:   result.UpdatedAt,
		},
	})
}

func (c *ServiceAPIController) ListServiceVersions(g *gin.Context) {
	var request ListServiceVersionsByServiceIDRequest
	if err := utils.ParseRequest(g, &request); err != nil {
		err := api.APIError{Status: http.StatusBadRequest, Response: api.ErrorResponse{Details: err.Error(), Error: "bad request"}}
		g.AbortWithStatusJSON(err.Status, err.Response)
		return
	}
	result, err := c.service.ListServiceVersions(g, request.ToParams())
	if err != nil {
		err := api.APIError{Status: http.StatusBadRequest, Response: api.ErrorResponse{Details: err.Error(), Error: "bad request"}}
		g.AbortWithStatusJSON(err.Status, err.Response)
		return
	}
	versions := make([]ServiceVersion, 0, len(result.Data))
	for _, v := range result.Data {
		versions = append(versions, ServiceVersion{
			VersionID: v.VersionID,
			ServiceID: v.ServiceID,
			CreatedAt: v.CreatedAt,
			Active:    v.Active,
		})
	}
	g.JSON(http.StatusOK, ListServiceVersionsByServiceIDResponse{
		Data:  versions,
		Total: len(versions),
	})
}
