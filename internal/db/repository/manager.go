package repository

import "transport-service/internal/db/pg"

type Manager struct {
}

func NewRepositoryManager(pg *pg.Conn) Manager {
	return Manager{}
}
