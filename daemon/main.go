package daemon

import (
	"log"
	"poc/app"
	"poc/bus"
	"poc/config"
	"poc/daemon/client"
	cloud "poc/protos"
)

type (
	IDaemon interface {
		Start()
		Stop()
	}

	Daemon struct {
		EventBus bus.IEventBus
		incomingChan bus.DataChannel
		incomingTopic string
		outcomingTopic string
		nodeClients []client.INodeClient
	}
)

func NewDaemon(appContext app.IAppContext) IDaemon { 
	eventBus := appContext.Get("eventBus").(bus.IEventBus)
	incomingTopic := appContext.Get("daemonIncomingTopic").(string)
	outcomingTopic := appContext.Get("daemonOutcomigTopic").(string)
	config := appContext.Get("config").(*config.CloudConfig)

	daemon := &Daemon {
		EventBus: eventBus,
		incomingChan: make(bus.DataChannel),
		incomingTopic: incomingTopic,
		outcomingTopic: outcomingTopic,
		nodeClients: []client.INodeClient {},
	}

	for _, nodeConfig := range config.Server.Nodes {
		daemon.nodeClients = append(daemon.nodeClients, client.NewNodeClient(nodeConfig, appContext))
	}

	go daemon.startEventHandler()
	return daemon
}

func (d* Daemon) startEventHandler() {
	for event := range d.incomingChan {
		nodeClient := d.pickClient()
		cloudObj := event.Data.(*cloud.CloudObject)
		err := nodeClient.Send(cloudObj)

		if err != nil {
			log.Printf("Can not send %v to %v", cloudObj, nodeClient)
			//return unprocessed message back to application 
			d.EventBus.Publish(d.outcomingTopic, event.Data)
		}
	}
}

func (d* Daemon) Start() {
	for _, client := range d.nodeClients {
		err := client.Start()
		
		if err != nil {
			log.Printf("Cant start nodeClient: %v", err)
		}
	}
	
	d.EventBus.Subscribe(d.incomingTopic, d.incomingChan)
}

func (d* Daemon) Stop() {
	d.EventBus.Unsubscribe(d.incomingTopic, d.incomingChan)

	for _, client := range d.nodeClients {
		err := client.Stop()
		
		if err != nil {
			log.Printf("Cant stop nodeClient: %v", err)
		}
	}
}