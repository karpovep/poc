package server

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
		log.Printf("server port: %v\n", s.Config.Server.Port)
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
		log.Println("ERROR: stream.Recv error", err)
		return nil
	}

	// we expect first message to be SubscribeRequest
	var subscribeRequest cloud.SubscribeRequest
	if err := msg.Entity.UnmarshalTo(&subscribeRequest); err != nil {
		log.Println("ERROR: unmarshal SubscribeRequest from any field", err)
		return nil
	}

	clientCloseChan := make(chan bool)
	// subscribe client
	subscriptionId, err := s.SubscriptionManager.RegisterSubscription(subscribeRequest.Type, stream, clientCloseChan)
	if err != nil {
		log.Println("ERROR: RegisterSubscription", err)
		return nil
	}

	select {
	case <-clientCloseChan:
		log.Println("Closing stream for client:", subscriptionId)
	case <-stream.Context().Done():
		log.Println("stream.Context().Done()")
	}
	s.SubscriptionManager.UnregisterSubscription(subscribeRequest.Type, subscriptionId)
	return nil
}

func (s *GrpcServer) Save(ctx context.Context, cloudObj *cloud.CloudObject) (*cloud.OperationResult, error) {
	log.Println("incomingObject", cloudObj)
	cloudObj.Id = s.Utils.GenerateTimeUuid()
	iso := model.NewIsoFromCloudObject(cloudObj)
	iso.Metadata.InitialNodeId = s.Config.NodeId
	s.EventBus.Publish(s.inboundChannelName, iso)
	return &cloud.OperationResult{Status: cloud.OperationStatus_OK}, nil
}
