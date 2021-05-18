package daemon

import (
	"log"
	"poc/app"
	"poc/bus"
	"poc/model"
	"poc/nodes"
	nodes_protoc "poc/protos/nodes"
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
		nodeClientProvider     nodes.INodeClientProvider
	}
)

func NewDaemon(appContext app.IAppContext) IDaemon {
	eventBus := appContext.Get("eventBus").(bus.IEventBus)
	outboundChan := appContext.Get("outboundChan").(bus.DataChannel)
	outboundChannelName := appContext.Get(model.OUTBOUND_CHANNEL_NAME).(string)
	unprocessedChannelName := appContext.Get(model.UNPROCESSED_CHANNEL_NAME).(string)
	nodeClientProvider := appContext.Get("nodeClientProvider").(nodes.INodeClientProvider)

	daemon := &Daemon{
		EventBus:               eventBus,
		outboundChan:           outboundChan,
		outboundChannelName:    outboundChannelName,
		unprocessedChannelName: unprocessedChannelName,
		nodeClientProvider:     nodeClientProvider,
	}

	go daemon.startEventHandler()
	return daemon
}

func (d *Daemon) startEventHandler() {
	for event := range d.outboundChan {
		iso := event.Data.(*nodes_protoc.ISO)
		nodeClient := d.nodeClientProvider.PickClient(iso)
		if nodeClient == nil {
			d.EventBus.Publish(d.unprocessedChannelName, event.Data)
			continue
		}
		err := nodeClient.Transfer(iso)

		if err != nil {
			log.Printf("Can not send %v to %v, err = %v", iso, nodeClient, err)
			//return unprocessed message back to application
			d.EventBus.Publish(d.unprocessedChannelName, event.Data)
		}
	}
}

func (d *Daemon) Start() {
	d.EventBus.Subscribe(d.outboundChannelName, d.outboundChan)
}

func (d *Daemon) Stop() {
	d.EventBus.Unsubscribe(d.outboundChannelName, d.outboundChan)
}
