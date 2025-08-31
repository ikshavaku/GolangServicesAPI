package api

import (
	"github.com/gin-gonic/gin"
	server "github.com/ikshavaku/catalogue/utils"
)

type RouteOption func(server *gin.Engine)

func NewAPIServer(
	servicesControllers IServiceAPIController,
) *gin.Engine {
	server := server.Build()
	InitRoutes(server, WithPathController(
		"/v1/service", servicesControllers,
	))
	return server
}

func WithPathController(path string, ctrl Routable) func(server *gin.Engine) {
	return func(server *gin.Engine) {
		g := server.Group(path)
		ctrl.RegisterRoutes(g)
	}
}

func InitRoutes(server *gin.Engine, options ...RouteOption) {
	for _, o := range options {
		o(server)
	}
}
