package main

import (
	"context"
	"log/slog"
	"transport-service/config"
	"transport-service/internal/app"
	"transport-service/internal/db/pg"
	"transport-service/internal/helpers"
	"transport-service/internal/transport/router"
)

func main() {
	cfg, err := config.ReadConfig(helpers.DefaultConfigPath)
	if err != nil {
		slog.Error("Failed to load YAML config: %v", err)

		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pgConn, err := pg.New(ctx, cfg.Postgres)
	if err != nil {
		slog.Error("Failed to connect to database: %v", err)

		return
	}

	application := app.NewApp(cfg, pgConn)

	router.InitRoutes(application).Run(cfg.Server.Listen)
}
