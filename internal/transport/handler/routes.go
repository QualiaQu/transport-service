package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
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

func BookRoutes(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request model.BookingRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "неправильный формат запроса"})
			return
		}

		failedIDs, err := app.Services.RoutesService.Book(c, request.UserID, request.RouteIDs)
		if err != nil {
			slog.Error(err.Error())
			c.JSON(http.StatusInternalServerError, model.BookingResponse{
				Success:   false,
				FailedIDs: failedIDs,
				Message:   "ошибка бронирования некоторых рейсов",
			})
			return
		}

		c.JSON(http.StatusOK, model.BookingResponse{
			Success: true,
			Message: "все рейсы успешно забронированы",
		})
	}
}

func GetBookedRoutes(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr := c.Query("id")
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "неправильный формат ID пользователя"})

			return
		}

		routes, err := app.Services.RoutesService.GetBookedRoutes(c, userID)
		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка получения забронированных маршрутов"})

			return
		}

		c.JSON(http.StatusOK, routes)
	}
}
