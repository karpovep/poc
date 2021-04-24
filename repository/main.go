package repository

import (
	"poc/app"
	"poc/bus"
	"poc/config"
	"poc/model"
)

type (
	IRepository interface {
		Start()
		Stop()
	}

	Repository struct {
		EventBus           bus.IEventBus
		config             *config.CloudConfig
		inboundRepoChan    bus.DataChannel
		inboundChannelName string
	}
)

func NewRepository(appContext app.IAppContext) *Repository {
	eventBus := appContext.Get("eventBus").(bus.IEventBus)
	cfg := appContext.Get("config").(*config.CloudConfig)
	inboundRepoChan := appContext.Get("inboundRepoChan").(bus.DataChannel)
	inboundChannelName := appContext.Get(model.INBOUND_CHANNEL_NAME).(string)
	return &Repository{
		EventBus:           eventBus,
		config:             cfg,
		inboundRepoChan:    inboundRepoChan,
		inboundChannelName: inboundChannelName,
	}
}
