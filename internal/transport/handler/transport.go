package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"transport-service/internal/app"
)

func GetTransportTypes(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		transportTypes, err := app.Services.TransportService.GetTransportTypes(c)

		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка получения списка типов транспорта"})

			return
		}

		c.JSON(http.StatusOK, transportTypes)
	}
}

func GetTransportTypesAlongRoute(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
