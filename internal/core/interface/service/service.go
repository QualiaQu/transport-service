package service

import (
	"context"
	"transport-service/internal/model"
)

type TransportService interface {
	GetTransportTypes(ctx context.Context) ([]model.TransportHandler, error)
}
