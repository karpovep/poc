package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
	"io"
	"log"
	"net/http"
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

	leaderboard := map[string]int32{}

	go func() {
		time.Sleep(time.Second * 5)
		subscribeRequest := &cloud.SubscribeRequest{}
		dummyEntity := &cloud.ParticipantPointsEntity{}
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
			var ent cloud.ParticipantPointsEntity
			if err := obj.Entity.UnmarshalTo(&ent); err != nil {
				log.Fatalf("Could not unmarshal ParticipantPointsEntity from any field: %s", err)
			}

			log.Println("score message:", ent.Name, ent.Points)
			score, ok := leaderboard[ent.Name]
			if !ok {
				score = 0
			}
			leaderboard[ent.Name] = score + ent.Points

			time.Sleep(time.Second * 3)
			ack := &cloud.Acknowledge{FinalizeObject: true}
			serializedAck, _ := proto.Marshal(ack)
			msg := &anypb.Any{TypeUrl: string(ack.ProtoReflect().Descriptor().FullName()), Value: serializedAck}
			_ = stream.Send(&cloud.CloudObject{Entity: msg})
			log.Println("ACK SENT - ", ent.Name)
		}
	}()

	//broadcast leaderboard info
	go func() {
		for {
			time.Sleep(time.Second * 10)
			val := &cloud.Leaderboard{Participants: make([]*cloud.Participant, len(leaderboard))}
			for key, value := range leaderboard {
				val.Participants = append(val.Participants, &cloud.Participant{
					Name:  key,
					Score: value,
				})
			}
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

	go func() {
		myRouter := mux.NewRouter().StrictSlash(true)
		myRouter.HandleFunc("/leaderboard", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(leaderboard)
		})
		log.Fatal(http.ListenAndServe(":3000", myRouter))
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
