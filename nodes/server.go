package nodes

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"poc/app"
	"poc/bus"
	"poc/config"
	"poc/model"
	"poc/protos/cloud"
	"poc/protos/nodes"
)

type (
	INodeServer interface {
		Start()
		Stop()
	}

	NodeServer struct {
		nodes.UnimplementedNodeServer
		EventBus            bus.IEventBus
		Config              *config.CloudConfig
		errChan             chan error
		transferChannelName string
		server              *grpc.Server
	}
)

func NewNodeServer(appContext app.IAppContext) INodeServer {
	eventBus := appContext.Get("eventBus").(bus.IEventBus)
	cfg := appContext.Get("config").(*config.CloudConfig)
	errChan := appContext.Get("errChan").(chan error)
	inboundChannelName := appContext.Get(model.TRANSFER_CHANNEL_NAME).(string)
	return &NodeServer{
		EventBus:            eventBus,
		Config:              cfg,
		errChan:             errChan,
		transferChannelName: inboundChannelName,
	}
}

func (s *NodeServer) Start() {
	go func() {
		lis, err := net.Listen("tcp", s.Config.NodeServer.Port)
		if err != nil {
			s.errChan <- err
			return
		}
		log.Printf("node server started: %v\n", s.Config.NodeServer.Port)
		s.server = grpc.NewServer()
		nodes.RegisterNodeServer(s.server, s)
		if err := s.server.Serve(lis); err != nil {
			s.errChan <- err
		}
	}()
}

func (s *NodeServer) Stop() {
	s.server.GracefulStop()
}

func (s *NodeServer) Transfer(ctx context.Context, isoObj *nodes.InternalServerObject) (*nodes.Acknowledge, error) {
	log.Println("received obj from transfer", isoObj.Id)
	s.EventBus.Publish(s.transferChannelName, model.NewInternalServerObject(&cloud.CloudObject{
		Id:     isoObj.Id,
		Entity: isoObj.Entity,
	}))
	return &nodes.Acknowledge{}, nil
}
