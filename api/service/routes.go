package service

import (
	"github.com/gin-gonic/gin"
)

func (c *ServiceAPIController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/", c.ListAllServices)
	router.GET("/:service_id", c.GetServiceByID)
	router.GET("/:service_id/versions", c.ListServiceVersions)
}
