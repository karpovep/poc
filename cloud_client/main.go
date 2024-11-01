package main

import (
	"context"
	"google.golang.org/protobuf/types/known/anypb"
	"io"
	"log"
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
	cfg := config.Init("config.yml").Client

	// Set up a connection to the server.
	conn, err := grpc.Dial(cfg.ServerAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := cloud.NewCloudClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	val := &cloud.TestEntity{Name: "First Cloud Client"}

	go func() {
		time.Sleep(time.Second * 5) // subscribe after 3 seconds
		subscribeRequest := &cloud.SubscribeRequest{}
		typeToSubscribeTo := string(val.ProtoReflect().Descriptor().FullName())
		subscribeRequest.Type = typeToSubscribeTo
		serializedReq, err := proto.Marshal(subscribeRequest)
		if err != nil {
			log.Fatal("could not serialize", err)
		}
		msg := &anypb.Any{TypeUrl: string(subscribeRequest.ProtoReflect().Descriptor().FullName()), Value: serializedReq}
		//subscribeContext, cancelSubCtx := context.WithCancel(context.Background())
		//defer cancelSubCtx()
		//stream, err := c.Subscribe(subscribeContext)
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
			var ent cloud.TestEntity
			if err := obj.Entity.UnmarshalTo(&ent); err != nil {
				log.Fatalf("Could not unmarshal TestEntity from any field: %s", err)
			}

			log.Println("received obj", ent.Name)

			time.Sleep(time.Second * 3)
			ack := &cloud.Acknowledge{}
			serializedAck, _ := proto.Marshal(ack)
			msg := &anypb.Any{TypeUrl: string(ack.ProtoReflect().Descriptor().FullName()), Value: serializedAck}
			_ = stream.Send(&cloud.CloudObject{Entity: msg})
			log.Println("ACK SENT - ", ent.Name)
		}
	}()

	time.Sleep(time.Second)
	serialized, err := proto.Marshal(val)
	if err != nil {
		log.Fatal("could not serialize", err)
	}
	msg := &anypb.Any{TypeUrl: string(val.ProtoReflect().Descriptor().FullName()), Value: serialized}
	opRes, err := c.Save(ctx, &cloud.CloudObject{Entity: msg})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("OperationResult: %s", opRes.Status)

	//hang the process
	done := make(chan bool, 1)
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interrupt
		done <- true
		os.Exit(0)
	}()
	<-done
}
