package subscriptions

import (
	log "github.com/sirupsen/logrus"
	"github.com/viney-shih/go-lock"
	"io"
	"poc/app"
	"poc/bus"
	"poc/model"
	"poc/protos/cloud"
	"poc/protos/nodes"
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
		managerChan          bus.DataChannel
		mx                   sync.RWMutex
		inboundChannelName   string
		transferChannelName  string
		cachedChannelName    string
		outboundChannelName  string
		processedChannelName string
	}
)

func NewSubscriptionManager(appContext app.IAppContext) ISubscriptionManager {
	eventBus := appContext.Get("eventBus").(bus.IEventBus)
	utls := appContext.Get("utils").(utils.IUtils)
	managerChan := eventBus.CreateDataChannel()
	inboundChannelName := appContext.Get(model.INBOUND_CHANNEL_NAME).(string)
	transferChannelName := appContext.Get(model.TRANSFER_CHANNEL_NAME).(string)
	cachedChannelName := appContext.Get(model.CACHED_CHANNEL_NAME).(string)
	outboundChannelName := appContext.Get(model.OUTBOUND_CHANNEL_NAME).(string)
	processedChannelName := appContext.Get(model.PROCESSED_CHANNEL_NAME).(string)
	sm := &SubscriptionManager{
		EventBus:             eventBus,
		Utils:                utls,
		subscriptions:        make(map[string]map[string]*Subscriber),
		managerChan:          managerChan,
		inboundChannelName:   inboundChannelName,
		transferChannelName:  transferChannelName,
		cachedChannelName:    cachedChannelName,
		outboundChannelName:  outboundChannelName,
		processedChannelName: processedChannelName,
	}
	go sm.setupIncomingHandler()
	sm.EventBus.Subscribe(inboundChannelName, managerChan)
	sm.EventBus.Subscribe(transferChannelName, managerChan)
	sm.EventBus.Subscribe(cachedChannelName, managerChan)
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
	log.WithFields(log.Fields{"subscriptionId": subId}).Debug("Subscription registered")
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
	for evnt := range sm.managerChan {
		iso := evnt.Data.(*nodes.ISO)
		if iso.CloudObj.IsFinal {
			// skip object if it is final
			continue
		}
		processed := sm.processObject(iso)
		if !processed {
			sm.EventBus.Publish(sm.outboundChannelName, iso)
		} else {
			sm.EventBus.Publish(sm.processedChannelName, iso)
		}
	}
}

func (sm *SubscriptionManager) processObject(obj *nodes.ISO) bool {
	objType := obj.CloudObj.Entity.TypeUrl
	for subId, subscriber := range sm.subscriptions[objType] {
		lockAcquired := subscriber.casMut.TryLock()
		if lockAcquired {
			defer subscriber.casMut.Unlock()
			// send object to client for the processing
			log.WithFields(log.Fields{"subscriptionId": subId, "id": obj.CloudObj.Id}).Debug("sending object to subscriber")
			err := subscriber.stream.Send(obj.CloudObj)
			if err != nil {
				log.WithFields(log.Fields{"error": err}).Fatal("subscriber.stream.Send error")
			}
			// wait for the Acknowledgment from the client
			log.WithFields(log.Fields{"subscriptionId": subId, "id": obj.CloudObj.Id}).Debug("waiting acknowledge from subscriber")
			encodedAck, err := subscriber.stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.WithFields(log.Fields{"error": err}).Fatal("stream.Recv error")
			}
			var ack cloud.Acknowledge
			if err := encodedAck.Entity.UnmarshalTo(&ack); err != nil {
				log.WithFields(log.Fields{"error": err}).Fatal("Could not unmarshal Acknowledge from any field")
			}
			log.WithFields(log.Fields{"subscriptionId": subId, "id": obj.CloudObj.Id}).Debug("received ACK from client")
		}
	}
	return true
}

func (sm *SubscriptionManager) Stop() {
	sm.EventBus.Unsubscribe(sm.inboundChannelName, sm.managerChan)
	sm.EventBus.Unsubscribe(sm.transferChannelName, sm.managerChan)
	sm.EventBus.Unsubscribe(sm.cachedChannelName, sm.managerChan)
	sm.mx.Lock()
	defer sm.mx.Unlock()
	for _, subscriberInfo := range sm.subscriptions {
		for _, subscriber := range subscriberInfo {
			subscriber.closeChan <- true
		}
	}
}
