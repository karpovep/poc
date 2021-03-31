package main

import (
	"context"
	"google.golang.org/protobuf/types/known/anypb"
	cloud "google.golang.org/protos"
	"io"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051" // todo - config
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := cloud.NewCloudClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	val := &cloud.TestEntity{Name: "First Cloud Client"}
	typeToSubscribeTo := "cloud/" + proto.MessageName(val)
	go func() {
		stream, err := c.Subscribe(context.Background(), &cloud.SubscribeRequest{Type: typeToSubscribeTo})
		if err != nil {
			log.Fatal("Subscribe error", err)
		}
		for {
			obj, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("stream.Recv error", err)
			}
			log.Println(obj)
		}
	}()

	serialized, err := proto.Marshal(val)
	if err != nil {
		log.Fatal("could not serialize", err)
	}
	msg := &anypb.Any{TypeUrl: "cloud/" + proto.MessageName(val), Value: serialized}
	opRes, err := c.Commit(ctx, &cloud.CloudObject{Entity: msg})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("OperationResult: %s", opRes.Status)

	//hang the process
	done := make(chan bool, 1)
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	go func() {
		<-interrupt
		done <- true
		os.Exit(0)
	}()
	<-done
}
