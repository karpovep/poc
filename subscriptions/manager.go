package subscriptions

import (
	"fmt"
	"log"
	"poc/app"
	"poc/bus"
	cloud "poc/protos"
	"poc/utils"
	"sync"
)

type (
	ISubscriptionManager interface {
		RegisterSubscription(objectType string, stream cloud.Cloud_SubscribeServer) error
	}

	Subscriber struct {
		stream cloud.Cloud_SubscribeServer
	}

	SubscriptionManager struct {
		EventBus bus.IEventBus
		Utils    utils.IUtils
		//[objectType] = [subscription-id]*Subscriber
		subscriptions map[string]map[string]*Subscriber
		incomingChan  bus.DataChannel
		mx            sync.RWMutex
	}
)

func NewSubscriptionManager(appContext app.IAppContext) ISubscriptionManager {
	eventBus := appContext.Get("eventBus").(bus.IEventBus)
	utls := appContext.Get("utils").(utils.IUtils)
	incomingChan := appContext.Get("incomingChan").(bus.DataChannel)
	incomingTopic := appContext.Get("incomingTopic").(string)
	sm := &SubscriptionManager{
		EventBus:      eventBus,
		Utils:         utls,
		subscriptions: make(map[string]map[string]*Subscriber),
		incomingChan:  incomingChan,
	}
	sm.setupIncomingHandler()
	sm.EventBus.Subscribe(incomingTopic, incomingChan)
	return sm
}

func (sm *SubscriptionManager) RegisterSubscription(objectType string, stream cloud.Cloud_SubscribeServer) error {
	sm.mx.Lock()
	defer sm.mx.Unlock()
	if _, found := sm.subscriptions[objectType]; !found {
		sm.subscriptions[objectType] = make(map[string]*Subscriber)
	}
	subId := sm.Utils.GenerateUuid()
	sm.subscriptions[objectType][subId] = &Subscriber{stream: stream}
	log.Println("Subscription registered", subId)
	return nil
}

func (sm *SubscriptionManager) setupIncomingHandler() {
	// todo implement me
	go func() {
		for evnt := range sm.incomingChan {
			//todo go over subscriptions by type, lock subscriber, send object via stream and wait for the response from the stream
			//todo if there is no specific subscription or subscriber - publish message via eventBus for the further processing
			fmt.Println(evnt)
		}
	}()
}
