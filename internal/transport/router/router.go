package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "transport-service/docs"

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
	router.Use(cors.Default())

	router.GET("/health", healthCheck())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := r.Group("/api/v1")

	routes := apiV1.Group("/routes")
	{
		routes.GET("/on-date", handler.GetRoutesOnDate(app))
		routes.POST("/book", handler.BookRoutes(app))
		routes.GET("/booked", handler.GetBookedRoutes(app))
	}

	transport := apiV1.Group("/transport")
	{
		transport.GET("/types", handler.GetTransportTypes(app))
		transport.GET("/types-between-cities", handler.GetTransportTypesBetweenCities(app))
	}

	return router
}

func healthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
