package bus

import (
	"fmt"
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
	subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

func NewEventBus() IEventBus {
	return &EventBus{
		subscribers: map[string]DataChannelSlice{},
	}
}

func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}
	eb.rm.Unlock()
}

func (eb *EventBus) Unsubscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	defer eb.rm.Unlock()
	if topicSubscribers, found := eb.subscribers[topic]; found {
		for idx, s := range topicSubscribers {
			if s == ch {
				eb.subscribers[topic] = append(topicSubscribers[0:idx], topicSubscribers[idx+1:]...)
			}
		}
	}
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	fmt.Println("-------->", topic)
	eb.rm.RLock()
	if chans, found := eb.subscribers[topic]; found {
		// this is done because the slices refer to same array even though they are passed by value
		// thus we are creating a new slice with our elements thus preserve locking correctly.
		channels := append(DataChannelSlice{}, chans...)
		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
	eb.rm.RUnlock()
}
