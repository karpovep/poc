package daemon

import (
	"log"
	"poc/app"
	"poc/bus"
	"poc/model"
	"poc/nodes"
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
		nodeClient := d.nodeClientProvider.PickClient()
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
	d.EventBus.Subscribe(d.outboundChannelName, d.outboundChan)
}

func (d *Daemon) Stop() {
	d.EventBus.Unsubscribe(d.outboundChannelName, d.outboundChan)
}
