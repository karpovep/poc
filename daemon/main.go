package daemon

import (
	"log"
	"poc/app"
	"poc/bus"
	"poc/config"
	"poc/daemon/client"
	"poc/model"
)

type (
	IDaemon interface {
		Start()
		Stop()
	}

	Daemon struct {
		EventBus               bus.IEventBus
		outboundChan           bus.DataChannel
		outboundChannelName    string
		unprocessedChannelName string
		nodeClients            []client.INodeClient
	}
)

func NewDaemon(appContext app.IAppContext) IDaemon {
	eventBus := appContext.Get("eventBus").(bus.IEventBus)
	outboundChan := appContext.Get("outboundChan").(bus.DataChannel)
	outboundChannelName := appContext.Get(model.OUTBOUND_CHANNEL_NAME).(string)
	unprocessedChannelName := appContext.Get(model.UNPROCESSED_CHANNEL_NAME).(string)
	cfg := appContext.Get("config").(*config.CloudConfig)

	daemon := &Daemon{
		EventBus:               eventBus,
		outboundChan:           outboundChan,
		outboundChannelName:    outboundChannelName,
		unprocessedChannelName: unprocessedChannelName,
		nodeClients:            []client.INodeClient{},
	}

	for _, nodeConfig := range cfg.Server.Nodes {
		daemon.nodeClients = append(daemon.nodeClients, client.NewNodeClient(nodeConfig, appContext))
	}

	go daemon.startEventHandler()
	return daemon
}

func (d *Daemon) startEventHandler() {
	for event := range d.outboundChan {
		nodeClient := d.pickClient()
		internalServerObject := event.Data.(*model.InternalServerObject)
		err := nodeClient.Send(internalServerObject)

		if err != nil {
			log.Printf("Can not send %v to %v", internalServerObject, nodeClient)
			//return unprocessed message back to application
			d.EventBus.Publish(d.unprocessedChannelName, event.Data)
		}
	}
}

func (d *Daemon) Start() {
	for _, nodeClient := range d.nodeClients {
		err := nodeClient.Start()

		if err != nil {
			log.Printf("Cant start nodeClient: %v", err)
		}
	}

	d.EventBus.Subscribe(d.outboundChannelName, d.outboundChan)
}

func (d *Daemon) Stop() {
	d.EventBus.Unsubscribe(d.outboundChannelName, d.outboundChan)

	for _, nodeClient := range d.nodeClients {
		err := nodeClient.Stop()

		if err != nil {
			log.Printf("Cant stop nodeClient: %v", err)
		}
	}
}
