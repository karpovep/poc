package cache

import (
	"log"
	"poc/app"
	"poc/bus"
	"poc/model"
	"poc/utils"
)

type (
	ICache interface {
		ScheduleProcessing(obj *model.InternalServerObject)
		Stop()
	}

	Cache struct {
		EventBus          bus.IEventBus
		retryChan         bus.DataChannel
		retryChannelName  string
		cachedChannelName string
		timer             utils.ICancellableTimer
	}
)

func NewCache(appContext app.IAppContext) ICache {
	eventBus := appContext.Get("eventBus").(bus.IEventBus)
	retryChan := appContext.Get("retryChan").(bus.DataChannel)
	retryChannelName := appContext.Get(model.RETRY_CHANNEL_NAME).(string)
	cachedChannelName := appContext.Get(model.CACHED_CHANNEL_NAME).(string)
	cacheTimer := appContext.Get("cacheTimer").(utils.ICancellableTimer)
	cache := &Cache{
		EventBus:          eventBus,
		retryChan:         retryChan,
		retryChannelName:  retryChannelName,
		cachedChannelName: cachedChannelName,
		timer:             cacheTimer,
	}
	go cache.setupRetryHandler()
	cache.EventBus.Subscribe(retryChannelName, retryChan)
	return cache
}

func (cache *Cache) setupRetryHandler() {
	for evnt := range cache.retryChan {
		internalServerObject := evnt.Data.(*model.InternalServerObject)
		go cache.ScheduleProcessing(internalServerObject)
	}
}

func (cache *Cache) ScheduleProcessing(obj *model.InternalServerObject) {
	timedOut := <-cache.timer.After(obj.Metadata.RetryIn)
	if timedOut {
		log.Println("cache: publish obj")
		cache.EventBus.Publish(cache.cachedChannelName, obj)
	}
}

func (cache *Cache) Stop() {
	cache.EventBus.Unsubscribe(cache.retryChannelName, cache.retryChan)
	cache.timer.Cancel()
}
