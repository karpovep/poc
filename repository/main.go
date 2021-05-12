package repository

import (
	"log"
	"poc/app"
	"poc/bus"
	"poc/config"
	"poc/model"
	"poc/protos/nodes"
	"poc/repository/impls"
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
		Impl               impls.IRepositoryImpl
	}
)

func NewRepository(appContext app.IAppContext) IRepository {
	eventBus := appContext.Get("eventBus").(bus.IEventBus)
	inboundRepoChan := appContext.Get("inboundRepoChan").(bus.DataChannel)
	inboundChannelName := appContext.Get(model.INBOUND_CHANNEL_NAME).(string)
	cfg := appContext.Get("config").(*config.CloudConfig)
	repoFactory := NewRepositoryFactory()
	repoImpl, err := repoFactory.CreateRepository(cfg.Repository.Type, appContext)
	if err != nil {
		log.Fatalln("repoFactory.CreateRepository error", err)
	}
	return &Repository{
		EventBus:           eventBus,
		config:             cfg,
		inboundRepoChan:    inboundRepoChan,
		inboundChannelName: inboundChannelName,
		Impl:               repoImpl,
	}
}

func (r *Repository) Stop() {
	r.EventBus.Unsubscribe(r.inboundChannelName, r.inboundRepoChan)
	r.Impl.Stop()
}

func (r *Repository) Start() {
	r.Impl.Start()
	go r.setupIncomingHandler()
	r.EventBus.Subscribe(r.inboundChannelName, r.inboundRepoChan)
}

func (r *Repository) setupIncomingHandler() {
	for evnt := range r.inboundRepoChan {
		internalServerObject := evnt.Data.(*nodes.ISO)
		err := r.Impl.SaveIso(internalServerObject)
		//todo handle errors
		if err != nil {
			log.Fatalln("SaveIso error", err)
		}
	}
}
