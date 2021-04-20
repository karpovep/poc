package bus

import (
	"poc/app/app_mock"
	"poc/utils/utils_mock"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldSendPublishedDataToAllSubscribers(t *testing.T) {
	// Given
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	
	mockUtils := utils_mock.NewMockIUtils(mockCtrl)
	mockSubId1 := "mock-sub-id-1"
	mockSubId2 := "mock-sub-id-2"
	gomock.InOrder(
		mockUtils.EXPECT().GenerateUuid().Return(mockSubId1),
		mockUtils.EXPECT().GenerateUuid().Return(mockSubId2),
	)
	mockAppContext := app_mock.NewMockIAppContext(mockCtrl)
	mockAppContext.EXPECT().Get("utils").Return(mockUtils)
	
	eb := NewEventBus(mockAppContext)
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

func Test_ShouldSubscriberShouldNotReceiveDataAfterUnsibscribed(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	
	mockUtils := utils_mock.NewMockIUtils(mockCtrl)
	mockSubId := "mock-sub-id"
	mockUtils.EXPECT().GenerateUuid().Return(mockSubId)
	mockAppContext := app_mock.NewMockIAppContext(mockCtrl)
	mockAppContext.EXPECT().Get("utils").Return(mockUtils)

	eb := NewEventBus(mockAppContext)
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
