package service

import (
	"context"
	"fmt"
	"transport-service/internal/core/interface/repository"
	"transport-service/internal/core/interface/service"
	"transport-service/internal/model"
)

type _transportService struct {
	repoPG repository.TransportRepository
}

func NewTransportService(repo repository.TransportRepository) service.TransportService {
	return _transportService{repoPG: repo}
}

func (service _transportService) GetTransportTypes(ctx context.Context) ([]model.TransportHandler, error) {
	transportTypesPG, err := service.repoPG.GetTransportTypes(ctx)

	if err != nil {
		return nil, fmt.Errorf("transportService GetTransportService: %w", err)
	}

	return transportPgToHandler(transportTypesPG), nil
}

func (service _transportService) GetTransportTypesBetweenCities(ctx context.Context,
	cityFrom string, cityTo string) ([]model.TransportHandler, error) {
	transportTypesPG, err := service.repoPG.GetTransportTypesBetweenCities(ctx, cityFrom, cityTo)

	if err != nil {
		return nil, fmt.Errorf("transportService GetTransportTypesBetweenCities: %w", err)
	}

	return transportPgToHandler(transportTypesPG), nil
}

func transportPgToHandler(transportTypesPG []model.TransportPG) []model.TransportHandler {
	transportTypes := make([]model.TransportHandler, len(transportTypesPG))
	for i, transportPG := range transportTypesPG {
		transportTypes[i] = transportPG.ToHandler()
	}

	return transportTypes
}
