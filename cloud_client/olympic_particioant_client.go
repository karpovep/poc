package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
	"io"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"poc/config"
	"poc/protos/cloud"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func main() {
	cfg := config.Init("config-client-leaderboard.yml").Client

	// Set up a connection to the server.
	conn, err := grpc.NewClient(cfg.ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println("conn.Close() err ", err.Error())
		}
	}(conn)
	c := cloud.NewCloudClient(conn)

	go func() {
		time.Sleep(time.Second)
		subscribeRequest := &cloud.SubscribeRequest{}
		dummyEntity := &cloud.Leaderboard{}
		typeToSubscribeTo := string(dummyEntity.ProtoReflect().Descriptor().FullName())
		subscribeRequest.Type = typeToSubscribeTo
		serializedReq, err := proto.Marshal(subscribeRequest)
		if err != nil {
			log.Fatal("could not serialize", err)
		}
		msg := &anypb.Any{TypeUrl: string(subscribeRequest.ProtoReflect().Descriptor().FullName()), Value: serializedReq}
		stream, err := c.Subscribe(context.Background())
		if err != nil {
			log.Fatal("Subscribe error", err)
		}
		err = stream.Send(&cloud.CloudObject{Entity: msg})
		if err != nil {
			log.Fatalf("stream.Send: %v", err)
		}
		for {
			obj, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal("stream.Recv error", err)
			}
			var ent cloud.Leaderboard
			if err := obj.Entity.UnmarshalTo(&ent); err != nil {
				log.Fatalf("Could not unmarshal Leaderboard from any field: %s", err)
			}

			log.Println("Leaderboard message:", ent)
			ack := &cloud.Acknowledge{}
			serializedAck, _ := proto.Marshal(ack)
			msg := &anypb.Any{TypeUrl: string(ack.ProtoReflect().Descriptor().FullName()), Value: serializedAck}
			_ = stream.Send(&cloud.CloudObject{Entity: msg})
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second / 2)
			val := &cloud.ParticipantPointsEntity{Name: cfg.ClientId, Points: int32(rand.Intn(100))}
			serialized, err := proto.Marshal(val)
			if err != nil {
				log.Fatal("could not serialize", err)
			}
			msg := &anypb.Any{TypeUrl: string(val.ProtoReflect().Descriptor().FullName()), Value: serialized}
			opRes, err := c.Save(context.Background(), &cloud.CloudObject{Entity: msg})
			if err != nil {
				log.Fatalf("could not save: %v", err)
			}
			log.Printf("OperationResult: %s", opRes.Status)
		}
	}()

	//hang the process
	done := make(chan bool, 1)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interrupt
		done <- true
		os.Exit(0)
	}()
	<-done
}
