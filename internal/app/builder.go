package app

import (
	"transport-service/config"
	"transport-service/internal/core/service"
	"transport-service/internal/db/repository"
)

type Builder struct {
	App
}

func newAppBuilder() *Builder {
	return &Builder{
		App: App{},
	}
}

func (b *Builder) build() *App {
	return &b.App
}

func (b *Builder) setConfig(cfg config.Config) *Builder {
	b.Cfg = cfg

	return b
}

func (b *Builder) setRepos(repoManager repository.Manager) *Builder {
	b.Repos = repoManager

	return b
}

func (b *Builder) setServices(serviceManager service.Manager) *Builder {
	b.Services = serviceManager

	return b
}
