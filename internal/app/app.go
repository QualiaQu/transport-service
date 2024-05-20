package app

import (
	"core-adapter-mini-billing/config"
	"transport-service/internal/core/service"
	"transport-service/internal/db/repository"
)

type App struct {
	Cfg config.Config

	Repo repository.Manager

	Services service.Manager
}
