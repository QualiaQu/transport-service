package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
