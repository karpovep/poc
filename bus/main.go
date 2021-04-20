package bus

import (
	"poc/app"
	"poc/utils"
	"sync"
)

type DataEvent struct {
	Data  interface{}
	Topic string
}

// DataChannel is a channel which can accept an DataEvent
type DataChannel chan DataEvent

// DataChannelSlice is a slice of DataChannels
type DataChannelSlice []DataChannel

type IEventBus interface {
	Subscribe(topic string, ch DataChannel)
	Unsubscribe(topic string, ch DataChannel)
	Publish(topic string, data interface{})
}

// EventBus stores the information about subscribers interested for // a particular topic
type EventBus struct {
	subscribers map[string]map[string]DataChannel
	rm          sync.RWMutex
	utils       utils.IUtils
}

func NewEventBus(appContext app.IAppContext) IEventBus {
	utls := appContext.Get("utils").(utils.IUtils)
	return &EventBus{
		subscribers: map[string]map[string]DataChannel{},
		utils: utls,
	}
}

func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	defer eb.rm.Unlock()
	uuid := eb.utils.GenerateUuid()
	if _, found := eb.subscribers[topic]; !found {
		eb.subscribers[topic] = make(map[string]DataChannel)
	}
	eb.subscribers[topic][uuid] = ch
}

func (eb *EventBus) Unsubscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	defer eb.rm.Unlock()
	if topicSubscribers, found := eb.subscribers[topic]; found {
		for id, s := range topicSubscribers {
			if s == ch {
				delete(topicSubscribers, id)
			}
		}
	} 
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.rm.RLock()
	if chans, found := eb.subscribers[topic]; found {
		// this is done because the slices refer to same array even though they are passed by value
		// thus we are creating a new slice with our elements thus preserve locking correctly.
		// channels := append(DataChannelSlice{}, chans...)
		channels := make([]DataChannel, 0, len(chans))
		for _, v := range chans {
			channels = append(channels, v)
		}

		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
	eb.rm.RUnlock()
}
