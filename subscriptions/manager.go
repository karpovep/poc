package subscriptions

import (
	"github.com/viney-shih/go-lock"
	"io"
	"log"
	"poc/app"
	"poc/bus"
	"poc/model"
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
		casMut    lock.CASMutex
	}

	SubscriptionManager struct {
		EventBus             bus.IEventBus
		Utils                utils.IUtils
		subscriptions        map[string]map[string]*Subscriber //[objectType][subscriptionId]*Subscriber
		inboundChan          bus.DataChannel
		mx                   sync.RWMutex
		inboundChannelName   string
		cachedChannelName    string
		outboundChannelName  string
		processedChannelName string
	}
)

func NewSubscriptionManager(appContext app.IAppContext) ISubscriptionManager {
	eventBus := appContext.Get("eventBus").(bus.IEventBus)
	utls := appContext.Get("utils").(utils.IUtils)
	inboundChan := appContext.Get("inboundChan").(bus.DataChannel)
	inboundChannelName := appContext.Get(model.INBOUND_CHANNEL_NAME).(string)
	cachedChannelName := appContext.Get(model.CACHED_CHANNEL_NAME).(string)
	outboundChannelName := appContext.Get(model.OUTBOUND_CHANNEL_NAME).(string)
	processedChannelName := appContext.Get(model.PROCESSED_CHANNEL_NAME).(string)
	sm := &SubscriptionManager{
		EventBus:             eventBus,
		Utils:                utls,
		subscriptions:        make(map[string]map[string]*Subscriber),
		inboundChan:          inboundChan,
		inboundChannelName:   inboundChannelName,
		cachedChannelName:    cachedChannelName,
		outboundChannelName:  outboundChannelName,
		processedChannelName: processedChannelName,
	}
	sm.setupIncomingHandler()
	sm.EventBus.Subscribe(inboundChannelName, inboundChan)
	sm.EventBus.Subscribe(cachedChannelName, inboundChan)
	return sm
}

func (sm *SubscriptionManager) RegisterSubscription(objectType string, stream cloud.Cloud_SubscribeServer, closeCh chan bool) (string, error) {
	sm.mx.Lock()
	defer sm.mx.Unlock()
	if _, found := sm.subscriptions[objectType]; !found {
		sm.subscriptions[objectType] = make(map[string]*Subscriber)
	}
	subId := sm.Utils.GenerateUuid()
	sm.subscriptions[objectType][subId] = &Subscriber{stream: stream, closeChan: closeCh, casMut: lock.NewCASMutex()}
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
	go func() {
		for evnt := range sm.inboundChan {
			internalServerObject := evnt.Data.(*model.InternalServerObject)
			processed := sm.processObject(internalServerObject)
			if !processed {
				sm.EventBus.Publish(sm.outboundChannelName, internalServerObject)
			} else {
				internalServerObject.Metadata.Status = model.PROCESSED
				sm.EventBus.Publish(sm.processedChannelName, internalServerObject)
			}
		}
	}()
}

func (sm *SubscriptionManager) processObject(obj *model.InternalServerObject) bool {
	objType := obj.Object.Entity.TypeUrl
	for subId, subscriber := range sm.subscriptions[objType] {
		lockAcquired := subscriber.casMut.TryLock()
		if lockAcquired {
			defer subscriber.casMut.Unlock()
			// send object to client for the processing
			log.Println("sending object to subscriber:", subId)
			err := subscriber.stream.Send(obj.Object)
			if err != nil {
				log.Fatal("Send error:", err)
			}
			// wait for the Acknowledgment from the client
			log.Println("waiting acknowledge from subscriber:", subId)
			encodedAck, err := subscriber.stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal("stream.Recv error", err)
			}
			var ack cloud.Acknowledge
			if err := encodedAck.Entity.UnmarshalTo(&ack); err != nil {
				log.Fatalf("Could not unmarshal Acknowledge from any field: %s", err)
			}

			return true
		}
	}
	return false
}

func (sm *SubscriptionManager) Stop() {
	sm.EventBus.Unsubscribe(sm.inboundChannelName, sm.inboundChan)
	sm.EventBus.Unsubscribe(sm.cachedChannelName, sm.inboundChan)
	sm.mx.Lock()
	defer sm.mx.Unlock()
	for _, subscriberInfo := range sm.subscriptions {
		for _, subscriber := range subscriberInfo {
			subscriber.closeChan <- true
		}
	}
}
