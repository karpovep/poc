package retry

import (
	"poc/app"
	"poc/bus"
	"poc/model"
	"time"
)

type (
	IRetryResolver interface {
		ProcessRetryableObject(obj *model.InternalServerObject)
		Stop()
	}

	RetryResolver struct {
		EventBus               bus.IEventBus
		unprocessedChan        bus.DataChannel
		unprocessedChannelName string
		retryChannelName       string
	}
)

func NewRetryResolver(appContext app.IAppContext) IRetryResolver {
	eventBus := appContext.Get("eventBus").(bus.IEventBus)
	unprocessedChan := appContext.Get("unprocessedChan").(bus.DataChannel)
	unprocessedChannelName := appContext.Get(model.UNPROCESSED_CHANNEL_NAME).(string)
	retryChannelName := appContext.Get(model.RETRY_CHANNEL_NAME).(string)
	r := &RetryResolver{
		EventBus:               eventBus,
		unprocessedChan:        unprocessedChan,
		unprocessedChannelName: unprocessedChannelName,
		retryChannelName:       retryChannelName,
	}
	go r.setupUnprocessedHandler()
	r.EventBus.Subscribe(unprocessedChannelName, unprocessedChan)
	return r
}

func (r *RetryResolver) setupUnprocessedHandler() {
	for evnt := range r.unprocessedChan {
		internalServerObject := evnt.Data.(*model.InternalServerObject)
		r.ProcessRetryableObject(internalServerObject)
	}
}

func (r *RetryResolver) ProcessRetryableObject(obj *model.InternalServerObject) {
	// todo - implement retry logic based on reason provided by other processors and historical data
	obj.Metadata.RetryIn = time.Second
	r.EventBus.Publish(r.retryChannelName, obj)
}

func (r *RetryResolver) Stop() {
	r.EventBus.Unsubscribe(r.unprocessedChannelName, r.unprocessedChan)
}
