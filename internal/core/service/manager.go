package service

import "transport-service/internal/db/repository"

type Manager struct{}

func NewServiceManager(repo repository.Manager) Manager {
	return Manager{}
}
