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

func GetTransportTypesBetweenCities(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		cityFrom := c.Query("cityFrom")
		cityTo := c.Query("cityTo")

		if cityFrom == "" || cityTo == "" {
			slog.Error("отсутвуют параметры cityFrom или cityTo")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "необходимо указать параметры cityFrom и cityTo"})
			return
		}

		transportTypes, err := app.Services.TransportService.GetTransportTypesBetweenCities(c, cityFrom, cityTo)

		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка получения списка типов транспорта"})

			return
		}

		c.JSON(http.StatusOK, transportTypes)
	}
}
