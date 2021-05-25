package nodes

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"poc/app"
	"poc/bus"
	"poc/config"
	"poc/model"
	"poc/protos/nodes"
	"poc/repository"
)

type (
	INodeServer interface {
		Start()
		Stop()
	}

	NodeServer struct {
		nodes.UnimplementedNodeServer
		EventBus            bus.IEventBus
		Repo                repository.IRepository
		Config              *config.CloudConfig
		errChan             chan error
		transferChannelName string
		server              *grpc.Server
	}
)

func NewNodeServer(appContext app.IAppContext) INodeServer {
	eventBus := appContext.Get("eventBus").(bus.IEventBus)
	repo := appContext.Get("repository").(repository.IRepository)
	cfg := appContext.Get("config").(*config.CloudConfig)
	errChan := appContext.Get("errChan").(chan error)
	inboundChannelName := appContext.Get(model.TRANSFER_CHANNEL_NAME).(string)
	return &NodeServer{
		EventBus:            eventBus,
		Repo:                repo,
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
		log.WithFields(log.Fields{"port": s.Config.NodeServer.Port}).Info("GRPC Node Server started")
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

func (s *NodeServer) Transfer(ctx context.Context, iso *nodes.ISO) (*nodes.Acknowledge, error) {
	log.WithFields(log.Fields{"id": iso.CloudObj.Id}).Debug("received obj from transfer")
	err := s.Repo.ResetActiveIsoNodeId(iso)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("ResetActiveIsoNodeId error")
		return nil, err
	}
	s.EventBus.Publish(s.transferChannelName, iso)
	return &nodes.Acknowledge{}, nil
}

func (s *NodeServer) GetInfo(ctx context.Context, infoReq *nodes.NodeInfoRequest) (*nodes.NodeInfoResponse, error) {
	log.Debug("GetInfo request")
	return &nodes.NodeInfoResponse{
		Id: s.Config.NodeId,
	}, nil
}
