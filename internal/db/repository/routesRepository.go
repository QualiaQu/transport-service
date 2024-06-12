package repository

import (
	"context"
	"github.com/lib/pq"
	"transport-service/internal/db/pg"
	"transport-service/internal/model"
)

type RoutesRepository struct {
	pg pg.Conn
}

func NewRoutesRepository(pg *pg.Conn) RoutesRepository {
	return RoutesRepository{
		*pg,
	}
}

func (repo RoutesRepository) GetRoutesOnDate(ctx context.Context, request model.RouteRequest) ([]model.RoutePG, error) {
	var routes []model.RoutePG

	err := repo.pg.SelectContext(
		ctx,
		&routes,
		`SELECT id, transport_type, price, departure_datetime, arrival_datetime
		FROM routes
		WHERE origin = (SELECT id FROM cities WHERE name = $1)
		  AND destination = (SELECT id FROM cities WHERE name = $2)
		  AND date(departure_datetime) = $3
		  AND transport_type = ANY($4)
		  AND is_booked = false`,
		request.Origin,
		request.Destination,
		request.Date,
		pq.Array(request.PreferredTransport))

	if err != nil {
		return nil, err
	}

	return routes, nil
}
