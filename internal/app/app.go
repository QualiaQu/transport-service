package app

import (
	"transport-service/config"
	"transport-service/internal/core/service"
	"transport-service/internal/db/pg"
	"transport-service/internal/db/repository"
)

type App struct {
	Cfg config.Config

	Repos repository.Manager

	Services service.Manager
}

func NewApp(cfg config.Config, pg *pg.Conn) *App {
	return newAppBuilder().
		setConfig(cfg).
		setServices(service.NewServiceManager(repository.NewRepositoryManager(pg))).
		build()
}
