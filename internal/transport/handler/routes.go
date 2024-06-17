package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
	"transport-service/internal/app"
	"transport-service/internal/model"
)

// GetRoutesOnDate godoc
// @Summary Получить маршруты на дату
// @Description Получить доступные маршруты на заданную дату и предпочтительные виды транспорта
// @Tags Маршруты
// @Accept  json
// @Produce  json
// @Param request body model.RouteRequest true "Данные запроса"
// @Success 200 {object} []model.RoutePG "Список маршрутов"
// @Failure 400 {object} gin.H "Неправильный формат запроса"
// @Failure 500 {object} gin.H "Ошибка получения рейсов на день"
// @Router /routes/on-date [get]
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

// BookRoutes godoc
// @Summary Забронировать маршруты
// @Description Забронировать указанные маршруты для пользователя
// @Tags Маршруты
// @Accept  json
// @Produce  json
// @Param request body model.BookingRequest true "Данные запроса на бронирование"
// @Success 200 {object} model.BookingResponse "Результат бронирования"
// @Failure 400 {object} gin.H "Неправильный формат запроса"
// @Failure 500 {object} model.BookingResponse "Ошибка бронирования некоторых рейсов"
// @Router /routes/book [post]
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

// GetBookedRoutes godoc
// @Summary Получить забронированные маршруты
// @Description Получить все забронированные маршруты для конкретного пользователя
// @Tags Маршруты
// @Accept  json
// @Produce  json
// @Param id query int true "ID пользователя"
// @Success 200 {object} []model.RoutePG "Список забронированных маршрутов"
// @Failure 400 {object} gin.H "Неправильный формат ID пользователя"
// @Failure 500 {object} gin.H "Ошибка получения забронированных маршрутов"
// @Router /routes/booked [get]
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
