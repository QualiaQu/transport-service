package pg

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"transport-service/config"
)

type Conn struct {
	*sqlx.DB
}

func New(ctx context.Context, conf config.Postgres) (*Conn, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		conf.User, conf.Password, conf.DBName, conf.Host)

	db, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &Conn{db}, nil
}
