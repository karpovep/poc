package nodes

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"poc/model"
	"poc/protos/nodes"
	"time"
)

type (
	INodeClient interface {
		Transfer(obj *model.InternalServerObject) error
		Start()
		Stop()
	}

	NodeClient struct {
		addr       string
		grpcClient nodes.NodeClient
		conn       *grpc.ClientConn
	}
)

func NewNodeClient(address string) INodeClient {
	nc := &NodeClient{
		addr: address,
	}

	return nc
}

func (c *NodeClient) Transfer(obj *model.InternalServerObject) error {
	if c.grpcClient == nil {
		return fmt.Errorf("cannot transfer object, grpcClient is not connected to another nodeServer, ID=%s", obj.Object.Id)
	}
	log.Println("transfer obj", obj.Object.Id)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := c.grpcClient.Transfer(ctx, &nodes.InternalServerObject{
		Id:     obj.Object.Id,
		Entity: obj.Object.Entity,
	})
	if err != nil {
		log.Println("error transferring object", err)
		return err
	}
	return nil
}

func (c *NodeClient) Start() {
	// Set up a connection to the server.
	// todo run establishing connection in background, handle reconnection logic
	conn, err := grpc.Dial(c.addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("nodeClient can't connect to another nodeServer, address =", c.addr)
	}
	c.conn = conn
	c.grpcClient = nodes.NewNodeClient(c.conn)
	log.Println("nodeClient established connection to", c.addr)
}

func (c *NodeClient) Stop() {
	if c.conn == nil {
		return
	}
	err := c.conn.Close()
	if err != nil {
		log.Println("nodeClient: error closing connection to another node, address = ", c.addr)
	}
	log.Println("nodeClient closed connection to ", c.addr)
}
