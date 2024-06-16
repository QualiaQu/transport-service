package repository

import (
	"context"
	"database/sql"
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

func (repo RoutesRepository) BookRoutes(ctx context.Context, userID int, routeIDs []int) ([]int, error) {
	var failedIDs []int

	tx, err := repo.pg.BeginTxx(ctx, nil)
	if err != nil {
		return routeIDs, err
	}

	for _, id := range routeIDs {
		res, err := tx.ExecContext(ctx, `UPDATE routes SET is_booked = true WHERE id = $1 AND is_booked = false`, id)
		if err != nil {
			failedIDs = append(failedIDs, id)

			continue
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil || rowsAffected == 0 {
			failedIDs = append(failedIDs, id)

			continue
		}

		_, err = tx.ExecContext(ctx, `INSERT INTO bookings (user_id, route_id) VALUES ($1, $2)`, userID, id)
		if err != nil {
			failedIDs = append(failedIDs, id)
		}
	}

	if len(failedIDs) > 0 {
		err = tx.Rollback()
		if err != nil {
			return nil, err
		}

		return failedIDs, sql.ErrTxDone
	}

	if err = tx.Commit(); err != nil {
		return routeIDs, err
	}

	return nil, nil
}

func (repo RoutesRepository) GetBookedRoutes(ctx context.Context, userID int) ([]model.RoutePG, error) {
	var routes []model.RoutePG

	err := repo.pg.SelectContext(
		ctx,
		&routes,
		`SELECT r.id, r.transport_type, r.price, r.departure_datetime, r.arrival_datetime
		FROM routes r
		JOIN bookings b ON r.id = b.route_id
		WHERE b.user_id = $1`,
		userID)

	if err != nil {
		return nil, err
	}

	return routes, nil
}
