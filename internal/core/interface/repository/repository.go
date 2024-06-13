package repository

import (
	"context"
	"transport-service/internal/model"
)

type TransportRepository interface {
	GetTransportTypes(ctx context.Context) ([]model.TransportPG, error)
	GetTransportTypesBetweenCities(ctx context.Context, cityFrom string, cityTo string) ([]model.TransportPG, error)
}

type RoutesRepository interface {
	GetRoutesOnDate(ctx context.Context, request model.RouteRequest) ([]model.RoutePG, error)
	BookRoutes(ctx context.Context, userID int, routeIDs []int) ([]int, error)
}
