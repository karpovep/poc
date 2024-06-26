package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"poc/app"
	"poc/bus"
	cache2 "poc/cache"
	"poc/config"
	daemon2 "poc/daemon"
	"poc/logger"
	"poc/nodes"
	"poc/repository"
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

	cfg := config.Init("config-server-node-3.yml")
	logger.Init(cfg.Logger)

	// start the app
	log.Info("Starting the app...")

	appContext := app.NewApplicationContext()
	utils := utilsPkg.NewUtils()
	eventBus := bus.NewEventBus(appContext)

	appContext.Set("errChan", errChan)
	appContext.Set("config", cfg)
	appContext.Set("utils", utils)
	appContext.Set("eventBus", eventBus)

	cancellableTimer := utilsPkg.NewCancellableTimer()
	appContext.Set("cacheTimer", cancellableTimer)
	cache := cache2.NewCache(appContext)
	appContext.Set("cache", cache)

	nodeClientProvider := nodes.NewNodeClientProvider(appContext)
	appContext.Set("nodeClientProvider", nodeClientProvider)
	nodeClientProvider.Start()
	daemon := daemon2.NewDaemon(appContext)
	appContext.Set("daemon", daemon)
	daemon.Start()

	retryResolver := retry.NewRetryResolver(appContext)
	appContext.Set("retryResolver", retryResolver)

	subscriptionManager := subscriptions.NewSubscriptionManager(appContext)
	appContext.Set("subscriptionManager", subscriptionManager)

	repo := repository.NewRepository(appContext)
	appContext.Set("repository", repo)
	repo.Start()

	nodeServer := nodes.NewNodeServer(appContext)
	nodeServer.Start()
	grpcServer := server.NewGrpcServer(appContext)

	defer func() {
		log.Info("Stopping the app...")
		// do graceful stop of required resources here in right order
		nodeServer.Stop()
		cache.Stop()
		retryResolver.Stop()
		subscriptionManager.Stop()
		daemon.Stop()
		nodeClientProvider.Stop()
		//cassandraRepository.Stop()
		grpcServer.Stop()
		log.Info("App has been stopped")
	}()

	grpcServer.Start()

	// block until either OS signal, or fatal error
	select {
	case err := <-errChan:
		log.WithFields(log.Fields{"error": err}).Fatal("Fatal error")
	case sig := <-stopChan:
		log.WithFields(log.Fields{"sig": sig}).Info("Received OS signal")
	}
}
