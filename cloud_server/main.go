package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"poc/config"
	cloud "poc/protos"
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

func (s *server) Subscribe(stream cloud.Cloud_SubscribeServer) error {
	for {
		obj, err := stream.Recv()
		if err == io.EOF {
			//todo implement logic of closing the stream from client side
			break
		}
		if err != nil {
			log.Fatal("stream.Recv error", err)
		}
		//var ent cloud.TestEntity
		//if err := obj.Entity.UnmarshalTo(&ent); err != nil {
		//	log.Fatalf("Could not unmarshal TestEntity from any field: %s", err)
		//}

		log.Println(obj)
		msg, err := obj.Entity.UnmarshalNew()
		if err != nil {
			log.Fatalf("UnmarshalNew: %s", err)
		}
		log.Println(msg.ProtoReflect().Type())

		//
		//
		//time.Sleep(time.Second * 2)
		//fmt.Println("checking subscriptions...")
		//for _, obj := range s.objects {
		//	fmt.Println("obj", obj.Entity)
		//	if obj.Entity.TypeUrl == in.Type {
		//		err := stream.Send(obj)
		//		if err != nil {
		//			log.Fatalf("stream.Send: %v", err)
		//		}
		//	}
		//}
	}
	return nil
}

func main() {
	cfg := config.Init("config.yml").Server

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	cloud.RegisterCloudServer(s, NewServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
