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

func (repo TransportRepository) GetTransportTypesBetweenCities(ctx context.Context,
	cityFrom string, cityTo string) ([]model.TransportPG, error) {
	var transportTypes []model.TransportPG

	err := repo.pg.SelectContext(
		ctx,
		&transportTypes,
		`
		SELECT DISTINCT tt.id, tt.name
		FROM routes r
		JOIN transport_types tt ON r.transport_type = tt.id
		JOIN cities c1 ON r.origin = c1.id
		JOIN cities c2 ON r.destination = c2.id
		WHERE c1.name = $1 AND c2.name = $2`,
		cityFrom,
		cityTo,
	)

	if err != nil {
		return nil, err
	}

	return transportTypes, nil
}
