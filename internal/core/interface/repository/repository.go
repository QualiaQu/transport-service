package repository

import (
	"context"
	"transport-service/internal/model"
)

type TransportRepository interface {
	GetTransportTypes(ctx context.Context) ([]model.TransportPG, error)
}
