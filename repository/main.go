package repository

import (
	"poc/bus"
	"poc/config"
)

type (
	IRepository interface {
		Start()
		Stop()
	}

	Repository struct {
		EventBus bus.IEventBus
		config   *config.CloudConfig
	}
)

func (r *Repository) Start() {
	panic("implement me")
}

func (r *Repository) Stop() {
	panic("implement me")
}
