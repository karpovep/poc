package client

import (
	"fmt"
	"poc/app"
	"poc/model"
)

type (
	INodeClient interface {
		Send(obj *model.InternalServerObject) error
		Start() error
		Stop() error
	}

	NodeClient struct {
		addr string
	}
)

func NewNodeClient(address string, appContext app.IAppContext) INodeClient {
	nc := &NodeClient{
		addr: address,
	}

	return nc
}

func (c *NodeClient) String() string {
	return fmt.Sprintf("NodeClient [addr: %v]", c.addr)
}

func (c *NodeClient) Send(obj *model.InternalServerObject) error {
	return fmt.Errorf("stub error. %v", c)
}

func (c *NodeClient) Start() error {
	return fmt.Errorf("stub error. %v", c)
}

func (c *NodeClient) Stop() error {
	return fmt.Errorf("stub error. %v", c)
}
