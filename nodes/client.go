package nodes

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"poc/protos/nodes"
	"time"
)

type (
	INodeClient interface {
		Transfer(obj *nodes.ISO) error
		GetServerNodeId() string
		Start()
		Stop()
	}

	NodeClient struct {
		addr         string
		nodeId       string
		ServerNodeId string
		grpcClient   nodes.NodeClient
		conn         *grpc.ClientConn
	}
)

func NewNodeClient(address string, nodeId string) INodeClient {
	return &NodeClient{
		addr:   address,
		nodeId: nodeId,
	}
}

func (c *NodeClient) Transfer(iso *nodes.ISO) error {
	if c.grpcClient == nil {
		return fmt.Errorf("cannot transfer object, grpcClient is not connected to another nodeServer, ID=%s", iso.CloudObj.Id)
	}
	log.Println("transfer obj", iso.CloudObj.Id)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	iso.TransferredByNodes[c.nodeId] = time.Now().Unix()
	_, err := c.grpcClient.Transfer(ctx, iso)
	if err != nil {
		log.Println("error transferring object", err)
		return err
	}
	return nil
}

func (c *NodeClient) GetServerNodeId() string {
	return c.ServerNodeId
}

func (c *NodeClient) Start() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(c.addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("nodeClient can't connect to another nodeServer, address =", c.addr)
	}
	c.conn = conn
	c.grpcClient = nodes.NewNodeClient(c.conn)
	log.Println("nodeClient established connection to", c.addr)
	// request for server info
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	infoResponse, err := c.grpcClient.GetInfo(ctx, &nodes.NodeInfoRequest{})
	if err != nil {
		log.Fatalln("GetInfo error", err)
	}
	c.ServerNodeId = infoResponse.Id
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
