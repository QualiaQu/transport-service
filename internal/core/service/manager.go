package service

import (
	"transport-service/internal/core/interface/service"
	"transport-service/internal/db/repository"
)

type Manager struct {
	TransportService service.TransportService
}

func NewServiceManager(repo repository.Manager) Manager {
	return Manager{
		NewTransportService(repo.TransportRepository),
	}
}
