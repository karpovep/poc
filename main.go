package main

import (
	"log"
	"os"
	"os/signal"
	"poc/app"
	"poc/bus"
	cache2 "poc/cache"
	"poc/config"
	"poc/model"
	"poc/retry"
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
	inboundChannelName := "inbound"
	outboundChannelName := "outbound"
	processedChannelName := "processed"
	unprocessedChannelName := "unprocessed"
	retryChannelName := "retry"
	cachedChannelName := "cached"
	inboundChan := make(bus.DataChannel)
	unprocessedChan := make(bus.DataChannel)
	retryChan := make(bus.DataChannel)

	appContext.Set("errChan", errChan)
	appContext.Set("config", cfg)
	appContext.Set("utils", utils)
	appContext.Set("eventBus", eventBus)
	appContext.Set(model.INBOUND_CHANNEL_NAME, inboundChannelName)
	appContext.Set(model.OUTBOUND_CHANNEL_NAME, outboundChannelName)
	appContext.Set(model.PROCESSED_CHANNEL_NAME, processedChannelName)
	appContext.Set(model.UNPROCESSED_CHANNEL_NAME, unprocessedChannelName)
	appContext.Set(model.RETRY_CHANNEL_NAME, retryChannelName)
	appContext.Set(model.CACHED_CHANNEL_NAME, cachedChannelName)
	appContext.Set("inboundChan", inboundChan)
	appContext.Set("unprocessedChan", unprocessedChan)
	appContext.Set("retryChan", retryChan)

	cancellableTimer := utilsPkg.NewCancellableTimer()
	appContext.Set("cacheTimer", cancellableTimer)
	cache := cache2.NewCache(appContext)
	appContext.Set("cache", cache)

	retryResolver := retry.NewRetryResolver(appContext)
	appContext.Set("retryResolver", retryResolver)

	subscriptionManager := subscriptions.NewSubscriptionManager(appContext)
	appContext.Set("subscriptionManager", subscriptionManager)

	grpcServer := server.NewGrpcServer(appContext)

	defer func() {
		log.Println("Stopping the app...")
		// do graceful stop of required resources here in right order
		cache.Stop()
		retryResolver.Stop()
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
