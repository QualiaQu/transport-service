package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"transport-service/internal/app"
)

// GetTransportTypes godoc
// @Summary Получить все типы транспорта
// @Description Получить список всех доступных типов транспорта
// @Tags Транспорт
// @Accept  json
// @Produce  json
// @Success 200 {object} []model.TransportHandler "Список типов транспорта"
// @Failure 500 {object} gin.H "Ошибка получения списка типов транспорта"
// @Router /transport/types [get]
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

// GetTransportTypesBetweenCities godoc
// @Summary Получить типы транспорта между городами
// @Description Получить список типов транспорта, доступных для перемещения между двумя городами
// @Tags Транспорт
// @Accept  json
// @Produce  json
// @Param cityFrom query string true "Город отправления"
// @Param cityTo query string true "Город назначения"
// @Success 200 {object} []model.TransportHandler "Список типов транспорта"
// @Failure 400 {object} gin.H "Необходимо указать параметры cityFrom и cityTo"
// @Failure 500 {object} gin.H "Ошибка получения списка типов транспорта"
// @Router /transport/types-between-cities [get]
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
