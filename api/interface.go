package api

import "github.com/gin-gonic/gin"

type Routable interface {
	RegisterRoutes(*gin.RouterGroup)
}

type IServiceAPIController interface {
	Routable
	ListAllServices(*gin.Context)
	GetServiceByID(*gin.Context)
	ListServiceVersions(*gin.Context)
}
