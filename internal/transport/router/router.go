package router

import (
	"github.com/gin-gonic/gin"
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

	router.GET("/health", handler.HealthCheck())

	return router
}
