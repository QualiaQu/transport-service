package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"transport-service/internal/app"
	"transport-service/internal/model"
)

func GetRoutesOnDate(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var routeRequest model.RouteRequest

		if err := c.ShouldBindJSON(&routeRequest); err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "неправильный формат запроса"})

			return
		}

		if routeRequest.Origin == "" || routeRequest.Destination == "" || routeRequest.Date == "" || len(routeRequest.PreferredTransport) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "не все поля тела запроса заполнены"})

			return
		}

		routes, err := app.Services.RoutesService.GetRoutesOnDate(c, routeRequest)

		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка получения рейсов на день"})

			return
		}

		c.JSON(http.StatusOK, routes)
	}
}
