package client

import (
	"fmt"
	"poc/app"
	cloud "poc/protos"
)

type (
	INodeClient interface {
		Send(obj *cloud.CloudObject) error
		Start() error
		Stop() error
	}

	NodeClient struct {
		addr string
	}
)

func NewNodeClient(address string, appContext app.IAppContext) INodeClient {
	nc := &NodeClient {
		addr: address,
	}
	
	return nc
}

func (c *NodeClient) String() string {
	return fmt.Sprintf("NodeClient [addr: %v]", c.addr)
}

func (c *NodeClient) Send(obj *cloud.CloudObject) error {
	return fmt.Errorf("stub error. %v", c)
}

func (c *NodeClient) Start() error {
	return fmt.Errorf("stub error. %v", c)
}

func (c *NodeClient) Stop() error {
	return fmt.Errorf("stub error. %v", c)
}