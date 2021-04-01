package bus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldSendPublishedDataToAllSubscribers(t *testing.T) {
	// Given
	eb := NewEventBus()
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
		select {
		case receivedEvent := <-ch:
			assert.Equal(t, event, receivedEvent.Data, "receivedEvent should be equal to the sent one")
		}
	}
}
