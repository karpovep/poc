package nodes

import (
	"poc/app"
	"poc/config"
	"poc/protos/nodes"
	"time"
)

const RETRANSFER_DELAY = 5 // 5 seconds to transfer object to the node which had transferred it before

type (
	INodeClientProvider interface {
		PickClient(iso *nodes.ISO) INodeClient
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
		client := NewNodeClient(nodeConfig, p.config.NodeId)
		p.clients = append(p.clients, client)
		go client.Start()
	}
}

func (p *NodeClientProvider) Stop() {
	for _, client := range p.clients {
		client.Stop()
	}
}

func (p *NodeClientProvider) PickClient(iso *nodes.ISO) INodeClient {
	// try to choose the node which hasn't received iso yet
	for _, c := range p.clients {
		if _, ok := iso.TransferredByNodes[c.GetServerNodeId()]; !ok {
			return c
		}
	}

	// try to choose the node which had transferred iso earlier than now - RETRANSFER_DELAY
	for _, c := range p.clients {
		if t, ok := iso.TransferredByNodes[c.GetServerNodeId()]; ok && t < time.Now().Unix()-RETRANSFER_DELAY {
			return c
		}
	}

	// cannot pick any client
	return nil
}
