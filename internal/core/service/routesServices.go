package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"transport-service/internal/core/interface/repository"
	"transport-service/internal/core/interface/service"
	"transport-service/internal/model"
)

type _routesService struct {
	repoPG repository.RoutesRepository
}

func NewRoutesService(repo repository.RoutesRepository) service.RoutesService {
	return _routesService{repoPG: repo}
}

func (service _routesService) GetRoutesOnDate(ctx context.Context, request model.RouteRequest) ([]model.RouteResponse, error) {
	routesPG, err := service.repoPG.GetRoutesOnDate(ctx, request)

	if err != nil {
		return nil, fmt.Errorf("routesService GetRoutesOnDate: %w", err)
	}

	return routesPgToResponse(routesPG), nil
}

func routesPgToResponse(routesPG []model.RoutePG) []model.RouteResponse {
	routesResponse := make([]model.RouteResponse, len(routesPG))
	for i, routePG := range routesPG {
		routesResponse[i] = routePG.ToResponse()
	}

	return routesResponse
}

func (service _routesService) Book(ctx *gin.Context, userID int, routesID []int) ([]int, error) {
	return service.repoPG.BookRoutes(ctx, userID, routesID)
}
