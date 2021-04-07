package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"poc/app"
	"poc/bus"
	"poc/config"
	cloud "poc/protos"
	"poc/subscriptions"
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
		errChan             chan error
		server              *grpc.Server
	}
)

func NewGrpcServer(appContext app.IAppContext) IGrpcServer {
	eventBus := appContext.Get("eventBus").(bus.IEventBus)
	subscriptionManager := appContext.Get("subscriptionManager").(subscriptions.ISubscriptionManager)
	cfg := appContext.Get("config").(*config.CloudConfig)
	errChan := appContext.Get("errChan").(chan error)
	return &GrpcServer{
		EventBus:            eventBus,
		SubscriptionManager: subscriptionManager,
		Config:              cfg,
		errChan:             errChan,
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

	// subscribe client
	err = s.SubscriptionManager.RegisterSubscription(subscribeRequest.Type, stream)
	if err != nil {
		log.Println("ERROR: RegisterSubscription", err)
		return nil
	}

	return nil
}

func (s *GrpcServer) Commit(ctx context.Context, in *cloud.CloudObject) (*cloud.OperationResult, error) {
	fmt.Println(in)
	return &cloud.OperationResult{Status: cloud.OperationStatus_OK}, nil
}
