package server

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"poc/app"
	"poc/bus"
	"poc/config"
	"poc/model"
	"poc/protos/cloud"
	"poc/subscriptions"
	"poc/utils"
)

type (
	IGrpcServer interface {
		Start()
		Stop()
	}

	GrpcServer struct {
		cloud.UnimplementedCloudServer
		EventBus            bus.IEventBus
		SubscriptionManager subscriptions.ISubscriptionManager
		Config              *config.CloudConfig
		Utils               utils.IUtils
		errChan             chan error
		server              *grpc.Server
		inboundChannelName  string
	}
)

func NewGrpcServer(appContext app.IAppContext) IGrpcServer {
	eventBus := appContext.Get("eventBus").(bus.IEventBus)
	subscriptionManager := appContext.Get("subscriptionManager").(subscriptions.ISubscriptionManager)
	utls := appContext.Get("utils").(utils.IUtils)
	cfg := appContext.Get("config").(*config.CloudConfig)
	errChan := appContext.Get("errChan").(chan error)
	inboundChannelName := appContext.Get(model.INBOUND_CHANNEL_NAME).(string)
	return &GrpcServer{
		EventBus:            eventBus,
		SubscriptionManager: subscriptionManager,
		Config:              cfg,
		Utils:               utls,
		errChan:             errChan,
		inboundChannelName:  inboundChannelName,
	}
}

func (s *GrpcServer) Start() {
	go func() {
		lis, err := net.Listen("tcp", s.Config.Server.Port)
		if err != nil {
			s.errChan <- err
			return
		}
		log.WithFields(log.Fields{"port": s.Config.Server.Port}).Info("GRPC Client Server started")
		s.server = grpc.NewServer()
		cloud.RegisterCloudServer(s.server, s)
		if err := s.server.Serve(lis); err != nil {
			s.errChan <- err
		}
	}()
}

func (s *GrpcServer) Stop() {
	s.server.GracefulStop()
}

func (s *GrpcServer) Subscribe(stream cloud.Cloud_SubscribeServer) error {
	// read message
	msg, err := stream.Recv()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("stream.Recv error")
		return nil
	}

	// we expect first message to be SubscribeRequest
	var subscribeRequest cloud.SubscribeRequest
	if err := msg.Entity.UnmarshalTo(&subscribeRequest); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("unmarshal SubscribeRequest from any field error")
		return nil
	}

	clientCloseChan := make(chan bool)
	// subscribe client
	subscriptionId, err := s.SubscriptionManager.RegisterSubscription(subscribeRequest.Type, stream, clientCloseChan)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("RegisterSubscription error")
		return nil
	}

	select {
	case <-clientCloseChan:
		log.WithFields(log.Fields{"subscriptionId": subscriptionId}).Debug("Closing stream to client")
	case <-stream.Context().Done():
		log.WithFields(log.Fields{"subscriptionId": subscriptionId}).Debug("subscriber disconnected")
	}
	s.SubscriptionManager.UnregisterSubscription(subscribeRequest.Type, subscriptionId)
	return nil
}

func (s *GrpcServer) Save(ctx context.Context, cloudObj *cloud.CloudObject) (*cloud.OperationResult, error) {
	log.WithFields(log.Fields{"cloudObj": cloudObj}).Debug("incoming object")
	cloudObj.Id = s.Utils.GenerateTimeUuid()
	iso := model.NewIsoFromCloudObject(cloudObj)
	iso.Metadata.InitialNodeId = s.Config.NodeId
	s.EventBus.Publish(s.inboundChannelName, iso)
	return &cloud.OperationResult{Status: cloud.OperationStatus_OK}, nil
}
