package main

import (
	"log"
	"os"
	"os/signal"
	"poc/app"
	"poc/bus"
	"poc/config"
	"poc/server"
	"poc/subscriptions"
	utilsPkg "poc/utils"
	"syscall"
)

func main() {
	errChan := make(chan error)
	stopChan := make(chan os.Signal, 1)

	// bind OS events to the signal channel
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	cfg := config.Init("config.yml")
	appContext := app.NewApplicationContext()
	utils := utilsPkg.NewUtils()
	eventBus := bus.NewEventBus()
	incomingTopic := "incoming"
	incomingChan := make(bus.DataChannel)

	appContext.Set("errChan", errChan)
	appContext.Set("config", cfg)
	appContext.Set("utils", utils)
	appContext.Set("eventBus", eventBus)
	appContext.Set("incomingTopic", incomingTopic)
	appContext.Set("incomingChan", incomingChan)

	subscriptionManager := subscriptions.NewSubscriptionManager(appContext)
	appContext.Set("subscriptionManager", subscriptionManager)

	grpcServer := server.NewGrpcServer(appContext)

	defer func() {
		log.Println("Stopping the app...")
		// do graceful stop of required resources here in right order
		subscriptionManager.Stop()
		grpcServer.Stop()
		log.Println("App has been stopped")
	}()

	// start the app
	log.Println("Starting the app...")
	grpcServer.Start()
	log.Println("App has been started")

	// block until either OS signal, or fatal error
	select {
	case err := <-errChan:
		log.Printf("Fatal error: %v\n", err)
	case sig := <-stopChan:
		log.Printf("Received OS signal: %v\n", sig)
	}
}
