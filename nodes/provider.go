package nodes

import (
	"log"
	"math/rand"
	"poc/app"
	"poc/config"
)

type (
	INodeClientProvider interface {
		PickClient() INodeClient
		Start()
		Stop()
	}

	NodeClientProvider struct {
		config  *config.CloudConfig
		clients []INodeClient
	}
)

func NewNodeClientProvider(appContext app.IAppContext) INodeClientProvider {
	cfg := appContext.Get("config").(*config.CloudConfig)
	return &NodeClientProvider{
		config:  cfg,
		clients: []INodeClient{},
	}
}

func (p *NodeClientProvider) Start() {
	for _, nodeConfig := range p.config.Server.Nodes {
		client := NewNodeClient(nodeConfig)
		p.clients = append(p.clients, client)
		err := client.Start()
		if err != nil {
			log.Printf("Can't start nodeClient: %v", err)
		}
	}
}

func (p *NodeClientProvider) Stop() {
	for _, client := range p.clients {
		err := client.Stop()
		if err != nil {
			log.Printf("Can't stop nodeClient: %v", err)
		}
	}
}

func (p *NodeClientProvider) PickClient() INodeClient {
	return p.clients[rand.Intn(len(p.clients))] //take random client
}
