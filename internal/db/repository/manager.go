package repository

import (
	"transport-service/internal/core/interface/repository"
	"transport-service/internal/db/pg"
)

type Manager struct {
	repository.TransportRepository
	repository.RoutesRepository
}

func NewRepositoryManager(pg *pg.Conn) Manager {
	return Manager{
		NewTransportRepository(pg), NewRoutesRepository(pg),
	}
}
