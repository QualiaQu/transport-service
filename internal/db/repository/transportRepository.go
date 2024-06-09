package repository

import (
	"context"
	"transport-service/internal/db/pg"
	"transport-service/internal/model"
)

type TransportRepository struct {
	pg pg.Conn
}

func NewTransportRepository(pg *pg.Conn) TransportRepository {
	return TransportRepository{
		*pg,
	}
}

func (repo TransportRepository) GetTransportTypes(ctx context.Context) ([]model.TransportPG, error) {
	var transportTypes []model.TransportPG

	err := repo.pg.SelectContext(
		ctx,
		&transportTypes,
		`SELECT id, name FROM transport_types ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}

	return transportTypes, nil
}
