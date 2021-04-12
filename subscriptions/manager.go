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
		RegisterSubscription(objectType string, stream cloud.Cloud_SubscribeServer, closeCh chan bool) (string, error)
		UnregisterSubscription(objectType string, subId string)
		Stop()
	}

	Subscriber struct {
		stream    cloud.Cloud_SubscribeServer
		closeChan chan bool
	}

	SubscriptionManager struct {
		EventBus      bus.IEventBus
		Utils         utils.IUtils
		subscriptions map[string]map[string]*Subscriber //[objectType][subscriptionId]*Subscriber
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

func (sm *SubscriptionManager) RegisterSubscription(objectType string, stream cloud.Cloud_SubscribeServer, closeCh chan bool) (string, error) {
	sm.mx.Lock()
	defer sm.mx.Unlock()
	if _, found := sm.subscriptions[objectType]; !found {
		sm.subscriptions[objectType] = make(map[string]*Subscriber)
	}
	subId := sm.Utils.GenerateUuid()
	sm.subscriptions[objectType][subId] = &Subscriber{stream: stream, closeChan: closeCh}
	log.Println("Subscription registered", subId)
	return subId, nil
}

func (sm *SubscriptionManager) UnregisterSubscription(objectType string, subId string) {
	sm.mx.Lock()
	defer sm.mx.Unlock()
	if _, found := sm.subscriptions[objectType]; found {
		delete(sm.subscriptions[objectType], subId)
		if len(sm.subscriptions[objectType]) == 0 {
			delete(sm.subscriptions, objectType)
		}
	}
}

func (sm *SubscriptionManager) setupIncomingHandler() {
	// todo implement me
	go func() {
		for evnt := range sm.incomingChan {
			//todo go over subscriptions by type, lock subscriber, send object via stream and wait for the response from the stream
			//todo if there is no specific subscription or subscriber - publish message via eventBus for the further processing
			fmt.Println(evnt)
			for objType, subscriberInfo := range sm.subscriptions {
				for subId, subscriber := range subscriberInfo {
					fmt.Println("info", objType, subId)
					cloudObj := evnt.Data.(*cloud.CloudObject)
					err := subscriber.stream.Send(cloudObj)
					if err != nil {
						fmt.Println("Send error:", err)
					}
				}
			}
		}
	}()
}

func (sm *SubscriptionManager) Stop() {
	sm.mx.Lock()
	defer sm.mx.Unlock()
	for _, subscriberInfo := range sm.subscriptions {
		for _, subscriber := range subscriberInfo {
			subscriber.closeChan <- true
		}
	}
}
