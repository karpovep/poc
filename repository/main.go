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
		ResetActiveIsoNodeId(iso *nodes.ISO) error
	}

	Repository struct {
		EventBus               bus.IEventBus
		config                 *config.CloudConfig
		inboundChan            bus.DataChannel
		inboundChannelName     string
		unprocessedChannelName string
		Impl                   impls.IRepositoryImpl
	}
)

func NewRepository(appContext app.IAppContext) IRepository {
	eventBus := appContext.Get("eventBus").(bus.IEventBus)
	inboundRepoChan := eventBus.CreateDataChannel()
	inboundChannelName := appContext.Get(model.INBOUND_CHANNEL_NAME).(string)
	unprocessedChannelName := appContext.Get(model.UNPROCESSED_CHANNEL_NAME).(string)
	cfg := appContext.Get("config").(*config.CloudConfig)
	repoFactory := NewRepositoryFactory()
	repoImpl, err := repoFactory.CreateRepository(cfg.Repository.Type, appContext)
	if err != nil {
		log.Fatalln("repoFactory.CreateRepository error", err)
	}
	return &Repository{
		EventBus:               eventBus,
		config:                 cfg,
		inboundChan:            inboundRepoChan,
		inboundChannelName:     inboundChannelName,
		unprocessedChannelName: unprocessedChannelName,
		Impl:                   repoImpl,
	}
}

func (r *Repository) Stop() {
	r.EventBus.Unsubscribe(r.inboundChannelName, r.inboundChan)
	r.Impl.Stop()
}

func (r *Repository) Start() {
	r.Impl.Start()
	go r.setupIncomingHandler()
	go r.loadActiveIso()
	r.EventBus.Subscribe(r.inboundChannelName, r.inboundChan)
}

func (r *Repository) ResetActiveIsoNodeId(iso *nodes.ISO) error {
	return r.Impl.ResetActiveIsoNodeId(iso)
}

func (r *Repository) setupIncomingHandler() {
	for evnt := range r.inboundChan {
		internalServerObject := evnt.Data.(*nodes.ISO)
		err := r.Impl.SaveIso(internalServerObject)
		//todo handle errors
		if err != nil {
			log.Fatalln("SaveIso error", err)
		}
	}
}

// loads all active ISOs related to configured NodeId and checks whether it is processed or not. Removes active ISO
// it it was already processed and publishes it to unprocessed channel for processing otherwise
func (r *Repository) loadActiveIso() {
	log.Println("Start loading active ISOs from DB...")
	var nextPage []byte
	var activeIsoList []*nodes.ISO
	var err error
	for ok := true; ok; ok = len(nextPage) > 0 {
		activeIsoList, nextPage, err = r.Impl.ListActiveIso(r.config.NodeId, 10, nil)
		if err != nil {
			log.Fatalln("r.Impl.ListActiveIso error", err)
		}

		for _, activeIso := range activeIsoList {
			iso, err := r.Impl.FindIsoByTypeAndId(activeIso.CloudObj.Entity.TypeUrl, activeIso.CloudObj.Id)
			if err != nil {
				log.Fatalln("r.Impl.FindIsoByTypeAndId error", err)
			}
			if iso.CloudObj.IsFinal {
				err = r.Impl.DeleteActiveIso(iso)
				if err != nil {
					log.Fatalln("r.Impl.DeleteActiveIso error", err)
				}
			} else {
				r.EventBus.Publish(r.unprocessedChannelName, iso)
			}
		}
	}
	log.Println("Loading of active ISOs from DB has been finished")
}
