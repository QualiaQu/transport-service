package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"transport-service/internal/model"
)

type TransportService interface {
	GetTransportTypes(ctx context.Context) ([]model.TransportHandler, error)
	GetTransportTypesBetweenCities(ctx context.Context, cityFrom string, cityTo string) ([]model.TransportHandler, error)
}

type RoutesService interface {
	GetRoutesOnDate(ctx context.Context, request model.RouteRequest) ([]model.RouteResponse, error)
	Book(ctx *gin.Context, userID int, routesID []int) ([]int, error)
}
