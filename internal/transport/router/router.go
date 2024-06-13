package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"transport-service/internal/app"
	"transport-service/internal/transport/handler"
)

type Router struct {
	*gin.Engine
}

func InitRoutes(app *app.App) *Router {
	r := gin.Default()

	router := &Router{
		r,
	}

	router.GET("/health", HealthCheck())

	apiV1 := r.Group("/api/v1")

	routes := apiV1.Group("/routes")
	{
		routes.GET("/on-date", handler.GetRoutesOnDate(app))
		routes.POST("/book", handler.BookRoutes(app))
	}

	transport := apiV1.Group("/transport")
	{
		transport.GET("/types", handler.GetTransportTypes(app))
		transport.GET("/types-between-cities", handler.GetTransportTypesBetweenCities(app))
	}

	return router
}

func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
