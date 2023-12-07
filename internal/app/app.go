package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lucdoe/capstone_gateway/internal"
	"github.com/lucdoe/capstone_gateway/internal/middlewares"
)

func handleJSONValidation(body internal.BodyField) gin.HandlerFunc {
	return func(c *gin.Context) {
		if body.ApplyValidation {
			validator := middlewares.ValidateJSONFields(body)
			validator(c)
		}
		if !c.IsAborted() {
			c.Next()
		}
	}
}

func proxyRequest(URL string) gin.HandlerFunc {
	return middlewares.ReverseProxy(URL)
}

func SetupRoutes(r internal.RouterInterface, config *internal.Config) {
	for _, service := range config.Services {
		for _, endpoint := range service.Endpoints {
			endpointURL := service.URL + endpoint.Path
			relativePath := endpoint.Path

			switch endpoint.HTTPMethod {
			case "GET":
				r.GET(relativePath, handleJSONValidation(endpoint.Body), proxyRequest(endpointURL))
			case "POST":
				r.POST(relativePath, handleJSONValidation(endpoint.Body), proxyRequest(endpointURL))
			case "PUT":
				r.PUT(relativePath, handleJSONValidation(endpoint.Body), proxyRequest(endpointURL))
			case "PATCH":
				r.PATCH(relativePath, handleJSONValidation(endpoint.Body), proxyRequest(endpointURL))
			}
		}
	}
}

func APIGatewayApp(router internal.RouterInterface, config *internal.Config) (*internal.App, error) {
	middlewares.InitilizeMiddlewares(router, config)
	SetupRoutes(router, config)
	return &internal.App{Router: router}, nil
}
