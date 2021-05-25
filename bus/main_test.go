package bus

import (
	"poc/app"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldSendPublishedDataToAllSubscribers(t *testing.T) {
	// Given
	appContext := app.NewApplicationContext()
	eb := NewEventBus(appContext)
	topic := "test-topic"
	ch := make(chan DataEvent)
	event := "Some test data to be published"
	subscribesCount := 2

	// subscribes to topic
	for i := 0; i < subscribesCount; i++ {
		eb.Subscribe(topic, ch)
	}

	// When
	eb.Publish(topic, event)

	// Then
	for i := 0; i < subscribesCount; i++ {
		assert.Equal(t, event, (<-ch).Data, "receivedEvent should be equal to the sent one")
	}
}

func Test_ShouldSubscriberShouldNotReceiveDataAfterItUnsubscribed(t *testing.T) {
	// Given
	appContext := app.NewApplicationContext()
	eb := NewEventBus(appContext)
	topic := "testTopic"
	ch := make(chan DataEvent)
	undeliverableEvent := "undeliverableEvent"

	eb.Subscribe(topic, ch)
	eb.Unsubscribe(topic, ch)
	eb.Publish(topic, undeliverableEvent)
	go func() {
		time.Sleep(50 * time.Millisecond)
		close(ch)
	}()

	e, ok := <-ch

	assert.False(t, ok, "Channel shoud be closed")
	assert.True(t, e == DataEvent{}, "Channel shoud be empty")
}
