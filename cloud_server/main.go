package main

import (
	"context"
	"fmt"
	cloud "google.golang.org/protos"
	"log"
	"net"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const (
	port = ":50051" // todo - config
)

// server is used to implement Cloud
type server struct {
	cloud.UnimplementedCloudServer
	objects map[string]*cloud.CloudObject
}

func NewServer() *server {
	return &server{
		objects: make(map[string]*cloud.CloudObject),
	}
}

func (s *server) Commit(ctx context.Context, in *cloud.CloudObject) (*cloud.OperationResult, error) {
	fmt.Println(in)
	s.objects[uuid.New().String()] = in
	return &cloud.OperationResult{Status: cloud.OperationStatus_OK}, nil
}

func (s *server) Subscribe(in *cloud.SubscribeRequest, stream cloud.Cloud_SubscribeServer) error {
	for {
		timer := time.NewTimer(time.Second * 2)
		select {
		case <-timer.C:
			fmt.Println("checking subscriptions...")
			for _, obj := range s.objects {
				fmt.Println("obj", obj.Entity)
				if obj.Entity.TypeUrl == in.Type {
					err := stream.Send(obj)
					if err != nil {
						log.Fatalf("stream.Send: %v", err)
					}
				}
			}
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	cloud.RegisterCloudServer(s, NewServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
